package resources

type AwsOrganizationalUnit struct {
	OrganizationName string                            `hcl:"organization_name,attr"`
	Accounts         map[string]AwsOrganizationAccount `hcl:"accounts,optional"`
	DependsOn        []string                          `hcl:"depends_on,optional"`
}

type AwsOrganizationAccount struct {
	Email string `cty:"email"`
}
