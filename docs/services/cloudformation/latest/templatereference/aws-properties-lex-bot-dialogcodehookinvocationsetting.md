This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot DialogCodeHookInvocationSetting

Settings that specify the dialog code hook that is
called by Amazon Lex at a step of the conversation.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EnableCodeHookInvocation" : Boolean,
  "InvocationLabel" : String,
  "IsActive" : Boolean,
  "PostCodeHookSpecification" : PostDialogCodeHookInvocationSpecification
}

```

### YAML

```yaml

  EnableCodeHookInvocation: Boolean
  InvocationLabel: String
  IsActive: Boolean
  PostCodeHookSpecification:
    PostDialogCodeHookInvocationSpecification

```

## Properties

`EnableCodeHookInvocation`

Indicates whether a Lambda function should be invoked
for the dialog.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InvocationLabel`

A label that indicates the dialog step from which the dialog code
hook is happening.

_Required_: No

_Type_: String

_Pattern_: `^([0-9a-zA-Z][_-]?)+$`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IsActive`

Determines whether a dialog code hook is used when the intent is
activated.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PostCodeHookSpecification`

Contains the responses and actions that Amazon Lex takes
after the Lambda function is complete.

_Required_: Yes

_Type_: [PostDialogCodeHookInvocationSpecification](aws-properties-lex-bot-postdialogcodehookinvocationspecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DialogAction

DialogCodeHookSetting

All content copied from https://docs.aws.amazon.com/.
