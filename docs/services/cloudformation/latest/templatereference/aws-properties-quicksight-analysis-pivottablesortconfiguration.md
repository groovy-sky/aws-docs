This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis PivotTableSortConfiguration

The sort configuration for a `PivotTableVisual`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FieldSortOptions" : [ PivotFieldSortOptions, ... ]
}

```

### YAML

```yaml

  FieldSortOptions:
    - PivotFieldSortOptions

```

## Properties

`FieldSortOptions`

The field sort options for a pivot table sort configuration.

_Required_: No

_Type_: [Array](aws-properties-quicksight-analysis-fieldsortoptions.md) of [PivotFieldSortOptions](aws-properties-quicksight-analysis-pivotfieldsortoptions.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PivotTableSortBy

PivotTableTotalOptions

All content copied from https://docs.aws.amazon.com/.
