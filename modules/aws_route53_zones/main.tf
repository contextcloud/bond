resource "aws_route53_zone" "this" {
  for_each = local.zones_map

  name              = each.value.name
  comment           = each.value.comment
  force_destroy     = each.value.force_destroy
  delegation_set_id = each.value.delegation_set_id

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

locals {
  zones_map = { for z in var.zones : z.name => z }
}
