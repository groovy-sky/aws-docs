This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel AvailBlanking

The configuration of ad avail blanking in the output.

The parent of this entity is EncoderSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AvailBlankingImage" : InputLocation,
  "State" : String
}

```

### YAML

```yaml

  AvailBlankingImage:
    InputLocation
  State: String

```

## Properties

`AvailBlankingImage`

The blanking image to be used. Keep empty for solid black. Only .bmp and .png images are supported.

_Required_: No

_Type_: [InputLocation](aws-properties-medialive-channel-inputlocation.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`State`

When set to enabled, the video, audio, and captions are blanked when insertion metadata is added.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Av1Settings

AvailConfiguration

All content copied from https://docs.aws.amazon.com/.
