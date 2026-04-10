This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template SectionBasedLayoutPaperCanvasSizeOptions

The options for a paper canvas of a section-based layout.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PaperMargin" : Spacing,
  "PaperOrientation" : String,
  "PaperSize" : String
}

```

### YAML

```yaml

  PaperMargin:
    Spacing
  PaperOrientation: String
  PaperSize: String

```

## Properties

`PaperMargin`

Defines the spacing between the canvas content and the top, bottom, left, and right edges.

_Required_: No

_Type_: [Spacing](aws-properties-quicksight-template-spacing.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PaperOrientation`

The paper orientation that
is used to define canvas dimensions. Choose one of the following
options:

- PORTRAIT

- LANDSCAPE

_Required_: No

_Type_: String

_Allowed values_: `PORTRAIT | LANDSCAPE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PaperSize`

The paper size that is used to define canvas dimensions.

_Required_: No

_Type_: String

_Allowed values_: `US_LETTER | US_LEGAL | US_TABLOID_LEDGER | A0 | A1 | A2 | A3 | A4 | A5 | JIS_B4 | JIS_B5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SectionBasedLayoutConfiguration

SectionLayoutConfiguration

All content copied from https://docs.aws.amazon.com/.
