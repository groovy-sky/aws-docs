This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel Fmp4HlsSettings

Settings for the fMP4 containers.

The parent of this entity is HlsSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AudioRenditionSets" : String,
  "NielsenId3Behavior" : String,
  "TimedMetadataBehavior" : String
}

```

### YAML

```yaml

  AudioRenditionSets: String
  NielsenId3Behavior: String
  TimedMetadataBehavior: String

```

## Properties

`AudioRenditionSets`

List all the audio groups that are used with the video output stream. Input all the audio GROUP-IDs that are associated to the video, separate by ','.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NielsenId3Behavior`

If set to passthrough, Nielsen inaudible tones for media tracking will be detected in the input audio and an equivalent ID3 tag will be inserted in the output.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimedMetadataBehavior`

When set to passthrough, timed metadata is passed through from input to output.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FecOutputSettings

FollowerChannelSettings

All content copied from https://docs.aws.amazon.com/.
