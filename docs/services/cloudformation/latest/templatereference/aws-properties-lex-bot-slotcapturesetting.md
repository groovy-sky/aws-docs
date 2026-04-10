This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot SlotCaptureSetting

Settings used when Amazon Lex successfully captures a slot
value from a user.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CaptureConditional" : ConditionalSpecification,
  "CaptureNextStep" : DialogState,
  "CaptureResponse" : ResponseSpecification,
  "CodeHook" : DialogCodeHookInvocationSetting,
  "ElicitationCodeHook" : ElicitationCodeHookInvocationSetting,
  "FailureConditional" : ConditionalSpecification,
  "FailureNextStep" : DialogState,
  "FailureResponse" : ResponseSpecification
}

```

### YAML

```yaml

  CaptureConditional:
    ConditionalSpecification
  CaptureNextStep:
    DialogState
  CaptureResponse:
    ResponseSpecification
  CodeHook:
    DialogCodeHookInvocationSetting
  ElicitationCodeHook:
    ElicitationCodeHookInvocationSetting
  FailureConditional:
    ConditionalSpecification
  FailureNextStep:
    DialogState
  FailureResponse:
    ResponseSpecification

```

## Properties

`CaptureConditional`

A list of conditional branches to evaluate after the slot value is
captured.

_Required_: No

_Type_: [ConditionalSpecification](aws-properties-lex-bot-conditionalspecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CaptureNextStep`

Specifies the next step that the bot runs when the slot value is
captured before the code hook times out.

_Required_: No

_Type_: [DialogState](aws-properties-lex-bot-dialogstate.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CaptureResponse`

Specifies a list of message groups that Amazon Lex uses to respond the
user input.

_Required_: No

_Type_: [ResponseSpecification](aws-properties-lex-bot-responsespecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CodeHook`

Code hook called after Amazon Lex successfully captures a
slot value.

_Required_: No

_Type_: [DialogCodeHookInvocationSetting](aws-properties-lex-bot-dialogcodehookinvocationsetting.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ElicitationCodeHook`

Code hook called when Amazon Lex doesn't capture a slot
value.

_Required_: No

_Type_: [ElicitationCodeHookInvocationSetting](aws-properties-lex-bot-elicitationcodehookinvocationsetting.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FailureConditional`

A list of conditional branches to evaluate when the slot value isn't
captured.

_Required_: No

_Type_: [ConditionalSpecification](aws-properties-lex-bot-conditionalspecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FailureNextStep`

Specifies the next step that the bot runs when the slot value code
is not recognized.

_Required_: No

_Type_: [DialogState](aws-properties-lex-bot-dialogstate.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FailureResponse`

Specifies a list of message groups that Amazon Lex uses to respond the
user input when the slot fails to be captured.

_Required_: No

_Type_: [ResponseSpecification](aws-properties-lex-bot-responsespecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Slot

SlotDefaultValue

All content copied from https://docs.aws.amazon.com/.
