This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SMSVOICE::Pool MandatoryKeywords

The manadatory keywords, `HELP` and `STOP` to add to the pool. For more information, see
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

Specifies the pool's `HELP` keyword. For more information, see
[Opt out list required keywords](../../../sms-voice/latest/userguide/opt-out-list-keywords.md) in the AWS End User Messaging SMS User Guide.

_Required_: Yes

_Type_: [MandatoryKeyword](aws-properties-smsvoice-pool-mandatorykeyword.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`STOP`

Specifies the pool's opt-out keyword. For more information, see
[Required opt-out keywords](../../../sms-voice/latest/userguide/keywords-required.md) in the AWS End User Messaging SMS User Guide.

_Required_: Yes

_Type_: [MandatoryKeyword](aws-properties-smsvoice-pool-mandatorykeyword.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MandatoryKeyword

OptionalKeyword

All content copied from https://docs.aws.amazon.com/.
