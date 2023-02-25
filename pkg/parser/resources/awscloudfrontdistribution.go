package resources

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/zclconf/go-cty/cty"
)

var schema = &hcl.BodySchema{
	Attributes: []hcl.AttributeSchema{
		{
			Name: "aliases",
		},
		{
			Name: "comment",
		},
		{
			Name: "price_class",
		},
		{
			Name: "origin",
		},
		{
			Name: "viewer_certificate",
		},
		{
			Name: "geo_restriction",
		},
		{
			Name: "tags",
		},
		{
			Name: "default_cache_behavior",
		},
		{
			Name: "ordered_cache_behavior",
		},
	},
}

type AwsCloudfrontDistribution struct {
	Aliases              []string                                          `hcl:"aliases,optional"`
	Comment              *string                                           `hcl:"comment,optional"`
	PriceClass           *string                                           `hcl:"price_class,optional"`
	Origin               []*AwsCloudfrontDistribution_Origin               `hcl:"origin"`
	ViewerCertificate    *AwsCloudfrontDistribution_ViewerCertificate      `hcl:"viewer_certificate"`
	Tags                 map[string]string                                 `hcl:"tags,optional"`
	GeoRestriction       *AwsCloudfrontDistribution_GeoRestriction         `hcl:"geo_restriction"`
	DefaultCacheBehavior *AwsCloudfrontDistribution_DefaultCacheBehavior   `hcl:"default_cache_behavior"`
	OrderedCacheBehavior []*AwsCloudfrontDistribution_OrderedCacheBehavior `hcl:"ordered_cache_behavior"`
}

type AwsCloudfrontDistribution_Origin struct {
	DomainName         *string                                              `cty:"domain_name"`
	OriginId           *string                                              `cty:"origin_id"`
	OriginPath         *string                                              `cty:"origin_path"`
	ConnectionAttempts *int64                                               `cty:"connection_attempts"`
	ConnectionTimeout  *int64                                               `cty:"connection_timeout"`
	CustomOriginConfig *AwsCloudfrontDistribution_Origin_CustomOriginConfig `cty:"custom_origin_config"`
	CustomHeaders      []*AwsCloudfrontDistribution_Origin_CustomHeader     `cty:"custom_header"`
	OriginShield       *AwsCloudfrontDistribution_Origin_OriginShield       `cty:"origin_shield"`
}

type AwsCloudfrontDistribution_Origin_CustomOriginConfig struct {
	HttpPort               *int64  `cty:"http_port"`
	HttpsPort              *int64  `cty:"https_port"`
	OriginProtocolPolicy   *string `cty:"origin_protocol_policy"`
	OriginSslProtocols     []string
	OriginReadTimeout      *int64 `cty:"origin_read_timeout"`
	OriginKeepaliveTimeout *int64 `cty:"origin_keepalive_timeout"`
}

type AwsCloudfrontDistribution_Origin_CustomHeader struct {
	Name  string `cty:"name"`
	Value string `cty:"value"`
}

type AwsCloudfrontDistribution_Origin_OriginShield struct {
	Enabled            bool   `cty:"enabled"`
	OriginShieldRegion string `cty:"origin_shield_region"`
}

type AwsCloudfrontDistribution_ViewerCertificate struct {
	AcmCertificateArn            *string `cty:"acm_certificate_arn"`
	CloudfrontDefaultCertificate *bool   `cty:"cloudfront_default_certificate"`
	IamCertificateId             *string `cty:"iam_certificate_id"`
	MinimumProtocolVersion       *string `cty:"minimum_protocol_version"`
	SslSupportMethod             *string `cty:"ssl_support_method"`
}

type AwsCloudfrontDistribution_GeoRestriction struct {
	Locations       []string `cty:"locations"`
	RestrictionType string   `cty:"restriction_type"`
}

type AwsCloudfrontDistribution_DefaultCacheBehavior struct {
	AllowedMethods             []string `cty:"allowed_methods"`
	CachedMethods              []string `cty:"cached_methods"`
	CachePolicyId              *string  `cty:"cache_policy_id"`
	Compress                   *bool    `cty:"compress"`
	DefaultTtl                 *int64   `cty:"default_ttl"`
	FieldLevelEncryptionId     *string  `cty:"field_level_encryption_id"`
	LambdaFunctionAssociations []*AwsCloudfrontDistribution_CacheBehavior_LambdaFunctionAssociation
	FunctionAssociations       []*AwsCloudfrontDistribution_CacheBehavior_FunctionAssociation
	MaxTtl                     *int64   `cty:"max_ttl"`
	MinTtl                     *int64   `cty:"min_ttl"`
	OriginRequestPolicyId      *string  `cty:"origin_request_policy_id"`
	RealtimeLogConfigArn       *string  `cty:"realtime_log_config_arn"`
	ResponseHeadersPolicyId    *string  `cty:"response_headers_policy_id"`
	SmoothStreaming            *bool    `cty:"smooth_streaming"`
	TargetOriginId             string   `cty:"target_origin_id"`
	TrustedKeyGroups           []string `cty:"trusted_key_groups"`
	TrustedSigners             []string `cty:"trusted_signers"`
	ViewerProtocolPolicy       string   `cty:"viewer_protocol_policy"`
}

type AwsCloudfrontDistribution_OrderedCacheBehavior struct {
	AllowedMethods             []string `cty:"allowed_methods"`
	CachedMethods              []string `cty:"cached_methods"`
	CachePolicyId              *string  `cty:"cache_policy_id"`
	Compress                   *bool    `cty:"compress"`
	DefaultTtl                 *int64   `cty:"default_ttl"`
	FieldLevelEncryptionId     *string  `cty:"field_level_encryption_id"`
	LambdaFunctionAssociations []*AwsCloudfrontDistribution_CacheBehavior_LambdaFunctionAssociation
	FunctionAssociations       []*AwsCloudfrontDistribution_CacheBehavior_FunctionAssociation
	MaxTtl                     *int64   `cty:"max_ttl"`
	MinTtl                     *int64   `cty:"min_ttl"`
	OriginRequestPolicyId      *string  `cty:"origin_request_policy_id"`
	PathPattern                string   `cty:"path_pattern"`
	RealtimeLogConfigArn       *string  `cty:"realtime_log_config_arn"`
	ResponseHeadersPolicyId    *string  `cty:"response_headers_policy_id"`
	SmoothStreaming            *bool    `cty:"smooth_streaming"`
	TargetOriginId             string   `cty:"target_origin_id"`
	TrustedKeyGroups           []string `cty:"trusted_key_groups"`
	TrustedSigners             []string `cty:"trusted_signers"`
	ViewerProtocolPolicy       string   `cty:"viewer_protocol_policy"`
}

type AwsCloudfrontDistribution_CacheBehavior_LambdaFunctionAssociation struct {
	EventType   string `cty:"event_type"`
	LambdaArn   string `cty:"lambda_arn"`
	IncludeBody *bool  `cty:"include_body"`
}

type AwsCloudfrontDistribution_CacheBehavior_FunctionAssociation struct {
	EventType   string `cty:"event_type"`
	FunctionArn string `cty:"function_arn"`
}

func CustomOriginConfigFactory(v cty.Value) *AwsCloudfrontDistribution_Origin_CustomOriginConfig {
	m := v.AsValueMap()

	out := &AwsCloudfrontDistribution_Origin_CustomOriginConfig{}
	for k, v := range m {
		switch k {
		case "http_port":
			n, _ := v.AsBigFloat().Int64()
			out.HttpPort = &n
		case "https_port":
			n, _ := v.AsBigFloat().Int64()
			out.HttpsPort = &n
		case "origin_protocol_policy":
			str := v.AsString()
			out.OriginProtocolPolicy = &str
		case "origin_ssl_protocols":
			l := v.AsValueSlice()
			out.OriginSslProtocols = make([]string, len(l))
			for i, v := range l {
				out.OriginSslProtocols[i] = v.AsString()
			}
		case "origin_read_timeout":
			n, _ := v.AsBigFloat().Int64()
			out.OriginReadTimeout = &n
		case "origin_keepalive_timeout":
			n, _ := v.AsBigFloat().Int64()
			out.OriginKeepaliveTimeout = &n
		}
	}

	return out
}
func CustomHeadersFactory(v cty.Value) []*AwsCloudfrontDistribution_Origin_CustomHeader {
	l := v.AsValueSlice()
	out := make([]*AwsCloudfrontDistribution_Origin_CustomHeader, len(l))
	for i, v := range l {
		m := v.AsValueMap()
		out[i] = &AwsCloudfrontDistribution_Origin_CustomHeader{}
		for k, v := range m {
			switch k {
			case "name":
				out[i].Name = v.AsString()
			case "value":
				out[i].Value = v.AsString()
			}
		}
	}
	return out
}
func OriginShieldFactory(v cty.Value) *AwsCloudfrontDistribution_Origin_OriginShield {
	m := v.AsValueMap()
	out := &AwsCloudfrontDistribution_Origin_OriginShield{}
	for k, v := range m {
		switch k {
		case "enabled":
			out.Enabled = v.True()
		case "origin_shield_region":
			out.OriginShieldRegion = v.AsString()
		}
	}
	return out
}
func DefaultCacheBehaviorFactory(v cty.Value) *AwsCloudfrontDistribution_DefaultCacheBehavior {
	m := v.AsValueMap()
	out := &AwsCloudfrontDistribution_DefaultCacheBehavior{}
	for k, v := range m {
		switch k {
		case "allowed_methods":
			l := v.AsValueSlice()
			out.AllowedMethods = make([]string, len(l))
			for i, v := range l {
				out.AllowedMethods[i] = v.AsString()
			}
		case "cached_methods":
			l := v.AsValueSlice()
			out.CachedMethods = make([]string, len(l))
			for i, v := range l {
				out.CachedMethods[i] = v.AsString()
			}
		case "cache_policy_id":
			str := v.AsString()
			out.CachePolicyId = &str
		case "compress":
			b := v.True()
			out.Compress = &b
		case "default_ttl":
			n, _ := v.AsBigFloat().Int64()
			out.DefaultTtl = &n
		case "field_level_encryption_id":
			str := v.AsString()
			out.FieldLevelEncryptionId = &str
		case "lambda_function_associations":
			l := v.AsValueSlice()
			out.LambdaFunctionAssociations = make([]*AwsCloudfrontDistribution_CacheBehavior_LambdaFunctionAssociation, len(l))
			for i, v := range l {
				m := v.AsValueMap()
				out.LambdaFunctionAssociations[i] = &AwsCloudfrontDistribution_CacheBehavior_LambdaFunctionAssociation{}
				for k, v := range m {
					switch k {
					case "event_type":
						out.LambdaFunctionAssociations[i].EventType = v.AsString()
					case "lambda_arn":
						out.LambdaFunctionAssociations[i].LambdaArn = v.AsString()
					case "include_body":
						b := v.True()
						out.LambdaFunctionAssociations[i].IncludeBody = &b
					}
				}
			}
		case "function_associations":
			l := v.AsValueSlice()
			out.FunctionAssociations = make([]*AwsCloudfrontDistribution_CacheBehavior_FunctionAssociation, len(l))
			for i, v := range l {
				m := v.AsValueMap()
				out.FunctionAssociations[i] = &AwsCloudfrontDistribution_CacheBehavior_FunctionAssociation{}
				for k, v := range m {
					switch k {
					case "event_type":
						out.FunctionAssociations[i].EventType = v.AsString()
					case "function_arn":
						out.FunctionAssociations[i].FunctionArn = v.AsString()
					}
				}
			}
		case "max_ttl":
			n, _ := v.AsBigFloat().Int64()
			out.MaxTtl = &n
		case "min_ttl":
			n, _ := v.AsBigFloat().Int64()
			out.MinTtl = &n
		case "origin_request_policy_id":
			str := v.AsString()
			out.OriginRequestPolicyId = &str
		case "realtime_log_config_arn":
			str := v.AsString()
			out.RealtimeLogConfigArn = &str
		case "response_headers_policy_id":
			str := v.AsString()
			out.ResponseHeadersPolicyId = &str
		case "smooth_streaming":
			b := v.True()
			out.SmoothStreaming = &b
		case "target_origin_id":
			out.TargetOriginId = v.AsString()
		case "trusted_key_groups":
			l := v.AsValueSlice()
			out.TrustedKeyGroups = make([]string, len(l))
			for i, v := range l {
				out.TrustedKeyGroups[i] = v.AsString()
			}
		case "trusted_signers":
			l := v.AsValueSlice()
			out.TrustedSigners = make([]string, len(l))
			for i, v := range l {
				out.TrustedSigners[i] = v.AsString()
			}
		case "viewer_protocol_policy":
			out.ViewerProtocolPolicy = v.AsString()
		}
	}
	return out
}
func OrderedCacheBehaviorFactory(v cty.Value) *AwsCloudfrontDistribution_OrderedCacheBehavior {
	m := v.AsValueMap()
	out := &AwsCloudfrontDistribution_OrderedCacheBehavior{}
	for k, v := range m {
		switch k {
		case "allowed_methods":
			l := v.AsValueSlice()
			out.AllowedMethods = make([]string, len(l))
			for i, v := range l {
				out.AllowedMethods[i] = v.AsString()
			}
		case "cached_methods":
			l := v.AsValueSlice()
			out.CachedMethods = make([]string, len(l))
			for i, v := range l {
				out.CachedMethods[i] = v.AsString()
			}
		case "cache_policy_id":
			str := v.AsString()
			out.CachePolicyId = &str
		case "compress":
			b := v.True()
			out.Compress = &b
		case "default_ttl":
			n, _ := v.AsBigFloat().Int64()
			out.DefaultTtl = &n
		case "field_level_encryption_id":
			str := v.AsString()
			out.FieldLevelEncryptionId = &str
		case "lambda_function_associations":
			l := v.AsValueSlice()
			out.LambdaFunctionAssociations = make([]*AwsCloudfrontDistribution_CacheBehavior_LambdaFunctionAssociation, len(l))
			for i, v := range l {
				m := v.AsValueMap()
				out.LambdaFunctionAssociations[i] = &AwsCloudfrontDistribution_CacheBehavior_LambdaFunctionAssociation{}
				for k, v := range m {
					switch k {
					case "event_type":
						out.LambdaFunctionAssociations[i].EventType = v.AsString()
					case "lambda_arn":
						out.LambdaFunctionAssociations[i].LambdaArn = v.AsString()
					case "include_body":
						b := v.True()
						out.LambdaFunctionAssociations[i].IncludeBody = &b
					}
				}
			}
		case "function_associations":
			l := v.AsValueSlice()
			out.FunctionAssociations = make([]*AwsCloudfrontDistribution_CacheBehavior_FunctionAssociation, len(l))
			for i, v := range l {
				m := v.AsValueMap()
				out.FunctionAssociations[i] = &AwsCloudfrontDistribution_CacheBehavior_FunctionAssociation{}
				for k, v := range m {
					switch k {
					case "event_type":
						out.FunctionAssociations[i].EventType = v.AsString()
					case "function_arn":
						out.FunctionAssociations[i].FunctionArn = v.AsString()
					}
				}
			}
		case "max_ttl":
			n, _ := v.AsBigFloat().Int64()
			out.MaxTtl = &n
		case "min_ttl":
			n, _ := v.AsBigFloat().Int64()
			out.MinTtl = &n
		case "origin_request_policy_id":
			str := v.AsString()
			out.OriginRequestPolicyId = &str
		case "path_pattern":
			out.PathPattern = v.AsString()
		case "realtime_log_config_arn":
			str := v.AsString()
			out.RealtimeLogConfigArn = &str
		case "response_headers_policy_id":
			str := v.AsString()
			out.ResponseHeadersPolicyId = &str
		case "smooth_streaming":
			b := v.True()
			out.SmoothStreaming = &b
		case "target_origin_id":
			out.TargetOriginId = v.AsString()
		case "trusted_key_groups":
			l := v.AsValueSlice()
			out.TrustedKeyGroups = make([]string, len(l))
			for i, v := range l {
				out.TrustedKeyGroups[i] = v.AsString()
			}
		case "trusted_signers":
			l := v.AsValueSlice()
			out.TrustedSigners = make([]string, len(l))
			for i, v := range l {
				out.TrustedSigners[i] = v.AsString()
			}
		case "viewer_protocol_policy":
			out.ViewerProtocolPolicy = v.AsString()
		}
	}
	return out
}

func AwsCloudfrontDistributionFactory(body hcl.Body) (Resource, error) {
	var out AwsCloudfrontDistribution

	var diags hcl.Diagnostics
	partialContent, _, partialContentDiags := body.PartialContent(schema)
	diags = append(diags, partialContentDiags...)

	for _, attr := range partialContent.Attributes {
		switch attr.Name {
		case "aliases":
			v, vDiags := attr.Expr.Value(nil)
			if vDiags.HasErrors() {
				diags = append(diags, vDiags...)
				continue
			}

			l := v.AsValueSlice()
			out.Aliases = make([]string, len(l))
			for i, v := range l {
				out.Aliases[i] = v.AsString()
			}
		case "comment":
			v, vDiags := attr.Expr.Value(nil)
			if vDiags.HasErrors() {
				diags = append(diags, vDiags...)
				continue
			}
			str := v.AsString()
			out.Comment = &str
		case "price_class":
			v, vDiags := attr.Expr.Value(nil)
			if vDiags.HasErrors() {
				diags = append(diags, vDiags...)
				continue
			}
			str := v.AsString()
			out.PriceClass = &str
		case "origin":
			v, vDiags := attr.Expr.Value(nil)
			if vDiags.HasErrors() {
				diags = append(diags, vDiags...)
				continue
			}

			l := v.AsValueSlice()
			out.Origin = make([]*AwsCloudfrontDistribution_Origin, len(l))
			for i, v := range l {
				out.Origin[i] = &AwsCloudfrontDistribution_Origin{}

				m := v.AsValueMap()
				for k, v := range m {
					switch k {
					case "domain_name":
						str := v.AsString()
						out.Origin[i].DomainName = &str
					case "origin_id":
						str := v.AsString()
						out.Origin[i].OriginId = &str
					case "origin_path":
						str := v.AsString()
						out.Origin[i].OriginPath = &str
					case "connection_attempts":
						n, _ := v.AsBigFloat().Int64()
						out.Origin[i].ConnectionAttempts = &n
					case "connection_timeout":
						n, _ := v.AsBigFloat().Int64()
						out.Origin[i].ConnectionTimeout = &n
					case "custom_origin_config":
						out.Origin[i].CustomOriginConfig = CustomOriginConfigFactory(v)
					case "custom_header":
						out.Origin[i].CustomHeaders = CustomHeadersFactory(v)
					case "origin_shield":
						out.Origin[i].OriginShield = OriginShieldFactory(v)
					}
				}
			}
		case "viewer_certificate":
			v, vDiags := attr.Expr.Value(nil)
			if vDiags.HasErrors() {
				diags = append(diags, vDiags...)
				continue
			}

			out.ViewerCertificate = &AwsCloudfrontDistribution_ViewerCertificate{}
			for k, v := range v.AsValueMap() {
				switch k {
				case "cloudfront_default_certificate":
					b := v.True()
					out.ViewerCertificate.CloudfrontDefaultCertificate = &b
				case "acm_certificate_arn":
					str := v.AsString()
					out.ViewerCertificate.AcmCertificateArn = &str
				case "iam_certificate_id":
					str := v.AsString()
					out.ViewerCertificate.IamCertificateId = &str
				case "minimum_protocol_version":
					str := v.AsString()
					out.ViewerCertificate.MinimumProtocolVersion = &str
				case "ssl_support_method":
					str := v.AsString()
					out.ViewerCertificate.SslSupportMethod = &str
				}
			}
		case "tags":
			v, vDiags := attr.Expr.Value(nil)
			if vDiags.HasErrors() {
				diags = append(diags, vDiags...)
				continue
			}
			m := v.AsValueMap()
			out.Tags = make(map[string]string, len(m))
			for k, v := range m {
				out.Tags[k] = v.AsString()
			}
		case "geo_restriction":
			v, vDiags := attr.Expr.Value(nil)
			if vDiags.HasErrors() {
				diags = append(diags, vDiags...)
				continue
			}
			m := v.AsValueMap()
			out.GeoRestriction = &AwsCloudfrontDistribution_GeoRestriction{}
			for k, v := range m {
				switch k {
				case "restriction_type":
					out.GeoRestriction.RestrictionType = v.AsString()
				case "locations":
					l := v.AsValueSlice()
					out.GeoRestriction.Locations = make([]string, len(l))
					for i, v := range l {
						out.GeoRestriction.Locations[i] = v.AsString()
					}
				}
			}
		case "default_cache_behavior":
			v, vDiags := attr.Expr.Value(nil)
			if vDiags.HasErrors() {
				diags = append(diags, vDiags...)
				continue
			}
			out.DefaultCacheBehavior = DefaultCacheBehaviorFactory(v)
		case "ordered_cache_behavior":
			v, vDiags := attr.Expr.Value(nil)
			if vDiags.HasErrors() {
				diags = append(diags, vDiags...)
				continue
			}
			l := v.AsValueSlice()
			out.OrderedCacheBehavior = make([]*AwsCloudfrontDistribution_OrderedCacheBehavior, len(l))
			for i, v := range l {
				out.OrderedCacheBehavior[i] = OrderedCacheBehaviorFactory(v)
			}
		}
	}

	// if decodeDiags := gohcl.DecodeBody(remainBody, nil, &out); decodeDiags != nil && decodeDiags.HasErrors() {
	// 	diags = append(diags, decodeDiags...)
	// }

	if diags.HasErrors() {
		return nil, diags
	}
	return out, nil
}
