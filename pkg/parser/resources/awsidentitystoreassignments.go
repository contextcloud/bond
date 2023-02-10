package resources

type AwsIdentitystoreAssignments struct {
	Assignments []AwsIdentitystoreAssignment `hcl:"assignments"`
}

type AwsIdentitystoreAssignment struct {
	AccountName       string `cty:"account_name"`
	GroupName         string `cty:"group_name"`
	PermissionSetName string `cty:"permission_set_name"`
}
