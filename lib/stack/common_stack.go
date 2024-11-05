package stacks

import (
	"golang_cdk_resource/lib/resource"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

type CommonStackProps struct {
	awscdk.StackProps
  ProjectName string
  Env string
}

func CommonStackConstruct(scope constructs.Construct, id string, props *CommonStackProps ) {
  var sprops awscdk.StackProps
  if props != nil {
    sprops = props.StackProps
  }
  stack := awscdk.NewStack(scope, &id, &sprops)

  resource.VpcConstruct(resource.VpcConstructProps{
    Scppe: stack,
    ProjectName: props.ProjectName + props.Env,
    VpcCidr: "192.168.0.0/24",
  })
}