package resources

import "github.com/hashicorp/hcl/v2"

type AwsRoute53Records struct {
	ZoneId      *string            `hcl:"zone_id"`
	ZoneName    *string            `hcl:"zone_name"`
	PrivateZone *bool              `hcl:"private_zone"`
	Records     []AwsRoute53Record `hcl:"records"`
}

type AwsRoute53Record struct {
	Name                          string                                    `cty:"name"`
	Type                          string                                    `cty:"type"`
	TTL                           *int64                                    `cty:"ttl"`
	Records                       []string                                  `cty:"records"`
	SetIdentifier                 *string                                   `cty:"set_identifier"`
	HealthCheckId                 *string                                   `cty:"health_check_id"`
	MultivalueAnswerRoutingPolicy *bool                                     `cty:"multivalue_answer_routing_policy"`
	AllowOverwrite                *bool                                     `cty:"allow_overwrite"`
	Alias                         *AwsRoute53RecordAlias                    `cty:"alias"`
	FailoverRoutingPolicy         *AwsRoute53RecordFailoverRoutingPolicy    `cty:"failover_routing_policy"`
	LatencyRoutingPolicy          *AwsRoute53RecordLatencyRoutingPolicy     `cty:"latency_routing_policy"`
	WeightedRoutingPolicy         *AwsRoute53RecordWeightedRoutingPolicy    `cty:"weighted_routing_policy"`
	GeolocationRoutingPolicy      *AwsRoute53RecordGeolocationRoutingPolicy `cty:"geolocation_routing_policy"`
}

type AwsRoute53RecordAlias struct {
	Name                 string `cty:"name"`
	ZoneId               string `cty:"zone_id"`
	EvaluateTargetHealth bool   `cty:"evaluate_target_health"`
}

type AwsRoute53RecordFailoverRoutingPolicy struct {
	Type string `cty:"type"`
}

type AwsRoute53RecordLatencyRoutingPolicy struct {
	Region string `cty:"region"`
}

type AwsRoute53RecordWeightedRoutingPolicy struct {
	Weight int `cty:"weight"`
}

type AwsRoute53RecordGeolocationRoutingPolicy struct {
	Continent   string `cty:"continent"`
	Country     string `cty:"country"`
	Subdivision string `cty:"subdivision"`
}

func AwsRoute53RecordsFactory(body hcl.Body) (Resource, error) {
	var out AwsRoute53Records

	var diags hcl.Diagnostics
	attrs, attrsDiags := body.JustAttributes()
	diags = append(diags, attrsDiags...)

	for _, attr := range attrs {
		switch attr.Name {
		case "zone_id":
			v, vDiags := attr.Expr.Value(nil)
			if vDiags.HasErrors() {
				diags = append(diags, vDiags...)
				continue
			}
			str := v.AsString()
			out.ZoneId = &str
		case "zone_name":
			v, vDiags := attr.Expr.Value(nil)
			if vDiags.HasErrors() {
				diags = append(diags, vDiags...)
				continue
			}
			str := v.AsString()
			out.ZoneName = &str
		case "private_zone":
			v, vDiags := attr.Expr.Value(nil)
			if vDiags.HasErrors() {
				diags = append(diags, vDiags...)
				continue
			}
			b := v.True()
			out.PrivateZone = &b
		case "records":
			v, vDiags := attr.Expr.Value(nil)
			if vDiags.HasErrors() {
				diags = append(diags, vDiags...)
				continue
			}

			l := v.AsValueSlice()
			out.Records = make([]AwsRoute53Record, len(l))
			for i, v := range l {
				m := v.AsValueMap()
				for k, v := range m {
					switch k {
					case "name":
						out.Records[i].Name = v.AsString()
					case "type":
						out.Records[i].Type = v.AsString()
					case "ttl":
						bf, _ := v.AsBigFloat().Int64()
						out.Records[i].TTL = &bf
					case "records":
						l := v.AsValueSlice()
						out.Records[i].Records = make([]string, len(l))
						for j, v := range l {
							out.Records[i].Records[j] = v.AsString()
						}
					case "set_identifier":
						str := v.AsString()
						out.Records[i].SetIdentifier = &str
					case "health_check_id":
						str := v.AsString()
						out.Records[i].HealthCheckId = &str
					case "multivalue_answer_routing_policy":
						b := v.True()
						out.Records[i].MultivalueAnswerRoutingPolicy = &b
					case "allow_overwrite":
						b := v.True()
						out.Records[i].AllowOverwrite = &b
					case "alias":
						m := v.AsValueMap()
						for k, v := range m {
							switch k {
							case "name":
								out.Records[i].Alias.Name = v.AsString()
							case "zone_id":
								out.Records[i].Alias.ZoneId = v.AsString()
							case "evaluate_target_health":
								out.Records[i].Alias.EvaluateTargetHealth = v.True()
							}
						}
					case "failover_routing_policy":
						m := v.AsValueMap()
						for k, v := range m {
							switch k {
							case "type":
								out.Records[i].FailoverRoutingPolicy.Type = v.AsString()
							}
						}
					case "latency_routing_policy":
						m := v.AsValueMap()
						for k, v := range m {
							switch k {
							case "region":
								out.Records[i].LatencyRoutingPolicy.Region = v.AsString()
							}
						}
					case "weighted_routing_policy":
						m := v.AsValueMap()
						for k, v := range m {
							switch k {
							case "weight":
								bf, _ := v.AsBigFloat().Int64()
								out.Records[i].WeightedRoutingPolicy.Weight = int(bf)
							}
						}
					case "geolocation_routing_policy":
						m := v.AsValueMap()
						for k, v := range m {
							switch k {
							case "continent":
								out.Records[i].GeolocationRoutingPolicy.Continent = v.AsString()
							case "country":
								out.Records[i].GeolocationRoutingPolicy.Country = v.AsString()
							case "subdivision":
								out.Records[i].GeolocationRoutingPolicy.Subdivision = v.AsString()
							}
						}
					}
				}
			}
		}
	}

	if diags.HasErrors() {
		return nil, diags
	}
	return out, nil
}
