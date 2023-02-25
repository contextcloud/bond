resource "aws_acm_certificate" "this" {
  domain_name               = var.domain_name
  subject_alternative_names = var.subject_alternative_names
  validation_method         = "DNS"
  tags                      = var.tags

  options {
    certificate_transparency_logging_preference = "ENABLED"
  }

  lifecycle {
    create_before_destroy = true
  }
}

data "aws_route53_zone" "this" {
  for_each = toset(local.unique_zones)

  zone_id  = var.zone_id
  name     = try(length(var.zone_id), 0) == 0 ? (var.zone_name == "" ? each.key : var.zone_name) : null
  provider = aws.dns
}

resource "aws_route53_record" "this" {
  for_each = {
    for dvo in aws_acm_certificate.this.domain_validation_options : dvo.domain_name => {
      name   = dvo.resource_record_name
      record = dvo.resource_record_value
      type   = dvo.resource_record_type
    }
  }
  zone_id         = data.aws_route53_zone.this[local.domain_to_zone[each.key]].id
  ttl             = 60
  allow_overwrite = true
  name            = each.value.name
  type            = each.value.type
  records         = [each.value.record]
  provider        = aws.dns
}

resource "aws_acm_certificate_validation" "this" {
  certificate_arn         = aws_acm_certificate.this.arn
  validation_record_fqdns = [for record in aws_route53_record.this : record.fqdn]
}

locals {
  all_domains = concat(
    [var.domain_name],
    var.subject_alternative_names
  )

  domain_to_zone = {
    for domain in local.all_domains :
    domain => join(".", slice(split(".", domain), 1, length(split(".", domain))))
  }

  unique_zones = distinct(values(local.domain_to_zone))
}
