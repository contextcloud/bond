output "acm_certificate_arn" {
  description = "The ARN of the certificate"
  value       = aws_acm_certificate.this.arn
}

output "acm_certificate_domain_validation_options" {
  description = "A list of attributes to feed into other resources to complete certificate validation. Can have more than one element, e.g. if SANs are defined. Only set if DNS-validation was used."
  value       = flatten(aws_acm_certificate.this.domain_validation_options)
}

output "acm_certificate_status" {
  description = "Status of the certificate."
  value       = try(aws_acm_certificate.this.status, "")
}

output "validation_route53_record_fqdns" {
  description = "List of FQDNs built using the zone domain and name."
  value       = [for record in aws_route53_record.this : record.fqdn]
} 

output "all_domains" {
  description = "All domains (including the primary domain)."
  value       = local.all_domains
}

output "domain_to_zone" {
  description = "Domain to zone mapping."
  value       = local.domain_to_zone
}

output "unique_zones" {
  description = "Unique zones."
  value       = local.unique_zones
}