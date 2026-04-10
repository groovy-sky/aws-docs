This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Flow FlowDataConnectionConfiguration

The configuration of a connection originating from a node that isn't a Condition node.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SourceOutput" : String,
  "TargetInput" : String
}

```

### YAML

```yaml

  SourceOutput: String
  TargetInput: String

```

## Properties

`SourceOutput`

The name of the output in the source node that the connection begins from.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z]([_]?[0-9a-zA-Z]){1,50}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetInput`

The name of the input in the target node that the connection ends at.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z]([_]?[0-9a-zA-Z]){1,50}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FlowConnectionConfiguration

FlowDefinition

All content copied from https://docs.aws.amazon.com/.
