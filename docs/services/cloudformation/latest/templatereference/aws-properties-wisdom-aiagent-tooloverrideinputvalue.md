This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Wisdom::AIAgent ToolOverrideInputValue

An input value override for tools.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "JsonPath" : String,
  "Value" : ToolOverrideInputValueConfiguration
}

```

### YAML

```yaml

  JsonPath: String
  Value:
    ToolOverrideInputValueConfiguration

```

## Properties

`JsonPath`

The JSON path for the input value override.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The override input value.

_Required_: Yes

_Type_: [ToolOverrideInputValueConfiguration](aws-properties-wisdom-aiagent-tooloverrideinputvalueconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ToolOverrideConstantInputValue

ToolOverrideInputValueConfiguration

All content copied from https://docs.aws.amazon.com/.
