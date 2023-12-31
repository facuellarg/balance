terraform {
  required_providers {
    aws = {
      source = "hashicorp/aws"
    }
    archive = {
      source = "hashicorp/archive"
    }
    null = {
      source = "hashicorp/null"
    }
  }

  required_version = ">= 1.3.7"
}


module "transactions" {
  source = "./terraform"
}



provider "aws" {
  region = "us-east-1"
  # profile = "tutorial-terraform-profile"

  default_tags {
    tags = {
      app = "tutorial-terraform"
    }
  }
}

output "api_url" {
  value = "${module.transactions.api_url}/transactions"
}
