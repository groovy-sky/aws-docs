This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot PostDialogCodeHookInvocationSpecification

Specifies next steps to run after the dialog code hook
finishes.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FailureConditional" : ConditionalSpecification,
  "FailureNextStep" : DialogState,
  "FailureResponse" : ResponseSpecification,
  "SuccessConditional" : ConditionalSpecification,
  "SuccessNextStep" : DialogState,
  "SuccessResponse" : ResponseSpecification,
  "TimeoutConditional" : ConditionalSpecification,
  "TimeoutNextStep" : DialogState,
  "TimeoutResponse" : ResponseSpecification
}

```

### YAML

```yaml

  FailureConditional:
    ConditionalSpecification
  FailureNextStep:
    DialogState
  FailureResponse:
    ResponseSpecification
  SuccessConditional:
    ConditionalSpecification
  SuccessNextStep:
    DialogState
  SuccessResponse:
    ResponseSpecification
  TimeoutConditional:
    ConditionalSpecification
  TimeoutNextStep:
    DialogState
  TimeoutResponse:
    ResponseSpecification

```

## Properties

`FailureConditional`

A list of conditional branches to evaluate after the dialog code
hook throws an exception or returns with the `State` field
of the `Intent` object set to `Failed`.

_Required_: No

_Type_: [ConditionalSpecification](aws-properties-lex-bot-conditionalspecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FailureNextStep`

Specifies the next step the bot runs after the dialog code hook
throws an exception or returns with the `State` field of the
`Intent` object set to `Failed`.

_Required_: No

_Type_: [DialogState](aws-properties-lex-bot-dialogstate.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FailureResponse`

Specifies a list of message groups that Amazon Lex uses to respond the
user input when the code hook fails.

_Required_: No

_Type_: [ResponseSpecification](aws-properties-lex-bot-responsespecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SuccessConditional`

A list of conditional branches to evaluate after the dialog code
hook finishes successfully.

_Required_: No

_Type_: [ConditionalSpecification](aws-properties-lex-bot-conditionalspecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SuccessNextStep`

Specifics the next step the bot runs after the dialog code hook
finishes successfully.

_Required_: No

_Type_: [DialogState](aws-properties-lex-bot-dialogstate.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SuccessResponse`

Specifies a list of message groups that Amazon Lex uses to respond when
the code hook succeeds.

_Required_: No

_Type_: [ResponseSpecification](aws-properties-lex-bot-responsespecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeoutConditional`

A list of conditional branches to evaluate if the code hook times
out.

_Required_: No

_Type_: [ConditionalSpecification](aws-properties-lex-bot-conditionalspecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeoutNextStep`

Specifies the next step that the bot runs when the code hook times
out.

_Required_: No

_Type_: [DialogState](aws-properties-lex-bot-dialogstate.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeoutResponse`

Specifies a list of message groups that Amazon Lex uses to respond to the
user input when the code hook times out.

_Required_: No

_Type_: [ResponseSpecification](aws-properties-lex-bot-responsespecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PlainTextMessage

PostFulfillmentStatusSpecification

All content copied from https://docs.aws.amazon.com/.
