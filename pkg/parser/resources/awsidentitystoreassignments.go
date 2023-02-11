package resources

type AwsIdentitystoreAssignments struct {
	Assignments []AwsIdentitystoreAssignment `hcl:"assignments"`
	DependsOn   []string                     `hcl:"depends_on,optional"`
}

type AwsIdentitystoreAssignment struct {
	AccountName       string `cty:"account_name"`
	GroupName         string `cty:"group_name"`
	PermissionSetName string `cty:"permission_set_name"`
}
