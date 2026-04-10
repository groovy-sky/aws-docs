This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template DataSetConfiguration

Dataset configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ColumnGroupSchemaList" : [ ColumnGroupSchema, ... ],
  "DataSetSchema" : DataSetSchema,
  "Placeholder" : String
}

```

### YAML

```yaml

  ColumnGroupSchemaList:
    - ColumnGroupSchema
  DataSetSchema:
    DataSetSchema
  Placeholder: String

```

## Properties

`ColumnGroupSchemaList`

A structure containing the list of column group schemas.

_Required_: No

_Type_: Array of [ColumnGroupSchema](aws-properties-quicksight-template-columngroupschema.md)

_Minimum_: `0`

_Maximum_: `500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataSetSchema`

Dataset schema.

_Required_: No

_Type_: [DataSetSchema](aws-properties-quicksight-template-datasetschema.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Placeholder`

Placeholder.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataPathValue

DataSetReference

All content copied from https://docs.aws.amazon.com/.
