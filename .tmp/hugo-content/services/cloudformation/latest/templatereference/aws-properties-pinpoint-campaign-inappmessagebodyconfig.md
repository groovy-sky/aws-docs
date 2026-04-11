This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pinpoint::Campaign InAppMessageBodyConfig

Specifies the configuration of main body text of the in-app message.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Alignment" : String,
  "Body" : String,
  "TextColor" : String
}

```

### YAML

```yaml

  Alignment: String
  Body: String
  TextColor: String

```

## Properties

`Alignment`

The text alignment of the main body text of the message. Acceptable values:
`LEFT`, `CENTER`, `RIGHT`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Body`

The main body text of the message.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TextColor`

The color of the body text, expressed as a string consisting of a hex color
code (such as "#000000" for black).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EventDimensions

InAppMessageButton

All content copied from https://docs.aws.amazon.com/.
