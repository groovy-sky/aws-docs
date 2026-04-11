This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel H265FilterSettings

Settings to configure video filters that apply to the H265 codec.

The parent of this entity is H265Settings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BandwidthReductionFilterSettings" : BandwidthReductionFilterSettings,
  "TemporalFilterSettings" : TemporalFilterSettings
}

```

### YAML

```yaml

  BandwidthReductionFilterSettings:
    BandwidthReductionFilterSettings
  TemporalFilterSettings:
    TemporalFilterSettings

```

## Properties

`BandwidthReductionFilterSettings`

Property description not available.

_Required_: No

_Type_: [BandwidthReductionFilterSettings](aws-properties-medialive-channel-bandwidthreductionfiltersettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TemporalFilterSettings`

Settings for applying the temporal filter to the video.

_Required_: No

_Type_: [TemporalFilterSettings](aws-properties-medialive-channel-temporalfiltersettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

H265ColorSpaceSettings

H265Settings

All content copied from https://docs.aws.amazon.com/.
