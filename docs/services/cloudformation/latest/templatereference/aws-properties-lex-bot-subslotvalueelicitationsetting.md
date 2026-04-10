This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot SubSlotValueElicitationSetting

Subslot elicitation settings.

`DefaultValueSpecification` is a list of default values for a constituent sub slot in a composite slot. Default values are used when
Amazon Lex hasn't determined a value for a slot. You can specify default values from context variables,
session attributes, and defined values. This is similar to `DefaultValueSpecification` for slots.

`PromptSpecification` is the prompt that Amazon Lex uses to elicit the sub slot value from the user.
This is similar to `PromptSpecification` for slots.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DefaultValueSpecification" : SlotDefaultValueSpecification,
  "PromptSpecification" : PromptSpecification,
  "SampleUtterances" : [ SampleUtterance, ... ],
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
  WaitAndContinueSpecification:
    WaitAndContinueSpecification

```

## Properties

`DefaultValueSpecification`

Property description not available.

_Required_: No

_Type_: [SlotDefaultValueSpecification](aws-properties-lex-bot-slotdefaultvaluespecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PromptSpecification`

Property description not available.

_Required_: No

_Type_: [PromptSpecification](aws-properties-lex-bot-promptspecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SampleUtterances`

If you know a specific pattern that users might respond to an Amazon Lex request for a sub slot value,
you can provide those utterances to improve accuracy. This is optional. In most cases Amazon Lex is capable
of understanding user utterances. This is similar to `SampleUtterances` for slots.

_Required_: No

_Type_: Array of [SampleUtterance](aws-properties-lex-bot-sampleutterance.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WaitAndContinueSpecification`

Property description not available.

_Required_: No

_Type_: [WaitAndContinueSpecification](aws-properties-lex-bot-waitandcontinuespecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SubSlotTypeComposition

Tag

All content copied from https://docs.aws.amazon.com/.
