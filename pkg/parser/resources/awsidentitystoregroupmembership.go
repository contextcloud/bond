package resources

type AwsIdentitystoreGroupMemberships struct {
	Members   []AwsIdentitystoreGroupMembership `hcl:"members"`
	DependsOn []string                          `hcl:"depends_on,optional"`
}

type AwsIdentitystoreGroupMembership struct {
	UserName  string `cty:"user_name"`
	GroupName string `cty:"group_name"`
}
