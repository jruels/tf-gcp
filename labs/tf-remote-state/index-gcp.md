# Terraform - remote state configuration

## Overview 
In this lab, you will create a Google Cloud Storage bucket and migrate the Terraform state to a remote backend. 

## Create a GCS bucket 
GCP requires every Cloud Storage bucket to have a unique name. For this reason, add your initials to the end of the bucket. The example below uses `jrs` as the initials.

For the following steps, replace all references to `remote-state-jrs` with your bucket name.

In the `tf-lab3/learn-terraform-variables` directory create a new file `gcs.tf` with the following: 

```hcl
resource "google_storage_bucket" "remote_state" {
    name          = "remote-state-jrs"
    force_destroy = true
    location      = "US"
    
    uniform_bucket_level_access = true
    
    labels = {
        name = "remote state backend"
    }
}
```

So that we can easily retrieve the name of the bucket in the future add the following to `outputs.tf`
```hcl
output "gcs_bucket" {
  description = "GCS bucket name"
  value       = google_storage_bucket.remote_state.name
}
```
Using Terraform apply the changes. 

## Migrate the state
Now that we've created a Cloud Storage bucket, we need to migrate the state to the remote backend. 

When creating the backend configuration remember to replace the `bucket` with the name of the bucket you created. 

Create `backend.tf` with the following:
```hcl
terraform {
  backend "gcs" {
    bucket = "remote-state-jrs"
    prefix = "terraform/state"
  }
}
```

## Reinitialize Terraform 
Now that you have created the GCS bucket and configured the `backend.tf` you must run `terraform init` to migrate the state to the new remote backend. 

If prompted to migrate the existing state type `yes`

If everything is successful, you should see a message that the backend was migrated. 

## Cleanup

Run the following to clean up the resources

```
terraform destroy -auto-approve
```

## Congratulations

You have successfully created a Cloud Storage bucket and migrated to a remote backend state. 