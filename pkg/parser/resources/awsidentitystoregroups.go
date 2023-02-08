package resources

type AwsIdentitystoreGroups struct {
	Groups []AwsIdentitystoreUser `hcl:"groups"`
}

type AwsIdentitystoreGroup struct {
	DisplayName string `cty:"display_name"`
	Description string `cty:"description"`
}
