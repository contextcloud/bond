package resources

import (
	"github.com/hashicorp/hcl/v2"
)

type AwsRoute53Zones struct {
	Zones []AwsRoute53Zone  `hcl:"zones"`
	Tags  map[string]string `hcl:"tags"`
}

type AwsRoute53Zone struct {
	Name            string              `cty:"name"`
	Comment         *string             `cty:"comment"`
	ForceDestroy    *bool               `cty:"force_destroy"`
	DelegationSetId *string             `cty:"delegation_set_id"`
	Tags            map[string]string   `cty:"tags"`
	Vpc             []AwsRoute53ZoneVpc `cty:"vpc"`
}

type AwsRoute53ZoneVpc struct {
	VpcId     string  `cty:"vpc_id"`
	VpcRegion *string `cty:"vpc_region"`
}

func AwsRoute53ZonesFactory(body hcl.Body) (Resource, error) {
	var out AwsRoute53Zones

	var diags hcl.Diagnostics
	attrs, attrsDiags := body.JustAttributes()
	diags = append(diags, attrsDiags...)

	for _, attr := range attrs {
		switch attr.Name {
		case "tags":
			v, vDiags := attr.Expr.Value(nil)
			if vDiags.HasErrors() {
				diags = append(diags, vDiags...)
				continue
			}

			m := v.AsValueMap()
			out.Tags = make(map[string]string)
			for k, v := range m {
				out.Tags[k] = v.AsString()
			}

		case "zones":
			v, vDiags := attr.Expr.Value(nil)
			if vDiags.HasErrors() {
				diags = append(diags, vDiags...)
				continue
			}

			l := v.AsValueSlice()
			out.Zones = make([]AwsRoute53Zone, len(l))
			for i, v := range l {
				m := v.AsValueMap()
				for k, v := range m {
					switch k {
					case "name":
						out.Zones[i].Name = v.AsString()
					case "comment":
						str := v.AsString()
						out.Zones[i].Comment = &str
					case "force_destroy":
						b := v.True()
						out.Zones[i].ForceDestroy = &b
					case "delegation_set_id":
						str := v.AsString()
						out.Zones[i].DelegationSetId = &str
					case "tags":
						m := v.AsValueMap()
						out.Zones[i].Tags = make(map[string]string)
						for k, v := range m {
							out.Zones[i].Tags[k] = v.AsString()
						}
					case "vpc":
						l := v.AsValueSlice()
						out.Zones[i].Vpc = make([]AwsRoute53ZoneVpc, len(l))
						for j, v := range l {
							m := v.AsValueMap()
							for k, v := range m {
								switch k {
								case "vpc_id":
									out.Zones[i].Vpc[j].VpcId = v.AsString()
								case "vpc_region":
									str := v.AsString()
									out.Zones[i].Vpc[j].VpcRegion = &str
								}
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
