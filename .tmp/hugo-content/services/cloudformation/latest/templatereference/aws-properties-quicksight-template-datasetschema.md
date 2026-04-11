This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template DataSetSchema

Dataset schema.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ColumnSchemaList" : [ ColumnSchema, ... ]
}

```

### YAML

```yaml

  ColumnSchemaList:
    - ColumnSchema

```

## Properties

`ColumnSchemaList`

A structure containing the list of column schemas.

_Required_: No

_Type_: Array of [ColumnSchema](aws-properties-quicksight-template-columnschema.md)

_Minimum_: `0`

_Maximum_: `500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataSetReference

DateAxisOptions

All content copied from https://docs.aws.amazon.com/.
