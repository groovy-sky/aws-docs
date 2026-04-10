This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot WaitAndContinueSpecification

Specifies the prompts that Amazon Lex uses while a bot is waiting for
customer input.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ContinueResponse" : ResponseSpecification,
  "IsActive" : Boolean,
  "StillWaitingResponse" : StillWaitingResponseSpecification,
  "WaitingResponse" : ResponseSpecification
}

```

### YAML

```yaml

  ContinueResponse:
    ResponseSpecification
  IsActive: Boolean
  StillWaitingResponse:
    StillWaitingResponseSpecification
  WaitingResponse:
    ResponseSpecification

```

## Properties

`ContinueResponse`

The response that Amazon Lex sends to indicate that the bot is ready to
continue the conversation.

_Required_: Yes

_Type_: [ResponseSpecification](aws-properties-lex-bot-responsespecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IsActive`

Specifies whether the bot will wait for a user to respond. When this
field is false, wait and continue responses for a slot aren't used. If
the `IsActive` field isn't specified, the default is
true.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StillWaitingResponse`

A response that Amazon Lex sends periodically to the user to indicate
that the bot is still waiting for input from the user.

_Required_: No

_Type_: [StillWaitingResponseSpecification](aws-properties-lex-bot-stillwaitingresponsespecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WaitingResponse`

The response that Amazon Lex sends to indicate that the bot is waiting
for the conversation to continue.

_Required_: Yes

_Type_: [ResponseSpecification](aws-properties-lex-bot-responsespecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VoiceSettings

AWS::Lex::BotAlias

All content copied from https://docs.aws.amazon.com/.
