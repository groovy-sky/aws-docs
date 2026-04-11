This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template PivotTablePaginatedReportOptions

The paginated report options for a pivot table visual.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "OverflowColumnHeaderVisibility" : String,
  "VerticalOverflowVisibility" : String
}

```

### YAML

```yaml

  OverflowColumnHeaderVisibility: String
  VerticalOverflowVisibility: String

```

## Properties

`OverflowColumnHeaderVisibility`

The visibility of the repeating header rows on each page.

_Required_: No

_Type_: String

_Allowed values_: `HIDDEN | VISIBLE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VerticalOverflowVisibility`

The visibility of the printing table overflow across pages.

_Required_: No

_Type_: String

_Allowed values_: `HIDDEN | VISIBLE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PivotTableOptions

PivotTableRowsLabelOptions

All content copied from https://docs.aws.amazon.com/.
