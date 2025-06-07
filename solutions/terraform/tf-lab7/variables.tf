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