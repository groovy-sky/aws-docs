This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template PivotTableConfiguration

The configuration for a `PivotTableVisual`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FieldOptions" : PivotTableFieldOptions,
  "FieldWells" : PivotTableFieldWells,
  "Interactions" : VisualInteractionOptions,
  "PaginatedReportOptions" : PivotTablePaginatedReportOptions,
  "SortConfiguration" : PivotTableSortConfiguration,
  "TableOptions" : PivotTableOptions,
  "TotalOptions" : PivotTableTotalOptions
}

```

### YAML

```yaml

  FieldOptions:
    PivotTableFieldOptions
  FieldWells:
    PivotTableFieldWells
  Interactions:
    VisualInteractionOptions
  PaginatedReportOptions:
    PivotTablePaginatedReportOptions
  SortConfiguration:
    PivotTableSortConfiguration
  TableOptions:
    PivotTableOptions
  TotalOptions:
    PivotTableTotalOptions

```

## Properties

`FieldOptions`

The field options for a pivot table visual.

_Required_: No

_Type_: [PivotTableFieldOptions](aws-properties-quicksight-template-pivottablefieldoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FieldWells`

The field wells of the visual.

_Required_: No

_Type_: [PivotTableFieldWells](aws-properties-quicksight-template-pivottablefieldwells.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Interactions`

The general visual interactions setup for a visual.

_Required_: No

_Type_: [VisualInteractionOptions](aws-properties-quicksight-template-visualinteractionoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PaginatedReportOptions`

The paginated report options for a pivot table visual.

_Required_: No

_Type_: [PivotTablePaginatedReportOptions](aws-properties-quicksight-template-pivottablepaginatedreportoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SortConfiguration`

The sort configuration for a `PivotTableVisual`.

_Required_: No

_Type_: [PivotTableSortConfiguration](aws-properties-quicksight-template-pivottablesortconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TableOptions`

The table options for a pivot table visual.

_Required_: No

_Type_: [PivotTableOptions](aws-properties-quicksight-template-pivottableoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TotalOptions`

The total options for a pivot table visual.

_Required_: No

_Type_: [PivotTableTotalOptions](aws-properties-quicksight-template-pivottabletotaloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PivotTableConditionalFormattingScope

PivotTableDataPathOption

All content copied from https://docs.aws.amazon.com/.
