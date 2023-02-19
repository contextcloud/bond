variable "zones" {
  type = list(object({
    name              = string
    comment           = string
    force_destroy     = bool
    delegation_set_id = optional(string)
    tags              = optional(map(any))
    vpc = optional(list(object({
      vpc_id     = string
      vpc_region = optional(string)
    })))
  }))
  description = "List of zones to create"
}

variable "tags" {
  description = "Tags added to all zones. Will take precedence over tags from the 'zones' variable"
  type        = map(any)
  default     = {}
}