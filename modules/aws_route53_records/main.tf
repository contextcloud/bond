resource "aws_route53_record" "this" {
  for_each = local.records_map

  zone_id                          = data.aws_route53_zone.this.zone_id
  name                             = each.value.name
  type                             = each.value.type
  ttl                              = each.value.ttl
  records                          = each.value.records
  set_identifier                   = each.value.set_identifier
  health_check_id                  = each.value.health_check_id
  multivalue_answer_routing_policy = each.value.multivalue_answer_routing_policy
  allow_overwrite                  = each.value.allow_overwrite

  dynamic "alias" {
    for_each = each.value.alias == null ? [] : [true]

    content {
      name                   = each.value.alias.name
      zone_id                = each.value.alias.zone_id
      evaluate_target_health = each.value.alias.evaluate_target_health
    }
  }

  dynamic "failover_routing_policy" {
    for_each = each.value.failover_routing_policy == null ? [] : [true]

    content {
      type = each.value.failover_routing_policy.type
    }
  }

  dynamic "latency_routing_policy" {
    for_each = each.value.latency_routing_policy == null ? [] : [true]

    content {
      region = each.value.latency_routing_policy.region
    }
  }

  dynamic "weighted_routing_policy" {
    for_each = each.value.weighted_routing_policy == null ? [] : [true]

    content {
      weight = each.value.weighted_routing_policy.weight
    }
  }

  dynamic "geolocation_routing_policy" {
    for_each = each.value.geolocation_routing_policy == null ? [] : [true]

    content {
      continent   = each.value.geolocation_routing_policy.continent
      country     = each.value.geolocation_routing_policy.country
      subdivision = each.value.geolocation_routing_policy.subdivision
    }
  }
}

data "aws_route53_zone" "this" {
  zone_id      = var.zone_id
  name         = var.zone_name
  private_zone = var.private_zone
}

locals {
  records_map = { for r in var.records : r.name => r }
}
