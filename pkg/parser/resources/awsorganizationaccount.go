package resources

type AwsOrganizationAccount struct {
	Name               string `hcl:"name"`
	Email              string `hcl:"email"`
	OrganizationalUnit string `hcl:"organizational_unit"`
}
