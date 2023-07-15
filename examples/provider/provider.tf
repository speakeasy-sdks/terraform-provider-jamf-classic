terraform {
  required_providers {
    jamf = {
      source  = "halosync/jamf"
      version = "1.0.0"
    }
  }
}

provider "jamf" {
  # Configuration options
}