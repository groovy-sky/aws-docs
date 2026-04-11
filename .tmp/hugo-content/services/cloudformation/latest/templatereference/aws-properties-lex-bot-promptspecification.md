This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot PromptSpecification

Specifies a list of message groups that Amazon Lex sends to a user to
elicit a response.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AllowInterrupt" : Boolean,
  "MaxRetries" : Integer,
  "MessageGroupsList" : [ MessageGroup, ... ],
  "MessageSelectionStrategy" : String,
  "PromptAttemptsSpecification" : {Key: Value, ...}
}

```

### YAML

```yaml

  AllowInterrupt: Boolean
  MaxRetries: Integer
  MessageGroupsList:
    - MessageGroup
  MessageSelectionStrategy: String
  PromptAttemptsSpecification:
    Key: Value

```

## Properties

`AllowInterrupt`

Indicates whether the user can interrupt a speech prompt from the
bot.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxRetries`

The maximum number of times the bot tries to elicit a response from
the user using this prompt.

_Required_: Yes

_Type_: Integer

_Minimum_: `0`

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MessageGroupsList`

A collection of messages that Amazon Lex can send to the user. Amazon Lex
chooses the actual message to send at runtime.

_Required_: Yes

_Type_: Array of [MessageGroup](aws-properties-lex-bot-messagegroup.md)

_Minimum_: `1`

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MessageSelectionStrategy`

Indicates how a message is selected from a message group among retries.

_Required_: No

_Type_: String

_Allowed values_: `Random | Ordered`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PromptAttemptsSpecification`

Specifies the advanced settings on each attempt of the prompt.

_Required_: No

_Type_: Object of [PromptAttemptSpecification](aws-properties-lex-bot-promptattemptspecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PromptAttemptSpecification

QInConnectAssistantConfiguration

All content copied from https://docs.aws.amazon.com/.
