This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Wisdom::AIAgent ToolOutputConfiguration

Configuration for tool output handling.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "OutputVariableNameOverride" : String,
  "SessionDataNamespace" : String
}

```

### YAML

```yaml

  OutputVariableNameOverride: String
  SessionDataNamespace: String

```

## Properties

`OutputVariableNameOverride`

Override the tool output results to different variable name.

_Required_: No

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SessionDataNamespace`

The session data namespace for tool output.

_Required_: No

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ToolInstruction

ToolOutputFilter

All content copied from https://docs.aws.amazon.com/.
