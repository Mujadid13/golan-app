terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "3.48.0"
    }
  }

  backend "azurerm" {
    resource_group_name = "dev-opps"
    storage_account_name = "mujadidstorage"
    container_name = "terraformstate"
    key = "terraform.tfstate"
    
  }
}



provider "azurerm" {
  features {
    
  }
}

