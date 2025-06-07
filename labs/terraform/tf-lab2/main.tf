terraform {
  required_providers {
    google = {
      source = "hashicorp/google"
    }
  }
}

provider "google" {
  project = "iis-tf-dev"
  region  = "us-west1"
}

resource "google_compute_instance" "lab2-tf-example" {
  name         = var.instance_name
  machine_type = "e2-micro"
  zone         = "us-west1-a"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-11"
    }
  }

  network_interface {
    network = "default"
    access_config {
      // Ephemeral public IP
    }
  }

  labels = {
    name = var.instance_name
  }
}