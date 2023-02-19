variable "domain_name" {
  description = "A domain name for which the certificate should be issued"
  type        = string
  default     = ""
}

variable "subject_alternative_names" {
  description = "A list of domains that should be SANs in the issued certificate"
  type        = list(string)
  default     = []
}

variable "zone_id" {
  description = "The ID of the hosted zone to contain this record. Required when validating via Route53"
  type        = string
  default     = ""
}

variable "zone_name" {
  description = "The name of the hosted zone to contain this record. Required when validating via Route53"
  type        = string
  default     = ""
}

variable "tags" {
  description = "A mapping of tags to assign to the resource"
  type        = map(string)
  default     = {}
}

variable "validation_record_fqdns" {
  description = "When validation is set to DNS and the DNS validation records are set externally, provide the fqdns for the validation"
  type        = list(string)
  default     = []
}

variable "acm_certificate_domain_validation_options" {
  description = "A list of domain_validation_options created by the ACM certificate to create required Route53 records from it (used when create_route53_records_only is set to true)"
  type        = any
  default     = {}
}