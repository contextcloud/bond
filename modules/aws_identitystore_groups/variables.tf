variable "groups" {
  type = list(object({
    display_name = string
    description   = string
  }))
  description = "A list of groups to create in the AWS SSO Identity Store."
}
