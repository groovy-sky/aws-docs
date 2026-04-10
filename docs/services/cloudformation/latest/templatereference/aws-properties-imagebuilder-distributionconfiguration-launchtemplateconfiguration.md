This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ImageBuilder::DistributionConfiguration LaunchTemplateConfiguration

Identifies an Amazon EC2 launch template to use for a specific account.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AccountId" : String,
  "LaunchTemplateId" : String,
  "SetDefaultVersion" : Boolean
}

```

### YAML

```yaml

  AccountId: String
  LaunchTemplateId: String
  SetDefaultVersion: Boolean

```

## Properties

`AccountId`

The account ID that this configuration applies to.

_Required_: No

_Type_: String

_Pattern_: `^[0-9]{12}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LaunchTemplateId`

Identifies the Amazon EC2 launch template to use.

_Required_: No

_Type_: String

_Pattern_: `^lt-[a-z0-9-_]{17}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SetDefaultVersion`

Set the specified Amazon EC2 launch template as the default launch template for the
specified account.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LaunchPermissionConfiguration

SsmParameterConfiguration

All content copied from https://docs.aws.amazon.com/.
