This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Wisdom::AIAgent ToolOverrideConstantInputValue

A constant input value for tool override.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : String,
  "Value" : String
}

```

### YAML

```yaml

  Type: String
  Value: String

```

## Properties

`Type`

Override tool input value with constant values

_Required_: Yes

_Type_: String

_Allowed values_: `STRING | NUMBER | JSON_STRING`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The constant input override value.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ToolOutputFilter

ToolOverrideInputValue

All content copied from https://docs.aws.amazon.com/.
