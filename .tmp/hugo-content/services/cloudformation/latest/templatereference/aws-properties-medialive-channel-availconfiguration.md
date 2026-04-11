This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel AvailConfiguration

The setup of ad avail handling in the output.

The parent of this entity is EncoderSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AvailSettings" : AvailSettings,
  "Scte35SegmentationScope" : String
}

```

### YAML

```yaml

  AvailSettings:
    AvailSettings
  Scte35SegmentationScope: String

```

## Properties

`AvailSettings`

The setup of ad avail handling in the output.

_Required_: No

_Type_: [AvailSettings](aws-properties-medialive-channel-availsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Scte35SegmentationScope`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AvailBlanking

AvailSettings

All content copied from https://docs.aws.amazon.com/.
