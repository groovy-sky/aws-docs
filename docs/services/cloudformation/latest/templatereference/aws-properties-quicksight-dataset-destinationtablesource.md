This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSet DestinationTableSource

Specifies the source of data for a destination table, including the transform operation and column mappings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "TransformOperationId" : String
}

```

### YAML

```yaml

  TransformOperationId: String

```

## Properties

`TransformOperationId`

The identifier of the transform operation that provides data to the destination table.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9a-zA-Z-]*$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DestinationTable

FieldFolder

All content copied from https://docs.aws.amazon.com/.
