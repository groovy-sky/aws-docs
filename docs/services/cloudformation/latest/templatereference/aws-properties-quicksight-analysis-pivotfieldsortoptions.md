This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis PivotFieldSortOptions

The field sort options for a pivot table sort configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FieldId" : String,
  "SortBy" : PivotTableSortBy
}

```

### YAML

```yaml

  FieldId: String
  SortBy:
    PivotTableSortBy

```

## Properties

`FieldId`

The field ID for the field sort options.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SortBy`

The sort by field for the field sort options.

_Required_: Yes

_Type_: [PivotTableSortBy](aws-properties-quicksight-analysis-pivottablesortby.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PieChartVisual

PivotTableAggregatedFieldWells

All content copied from https://docs.aws.amazon.com/.
