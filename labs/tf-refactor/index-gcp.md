# Terraform Refactor codebase

## Overview 
In this lab, you will refactor a monolithic Terraform configuration into a modular design. 

You will provision two instances of a web application hosted in a Cloud Storage bucket that represent production and development environments. The configuration you use to deploy the application will start as a monolith. You will modify it to step through the common phases of evolution for a Terraform project, until each environment has its own independent configuration and state.

## Start with a monolith configuration
### Create the Lab Directory

1. In **Visual Studio Code**, open the working directory created in the previous lab (`YYYYMMDD/terraform`).
2. Right-click in the **Explorer** pane and select **New Folder**.
3. Name the folder `tf-refactor`.
4. Create the following files in your `tf-refactor` directory:

`main.tf` - configures the resources that make up your infrastructure:

```hcl
terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "~> 4.0"
    }
    random = {
      source  = "hashicorp/random"
      version = "~> 3.1.0"
    }
  }
}

provider "google" {
  project = var.project_id
  region  = var.region
}

resource "random_pet" "petname" {
  length    = 3
  separator = "-"
}

# Production bucket
resource "google_storage_bucket" "prod" {
  name          = "${var.prod_prefix}-${random_pet.petname.id}"
  location      = var.region
  force_destroy = true

  uniform_bucket_level_access = true

  website {
    main_page_suffix = "index.html"
    not_found_page   = "404.html"
  }
}

resource "google_storage_bucket_iam_member" "prod" {
  bucket = google_storage_bucket.prod.name
  role   = "roles/storage.objectViewer"
  member = "allUsers"
}

resource "google_storage_bucket_object" "prod" {
  name          = "index.html"
  bucket        = google_storage_bucket.prod.name
  content       = file("${path.module}/assets/index.html")
  content_type  = "text/html"
}

# Development bucket
resource "google_storage_bucket" "dev" {
  name          = "${var.dev_prefix}-${random_pet.petname.id}"
  location      = var.region
  force_destroy = true

  uniform_bucket_level_access = true

  website {
    main_page_suffix = "index.html"
    not_found_page   = "404.html"
  }
}

resource "google_storage_bucket_iam_member" "dev" {
  bucket = google_storage_bucket.dev.name
  role   = "roles/storage.objectViewer"
  member = "allUsers"
}

resource "google_storage_bucket_object" "dev" {
  name          = "index.html"
  bucket        = google_storage_bucket.dev.name
  content       = file("${path.module}/assets/index.html")
  content_type  = "text/html"
}
```

`variables.tf` - declares input variables:

```hcl
variable "project_id" {
  description = "Google Cloud Project ID"
  type        = string
}

variable "region" {
  description = "Region for GCP resources"
  type        = string
  default     = "us-central1"
}

variable "prod_prefix" {
  description = "Prefix for production resources"
  type        = string
  default     = "prod"
}

variable "dev_prefix" {
  description = "Prefix for development resources"
  type        = string
  default     = "dev"
}
```

`terraform.tfvars` - defines your variables:

```hcl
project_id  = "YOUR_PROJECT_ID"
region      = "us-central1"
prod_prefix = "prod"
dev_prefix  = "dev"
```

`outputs.tf` - specifies the website endpoints:

```hcl
output "prod_website_url" {
  description = "URL of the production website"
  value       = "https://storage.googleapis.com/${google_storage_bucket.prod.name}/index.html"
}

output "dev_website_url" {
  description = "URL of the development website"
  value       = "https://storage.googleapis.com/${google_storage_bucket.dev.name}/index.html"
}
```

Create an `assets` directory and add an `index.html` file with some sample content:

```html
<!DOCTYPE html>
<html>
<head>
    <title>My Static Website</title>
</head>
<body>
    <h1>Welcome to my website!</h1>
</body>
</html>
```

Initialize and apply your Terraform configuration:

```sh
terraform init
terraform apply
```

Navigate to the web addresses from the Terraform output to display the deployments in a browser.

## Separate the configuration

Defining multiple environments in the same `main.tf` file may become hard to manage as you add more resources. HCL is meant to be human-readable and supports using multiple configuration files to help organize your infrastructure.

1. Copy `main.tf` to `dev.tf`
2. Rename `main.tf` to `prod.tf`
3. In `dev.tf`, remove all production-related resources (resources with `prod` in their names)
4. In `prod.tf`, remove all development-related resources (resources with `dev` in their names)

Your directory structure will look like: 
```
.
├── README.md
├── assets
│   └── index.html
├── dev.tf
├── outputs.tf
├── prod.tf
├── terraform.tfstate
├── terraform.tfvars
└── variables.tf
```

Comment out the following in `prod.tf` to avoid duplicate resource definitions:
- The `terraform` block
- The `provider` block
- The `random_pet` resource

## Simulate a hidden dependency

Update the `random_pet` resource in `dev.tf` to use 4 words instead of 3:

```hcl
resource "random_pet" "petname" {
  length    = 4
  separator = "-"
}
```

Apply the changes:

```sh
terraform apply
```

Notice that both environments are affected because they share the same `random_pet` resource.

Clean up before continuing:

```sh
terraform destroy
```

## Separate states

### Create environment directories

1. Create directories for each environment:
```sh
mkdir prod && mkdir dev
```

2. Move the files to their respective directories:
- Move `dev.tf` to `dev/main.tf`
- Move `prod.tf` to `prod/main.tf`
- Copy `variables.tf`, `terraform.tfvars`, and `outputs.tf` to both directories

3. Update the asset path in both `main.tf` files:
```hcl
content = file("${path.module}/../assets/index.html")
```

4. Remove production-related content from development files:
- Remove prod outputs from `dev/outputs.tf`
- Remove prod variables from `dev/variables.tf`
- Remove prod values from `dev/terraform.tfvars`

5. Remove development-related content from production files:
- Remove dev outputs from `prod/outputs.tf`
- Remove dev variables from `prod/variables.tf`
- Remove dev values from `prod/terraform.tfvars`

6. Uncomment the provider and random_pet blocks in `prod/main.tf`

Your final structure should look like:
```
.
├── assets
│   ├── index.html
├── prod
│   ├── main.tf
│   ├── variables.tf
│   ├── terraform.tfstate
│   └── terraform.tfvars
└── dev
    ├── main.tf
    ├── variables.tf
    ├── terraform.tfstate
    └── terraform.tfvars
```

### Deploy environments 

Deploy the development environment:
```sh
cd dev
terraform init
terraform apply
```

Deploy the production environment:
```sh
cd ../prod
terraform init
terraform apply
```

Visit both website URLs to confirm they're working.

## Cleanup

Clean up both environments:
```sh
# In prod directory
terraform destroy

cd ../dev
terraform destroy
```

## Congratulations!

You have successfully:
1. Created a monolithic configuration for multiple environments
2. Split the configuration into separate files
3. Separated the state files by using directories
4. Deployed independent development and production environments 