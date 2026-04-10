This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel AudioLanguageSelection

Information about the audio language to extract.

The parent of this entity is AudioSelectorSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LanguageCode" : String,
  "LanguageSelectionPolicy" : String
}

```

### YAML

```yaml

  LanguageCode: String
  LanguageSelectionPolicy: String

```

## Properties

`LanguageCode`

Selects a specific three-letter language code from within an audio source.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LanguageSelectionPolicy`

When set to "strict," the transport stream demux strictly identifies audio streams by their language descriptor.
If a PMT update occurs such that an audio stream matching the initially selected language is no longer present, then
mute is encoded until the language returns. If set to "loose," then on a PMT update the demux chooses another audio
stream in the program with the same stream type if it can't find one with the same language.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AudioHlsRenditionSelection

AudioNormalizationSettings

All content copied from https://docs.aws.amazon.com/.
