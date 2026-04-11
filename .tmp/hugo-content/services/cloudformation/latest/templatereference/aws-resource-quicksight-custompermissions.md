This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::CustomPermissions

Creates a custom permissions profile.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::QuickSight::CustomPermissions",
  "Properties" : {
      "AwsAccountId" : String,
      "Capabilities" : Capabilities,
      "CustomPermissionsName" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::QuickSight::CustomPermissions
Properties:
  AwsAccountId: String
  Capabilities:
    Capabilities
  CustomPermissionsName: String
  Tags:
    - Tag

```

## Properties

`AwsAccountId`

The ID of the AWS account that contains the custom permission configuration that you want to update.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9]{12}$`

_Minimum_: `12`

_Maximum_: `12`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Capabilities`

A set of actions in the custom permissions profile.

_Required_: No

_Type_: [Capabilities](aws-properties-quicksight-custompermissions-capabilities.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomPermissionsName`

The name of the custom permissions profile.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9+=,.@_-]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags to associate with the custom permissions profile.

_Required_: No

_Type_: Array of [Tag](aws-properties-quicksight-custompermissions-tag.md)

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`Arn`

The Amazon Resource Name (ARN) of the custom permissions profile.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

YAxisOptions

Capabilities

All content copied from https://docs.aws.amazon.com/.
