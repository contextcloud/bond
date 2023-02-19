variable "zone_id" {
  description = "ID of DNS zone"
  type        = string
  default     = null
}

variable "zone_name" {
  description = "Name of DNS zone"
  type        = string
  default     = null
}

variable "private_zone" {
  description = "Whether Route53 zone is private or public"
  type        = bool
  default     = false
}

variable "records" {
  description = "List of objects of DNS records"
  type        = list(object({
    name = string
    type = string
    ttl  = optional(number)
    records = optional(list(string))
    set_identifier = optional(string)
    health_check_id = optional(string)
    multivalue_answer_routing_policy = optional(bool)
    allow_overwrite = optional(bool)
    alias = optional(object({
      name = string
      zone_id = string
      evaluate_target_health = bool
    }))
    failover_routing_policy = optional(object({
      type = string
    }))
    latency_routing_policy = optional(object({
      region = string
    }))
    weighted_routing_policy = optional(object({
      weight = number
    }))
    geolocation_routing_policy = optional(object({
      continent = string
      country = string
      subdivision = string
    }))
  }))
  default     = []
}
