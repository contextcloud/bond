package terra

type TerraformBlock struct {
	RequiredVersion string `hcl:"required_version"`
}

func NewTerraformBlock() *TerraformBlock {
	return &TerraformBlock{
		RequiredVersion: ">= 0.13",
	}
}
