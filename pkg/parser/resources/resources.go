package resources

type Resource interface {
}

type Factory func() Resource
type Resources map[string]Factory

func NewResources() Resources {
	return map[string]Factory{
		"aws_s3_bucket":           func() Resource { return &AwsS3Bucket{} },
		"aws_organization":        func() Resource { return &AwsOrganization{} },
		"aws_organizational_unit": func() Resource { return &AwsOrganizationalUnit{} },
	}
}
