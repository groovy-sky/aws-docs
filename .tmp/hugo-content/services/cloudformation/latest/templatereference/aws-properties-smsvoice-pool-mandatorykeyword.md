This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SMSVOICE::Pool MandatoryKeyword

The keywords `HELP` and `STOP` are mandatory keywords that each phone number must have. For more information, see
[Keywords](../../../sms-voice/latest/userguide/keywords.md) in the AWS End User Messaging SMS User Guide.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Message" : String
}

```

### YAML

```yaml

  Message: String

```

## Properties

`Message`

The message associated with the keyword.

_Required_: Yes

_Type_: String

_Pattern_: `^(?!\s*$)[\s\S]+$`

_Maximum_: `1600`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::SMSVOICE::Pool

MandatoryKeywords

All content copied from https://docs.aws.amazon.com/.
