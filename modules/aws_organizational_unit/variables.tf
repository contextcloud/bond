variable "organization_name" {
  type        = string
  description = "Organization Name"
}

variable "accounts" {
  type = map(object({
    email = string
    tags  = optional(map(string))
  }))
  description = "Accounts to be created"
}
