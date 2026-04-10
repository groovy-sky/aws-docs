This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot NluImprovementSpecification

Configures the Assisted Natural Language Understanding (NLU) feature for your bot. This specification determines whether enhanced intent recognition and utterance understanding capabilities are active.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AssistedNluMode" : String,
  "Enabled" : Boolean,
  "IntentDisambiguationSettings" : IntentDisambiguationSettings
}

```

### YAML

```yaml

  AssistedNluMode: String
  Enabled: Boolean
  IntentDisambiguationSettings:
    IntentDisambiguationSettings

```

## Properties

`AssistedNluMode`

Specifies the mode for Assisted NLU operation. Use `Primary` to make Assisted NLU the primary intent recognition method, or `Fallback` to use it only when standard NLU confidence is low.

_Required_: No

_Type_: String

_Allowed values_: `Primary | Fallback`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Enabled`

Determines whether the Assisted NLU feature is enabled for the bot. When set to `true`, Amazon Lex uses advanced models to improve intent recognition and slot resolution, with the default being `false`.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IntentDisambiguationSettings`

An object containing specifications for the Intent Disambiguation feature within the Assisted NLU settings. These settings determine how the bot handles ambiguous user inputs that could match multiple intents.

_Required_: No

_Type_: [IntentDisambiguationSettings](aws-properties-lex-bot-intentdisambiguationsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MultipleValuesSetting

ObfuscationSetting

All content copied from https://docs.aws.amazon.com/.
