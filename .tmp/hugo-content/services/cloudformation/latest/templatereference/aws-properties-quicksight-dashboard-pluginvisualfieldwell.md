This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard PluginVisualFieldWell

A collection of field wells for a plugin visual.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AxisName" : String,
  "Dimensions" : [ DimensionField, ... ],
  "Measures" : [ MeasureField, ... ],
  "Unaggregated" : [ UnaggregatedField, ... ]
}

```

### YAML

```yaml

  AxisName: String
  Dimensions:
    - DimensionField
  Measures:
    - MeasureField
  Unaggregated:
    - UnaggregatedField

```

## Properties

`AxisName`

The semantic axis name for the field well.

_Required_: No

_Type_: String

_Allowed values_: `GROUP_BY | VALUE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Dimensions`

A list of dimensions for the field well.

_Required_: No

_Type_: Array of [DimensionField](aws-properties-quicksight-dashboard-dimensionfield.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Measures`

A list of measures that exist in the field well.

_Required_: No

_Type_: Array of [MeasureField](aws-properties-quicksight-dashboard-measurefield.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Unaggregated`

A list of unaggregated fields that exist in the field well.

_Required_: No

_Type_: Array of [UnaggregatedField](aws-properties-quicksight-dashboard-unaggregatedfield.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PluginVisualConfiguration

PluginVisualItemsLimitConfiguration

All content copied from https://docs.aws.amazon.com/.
