package resources

type Factory func() interface{}
type Resources map[string]Factory

func NewResources() Resources {
	return map[string]Factory{
		"aws_s3_bucket":            func() interface{} { return &AwsS3Bucket{} },
		"aws_organization":         func() interface{} { return &AwsOrganization{} },
		"aws_organization_account": func() interface{} { return &AwsOrganizationAccount{} },
		"aws_organizational_unit":  func() interface{} { return &AwsOrganizationalUnit{} },
	}
}
