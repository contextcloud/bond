variable "assignments" {
  type = list(object({
    account_name        = string
    group_name          = string
    permission_set_name = string
  }))

  default = []
}
