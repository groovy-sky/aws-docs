This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Flow LoopFlowNodeConfiguration

Contains configurations for the nodes of a DoWhile loop in your flow.

A DoWhile loop is made up of the following nodes:

- `Loop` \- The container node that holds the loop's flow definition. This node encompasses the entire loop structure.

- `LoopInput` \- The entry point node for the loop. This node receives inputs from nodes outside the loop and from previous loop iterations.

- Body nodes - The processing nodes that execute within each loop iteration.
These can be nodes for handling data in your flow, such as a prompt or Lambda
function nodes. Some node types aren't supported inside a DoWhile loop body. For
more information, see [LoopIncompatibleNodeTypeFlowValidationDetails](../../../../reference/bedrock/latest/apireference/api-agent-loopincompatiblenodetypeflowvalidationdetails.md).

- `LoopController` \- The node that evaluates whether the loop should continue or exit based on a condition.

These nodes work together to create a loop that runs at least once and continues until a specified condition is met or a maximum number of iterations is reached.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Definition" : FlowDefinition
}

```

### YAML

```yaml

  Definition:
    FlowDefinition

```

## Properties

`Definition`

The definition of the DoWhile loop nodes and connections between nodes in the flow.

_Required_: Yes

_Type_: [FlowDefinition](aws-properties-bedrock-flow-flowdefinition.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LoopControllerFlowNodeConfiguration

MetadataConfigurationForReranking

All content copied from https://docs.aws.amazon.com/.
