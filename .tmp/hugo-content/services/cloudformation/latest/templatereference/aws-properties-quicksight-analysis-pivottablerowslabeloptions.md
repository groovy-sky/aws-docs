This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis PivotTableRowsLabelOptions

The options for the label thta is located above the row headers. This option is only applicable when `RowsLayout` is set to `HIERARCHY`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CustomLabel" : String,
  "Visibility" : String
}

```

### YAML

```yaml

  CustomLabel: String
  Visibility: String

```

## Properties

`CustomLabel`

The custom label string for the rows label.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Visibility`

The visibility of the rows label.

_Required_: No

_Type_: String

_Allowed values_: `HIDDEN | VISIBLE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PivotTablePaginatedReportOptions

PivotTableSortBy

All content copied from https://docs.aws.amazon.com/.
