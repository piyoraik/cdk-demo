package main

import (
	stacks "golang_cdk_resource/lib/stack"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"
)

func main() {
	defer jsii.Close()
	app := awscdk.NewApp(nil)

	projectName := "s-tanaka"
	env := "dev"

	stacks.CommonStackConstruct(app, "CommonStack", &stacks.CommonStackProps{
		ProjectName: projectName,
		Env:         env,
	})

	app.Synth(nil)
	awscdk.Tags_Of(app).Add(jsii.String("Owner"), jsii.String("tanaka-shun"), &awscdk.TagProps{})
}
