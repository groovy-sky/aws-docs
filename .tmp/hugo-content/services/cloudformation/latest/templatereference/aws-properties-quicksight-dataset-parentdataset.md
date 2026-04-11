This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSet ParentDataSet

References a parent dataset that serves as a data source, including its columns and metadata.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DataSetArn" : String,
  "InputColumns" : [ InputColumn, ... ]
}

```

### YAML

```yaml

  DataSetArn: String
  InputColumns:
    - InputColumn

```

## Properties

`DataSetArn`

The Amazon Resource Name (ARN) of the parent dataset.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InputColumns`

The list of input columns available from the parent dataset.

_Required_: Yes

_Type_: Array of [InputColumn](aws-properties-quicksight-dataset-inputcolumn.md)

_Minimum_: `0`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OverrideDatasetParameterOperation

PerformanceConfiguration

All content copied from https://docs.aws.amazon.com/.
