This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot IntentClosingSetting

Provides a statement the Amazon Lex conveys to the user when the intent
is successfully fulfilled.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ClosingResponse" : ResponseSpecification,
  "Conditional" : ConditionalSpecification,
  "IsActive" : Boolean,
  "NextStep" : DialogState
}

```

### YAML

```yaml

  ClosingResponse:
    ResponseSpecification
  Conditional:
    ConditionalSpecification
  IsActive: Boolean
  NextStep:
    DialogState

```

## Properties

`ClosingResponse`

The response that Amazon Lex sends to the user when the intent is
complete.

_Required_: No

_Type_: [ResponseSpecification](aws-properties-lex-bot-responsespecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Conditional`

A list of conditional branches associated with the intent's closing
response. These branches are executed when the `nextStep`
attribute is set to `EvalutateConditional`.

_Required_: No

_Type_: [ConditionalSpecification](aws-properties-lex-bot-conditionalspecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IsActive`

Specifies whether an intent's closing response is used. When this
field is false, the closing response isn't sent to the user. If the
`IsActive` field isn't specified, the default is
true.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NextStep`

Specifies the next step that the bot executes after playing the
intent's closing response.

_Required_: No

_Type_: [DialogState](aws-properties-lex-bot-dialogstate.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Intent

IntentConfirmationSetting

All content copied from https://docs.aws.amazon.com/.
