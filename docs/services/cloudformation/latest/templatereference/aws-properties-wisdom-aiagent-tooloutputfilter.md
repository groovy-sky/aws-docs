This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Wisdom::AIAgent ToolOutputFilter

Filter configuration for tool output.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "JsonPath" : String,
  "OutputConfiguration" : ToolOutputConfiguration
}

```

### YAML

```yaml

  JsonPath: String
  OutputConfiguration:
    ToolOutputConfiguration

```

## Properties

`JsonPath`

The JSON path for filtering tool output.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OutputConfiguration`

The output configuration for the filter.

_Required_: No

_Type_: [ToolOutputConfiguration](aws-properties-wisdom-aiagent-tooloutputconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ToolOutputConfiguration

ToolOverrideConstantInputValue

All content copied from https://docs.aws.amazon.com/.
