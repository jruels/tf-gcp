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