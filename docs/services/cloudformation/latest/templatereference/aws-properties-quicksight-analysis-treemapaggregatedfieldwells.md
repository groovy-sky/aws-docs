This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis TreeMapAggregatedFieldWells

Aggregated field wells of a tree map.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Colors" : [ MeasureField, ... ],
  "Groups" : [ DimensionField, ... ],
  "Sizes" : [ MeasureField, ... ]
}

```

### YAML

```yaml

  Colors:
    - MeasureField
  Groups:
    - DimensionField
  Sizes:
    - MeasureField

```

## Properties

`Colors`

The color field well of a tree map. Values are grouped by aggregations based on group by fields.

_Required_: No

_Type_: Array of [MeasureField](aws-properties-quicksight-analysis-measurefield.md)

_Minimum_: `0`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Groups`

The group by field well of a tree map. Values are grouped based on group by fields.

_Required_: No

_Type_: Array of [DimensionField](aws-properties-quicksight-analysis-dimensionfield.md)

_Minimum_: `0`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Sizes`

The size field well of a tree map. Values are aggregated based on group by fields.

_Required_: No

_Type_: Array of [MeasureField](aws-properties-quicksight-analysis-measurefield.md)

_Minimum_: `0`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TransposedTableOption

TreeMapConfiguration

All content copied from https://docs.aws.amazon.com/.
