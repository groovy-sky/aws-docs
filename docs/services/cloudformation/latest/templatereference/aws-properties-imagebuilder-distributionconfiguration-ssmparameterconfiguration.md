This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ImageBuilder::DistributionConfiguration SsmParameterConfiguration

Configuration for a single Parameter in the AWS Systems Manager (SSM) Parameter Store in
a given Region.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AmiAccountId" : String,
  "DataType" : String,
  "ParameterName" : String
}

```

### YAML

```yaml

  AmiAccountId: String
  DataType: String
  ParameterName: String

```

## Properties

`AmiAccountId`

Specify the account that will own the Parameter in a given Region. During distribution,
this account must be specified in distribution settings as a target account for the
Region.

_Required_: No

_Type_: String

_Pattern_: `^[0-9]{12}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataType`

The data type specifies what type of value the Parameter contains. We recommend that
you use data type `aws:ec2:image`.

_Required_: No

_Type_: String

_Allowed values_: `text | aws:ec2:image`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParameterName`

This is the name of the Parameter in the target Region or account. The image
distribution creates the Parameter if it doesn't already exist. Otherwise, it updates
the parameter.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_.\-\/]+$`

_Minimum_: `1`

_Maximum_: `1011`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LaunchTemplateConfiguration

TargetContainerRepository

All content copied from https://docs.aws.amazon.com/.
