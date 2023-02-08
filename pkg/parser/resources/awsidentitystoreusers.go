package resources

type AwsIdentitystoreUsers struct {
	Users []AwsIdentitystoreUser `hcl:"users"`
}

type AwsIdentitystoreUser struct {
	UserName    string `cty:"user_name"`
	DisplayName string `cty:"display_name"`
	GivenName   string `cty:"given_name"`
	FamilyName  string `cty:"family_name"`
	Email       string `cty:"email"`
}
