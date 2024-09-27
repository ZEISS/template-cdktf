package main

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdktf/cdktf-provider-local-go/local/v10/file"
	"github.com/cdktf/cdktf-provider-local-go/local/v10/provider"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type config struct {
	filename string
}

func NewNoopStack(scope constructs.Construct, id string, cfg *config) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, jsii.String(id))

	provider.NewLocalProvider(stack, jsii.String("local"), &provider.LocalProviderConfig{})

	file := file.NewFile(stack, jsii.String("file"), &file.FileConfig{
		Content:  jsii.String("Hello, World!"),
		Filename: jsii.String(cfg.filename),
	})

	cdktf.NewTerraformOutput(stack, jsii.String("filename"), &cdktf.TerraformOutputConfig{
		Value: file.Filename(),
	})

	return stack
}

func main() {
	app := cdktf.NewApp(nil)

	NewNoopStack(app, "stack", &config{
		filename: "hello.txt",
	})

	app.Synth()
}
