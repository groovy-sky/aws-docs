# VpcLatticeConfiguration

The VPC Lattice configuration for your service that holds the information for the
target group(s) Amazon ECS tasks will be registered to.

## Contents

**portName**

The name of the port mapping to register in the VPC Lattice target group. This is the
name of the `portMapping` you defined in your task definition.

Type: String

Required: Yes

**roleArn**

The ARN of the IAM role to associate with this VPC Lattice configuration. This is the
Amazon ECS infrastructure IAM role that is used to manage your VPC Lattice
infrastructure.

Type: String

Required: Yes

**targetGroupArn**

The full Amazon Resource Name (ARN) of the target group or groups associated with the
VPC Lattice configuration that the Amazon ECS tasks will be registered to.

Type: String

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ecs-2014-11-13/VpcLatticeConfiguration)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ecs-2014-11-13/VpcLatticeConfiguration)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ecs-2014-11-13/VpcLatticeConfiguration)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

VolumeFrom

API request throttling
