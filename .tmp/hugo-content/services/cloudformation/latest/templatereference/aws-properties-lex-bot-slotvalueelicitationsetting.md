This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot SlotValueElicitationSetting

Specifies the elicitation setting details eliciting a slot.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DefaultValueSpecification" : SlotDefaultValueSpecification,
  "PromptSpecification" : PromptSpecification,
  "SampleUtterances" : [ SampleUtterance, ... ],
  "SlotCaptureSetting" : SlotCaptureSetting,
  "SlotConstraint" : String,
  "WaitAndContinueSpecification" : WaitAndContinueSpecification
}

```

### YAML

```yaml

  DefaultValueSpecification:
    SlotDefaultValueSpecification
  PromptSpecification:
    PromptSpecification
  SampleUtterances:
    - SampleUtterance
  SlotCaptureSetting:
    SlotCaptureSetting
  SlotConstraint: String
  WaitAndContinueSpecification:
    WaitAndContinueSpecification

```

## Properties

`DefaultValueSpecification`

A list of default values for a slot. Default values are used when
Amazon Lex hasn't determined a value for a slot. You can specify default
values from context variables, session attributes, and defined
values.

_Required_: No

_Type_: [SlotDefaultValueSpecification](aws-properties-lex-bot-slotdefaultvaluespecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PromptSpecification`

The prompt that Amazon Lex uses to elicit the slot value from the
user.

_Required_: No

_Type_: [PromptSpecification](aws-properties-lex-bot-promptspecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SampleUtterances`

If you know a specific pattern that users might respond to an Amazon Lex
request for a slot value, you can provide those utterances to improve
accuracy. This is optional. In most cases, Amazon Lex is capable of
understanding user utterances.

_Required_: No

_Type_: Array of [SampleUtterance](aws-properties-lex-bot-sampleutterance.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SlotCaptureSetting`

Specifies the settings that Amazon Lex uses when a slot
value is successfully entered by a user.

_Required_: No

_Type_: [SlotCaptureSetting](aws-properties-lex-bot-slotcapturesetting.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SlotConstraint`

Specifies whether the slot is required or optional.

_Required_: Yes

_Type_: String

_Allowed values_: `Required | Optional`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WaitAndContinueSpecification`

Specifies the prompts that Amazon Lex uses while a bot is waiting for
customer input.

_Required_: No

_Type_: [WaitAndContinueSpecification](aws-properties-lex-bot-waitandcontinuespecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SlotValue

SlotValueOverride

All content copied from https://docs.aws.amazon.com/.
