terraform {
    backend "s3" {
        bucket = "gaming-admin-dev-terraform-state"
        key = "mattsRemoteState/dev.tfstate"
        profile="gaming-admin-dev"
        region="eu-west-1"
    }
}