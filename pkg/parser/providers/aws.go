package providers

type AWSProvider struct {
	Region     string         `hcl:"region"`
	AssumeRole *AWSAssumeRole `hcl:"assume_role,block"`
}

type AWSAssumeRole struct {
	RoleArn         *string `hcl:"role_arn,optional"`
	RoleSessionName *string `hcl:"role_session_name,optional"`
	ExternalID      *string `hcl:"external_id,optional"`
}
