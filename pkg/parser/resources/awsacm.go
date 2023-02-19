package resources

type AwsAcm struct {
	DomainName              string            `hcl:"domain_name"`
	SubjectAlternativeNames []string          `hcl:"subject_alternative_names"`
	ZoneId                  string            `hcl:"zone_id"`
	Tags                    map[string]string `hcl:"tags"`
}

// variable "validation_record_fqdns" {
//   description = "When validation is set to DNS and the DNS validation records are set externally, provide the fqdns for the validation"
//   type        = list(string)
//   default     = []
// }

// variable "acm_certificate_domain_validation_options" {
//   description = "A list of domain_validation_options created by the ACM certificate to create required Route53 records from it (used when create_route53_records_only is set to true)"
//   type        = any
//   default     = {}
// }
