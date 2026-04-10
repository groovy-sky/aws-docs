This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SMSVOICE::Pool OptionalKeyword

The pool's `OptionalKeyword ` configuration. For more information, see
[Keywords](../../../sms-voice/latest/userguide/keywords.md) in the AWS End User Messaging SMS User Guide.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Action" : String,
  "Keyword" : String,
  "Message" : String
}

```

### YAML

```yaml

  Action: String
  Keyword: String
  Message: String

```

## Properties

`Action`

The action to perform when the keyword is used.

_Required_: Yes

_Type_: String

_Allowed values_: `AUTOMATIC_RESPONSE | OPT_OUT | OPT_IN`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Keyword`

The new keyword to add.

_Required_: Yes

_Type_: String

_Pattern_: `^(?!(?i)(stop|help)$)[ \S]+`

_Maximum_: `30`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Message`

The message associated with the keyword.

_Required_: Yes

_Type_: String

_Pattern_: `^(?!\s*$)[\s\S]+$`

_Maximum_: `1600`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MandatoryKeywords

Tag

All content copied from https://docs.aws.amazon.com/.
