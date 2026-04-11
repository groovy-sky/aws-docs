This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot DefaultConditionalBranch

A set of actions that Amazon Lex should run if none of the
other conditions are met.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "NextStep" : DialogState,
  "Response" : ResponseSpecification
}

```

### YAML

```yaml

  NextStep:
    DialogState
  Response:
    ResponseSpecification

```

## Properties

`NextStep`

The next step in the conversation.

_Required_: No

_Type_: [DialogState](aws-properties-lex-bot-dialogstate.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Response`

Specifies a list of message groups that Amazon Lex uses to respond the
user input.

_Required_: No

_Type_: [ResponseSpecification](aws-properties-lex-bot-responsespecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeepgramSpeechModelConfig

DescriptiveBotBuilderSpecification

All content copied from https://docs.aws.amazon.com/.
