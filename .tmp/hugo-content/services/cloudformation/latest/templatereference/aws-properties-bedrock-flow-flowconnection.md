This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Flow FlowConnection

Contains information about a connection between two nodes in the flow.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Configuration" : FlowConnectionConfiguration,
  "Name" : String,
  "Source" : String,
  "Target" : String,
  "Type" : String
}

```

### YAML

```yaml

  Configuration:
    FlowConnectionConfiguration
  Name: String
  Source: String
  Target: String
  Type: String

```

## Properties

`Configuration`

The configuration of the connection.

_Required_: No

_Type_: [FlowConnectionConfiguration](aws-properties-bedrock-flow-flowconnectionconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

A name for the connection that you can reference.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z]([_]?[0-9a-zA-Z]){1,100}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Source`

The node that the connection starts at.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z]([_]?[0-9a-zA-Z]){1,50}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Target`

The node that the connection ends at.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z]([_]?[0-9a-zA-Z]){1,50}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

Whether the source node that the connection begins from is a condition node ( `Conditional`) or not ( `Data`).

_Required_: Yes

_Type_: String

_Allowed values_: `Data | Conditional`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FlowConditionalConnectionConfiguration

FlowConnectionConfiguration

All content copied from https://docs.aws.amazon.com/.
