This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeBuild::Project VpcConfig

`VpcConfig` is a property of the [AWS::CodeBuild::Project](../userguide/aws-resource-codebuild-project.md) resource
that enable AWS CodeBuild to access resources in an Amazon VPC. For more information, see
[Use AWS CodeBuild with Amazon Virtual Private Cloud](../../../codebuild/latest/userguide/vpc-support.md) in
the _AWS CodeBuild User Guide_.

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

A list of one or more security groups IDs in your Amazon VPC. The maximum count is 5.

_Required_: No

_Type_: Array of String

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Subnets`

A list of one or more subnet IDs in your Amazon VPC. The maximum count is 16.

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

## See also

- [VpcConfig](../../../../reference/codebuild/latest/apireference/api-vpcconfig.md) in the _AWS CodeBuild API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

WebhookFilter

All content copied from https://docs.aws.amazon.com/.
