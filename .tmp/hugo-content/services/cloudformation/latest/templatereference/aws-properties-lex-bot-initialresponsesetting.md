This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot InitialResponseSetting

Configuration setting for a response sent to the user before Amazon Lex starts eliciting slots.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CodeHook" : DialogCodeHookInvocationSetting,
  "Conditional" : ConditionalSpecification,
  "InitialResponse" : ResponseSpecification,
  "NextStep" : DialogState
}

```

### YAML

```yaml

  CodeHook:
    DialogCodeHookInvocationSetting
  Conditional:
    ConditionalSpecification
  InitialResponse:
    ResponseSpecification
  NextStep:
    DialogState

```

## Properties

`CodeHook`

Settings that specify the dialog code hook that is
called by Amazon Lex at a step of the conversation.

_Required_: No

_Type_: [DialogCodeHookInvocationSetting](aws-properties-lex-bot-dialogcodehookinvocationsetting.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Conditional`

Provides a list of conditional branches. Branches are evaluated in
the order that they are entered in the list. The first branch with a
condition that evaluates to true is executed. The last branch in the
list is the default branch. The default branch should not have any condition
expression. The default branch is executed if no other branch has a
matching condition.

_Required_: No

_Type_: [ConditionalSpecification](aws-properties-lex-bot-conditionalspecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InitialResponse`

Specifies a list of message groups that Amazon Lex uses to respond the
user input.

_Required_: No

_Type_: [ResponseSpecification](aws-properties-lex-bot-responsespecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NextStep`

The next step in the conversation.

_Required_: No

_Type_: [DialogState](aws-properties-lex-bot-dialogstate.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ImageResponseCard

InputContext

All content copied from https://docs.aws.amazon.com/.
