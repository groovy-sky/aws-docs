This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::Workspace PaletteHeader

Contains color configuration for header elements in a workspace theme.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Background" : String,
  "InvertActionsColors" : Boolean,
  "Text" : String,
  "TextHover" : String
}

```

### YAML

```yaml

  Background: String
  InvertActionsColors: Boolean
  Text: String
  TextHover: String

```

## Properties

`Background`

The background color of the header.

_Required_: No

_Type_: String

_Pattern_: `.*\S.*`

_Minimum_: `1`

_Maximum_: `127`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InvertActionsColors`

Whether to invert the colors of action buttons in the header.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Text`

The text color in the header.

_Required_: No

_Type_: String

_Pattern_: `.*\S.*`

_Minimum_: `1`

_Maximum_: `127`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TextHover`

The text color when hovering over header elements.

_Required_: No

_Type_: String

_Pattern_: `.*\S.*`

_Minimum_: `1`

_Maximum_: `127`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PaletteCanvas

PaletteNavigation

All content copied from https://docs.aws.amazon.com/.
