variable "members" {
  type = list(object({
    user_name  = string
    group_name = string
  }))
  description = "The list of members to add to the group."
}
