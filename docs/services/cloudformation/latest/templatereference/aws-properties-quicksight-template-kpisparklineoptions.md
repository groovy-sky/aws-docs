This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template KPISparklineOptions

The options that determine the visibility, color, type, and tooltip visibility of the sparkline of a KPI visual.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Color" : String,
  "TooltipVisibility" : String,
  "Type" : String,
  "Visibility" : String
}

```

### YAML

```yaml

  Color: String
  TooltipVisibility: String
  Type: String
  Visibility: String

```

## Properties

`Color`

The color of the sparkline.

_Required_: No

_Type_: String

_Pattern_: `^#[A-F0-9]{6}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TooltipVisibility`

The tooltip visibility of the sparkline.

_Required_: No

_Type_: String

_Allowed values_: `HIDDEN | VISIBLE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of the sparkline.

_Required_: Yes

_Type_: String

_Allowed values_: `LINE | AREA`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Visibility`

The visibility of the sparkline.

_Required_: No

_Type_: String

_Allowed values_: `HIDDEN | VISIBLE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

KPISortConfiguration

KPIVisual

All content copied from https://docs.aws.amazon.com/.
