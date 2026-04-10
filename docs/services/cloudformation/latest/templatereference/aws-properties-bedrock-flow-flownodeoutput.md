This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Flow FlowNodeOutput

Contains configurations for an output from a node.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String,
  "Type" : String
}

```

### YAML

```yaml

  Name: String
  Type: String

```

## Properties

`Name`

A name for the output that you can reference.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z]([_]?[0-9a-zA-Z]){1,50}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The data type of the output. If the output doesn't match this type at runtime, a validation error will be thrown.

_Required_: Yes

_Type_: String

_Allowed values_: `String | Number | Boolean | Object | Array`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FlowNodeInput

FlowValidation

All content copied from https://docs.aws.amazon.com/.
