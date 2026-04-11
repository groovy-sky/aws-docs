This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot CodeHookSpecification

Contains information about code hooks that Amazon Lex calls during a
conversation.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LambdaCodeHook" : LambdaCodeHook
}

```

### YAML

```yaml

  LambdaCodeHook:
    LambdaCodeHook

```

## Properties

`LambdaCodeHook`

Specifies a Lambda function that verifies requests to a bot or
fulfills the user's request to a bot.

_Required_: Yes

_Type_: [LambdaCodeHook](aws-properties-lex-bot-lambdacodehook.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CloudWatchLogGroupLogDestination

CompositeSlotTypeSetting

All content copied from https://docs.aws.amazon.com/.
