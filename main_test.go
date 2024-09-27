package main

import (
	"testing"

	"github.com/aws/jsii-runtime-go"
	"github.com/cdktf/cdktf-provider-local-go/local/v10/file"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/stretchr/testify/assert"
)

// The tests below are example tests, you can find more information at
// https://cdk.tf/testing

var (
	stack = NewNoopStack(cdktf.Testing_App(nil), "stack", &config{})
	synth = cdktf.Testing_Synth(stack, jsii.Bool(true))
)

func TestShouldContainLocalFile(t *testing.T) {
	assertion := cdktf.Testing_ToHaveResource(synth, file.File_TfResourceType())
	assert.True(t, *assertion)
}
