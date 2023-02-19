package resources

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
)

type Resource interface {
}

type Factory func(hcl.Body) (Resource, error)
type Resources map[string]Factory

func Basic(body hcl.Body, opts Resource) (Resource, error) {
	if err := gohcl.DecodeBody(body, nil, opts); err != nil {
		return nil, err
	}

	return opts, nil
}

func NewResources() Resources {
	return map[string]Factory{
		"aws_acm":                             func(b hcl.Body) (Resource, error) { return Basic(b, &AwsAcm{}) },
		"aws_cloudfront_distribution":         func(b hcl.Body) (Resource, error) { return Basic(b, &AwsCloudfrontDistribution{}) },
		"aws_identitystore_assignments":       func(b hcl.Body) (Resource, error) { return Basic(b, &AwsIdentitystoreAssignments{}) },
		"aws_identitystore_group_memberships": func(b hcl.Body) (Resource, error) { return Basic(b, &AwsIdentitystoreGroupMemberships{}) },
		"aws_identitystore_groups":            func(b hcl.Body) (Resource, error) { return Basic(b, &AwsIdentitystoreGroups{}) },
		"aws_identitystore_users":             func(b hcl.Body) (Resource, error) { return Basic(b, &AwsIdentitystoreUsers{}) },
		"aws_identitystore_permission_sets":   func(b hcl.Body) (Resource, error) { return Basic(b, &AwsIdentitystorePermissionSets{}) },
		"aws_route53_records":                 AwsRoute53RecordsFactory,
		"aws_route53_zones":                   AwsRoute53ZonesFactory,
		"aws_s3_bucket":                       func(b hcl.Body) (Resource, error) { return Basic(b, &AwsS3Bucket{}) },
		"aws_organization":                    func(b hcl.Body) (Resource, error) { return Basic(b, &AwsOrganization{}) },
		"aws_organizational_unit":             func(b hcl.Body) (Resource, error) { return Basic(b, &AwsOrganizationalUnit{}) },
	}
}
