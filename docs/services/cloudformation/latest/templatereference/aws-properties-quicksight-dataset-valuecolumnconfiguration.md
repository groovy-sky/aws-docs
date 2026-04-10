This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSet ValueColumnConfiguration

Configuration for how to handle value columns in pivot operations, including aggregation settings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AggregationFunction" : DataPrepAggregationFunction
}

```

### YAML

```yaml

  AggregationFunction:
    DataPrepAggregationFunction

```

## Properties

`AggregationFunction`

The aggregation function to apply when multiple values map to the same pivoted cell.

_Required_: No

_Type_: [DataPrepAggregationFunction](aws-properties-quicksight-dataset-dataprepaggregationfunction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UploadSettings

AWS::QuickSight::DataSource

All content copied from https://docs.aws.amazon.com/.
