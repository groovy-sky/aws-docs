This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Flow FlowNode

Contains configurations about a node in the flow.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Configuration" : FlowNodeConfiguration,
  "Inputs" : [ FlowNodeInput, ... ],
  "Name" : String,
  "Outputs" : [ FlowNodeOutput, ... ],
  "Type" : String
}

```

### YAML

```yaml

  Configuration:
    FlowNodeConfiguration
  Inputs:
    - FlowNodeInput
  Name: String
  Outputs:
    - FlowNodeOutput
  Type: String

```

## Properties

`Configuration`

Contains configurations for the node.

_Required_: No

_Type_: [FlowNodeConfiguration](aws-properties-bedrock-flow-flownodeconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Inputs`

An array of objects, each of which contains information about an input into the node.

_Required_: No

_Type_: Array of [FlowNodeInput](aws-properties-bedrock-flow-flownodeinput.md)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

A name for the node.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z]([_]?[0-9a-zA-Z]){1,50}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Outputs`

A list of objects, each of which contains information about an output from the node.

_Required_: No

_Type_: Array of [FlowNodeOutput](aws-properties-bedrock-flow-flownodeoutput.md)

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of node. This value must match the name of the key that you provide in the configuration you provide in the `FlowNodeConfiguration` field.

_Required_: Yes

_Type_: String

_Allowed values_: `Input | Output | KnowledgeBase | Condition | Lex | Prompt | LambdaFunction | Agent | Storage | Retrieval | Iterator | Collector | InlineCode | Loop | LoopInput | LoopController`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FlowDefinition

FlowNodeConfiguration

All content copied from https://docs.aws.amazon.com/.
