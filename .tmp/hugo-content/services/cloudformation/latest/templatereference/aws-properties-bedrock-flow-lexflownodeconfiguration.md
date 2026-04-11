This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Flow LexFlowNodeConfiguration

Contains configurations for a Lex node in the flow. You specify a Amazon Lex bot to invoke. This node takes an utterance as the input and returns as the output the intent identified by the Amazon Lex bot. For more information, see [Node types in a flow](../../../bedrock/latest/userguide/flows-nodes.md) in the Amazon Bedrock User Guide.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BotAliasArn" : String,
  "LocaleId" : String
}

```

### YAML

```yaml

  BotAliasArn: String
  LocaleId: String

```

## Properties

`BotAliasArn`

The Amazon Resource Name (ARN) of the Amazon Lex bot alias to invoke.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws(|-us-gov):lex:[a-z]{2}(-gov)?-[a-z]+-\d{1}:\d{12}:bot-alias/[0-9a-zA-Z]+/[0-9a-zA-Z]+$`

_Maximum_: `78`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LocaleId`

The Region to invoke the Amazon Lex bot in.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LambdaFunctionFlowNodeConfiguration

LoopControllerFlowNodeConfiguration

All content copied from https://docs.aws.amazon.com/.
