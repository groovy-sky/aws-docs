This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cloud9::EnvironmentEC2

The `AWS::Cloud9::EnvironmentEC2` resource creates an Amazon EC2 development environment in AWS Cloud9. For more information, see [Creating an Environment](https://docs.aws.amazon.com/cloud9/latest/user-guide/create-environment.html) in the _AWS Cloud9 User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Cloud9::EnvironmentEC2",
  "Properties" : {
      "AutomaticStopTimeMinutes" : Integer,
      "ConnectionType" : String,
      "Description" : String,
      "ImageId" : String,
      "InstanceType" : String,
      "Name" : String,
      "OwnerArn" : String,
      "Repositories" : [ Repository, ... ],
      "SubnetId" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Cloud9::EnvironmentEC2
Properties:
  AutomaticStopTimeMinutes: Integer
  ConnectionType: String
  Description: String
  ImageId: String
  InstanceType: String
  Name: String
  OwnerArn: String
  Repositories:
    - Repository
  SubnetId: String
  Tags:
    - Tag

```

## Properties

`AutomaticStopTimeMinutes`

The number of minutes until the running instance is shut down after the environment was last used.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `20160`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ConnectionType`

The connection type used for connecting to an Amazon EC2 environment. Valid values are `CONNECT_SSH` (default) and `CONNECT_SSM` (connected through AWS Systems Manager).

_Required_: No

_Type_: String

_Allowed values_: `CONNECT_SSH | CONNECT_SSM`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

The description of the environment to create.

_Required_: No

_Type_: String

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ImageId`

The identifier for the Amazon Machine Image (AMI) that's used to create the EC2 instance.
To choose an AMI for the instance, you must specify a valid AMI alias or a valid AWS Systems Manager path.

From December 04, 2023, you will be required to include the `ImageId` parameter
for the `CreateEnvironmentEC2` action. This change will be reflected across all
direct methods of communicating with the API, such as AWS SDK, AWS CLI and AWS CloudFormation. This change will only affect
direct API consumers, and not AWS Cloud9 console users.

Since Ubuntu 18.04 has ended standard support as of May 31, 2023, we recommend you choose Ubuntu 22.04.

**AMI aliases**

- Amazon Linux 2: `amazonlinux-2-x86_64`

- Amazon Linux 2023 (recommended): `amazonlinux-2023-x86_64`

- Ubuntu 18.04: `ubuntu-18.04-x86_64`

- Ubuntu 22.04: `ubuntu-22.04-x86_64`

**SSM paths**

- Amazon Linux 2: `resolve:ssm:/aws/service/cloud9/amis/amazonlinux-2-x86_64`

- Amazon Linux 2023 (recommended): `resolve:ssm:/aws/service/cloud9/amis/amazonlinux-2023-x86_64`

- Ubuntu 18.04: `resolve:ssm:/aws/service/cloud9/amis/ubuntu-18.04-x86_64`

- Ubuntu 22.04: `resolve:ssm:/aws/service/cloud9/amis/ubuntu-22.04-x86_64`

_Required_: Yes

_Type_: String

_Maximum_: `512`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`InstanceType`

The type of instance to connect to the environment (for example, `t2.micro`).

_Required_: Yes

_Type_: String

_Pattern_: `^[a-z]+[1-9][.][a-z0-9]+$`

_Minimum_: `5`

_Maximum_: `20`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the environment.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OwnerArn`

The Amazon Resource Name (ARN) of the environment owner. This ARN can be the ARN of any AWS Identity and Access Management principal. If this value is not specified, the ARN defaults to this environment's creator.

_Required_: No

_Type_: String

_Pattern_: `^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):(iam|sts)::\d+:(root|(user\/[\w+=/:,.@-]{1,64}|federated-user\/[\w+=/:,.@-]{2,32}|assumed-role\/[\w+=:,.@-]{1,64}\/[\w+=,.@-]{1,64}))$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Repositories`

Any AWS CodeCommit source code repositories to be cloned into the development environment.

_Required_: No

_Type_: Array of [Repository](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloud9-environmentec2-repository.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SubnetId`

The ID of the subnet in Amazon Virtual Private Cloud (Amazon VPC) that AWS Cloud9 will use to communicate with the Amazon Elastic Compute Cloud (Amazon EC2) instance.

_Required_: No

_Type_: String

_Pattern_: `^(subnet-[0-9a-f]{8}|subnet-[0-9a-f]{17})$`

_Minimum_: `15`

_Maximum_: `24`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

An array of key-value pairs that will be associated with the new AWS Cloud9 development
environment.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloud9-environmentec2-tag.html)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the development environment, such as `2bc3642873c342e485f7e0c561234567`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the development environment, such as `arn:aws:cloud9:us-east-2:123456789012:environment:2bc3642873c342e485f7e0c561234567`.

`Name`

The name of the environment.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS Cloud9

Repository
