This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot BotLocale

Provides configuration information for a locale.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CustomVocabulary" : CustomVocabulary,
  "Description" : String,
  "GenerativeAISettings" : GenerativeAISettings,
  "Intents" : [ Intent, ... ],
  "LocaleId" : String,
  "NluConfidenceThreshold" : Number,
  "SlotTypes" : [ SlotType, ... ],
  "SpeechDetectionSensitivity" : String,
  "SpeechRecognitionSettings" : SpeechRecognitionSettings,
  "UnifiedSpeechSettings" : UnifiedSpeechSettings,
  "VoiceSettings" : VoiceSettings
}

```

### YAML

```yaml

  CustomVocabulary:
    CustomVocabulary
  Description: String
  GenerativeAISettings:
    GenerativeAISettings
  Intents:
    - Intent
  LocaleId: String
  NluConfidenceThreshold: Number
  SlotTypes:
    - SlotType
  SpeechDetectionSensitivity: String
  SpeechRecognitionSettings:
    SpeechRecognitionSettings
  UnifiedSpeechSettings:
    UnifiedSpeechSettings
  VoiceSettings:
    VoiceSettings

```

## Properties

`CustomVocabulary`

Specifies a custom vocabulary to use with a specific locale.

_Required_: No

_Type_: [CustomVocabulary](aws-properties-lex-bot-customvocabulary.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description of the bot locale. Use this to help identify the bot
locale in lists.

_Required_: No

_Type_: String

_Maximum_: `2000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GenerativeAISettings`

Property description not available.

_Required_: No

_Type_: [GenerativeAISettings](aws-properties-lex-bot-generativeaisettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Intents`

One or more intents defined for the locale.

_Required_: No

_Type_: Array of [Intent](aws-properties-lex-bot-intent.md)

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LocaleId`

The identifier of the language and locale that the bot will be used
in. The string must match one of the supported locales.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NluConfidenceThreshold`

Determines the threshold where Amazon Lex will insert the
`AMAZON.FallbackIntent`, `AMAZON.KendraSearchIntent`,
or both when returning alternative intents. You must configure an
`AMAZON.FallbackIntent`. `AMAZON.KendraSearchIntent` is
only inserted if it is configured for the bot.

_Required_: Yes

_Type_: Number

_Minimum_: `0`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SlotTypes`

One or more slot types defined for the locale.

_Required_: No

_Type_: Array of [SlotType](aws-properties-lex-bot-slottype.md)

_Maximum_: `250`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SpeechDetectionSensitivity`

Property description not available.

_Required_: No

_Type_: String

_Allowed values_: `Default | HighNoiseTolerance | MaximumNoiseTolerance`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SpeechRecognitionSettings`

Property description not available.

_Required_: No

_Type_: [SpeechRecognitionSettings](aws-properties-lex-bot-speechrecognitionsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UnifiedSpeechSettings`

Property description not available.

_Required_: No

_Type_: [UnifiedSpeechSettings](aws-properties-lex-bot-unifiedspeechsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VoiceSettings`

Defines settings for using an Amazon Polly voice to communicate with a
user.

Valid values include:

- `standard`

- `neural`

- `long-form`

- `generative`

_Required_: No

_Type_: [VoiceSettings](aws-properties-lex-bot-voicesettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BotAliasLocaleSettingsItem

BuildtimeSettings

All content copied from https://docs.aws.amazon.com/.
