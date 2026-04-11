This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel AudioChannelMapping

The settings for remixing audio.

The parent of this entity is RemixSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InputChannelLevels" : [ InputChannelLevel, ... ],
  "OutputChannel" : Integer
}

```

### YAML

```yaml

  InputChannelLevels:
    - InputChannelLevel
  OutputChannel: Integer

```

## Properties

`InputChannelLevels`

The indices and gain values for each input channel that should be remixed into this output channel.

_Required_: No

_Type_: Array of [InputChannelLevel](aws-properties-medialive-channel-inputchannellevel.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OutputChannel`

The index of the output channel that is being produced.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ArchiveS3Settings

AudioCodecSettings

All content copied from https://docs.aws.amazon.com/.
