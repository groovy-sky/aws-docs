This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Flow LoopControllerFlowNodeConfiguration

Contains configurations for the controller node of a DoWhile loop in the flow.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ContinueCondition" : FlowCondition,
  "MaxIterations" : Number
}

```

### YAML

```yaml

  ContinueCondition:
    FlowCondition
  MaxIterations: Number

```

## Properties

`ContinueCondition`

Specifies the condition that determines when the flow exits the DoWhile loop. The loop
executes until this condition evaluates to true.

_Required_: Yes

_Type_: [FlowCondition](aws-properties-bedrock-flow-flowcondition.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxIterations`

Specifies the maximum number of times the DoWhile loop can iterate before the flow
exits the loop.

_Required_: No

_Type_: Number

_Minimum_: `1`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LexFlowNodeConfiguration

LoopFlowNodeConfiguration

All content copied from https://docs.aws.amazon.com/.
