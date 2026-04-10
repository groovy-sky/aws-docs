This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SMSVOICE::PhoneNumber MandatoryKeywords

The keywords `HELP` and `STOP` are mandatory keywords that each phone number must have. For more information, see
[Keywords](../../../sms-voice/latest/userguide/keywords.md) in the AWS End User Messaging SMS User Guide.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "HELP" : MandatoryKeyword,
  "STOP" : MandatoryKeyword
}

```

### YAML

```yaml

  HELP:
    MandatoryKeyword
  STOP:
    MandatoryKeyword

```

## Properties

`HELP`

Specifies the `HELP` keyword that customers use to obtain customer support for this phone number. For more information, see
[Keywords](../../../sms-voice/latest/userguide/keywords.md) in the AWS End User Messaging SMS User Guide.

_Required_: Yes

_Type_: [MandatoryKeyword](aws-properties-smsvoice-phonenumber-mandatorykeyword.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`STOP`

Specifies the `STOP` keyword that customers use to opt out of receiving messages from this phone number. For more information, see
[Required opt-out keywords](../../../sms-voice/latest/userguide/keywords-required.md) in the AWS End User Messaging SMS User Guide.

_Required_: Yes

_Type_: [MandatoryKeyword](aws-properties-smsvoice-phonenumber-mandatorykeyword.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MandatoryKeyword

OptionalKeyword

All content copied from https://docs.aws.amazon.com/.
