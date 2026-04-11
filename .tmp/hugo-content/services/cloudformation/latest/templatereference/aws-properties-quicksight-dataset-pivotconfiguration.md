This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSet PivotConfiguration

Configuration for a pivot operation, specifying which column contains labels and how to pivot them.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LabelColumnName" : String,
  "PivotedLabels" : [ PivotedLabel, ... ]
}

```

### YAML

```yaml

  LabelColumnName: String
  PivotedLabels:
    - PivotedLabel

```

## Properties

`LabelColumnName`

The name of the column that contains the labels to be pivoted into separate columns.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `127`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PivotedLabels`

The list of specific label values to pivot into separate columns.

_Required_: Yes

_Type_: Array of [PivotedLabel](aws-properties-quicksight-dataset-pivotedlabel.md)

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PhysicalTable

PivotedLabel

All content copied from https://docs.aws.amazon.com/.
