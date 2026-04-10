This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template FieldSortOptions

The field sort options in a chart configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ColumnSort" : ColumnSort,
  "FieldSort" : FieldSort
}

```

### YAML

```yaml

  ColumnSort:
    ColumnSort
  FieldSort:
    FieldSort

```

## Properties

`ColumnSort`

The sort configuration for a column that is not used in a field well.

_Required_: No

_Type_: [ColumnSort](aws-properties-quicksight-template-columnsort.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FieldSort`

The sort configuration for a field in a field well.

_Required_: No

_Type_: [FieldSort](aws-properties-quicksight-template-fieldsort.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FieldSort

FieldTooltipItem

All content copied from https://docs.aws.amazon.com/.
