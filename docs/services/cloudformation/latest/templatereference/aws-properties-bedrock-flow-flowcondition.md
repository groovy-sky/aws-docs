This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Flow FlowCondition

Defines a condition in the condition node.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Expression" : String,
  "Name" : String
}

```

### YAML

```yaml

  Expression: String
  Name: String

```

## Properties

`Expression`

Defines the condition. You must refer to at least one of the inputs in the condition. For more information, expand the Condition node section in [Node types in prompt flows](../../../bedrock/latest/userguide/flows-how-it-works.md#flows-nodes).

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

A name for the condition that you can reference.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z]([_]?[0-9a-zA-Z]){1,50}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FieldForReranking

FlowConditionalConnectionConfiguration

All content copied from https://docs.aws.amazon.com/.
