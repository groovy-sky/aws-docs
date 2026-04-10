This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template TablePaginatedReportOptions

The paginated report options for a table visual.

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

The visibility of repeating header rows on each page.

_Required_: No

_Type_: String

_Allowed values_: `HIDDEN | VISIBLE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VerticalOverflowVisibility`

The visibility of printing table overflow across pages.

_Required_: No

_Type_: String

_Allowed values_: `HIDDEN | VISIBLE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TableOptions

TablePinnedFieldOptions

All content copied from https://docs.aws.amazon.com/.
