variable "organization_name" {
  type        = string
  description = "Organization Name"
}

variable "accounts" {
  type = list(object({
    name  = string
    email = string
  }))
  description = "Accounts to be created"
}
