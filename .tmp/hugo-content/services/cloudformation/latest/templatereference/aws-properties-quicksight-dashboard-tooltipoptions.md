This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard TooltipOptions

The display options for the visual tooltip.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FieldBasedTooltip" : FieldBasedTooltip,
  "SelectedTooltipType" : String,
  "TooltipVisibility" : String
}

```

### YAML

```yaml

  FieldBasedTooltip:
    FieldBasedTooltip
  SelectedTooltipType: String
  TooltipVisibility: String

```

## Properties

`FieldBasedTooltip`

The setup for the detailed tooltip. The tooltip setup is always saved. The display type is decided based on the tooltip type.

_Required_: No

_Type_: [FieldBasedTooltip](aws-properties-quicksight-dashboard-fieldbasedtooltip.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SelectedTooltipType`

The selected type for the tooltip. Choose one of the following options:

- `BASIC`: A basic tooltip.

- `DETAILED`: A detailed tooltip.

_Required_: No

_Type_: String

_Allowed values_: `BASIC | DETAILED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TooltipVisibility`

Determines whether or not the tooltip is visible.

_Required_: No

_Type_: String

_Allowed values_: `HIDDEN | VISIBLE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TooltipItem

TopBottomFilter

All content copied from https://docs.aws.amazon.com/.
