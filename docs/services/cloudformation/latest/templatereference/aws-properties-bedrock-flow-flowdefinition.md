This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Flow FlowDefinition

The definition of the nodes and connections between nodes in the flow.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Connections" : [ FlowConnection, ... ],
  "Nodes" : [ FlowNode, ... ]
}

```

### YAML

```yaml

  Connections:
    - FlowConnection
  Nodes:
    - FlowNode

```

## Properties

`Connections`

An array of connection definitions in the flow.

_Required_: No

_Type_: Array of [FlowConnection](aws-properties-bedrock-flow-flowconnection.md)

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Nodes`

An array of node definitions in the flow.

_Required_: No

_Type_: Array of [FlowNode](aws-properties-bedrock-flow-flownode.md)

_Maximum_: `40`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FlowDataConnectionConfiguration

FlowNode

All content copied from https://docs.aws.amazon.com/.
