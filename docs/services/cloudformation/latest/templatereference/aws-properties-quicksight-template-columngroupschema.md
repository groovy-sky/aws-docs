This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template ColumnGroupSchema

The column group schema.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ColumnGroupColumnSchemaList" : [ ColumnGroupColumnSchema, ... ],
  "Name" : String
}

```

### YAML

```yaml

  ColumnGroupColumnSchemaList:
    - ColumnGroupColumnSchema
  Name: String

```

## Properties

`ColumnGroupColumnSchemaList`

A structure containing the list of schemas for column group columns.

_Required_: No

_Type_: Array of [ColumnGroupColumnSchema](aws-properties-quicksight-template-columngroupcolumnschema.md)

_Minimum_: `0`

_Maximum_: `500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the column group schema.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ColumnGroupColumnSchema

ColumnHierarchy

All content copied from https://docs.aws.amazon.com/.
