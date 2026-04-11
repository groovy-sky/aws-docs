This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot DialogAction

Defines the action that the bot executes at runtime when the
conversation reaches this step.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SlotToElicit" : String,
  "SuppressNextMessage" : Boolean,
  "Type" : String
}

```

### YAML

```yaml

  SlotToElicit: String
  SuppressNextMessage: Boolean
  Type: String

```

## Properties

`SlotToElicit`

If the dialog action is `ElicitSlot`, defines the slot to
elicit from the user.

_Required_: No

_Type_: String

_Pattern_: `^([0-9a-zA-Z][_-]?)+$`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SuppressNextMessage`

When true the next message for the intent is not used.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The action that the bot should execute.

_Required_: Yes

_Type_: String

_Allowed values_: `CloseIntent | ConfirmIntent | ElicitIntent | ElicitSlot | StartIntent | FulfillIntent | EndConversation | EvaluateConditional | InvokeDialogCodeHook`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DescriptiveBotBuilderSpecification

DialogCodeHookInvocationSetting

All content copied from https://docs.aws.amazon.com/.
