package resource

import (
	ec2 "github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type VpcConstructProps struct {
	Scppe   constructs.Construct
	ProjectName    string
	VpcCidr string
}

func VpcConstruct(props VpcConstructProps) *ec2.Vpc {
	vpc := ec2.NewVpc(props.Scppe, jsii.String("VPC"), &ec2.VpcProps{
		VpcName:                      jsii.String(props.ProjectName + "-" + "vpc"),
		IpAddresses:                  ec2.IpAddresses_Cidr(&props.VpcCidr),
		EnableDnsHostnames:           jsii.Bool(true),
		EnableDnsSupport:             jsii.Bool(true),
		MaxAzs:                       jsii.Number(2),
		NatGateways:                  jsii.Number(1),
		RestrictDefaultSecurityGroup: jsii.Bool(true),
                SubnetConfiguration: &[]*ec2.SubnetConfiguration{
                        {
                                Name: jsii.String("Public"),
                                SubnetType: ec2.SubnetType_PUBLIC,
                                CidrMask: jsii.Number(24),
                        },
                        {
                                Name: jsii.String(("Private")),
                                SubnetType: ec2.SubnetType_PRIVATE_WITH_NAT,
                                CidrMask: jsii.Number(24),
                        },
                },
	})

	return &vpc
}
