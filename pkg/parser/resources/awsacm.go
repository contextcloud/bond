package resources

type AwsAcm struct {
	DomainName              string            `hcl:"domain_name"`
	SubjectAlternativeNames []string          `hcl:"subject_alternative_names,optional"`
	ZoneId                  string            `hcl:"zone_id"`
	Tags                    map[string]string `hcl:"tags"`
}
