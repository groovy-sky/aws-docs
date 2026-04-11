This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Flow FlowNodeInput

Contains configurations for an input in an Amazon Bedrock Flows node.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Category" : String,
  "Expression" : String,
  "Name" : String,
  "Type" : String
}

```

### YAML

```yaml

  Category: String
  Expression: String
  Name: String
  Type: String

```

## Properties

`Category`

Specifies how input data flows between iterations in a DoWhile
loop.

- `LoopCondition` \- Controls whether the loop continues by evaluating condition expressions against the input data. Use this category to define
the condition that determines if the loop should continue.

- `ReturnValueToLoopStart` \- Defines data to pass back to the start
of the loop's next iteration. Use this category for variables that you
want to update for each loop iteration.

- `ExitLoop` \- Defines the value that's available once the loop ends.
Use this category to expose loop results to nodes outside the loop.

_Required_: No

_Type_: String

_Allowed values_: `LoopCondition | ReturnValueToLoopStart | ExitLoop`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Expression`

An expression that formats the input for the node. For an explanation of how to create expressions, see [Expressions in Prompt flows in Amazon Bedrock](../../../bedrock/latest/userguide/flows-expressions.md).

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

Specifies a name for the input that you can reference.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z]([_]?[0-9a-zA-Z]){1,50}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

Specifies the data type of the input. If the input doesn't match this type at runtime,
a validation error will be thrown.

_Required_: Yes

_Type_: String

_Allowed values_: `String | Number | Boolean | Object | Array`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FlowNodeConfiguration

FlowNodeOutput

All content copied from https://docs.aws.amazon.com/.
