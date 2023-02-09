package resources

type AwsOrganizationalUnit struct {
	OrganizationName string                   `hcl:"organization_name,attr"`
	Accounts         []AwsOrganizationAccount `hcl:"accounts,optional"`
	DependsOn        []string                 `hcl:"depends_on,optional"`
}

type AwsOrganizationAccount struct {
	Name  string `cty:"name"`
	Email string `cty:"email"`
}
