This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template SeriesItem

The series item configuration of a line chart.

This is a union type structure. For this structure to be valid, only one of the attributes can be defined.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DataFieldSeriesItem" : DataFieldSeriesItem,
  "FieldSeriesItem" : FieldSeriesItem
}

```

### YAML

```yaml

  DataFieldSeriesItem:
    DataFieldSeriesItem
  FieldSeriesItem:
    FieldSeriesItem

```

## Properties

`DataFieldSeriesItem`

The data field series item configuration of a line chart.

_Required_: No

_Type_: [DataFieldSeriesItem](aws-properties-quicksight-template-datafieldseriesitem.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FieldSeriesItem`

The field series item configuration of a line chart.

_Required_: No

_Type_: [FieldSeriesItem](aws-properties-quicksight-template-fieldseriesitem.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SelectedSheetsFilterScopeConfiguration

SetParameterValueConfiguration

All content copied from https://docs.aws.amazon.com/.
