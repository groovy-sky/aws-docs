This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot DeepgramSpeechModelConfig

Configuration settings for integrating Deepgram speech-to-text models with Amazon Lex.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ApiTokenSecretArn" : String,
  "ModelId" : String
}

```

### YAML

```yaml

  ApiTokenSecretArn: String
  ModelId: String

```

## Properties

`ApiTokenSecretArn`

The Amazon Resource Name (ARN) of the Secrets Manager secret that contains the Deepgram API token.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[A-Za-z-]*:secretsmanager:[a-z0-9-]{1,20}:[0-9]{12}:secret:[A-Za-z0-9/_+=.@-]{1,512}-[A-Za-z0-9]{6}$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ModelId`

The identifier of the Deepgram speech-to-text model to use for processing speech input.

_Required_: No

_Type_: String

_Pattern_: `[A-Za-z0-9-_]+`

_Minimum_: `4`

_Maximum_: `32`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataSourceConfiguration

DefaultConditionalBranch

All content copied from https://docs.aws.amazon.com/.
