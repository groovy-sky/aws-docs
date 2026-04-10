This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot DialogState

The current state of the conversation with the user.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DialogAction" : DialogAction,
  "Intent" : IntentOverride,
  "SessionAttributes" : [ SessionAttribute, ... ]
}

```

### YAML

```yaml

  DialogAction:
    DialogAction
  Intent:
    IntentOverride
  SessionAttributes:
    - SessionAttribute

```

## Properties

`DialogAction`

Defines the action that the bot executes at runtime when the
conversation reaches this step.

_Required_: No

_Type_: [DialogAction](aws-properties-lex-bot-dialogaction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Intent`

Override settings to configure the intent state.

_Required_: No

_Type_: [IntentOverride](aws-properties-lex-bot-intentoverride.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SessionAttributes`

Map of key/value pairs representing session-specific context
information. It contains application information passed between Amazon Lex and a client application.

_Required_: No

_Type_: Array of [SessionAttribute](aws-properties-lex-bot-sessionattribute.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DialogCodeHookSetting

DTMFSpecification

All content copied from https://docs.aws.amazon.com/.
