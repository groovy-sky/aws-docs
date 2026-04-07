This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSet DataPrepConfiguration

Configuration for data preparation operations, defining the complete pipeline from source tables
through transformations to destination tables.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DestinationTableMap" : {Key: Value, ...},
  "SourceTableMap" : {Key: Value, ...},
  "TransformStepMap" : {Key: Value, ...}
}

```

### YAML

```yaml

  DestinationTableMap:
    Key: Value
  SourceTableMap:
    Key: Value
  TransformStepMap:
    Key: Value

```

## Properties

`DestinationTableMap`

A map of destination tables that receive the final prepared data.

_Required_: Yes

_Type_: Object of [DestinationTable](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-dataset-destinationtable.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceTableMap`

A map of source tables that provide information about underlying sources.

_Required_: Yes

_Type_: Object of [SourceTable](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-dataset-sourcetable.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TransformStepMap`

A map of transformation steps that process the data.

_Required_: Yes

_Type_: Object of [TransformStep](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-dataset-transformstep.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DataPrepAggregationFunction

DataPrepListAggregationFunction
