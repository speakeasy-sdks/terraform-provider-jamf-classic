terraform {
  required_providers {
    jamf = {
      source  = "halosync/jamf"
      version = "1.1.0"
    }
  }
}

provider "jamf" {
  # Configuration options
}