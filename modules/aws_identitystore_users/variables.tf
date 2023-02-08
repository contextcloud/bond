variable "users" {
  type = list(object({
    user_name    = string
    display_name = string
    given_name   = string
    family_name  = string
    email        = string
  }))
  description = "A list of users to create in the AWS SSO Identity Store."
}
