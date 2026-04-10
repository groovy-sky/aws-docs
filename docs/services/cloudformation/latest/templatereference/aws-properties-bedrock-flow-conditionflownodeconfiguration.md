This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Flow ConditionFlowNodeConfiguration

Defines a condition node in your flow. You can specify conditions that determine which node comes next in the flow. For more information, see [Node types in a flow](../../../bedrock/latest/userguide/flows-nodes.md) in the Amazon Bedrock User Guide.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Conditions" : [ FlowCondition, ... ]
}

```

### YAML

```yaml

  Conditions:
    - FlowCondition

```

## Properties

`Conditions`

An array of conditions. Each member contains the name of a condition and an expression that defines the condition.

_Required_: Yes

_Type_: Array of [FlowCondition](aws-properties-bedrock-flow-flowcondition.md)

_Minimum_: `1`

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AgentFlowNodeConfiguration

FieldForReranking

All content copied from https://docs.aws.amazon.com/.
