This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot Button

Describes a button to use on a response card used to gather slot
values from a user.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Text" : String,
  "Value" : String
}

```

### YAML

```yaml

  Text: String
  Value: String

```

## Properties

`Text`

The text that appears on the button. Use this to tell the user what
value is returned when they choose this button.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The value returned to Amazon Lex when the user chooses this button. This
must be one of the slot values configured for the slot.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BuildtimeSettings

CloudWatchLogGroupLogDestination

All content copied from https://docs.aws.amazon.com/.
