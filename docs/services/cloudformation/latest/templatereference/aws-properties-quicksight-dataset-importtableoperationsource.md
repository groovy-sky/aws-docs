This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSet ImportTableOperationSource

Specifies the source table and column mappings for an import table operation.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ColumnIdMappings" : [ DataSetColumnIdMapping, ... ],
  "SourceTableId" : String
}

```

### YAML

```yaml

  ColumnIdMappings:
    - DataSetColumnIdMapping
  SourceTableId: String

```

## Properties

`ColumnIdMappings`

The mappings between source column identifiers and target column identifiers during the import.

_Required_: No

_Type_: Array of [DataSetColumnIdMapping](aws-properties-quicksight-dataset-datasetcolumnidmapping.md)

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceTableId`

The identifier of the source table to import data from.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9a-zA-Z-]*$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ImportTableOperation

IncrementalRefresh

All content copied from https://docs.aws.amazon.com/.
