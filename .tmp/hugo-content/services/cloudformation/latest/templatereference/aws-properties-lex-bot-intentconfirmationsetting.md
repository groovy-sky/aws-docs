This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot IntentConfirmationSetting

Provides a prompt for making sure that the user is ready for the
intent to be fulfilled.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CodeHook" : DialogCodeHookInvocationSetting,
  "ConfirmationConditional" : ConditionalSpecification,
  "ConfirmationNextStep" : DialogState,
  "ConfirmationResponse" : ResponseSpecification,
  "DeclinationConditional" : ConditionalSpecification,
  "DeclinationNextStep" : DialogState,
  "DeclinationResponse" : ResponseSpecification,
  "ElicitationCodeHook" : ElicitationCodeHookInvocationSetting,
  "FailureConditional" : ConditionalSpecification,
  "FailureNextStep" : DialogState,
  "FailureResponse" : ResponseSpecification,
  "IsActive" : Boolean,
  "PromptSpecification" : PromptSpecification
}

```

### YAML

```yaml

  CodeHook:
    DialogCodeHookInvocationSetting
  ConfirmationConditional:
    ConditionalSpecification
  ConfirmationNextStep:
    DialogState
  ConfirmationResponse:
    ResponseSpecification
  DeclinationConditional:
    ConditionalSpecification
  DeclinationNextStep:
    DialogState
  DeclinationResponse:
    ResponseSpecification
  ElicitationCodeHook:
    ElicitationCodeHookInvocationSetting
  FailureConditional:
    ConditionalSpecification
  FailureNextStep:
    DialogState
  FailureResponse:
    ResponseSpecification
  IsActive: Boolean
  PromptSpecification:
    PromptSpecification

```

## Properties

`CodeHook`

The `DialogCodeHookInvocationSetting` object associated
with intent's confirmation step. The dialog code hook is triggered
based on these invocation settings when the confirmation next step or
declination next step or failure next step is
`InvokeDialogCodeHook`.

_Required_: No

_Type_: [DialogCodeHookInvocationSetting](aws-properties-lex-bot-dialogcodehookinvocationsetting.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConfirmationConditional`

A list of conditional branches to evaluate after the intent is
closed.

_Required_: No

_Type_: [ConditionalSpecification](aws-properties-lex-bot-conditionalspecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConfirmationNextStep`

Specifies the next step that the bot executes when the customer
confirms the intent.

_Required_: No

_Type_: [DialogState](aws-properties-lex-bot-dialogstate.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConfirmationResponse`

Specifies a list of message groups that Amazon Lex uses to respond the
user input.

_Required_: No

_Type_: [ResponseSpecification](aws-properties-lex-bot-responsespecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeclinationConditional`

A list of conditional branches to evaluate after the intent is
declined.

_Required_: No

_Type_: [ConditionalSpecification](aws-properties-lex-bot-conditionalspecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeclinationNextStep`

Specifies the next step that the bot executes when the customer
declines the intent.

_Required_: No

_Type_: [DialogState](aws-properties-lex-bot-dialogstate.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeclinationResponse`

When the user answers "no" to the question defined in
`promptSpecification`, Amazon Lex responds with this response
to acknowledge that the intent was canceled.

_Required_: No

_Type_: [ResponseSpecification](aws-properties-lex-bot-responsespecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ElicitationCodeHook`

The `DialogCodeHookInvocationSetting` used when the code
hook is invoked during confirmation prompt retries.

_Required_: No

_Type_: [ElicitationCodeHookInvocationSetting](aws-properties-lex-bot-elicitationcodehookinvocationsetting.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FailureConditional`

Provides a list of conditional branches. Branches are evaluated in
the order that they are entered in the list. The first branch with a
condition that evaluates to true is executed. The last branch in the
list is the default branch. The default branch should not have any condition
expression. The default branch is executed if no other branch has a
matching condition.

_Required_: No

_Type_: [ConditionalSpecification](aws-properties-lex-bot-conditionalspecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FailureNextStep`

The next step to take in the conversation if the confirmation step
fails.

_Required_: No

_Type_: [DialogState](aws-properties-lex-bot-dialogstate.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FailureResponse`

Specifies a list of message groups that Amazon Lex uses to respond the
user input when the intent confirmation fails.

_Required_: No

_Type_: [ResponseSpecification](aws-properties-lex-bot-responsespecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IsActive`

Specifies whether the intent's confirmation is sent to the user.
When this field is false, confirmation and declination responses aren't
sent. If the `IsActive` field isn't specified, the default is
true.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PromptSpecification`

Prompts the user to confirm the intent. This question should have a
yes or no answer.

Amazon Lex uses this prompt to ensure that the user acknowledges that the
intent is ready for fulfillment. For example, with the
`OrderPizza` intent, you might want to confirm that the
order is correct before placing it. For other intents, such as intents
that simply respond to user questions, you might not need to ask the
user for confirmation before providing the information.

_Required_: Yes

_Type_: [PromptSpecification](aws-properties-lex-bot-promptspecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IntentClosingSetting

IntentDisambiguationSettings

All content copied from https://docs.aws.amazon.com/.
