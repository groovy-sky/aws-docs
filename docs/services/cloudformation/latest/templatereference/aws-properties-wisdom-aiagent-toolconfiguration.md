This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Wisdom::AIAgent ToolConfiguration

Configuration settings for a tool used by AI Agents.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Annotations" : Json,
  "Description" : String,
  "InputSchema" : Json,
  "Instruction" : ToolInstruction,
  "OutputFilters" : [ ToolOutputFilter, ... ],
  "OutputSchema" : Json,
  "OverrideInputValues" : [ ToolOverrideInputValue, ... ],
  "Title" : String,
  "ToolId" : String,
  "ToolName" : String,
  "ToolType" : String,
  "UserInteractionConfiguration" : UserInteractionConfiguration
}

```

### YAML

```yaml

  Annotations: Json
  Description: String
  InputSchema: Json
  Instruction:
    ToolInstruction
  OutputFilters:
    - ToolOutputFilter
  OutputSchema: Json
  OverrideInputValues:
    - ToolOverrideInputValue
  Title: String
  ToolId: String
  ToolName: String
  ToolType: String
  UserInteractionConfiguration:
    UserInteractionConfiguration

```

## Properties

`Annotations`

Annotations for the tool configuration.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the tool configuration.

_Required_: No

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InputSchema`

The input schema for the tool configuration.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Instruction`

Instructions for using the tool.

_Required_: No

_Type_: [ToolInstruction](aws-properties-wisdom-aiagent-toolinstruction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OutputFilters`

Output filters applies to the tool result.

_Required_: No

_Type_: Array of [ToolOutputFilter](aws-properties-wisdom-aiagent-tooloutputfilter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OutputSchema`

The output schema for the tool configuration.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OverrideInputValues`

Override input values for the tool configuration.

_Required_: No

_Type_: Array of [ToolOverrideInputValue](aws-properties-wisdom-aiagent-tooloverrideinputvalue.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Title`

The title of the tool configuration.

_Required_: No

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ToolId`

The identifier of the tool, for example toolName from Model Context Provider server.

_Required_: No

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ToolName`

The name of the tool.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ToolType`

The type of the tool.

_Required_: Yes

_Type_: String

_Allowed values_: `MODEL_CONTEXT_PROTOCOL | RETURN_TO_CONTROL | CONSTANT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserInteractionConfiguration`

Configuration for user interaction with the tool.

_Required_: No

_Type_: [UserInteractionConfiguration](aws-properties-wisdom-aiagent-userinteractionconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TagFilter

ToolInstruction

All content copied from https://docs.aws.amazon.com/.
