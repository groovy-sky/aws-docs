This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot VoiceSettings

Defines settings for using an Amazon Polly voice to communicate with a
user.

Valid values include:

- `standard`

- `neural`

- `long-form`

- `generative`

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Engine" : String,
  "VoiceId" : String
}

```

### YAML

```yaml

  Engine: String
  VoiceId: String

```

## Properties

`Engine`

Indicates the type of Amazon Polly voice that Amazon Lex should use for voice interaction with the user. For more information, see the [`engine` parameter of the `SynthesizeSpeech` operation](../../../polly/latest/dg/api-synthesizespeech.md#polly-SynthesizeSpeech-request-Engine) in the _Amazon Polly developer guide_.

If you do not specify a value, the default is `standard`.

_Required_: No

_Type_: String

_Allowed values_: `standard | neural | long-form | generative`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VoiceId`

The identifier of the Amazon Polly voice to use.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UnifiedSpeechSettings

WaitAndContinueSpecification

All content copied from https://docs.aws.amazon.com/.
