package resources

type AwsIdentitystoreGroupMemberships struct {
	Members []AwsIdentitystoreGroupMembership `hcl:"members"`
}

type AwsIdentitystoreGroupMembership struct {
	UserName  string `cty:"user_name"`
	GroupName string `cty:"group_name"`
}
