This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pinpoint::InAppTemplate DefaultButtonConfiguration

Specifies the default behavior of a button that appears in an in-app message. You can
optionally add button configurations that specifically apply to iOS, Android, or web
browser users.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BackgroundColor" : String,
  "BorderRadius" : Integer,
  "ButtonAction" : String,
  "Link" : String,
  "Text" : String,
  "TextColor" : String
}

```

### YAML

```yaml

  BackgroundColor: String
  BorderRadius: Integer
  ButtonAction: String
  Link: String
  Text: String
  TextColor: String

```

## Properties

`BackgroundColor`

The background color of a button, expressed as a hex color code (such as #000000 for
black).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BorderRadius`

The border radius of a button.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ButtonAction`

The action that occurs when a recipient chooses a button in an in-app message.
You can specify one of the following:

- `LINK` – A link to a web destination.

- `DEEP_LINK` – A link to a specific page in an
application.

- `CLOSE` – Dismisses the message.

_Required_: No

_Type_: String

_Allowed values_: `LINK | DEEP_LINK | CLOSE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Link`

The destination (such as a URL) for a button.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Text`

The text that appears on a button in an in-app message.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TextColor`

The color of the body text in a button, expressed as a hex color code (such as #000000
for black).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ButtonConfig

HeaderConfig

All content copied from https://docs.aws.amazon.com/.
