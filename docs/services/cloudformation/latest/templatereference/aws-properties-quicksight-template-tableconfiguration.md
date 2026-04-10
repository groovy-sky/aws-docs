This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template TableConfiguration

The configuration for a `TableVisual`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FieldOptions" : TableFieldOptions,
  "FieldWells" : TableFieldWells,
  "Interactions" : VisualInteractionOptions,
  "PaginatedReportOptions" : TablePaginatedReportOptions,
  "SortConfiguration" : TableSortConfiguration,
  "TableInlineVisualizations" : [ TableInlineVisualization, ... ],
  "TableOptions" : TableOptions,
  "TotalOptions" : TotalOptions
}

```

### YAML

```yaml

  FieldOptions:
    TableFieldOptions
  FieldWells:
    TableFieldWells
  Interactions:
    VisualInteractionOptions
  PaginatedReportOptions:
    TablePaginatedReportOptions
  SortConfiguration:
    TableSortConfiguration
  TableInlineVisualizations:
    - TableInlineVisualization
  TableOptions:
    TableOptions
  TotalOptions:
    TotalOptions

```

## Properties

`FieldOptions`

The field options for a table visual.

_Required_: No

_Type_: [TableFieldOptions](aws-properties-quicksight-template-tablefieldoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FieldWells`

The field wells of the visual.

_Required_: No

_Type_: [TableFieldWells](aws-properties-quicksight-template-tablefieldwells.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Interactions`

The general visual interactions setup for a visual.

_Required_: No

_Type_: [VisualInteractionOptions](aws-properties-quicksight-template-visualinteractionoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PaginatedReportOptions`

The paginated report options for a table visual.

_Required_: No

_Type_: [TablePaginatedReportOptions](aws-properties-quicksight-template-tablepaginatedreportoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SortConfiguration`

The sort configuration for a `TableVisual`.

_Required_: No

_Type_: [TableSortConfiguration](aws-properties-quicksight-template-tablesortconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TableInlineVisualizations`

A collection of inline visualizations to display within a chart.

_Required_: No

_Type_: Array of [TableInlineVisualization](aws-properties-quicksight-template-tableinlinevisualization.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TableOptions`

The table options for a table visual.

_Required_: No

_Type_: [TableOptions](aws-properties-quicksight-template-tableoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TotalOptions`

The total options for a table visual.

_Required_: No

_Type_: [TotalOptions](aws-properties-quicksight-template-totaloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TableConditionalFormattingOption

TableFieldCustomIconContent

All content copied from https://docs.aws.amazon.com/.
