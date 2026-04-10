This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis FieldBasedTooltip

The setup for the detailed tooltip.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AggregationVisibility" : String,
  "TooltipFields" : [ TooltipItem, ... ],
  "TooltipTitleType" : String
}

```

### YAML

```yaml

  AggregationVisibility: String
  TooltipFields:
    - TooltipItem
  TooltipTitleType: String

```

## Properties

`AggregationVisibility`

The visibility of `Show aggregations`.

_Required_: No

_Type_: String

_Allowed values_: `HIDDEN | VISIBLE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TooltipFields`

The fields configuration in the
tooltip.

_Required_: No

_Type_: Array of [TooltipItem](aws-properties-quicksight-analysis-tooltipitem.md)

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TooltipTitleType`

The type for the >tooltip title. Choose one of the following options:

- `NONE`: Doesn't use the primary value as the title.

- `PRIMARY_VALUE`: Uses primary value as the title.

_Required_: No

_Type_: String

_Allowed values_: `NONE | PRIMARY_VALUE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ExplicitHierarchy

FieldLabelType

All content copied from https://docs.aws.amazon.com/.
