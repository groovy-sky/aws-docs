This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot ElicitationCodeHookInvocationSetting

Settings that specify the dialog code hook that is called by Amazon Lex between eliciting slot values.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EnableCodeHookInvocation" : Boolean,
  "InvocationLabel" : String
}

```

### YAML

```yaml

  EnableCodeHookInvocation: Boolean
  InvocationLabel: String

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

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DTMFSpecification

ErrorLogSettings

All content copied from https://docs.aws.amazon.com/.
