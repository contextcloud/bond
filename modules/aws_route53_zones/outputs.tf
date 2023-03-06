output "route53_delegation_set_id" {
  description = "Delegation set ID of Route53 zone"
  value       = aws_route53_delegation_set.this.id
}

output "route53_delegation_set_arn" {
  description = "Delegation set ARN of Route53 zone"
  value       = aws_route53_delegation_set.this.arn
}

output "route53_delegation_set_name_servers" {
  description = "Name servers of Route53 delegation set"
  value       = aws_route53_delegation_set.this.name_servers
}

output "route53_zone_id" {
  description = "Zone ID of Route53 zone"
  value       = { for k, v in aws_route53_zone.this : k => v.zone_id }
}

output "route53_zone_arn" {
  description = "Zone ARN of Route53 zone"
  value       = { for k, v in aws_route53_zone.this : k => v.arn }
}

output "route53_zone_servers" {
  description = "Name servers of Route53 zone"
  value       = { for k, v in aws_route53_zone.this : k => v.name_servers }
}

output "route53_zone_name" {
  description = "Name of Route53 zone"
  value       = { for k, v in aws_route53_zone.this : k => v.name }
}