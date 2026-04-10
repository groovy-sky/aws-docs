This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template TableSortConfiguration

The sort configuration for a `TableVisual`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PaginationConfiguration" : PaginationConfiguration,
  "RowSort" : [ FieldSortOptions, ... ]
}

```

### YAML

```yaml

  PaginationConfiguration:
    PaginationConfiguration
  RowSort:
    - FieldSortOptions

```

## Properties

`PaginationConfiguration`

The pagination configuration (page size, page number) for the table.

_Required_: No

_Type_: [PaginationConfiguration](aws-properties-quicksight-template-paginationconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RowSort`

The field sort options for rows in the table.

_Required_: No

_Type_: Array of [FieldSortOptions](aws-properties-quicksight-template-fieldsortoptions.md)

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TableSideBorderOptions

TableStyleTarget

All content copied from https://docs.aws.amazon.com/.
