This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot LambdaCodeHook

Specifies a Lambda function that verifies requests to a bot or
fulfills the user's request to a bot.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CodeHookInterfaceVersion" : String,
  "LambdaArn" : String
}

```

### YAML

```yaml

  CodeHookInterfaceVersion: String
  LambdaArn: String

```

## Properties

`CodeHookInterfaceVersion`

The version of the request-response that you want Amazon Lex to use to
invoke your Lambda function.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LambdaArn`

The Amazon Resource Name (ARN) of the Lambda function.

_Required_: Yes

_Type_: String

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

KendraConfiguration

Message

All content copied from https://docs.aws.amazon.com/.
