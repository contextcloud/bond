package resources

type Resource interface {
}

type Factory func() Resource
type Resources map[string]Factory

func NewResources() Resources {
	return map[string]Factory{
		"aws_cloudfront_distribution":         func() Resource { return &AwsCloudfrontDistribution{} },
		"aws_identitystore_assignments":       func() Resource { return &AwsIdentitystoreAssignments{} },
		"aws_identitystore_group_memberships": func() Resource { return &AwsIdentitystoreGroupMemberships{} },
		"aws_identitystore_groups":            func() Resource { return &AwsIdentitystoreGroups{} },
		"aws_identitystore_users":             func() Resource { return &AwsIdentitystoreUsers{} },
		"aws_identitystore_permission_sets":   func() Resource { return &AwsIdentitystorePermissionSets{} },
		"aws_s3_bucket":                       func() Resource { return &AwsS3Bucket{} },
		"aws_organization":                    func() Resource { return &AwsOrganization{} },
		"aws_organizational_unit":             func() Resource { return &AwsOrganizationalUnit{} },
	}
}
