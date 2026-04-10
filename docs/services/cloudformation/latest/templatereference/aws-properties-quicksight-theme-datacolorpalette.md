This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Theme DataColorPalette

The theme colors that are used for data colors in charts. The colors description is a
hexadecimal color code that consists of six alphanumerical characters, prefixed with
`#`, for example #37BFF5.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Colors" : [ String, ... ],
  "EmptyFillColor" : String,
  "MinMaxGradient" : [ String, ... ]
}

```

### YAML

```yaml

  Colors:
    - String
  EmptyFillColor: String
  MinMaxGradient:
    - String

```

## Properties

`Colors`

The hexadecimal codes for the colors.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EmptyFillColor`

The hexadecimal code of a color that applies to charts where a lack of data is
highlighted.

_Required_: No

_Type_: String

_Pattern_: `^#[A-F0-9]{6}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinMaxGradient`

The minimum and maximum hexadecimal codes that describe a color gradient.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BorderStyle

Font

All content copied from https://docs.aws.amazon.com/.
