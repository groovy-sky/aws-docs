This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeBuild::Fleet VpcConfig

Information about the VPC configuration that AWS CodeBuild accesses.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SecurityGroupIds" : [ String, ... ],
  "Subnets" : [ String, ... ],
  "VpcId" : String
}

```

### YAML

```yaml

  SecurityGroupIds:
    - String
  Subnets:
    - String
  VpcId: String

```

## Properties

`SecurityGroupIds`

A list of one or more security groups IDs in your Amazon VPC.

_Required_: No

_Type_: Array of String

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Subnets`

A list of one or more subnet IDs in your Amazon VPC.

_Required_: No

_Type_: Array of String

_Maximum_: `16`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcId`

The ID of the Amazon VPC.

_Required_: No

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TargetTrackingScalingConfiguration

AWS::CodeBuild::Project

All content copied from https://docs.aws.amazon.com/.
