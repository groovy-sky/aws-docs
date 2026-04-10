This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Flow FlowConnectionConfiguration

The configuration of the connection.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Conditional" : FlowConditionalConnectionConfiguration,
  "Data" : FlowDataConnectionConfiguration
}

```

### YAML

```yaml

  Conditional:
    FlowConditionalConnectionConfiguration
  Data:
    FlowDataConnectionConfiguration

```

## Properties

`Conditional`

The configuration of a connection originating from a Condition node.

_Required_: No

_Type_: [FlowConditionalConnectionConfiguration](aws-properties-bedrock-flow-flowconditionalconnectionconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Data`

The configuration of a connection originating from a node that isn't a Condition node.

_Required_: No

_Type_: [FlowDataConnectionConfiguration](aws-properties-bedrock-flow-flowdataconnectionconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FlowConnection

FlowDataConnectionConfiguration

All content copied from https://docs.aws.amazon.com/.
