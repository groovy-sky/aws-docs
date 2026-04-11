This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSet SemanticTable

A semantic table that represents the final analytical structure of the data.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Alias" : String,
  "DestinationTableId" : String,
  "RowLevelPermissionConfiguration" : RowLevelPermissionConfiguration
}

```

### YAML

```yaml

  Alias: String
  DestinationTableId: String
  RowLevelPermissionConfiguration:
    RowLevelPermissionConfiguration

```

## Properties

`Alias`

Alias for the semantic table.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DestinationTableId`

The identifier of the destination table from data preparation that provides data to this semantic table.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9a-zA-Z-]*$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RowLevelPermissionConfiguration`

Configuration for row level security that control data access for this semantic table.

_Required_: No

_Type_: [RowLevelPermissionConfiguration](aws-properties-quicksight-dataset-rowlevelpermissionconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SemanticModelConfiguration

SourceTable

All content copied from https://docs.aws.amazon.com/.
