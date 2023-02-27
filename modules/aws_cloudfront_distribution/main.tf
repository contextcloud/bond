

resource "aws_cloudfront_cache_policy" "default" {
  name        = "default-cache-policy"
  comment     = "Default cache policy"
  min_ttl     = var.default_cache_behavior.min_ttl
  default_ttl = var.default_cache_behavior.default_ttl
  max_ttl     = var.default_cache_behavior.max_ttl

  parameters_in_cache_key_and_forwarded_to_origin {
    cookies_config {
      cookie_behavior = var.default_cache_behavior.cache_policy.cookie_behavior
      cookies {
        items = var.default_cache_behavior.cache_policy.cookie_items
      }
    }
    headers_config {
      header_behavior = var.default_cache_behavior.cache_policy.header_behavior
      headers {
        items = var.default_cache_behavior.cache_policy.header_items
      }
    }
    query_strings_config {
      query_string_behavior = var.default_cache_behavior.cache_policy.query_string_behavior
      query_strings {
        items = var.default_cache_behavior.cache_policy.query_string_items
      }
    }
  }
}

resource "aws_cloudfront_distribution" "this" {
  enabled             = true
  http_version        = "http2"
  is_ipv6_enabled     = true
  aliases             = var.aliases
  comment             = var.comment
  default_root_object = var.default_root_object
  price_class         = var.price_class
  web_acl_id          = var.web_acl_id
  tags                = var.tags
  retain_on_delete    = var.retain_on_delete
  wait_for_deployment = var.wait_for_deployment

  dynamic "origin" {
    for_each = var.origin

    content {
      domain_name         = origin.value.domain_name
      origin_id           = origin.value.origin_id
      origin_path         = origin.value.origin_path
      connection_attempts = origin.value.connection_attempts
      connection_timeout  = origin.value.connection_timeout
      /* origin_access_control_id = lookup(origin.value, "origin_access_control_id", lookup(lookup(aws_cloudfront_origin_access_control.this, lookup(origin.value, "origin_access_control", ""), {}), "id", null)) */

      dynamic "s3_origin_config" {
        for_each = origin.value.s3_origin_config == null ? [] : [origin.value.s3_origin_config]

        content {
          origin_access_identity = s3_origin_config.value.origin_access_identity
        }
      }

      dynamic "custom_origin_config" {
        for_each = origin.value.custom_origin_config == null ? [] : [origin.value.custom_origin_config]

        content {
          http_port                = custom_origin_config.value.http_port
          https_port               = custom_origin_config.value.https_port
          origin_protocol_policy   = custom_origin_config.value.origin_protocol_policy
          origin_ssl_protocols     = custom_origin_config.value.origin_ssl_protocols
          origin_keepalive_timeout = custom_origin_config.value.origin_keepalive_timeout
          origin_read_timeout      = custom_origin_config.value.origin_read_timeout
        }
      }

      dynamic "custom_header" {
        for_each = origin.value.custom_header

        content {
          name  = custom_header.value.name
          value = custom_header.value.value
        }
      }

      dynamic "origin_shield" {
        for_each = origin.value.origin_shield == null ? [] : [origin.value.origin_shield]

        content {
          enabled              = origin_shield.value.enabled
          origin_shield_region = origin_shield.value.origin_shield_region
        }
      }
    }
  }

  viewer_certificate {
    acm_certificate_arn            = local.acm_certificate_arn
    cloudfront_default_certificate = local.cloudfront_default_certificate
    ssl_support_method             = local.cloudfront_default_certificate ? null : "sni-only"
    minminimum_protocol_version    = local.cloudfront_default_certificate ? "TLSv1" : "TLSv1.2_2019"
  }

  restrictions {
    dynamic "geo_restriction" {
      for_each = var.geo_restriction == null ? [] : [var.geo_restriction]

      content {
        restriction_type = geo_restriction.value.restriction_type
        locations        = geo_restriction.value.locations
      }
    }
  }

  dynamic "default_cache_behavior" {
    for_each = toset([var.default_cache_behavior])
    iterator = i

    content {
      target_origin_id       = i.value.target_origin_id
      viewer_protocol_policy = i.value.viewer_protocol_policy

      allowed_methods           = i.value.allowed_methods
      cached_methods            = i.value.cached_methods
      compress                  = i.value.compress
      field_level_encryption_id = i.value.field_level_encryption_id
      smooth_streaming          = i.value.smooth_streaming
      trusted_signers           = i.value.trusted_signers
      trusted_key_groups        = i.value.trusted_key_groups

      cache_policy_id            = aws_cloudfront_cache_policy.default.id
      origin_request_policy_id   = i.value.origin_request_policy_id
      response_headers_policy_id = i.value.response_headers_policy_id
      realtime_log_config_arn    = i.value.realtime_log_config_arn

      min_ttl     = i.value.min_ttl
      default_ttl = i.value.default_ttl
      max_ttl     = i.value.max_ttl

      dynamic "lambda_function_association" {
        for_each = i.value.lambda_function_association == null ? [] : i.value.lambda_function_association
        iterator = l

        content {
          event_type   = l.value.event_type
          lambda_arn   = l.value.lambda_arn
          include_body = l.value.include_body
        }
      }

      dynamic "function_association" {
        for_each = i.value.function_association == null ? [] : i.value.function_association
        iterator = f

        content {
          event_type   = f.value.event_type
          function_arn = f.value.function_arn
        }
      }
    }
  }

  dynamic "ordered_cache_behavior" {
    for_each = toset(var.ordered_cache_behavior)
    iterator = i

    content {
      target_origin_id       = i.value.target_origin_id
      viewer_protocol_policy = i.value.viewer_protocol_policy

      allowed_methods           = i.value.allowed_methods
      cached_methods            = i.value.cached_methods
      compress                  = i.value.compress
      field_level_encryption_id = i.value.field_level_encryption_id
      smooth_streaming          = i.value.smooth_streaming
      trusted_signers           = i.value.trusted_signers
      trusted_key_groups        = i.value.trusted_key_groups

      /* cache_policy_id            = aws_cloudfront_cache_policy.default.id */
      origin_request_policy_id   = i.value.origin_request_policy_id
      path_pattern               = i.value.path_pattern
      response_headers_policy_id = i.value.response_headers_policy_id
      realtime_log_config_arn    = i.value.realtime_log_config_arn

      min_ttl     = i.value.min_ttl
      default_ttl = i.value.default_ttl
      max_ttl     = i.value.max_ttl

      dynamic "lambda_function_association" {
        for_each = i.value.lambda_function_association == null ? [] : i.value.lambda_function_association
        iterator = l

        content {
          event_type   = l.value.event_type
          lambda_arn   = l.value.lambda_arn
          include_body = l.value.include_body
        }
      }

      dynamic "function_association" {
        for_each = i.value.function_association == null ? [] : i.value.function_association
        iterator = f

        content {
          event_type   = f.value.event_type
          function_arn = f.value.function_arn
        }
      }
    }
  }
}

# Find a certificate that is issued
data "aws_acm_certificate" "issued" {
  count = length(coalesce(var.aliases, [])) > 0 ? 1 : 0

  domain      = var.aliases[0]
  statuses    = ["ISSUED"]
  most_recent = true
}

locals {
  acm_certificate_arn            = try(data.aws_acm_certificate.issued[0].arn, null)
  cloudfront_default_certificate = local.acm_certificate_arn == null ? true : false
}
