terraform {
  required_providers {
    myservice = {
      source  = "donis/myservices"
      version = "1.0.0"
    }
  }
}

provider "myservice" {
  # Optional: configure the service endpoint
  api_base_url = "http://localhost:8080"
}

# Step 1: Create the resource
resource "myservice_item" "example" {
  name = "Created by Terraform"
}

# Step 2: Capture output after first read
output "first_read_name" {
  value = myservice_item.example.name
}

# Step 3: Update the resource
resource "myservice_item" "example_update" {
  depends_on = [myservice_item.example]

  name = "Updated by Terraform"
}

# Step 4: Capture output after second read
output "second_read_name" {
  value = myservice_item.example_update.name
}
