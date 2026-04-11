This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSet SourceTable

A source table that provides initial data from either a physical table or parent dataset.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DataSet" : ParentDataSet,
  "PhysicalTableId" : String
}

```

### YAML

```yaml

  DataSet:
    ParentDataSet
  PhysicalTableId: String

```

## Properties

`DataSet`

A parent dataset that serves as the data source instead of a physical table.

_Required_: No

_Type_: [ParentDataSet](aws-properties-quicksight-dataset-parentdataset.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PhysicalTableId`

The identifier of the physical table that serves as the data source.

_Required_: No

_Type_: String

_Pattern_: `^[0-9a-zA-Z-]*$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SemanticTable

StringDatasetParameter

All content copied from https://docs.aws.amazon.com/.
