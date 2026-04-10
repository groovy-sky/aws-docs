This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSet DestinationTable

Defines a destination table in data preparation that receives the final transformed data.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Alias" : String,
  "Source" : DestinationTableSource
}

```

### YAML

```yaml

  Alias: String
  Source:
    DestinationTableSource

```

## Properties

`Alias`

Alias for the destination table.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Source`

The source configuration that specifies which transform operation provides data to this destination table.

_Required_: Yes

_Type_: [DestinationTableSource](aws-properties-quicksight-dataset-destinationtablesource.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DecimalDatasetParameterDefaultValues

DestinationTableSource

All content copied from https://docs.aws.amazon.com/.
