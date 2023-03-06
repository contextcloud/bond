resource "aws_route53_delegation_set" "this" {
}

resource "aws_route53_zone" "this" {
  for_each = local.zones_map

  name              = each.value.name
  comment           = each.value.comment
  force_destroy     = true
  delegation_set_id = aws_route53_delegation_set.this.id

  dynamic "vpc" {
    for_each = (each.value.vpc == null) ? [] : each.value.vpc

    content {
      vpc_id     = vpc.value.vpc_id
      vpc_region = vpc.value.vpc_region
    }
  }

  tags = merge(
    each.value.tags,
    var.tags
  )
}

data "aws_route53_zone" "ns" {
  for_each = local.ns_zones_map

  name     = each.key
  provider = aws.ns
}

resource "aws_route53_record" "ns" {
  for_each = local.ns_zones_map

  zone_id  = data.aws_route53_zone.ns[each.key].id
  name     = each.value.name
  type     = "NS"
  ttl      = "300"
  records  = aws_route53_zone.this[each.value.name].name_servers
  provider = aws.ns
}

locals {
  zones_map    = { for z in var.zones : z.name => z }
  ns_zones_map = { for z in var.zones : z.ns_zone_name => z if z.ns_zone_name != null }
}
