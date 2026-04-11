This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis ColorScale

Determines the color scale that is applied to the visual.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ColorFillType" : String,
  "Colors" : [ DataColor, ... ],
  "NullValueColor" : DataColor
}

```

### YAML

```yaml

  ColorFillType: String
  Colors:
    - DataColor
  NullValueColor:
    DataColor

```

## Properties

`ColorFillType`

Determines the color fill type.

_Required_: Yes

_Type_: String

_Allowed values_: `DISCRETE | GRADIENT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Colors`

Determines the list of colors that are applied to the visual.

_Required_: Yes

_Type_: Array of [DataColor](aws-properties-quicksight-analysis-datacolor.md)

_Minimum_: `2`

_Maximum_: `3`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NullValueColor`

Determines the color that is applied to null values.

_Required_: No

_Type_: [DataColor](aws-properties-quicksight-analysis-datacolor.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ClusterMarkerConfiguration

ColorsConfiguration

All content copied from https://docs.aws.amazon.com/.
