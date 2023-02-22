variable "aliases" {
  description = "Aliases of the distribution."
  type        = list(string)
  default     = []
}

variable "comment" {
  description = "Any comments you want to include about the distribution."
  type        = string
  default     = null
}

variable "default_root_object" {
  description = "The object that you want CloudFront to return (for example, index.html) when an end user requests the root URL."
  type        = string
  default     = null
}

variable "price_class" {
  description = "The price class for this distribution. One of PriceClass_All, PriceClass_200, PriceClass_100."
  type        = string
  default     = "PriceClass_All"
}

variable "web_acl_id" {
  description = "If you're using AWS WAF to filter CloudFront requests, the Id of the AWS WAF web ACL that is associated with the distribution. The WAF Web ACL must exist in the WAF Global (CloudFront) region and the credentials configuring this argument must have waf:GetWebACL permissions assigned. If using WAFv2, provide the ARN of the web ACL."
  type        = string
  default     = null
}

variable "tags" {
  description = "A map of tags to assign to the resource."
  type        = map(string)
  default     = null
}

variable "retain_on_delete" {
  description = "Disables the distribution instead of deleting it when destroying the resource through Terraform. If this is set, the distribution needs to be deleted manually afterwards."
  type        = bool
  default     = false
}

variable "wait_for_deployment" {
  description = "If enabled, the resource will wait for the distribution status to change from InProgress to Deployed. Setting this to false will skip the process."
  type        = bool
  default     = true
}

variable "origin" {
  description = "One or more origins for this distribution (multiples allowed)."
  type = list(object({
    domain_name         = string
    origin_id           = string
    origin_path         = optional(string)
    connection_attempts = optional(number)
    connection_timeout  = optional(number)
    custom_origin_config = optional(object({
      http_port                = optional(number)
      https_port               = optional(number)
      origin_protocol_policy   = optional(string)
      origin_ssl_protocols     = optional(list(string))
      origin_read_timeout      = optional(number)
      origin_keepalive_timeout = optional(number)
    }))
    origin_shield = optional(object({
      enabled              = bool
      origin_shield_region = string
    }))
    s3_origin_config = optional(object({
      origin_access_identity = optional(string)
    }))
    custom_header = optional(list(object({
      name  = string
      value = string
    })))
  }))
  default = null
}

variable "viewer_certificate" {
  description = "The SSL configuration for this distribution."
  type = object({
    cloudfront_default_certificate = optional(bool)
    acm_certificate_arn            = optional(string)
    iam_certificate_id             = optional(string)
    minimum_protocol_version       = optional(string)
    ssl_support_method             = optional(string)
  })
  default = null
}

variable "geo_restriction" {
  description = "A complex type that controls the countries in which your content is distributed. CloudFront determines the location of your users using MaxMind GeoIP databases."
  type = list(object({
    restriction_type = optional(string)
    locations        = optional(list(string))
  }))
  default = null
}
