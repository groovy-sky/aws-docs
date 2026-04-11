This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard FieldTooltipItem

The tooltip item for the fields.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FieldId" : String,
  "Label" : String,
  "TooltipTarget" : String,
  "Visibility" : String
}

```

### YAML

```yaml

  FieldId: String
  Label: String
  TooltipTarget: String
  Visibility: String

```

## Properties

`FieldId`

The unique ID of the field that is targeted by the tooltip.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Label`

The label of the tooltip item.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TooltipTarget`

Determines the target of the field tooltip item in a combo chart visual.

_Required_: No

_Type_: String

_Allowed values_: `BOTH | BAR | LINE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Visibility`

The visibility of the tooltip item.

_Required_: No

_Type_: String

_Allowed values_: `HIDDEN | VISIBLE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FieldSortOptions

FilledMapAggregatedFieldWells

All content copied from https://docs.aws.amazon.com/.
