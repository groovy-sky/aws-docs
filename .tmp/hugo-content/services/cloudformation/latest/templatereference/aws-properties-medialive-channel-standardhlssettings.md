This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel StandardHlsSettings

The configuration of an HLS output that is a standard output (not an audio-only output).

The parent of this entity is HlsSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AudioRenditionSets" : String,
  "M3u8Settings" : M3u8Settings
}

```

### YAML

```yaml

  AudioRenditionSets: String
  M3u8Settings:
    M3u8Settings

```

## Properties

`AudioRenditionSets`

Lists all the audio groups that are used with the video output stream. This inputs all the audio GROUP-IDs that
are associated with the video, separated by a comma (,).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`M3u8Settings`

Settings for the M3U8 container.

_Required_: No

_Type_: [M3u8Settings](aws-properties-medialive-channel-m3u8settings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SrtOutputSettings

StaticKeySettings

All content copied from https://docs.aws.amazon.com/.
