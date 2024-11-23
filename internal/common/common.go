package common

import "github.com/hashicorp/hcl/v2/hclwrite"

type IACType string

const (
	Terraform  IACType = "terraform"
	Terragrunt IACType = "terragrunt"
)

type Version struct {
	Major int
	Minor int
}

type TaggingArgs struct {
	Filter              string
	Skip                string
	Dir                 string
	File                string
	Tags                string
	Matches             []string
	IsSkipTerratagFiles bool
	Rename              bool
	DefaultToTerraform  bool
	IACType             IACType
}

type TerratagLocal struct {
	Found map[string]hclwrite.Tokens
	Added string
}
