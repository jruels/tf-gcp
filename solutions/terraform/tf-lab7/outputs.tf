output "prod_website_url" {
  description = "URL of the production website"
  value       = "https://storage.googleapis.com/${google_storage_bucket.prod.name}/index.html"
}

output "dev_website_url" {
  description = "URL of the development website"
  value       = "https://storage.googleapis.com/${google_storage_bucket.dev.name}/index.html"
}