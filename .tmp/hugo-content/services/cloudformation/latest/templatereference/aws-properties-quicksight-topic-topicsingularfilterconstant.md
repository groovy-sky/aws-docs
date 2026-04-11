This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Topic TopicSingularFilterConstant

A structure that represents a singular filter constant, used in filters to specify a single value to match against.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConstantType" : String,
  "SingularConstant" : String
}

```

### YAML

```yaml

  ConstantType: String
  SingularConstant: String

```

## Properties

`ConstantType`

The type of the singular filter constant. Valid values for this structure are `SINGULAR`.

_Required_: No

_Type_: String

_Allowed values_: `SINGULAR | RANGE | COLLECTIVE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SingularConstant`

The value of the singular filter constant.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TopicRelativeDateFilter

AWS::QuickSight::VPCConnection

All content copied from https://docs.aws.amazon.com/.
