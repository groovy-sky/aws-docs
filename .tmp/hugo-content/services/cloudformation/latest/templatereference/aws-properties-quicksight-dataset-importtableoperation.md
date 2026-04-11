This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSet ImportTableOperation

A transform operation that imports data from a source table.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Alias" : String,
  "Source" : ImportTableOperationSource
}

```

### YAML

```yaml

  Alias: String
  Source:
    ImportTableOperationSource

```

## Properties

`Alias`

Alias for this operation.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Source`

The source configuration that specifies which source table to import and any column mappings.

_Required_: Yes

_Type_: [ImportTableOperationSource](aws-properties-quicksight-dataset-importtableoperationsource.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GeoSpatialColumnGroup

ImportTableOperationSource

All content copied from https://docs.aws.amazon.com/.
