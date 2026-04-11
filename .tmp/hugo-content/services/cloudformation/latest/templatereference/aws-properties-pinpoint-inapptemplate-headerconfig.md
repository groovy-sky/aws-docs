This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pinpoint::InAppTemplate HeaderConfig

Specifies the configuration and content of the header or title text of the in-app
message.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Alignment" : String,
  "Header" : String,
  "TextColor" : String
}

```

### YAML

```yaml

  Alignment: String
  Header: String
  TextColor: String

```

## Properties

`Alignment`

The text alignment of the title of the message. Acceptable values: `LEFT`,
`CENTER`, `RIGHT`.

_Required_: No

_Type_: String

_Allowed values_: `LEFT | CENTER | RIGHT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Header`

The title text of the in-app message.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TextColor`

The color of the title text, expressed as a hex color code (such as #000000 for
black).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DefaultButtonConfiguration

InAppMessageContent

All content copied from https://docs.aws.amazon.com/.
