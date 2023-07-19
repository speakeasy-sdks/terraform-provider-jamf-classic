terraform {
  required_providers {
    jamf = {
      source  = "halosync/jamf"
      version = "1.1.2"
    }
  }
}

provider "jamf" {
  # Configuration options
}