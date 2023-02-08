package resources

type AwsIdentitystorePermissionSets struct {
	PermissionSets []AwsIdentitystorePermissionSet `hcl:"permission_sets"`
}

type AwsIdentitystorePermissionSet struct {
	Name                             string                            `cty:"name"`
	Description                      string                            `cty:"description"`
	RelayState                       string                            `cty:"relay_state"`
	SessionDuration                  string                            `cty:"session_duration"`
	Tags                             map[string]string                 `cty:"tags"`
	InlinePolicy                     string                            `cty:"inline_policy"`
	PolicyAttachments                []string                          `cty:"policy_attachments"`
	CustomerManagedPolicyAttachments []CustomerManagedPolicyAttachment `cty:"customer_managed_policy_attachments"`
}

type CustomerManagedPolicyAttachment struct {
	Name string `cty:"name"`
	Path string `cty:"path"`
}
