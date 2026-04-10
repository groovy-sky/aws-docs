This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot ImageResponseCard

A card that is shown to the user by a messaging platform. You define
the contents of the card, the card is displayed by the platform.

When you use a response card, the response from the user is
constrained to the text associated with a button on the card.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Buttons" : [ Button, ... ],
  "ImageUrl" : String,
  "Subtitle" : String,
  "Title" : String
}

```

### YAML

```yaml

  Buttons:
    - Button
  ImageUrl: String
  Subtitle: String
  Title: String

```

## Properties

`Buttons`

A list of buttons that should be displayed on the response card. The
arrangement of the buttons is determined by the platform that displays
the button.

_Required_: No

_Type_: Array of [Button](aws-properties-lex-bot-button.md)

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ImageUrl`

The URL of an image to display on the response card. The image URL
must be publicly available so that the platform displaying the response
card has access to the image.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `250`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Subtitle`

The subtitle to display on the response card. The format of the
subtitle is determined by the platform displaying the response
card.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `250`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Title`

The title to display on the response card. The format of the title
is determined by the platform displaying the response card.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `250`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GrammarSlotTypeSource

InitialResponseSetting

All content copied from https://docs.aws.amazon.com/.
