This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Flow FlowConditionalConnectionConfiguration

The configuration of a connection between a condition node and another node.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Condition" : String
}

```

### YAML

```yaml

  Condition: String

```

## Properties

`Condition`

The condition that triggers this connection. For more information about how to write conditions, see the **Condition** node type in the [Node types](../../../bedrock/latest/userguide/node-types.md) topic in the Amazon Bedrock User Guide.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z]([_]?[0-9a-zA-Z]){1,50}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FlowCondition

FlowConnection

All content copied from https://docs.aws.amazon.com/.
