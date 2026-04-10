This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel FrameCaptureCdnSettings

Settings to configure the destination of a Frame Capture output.

The parent of this entity is FrameCaptureGroupSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FrameCaptureS3Settings" : FrameCaptureS3Settings
}

```

### YAML

```yaml

  FrameCaptureS3Settings:
    FrameCaptureS3Settings

```

## Properties

`FrameCaptureS3Settings`

Sets up Amazon S3 as the destination for this Frame Capture output.

_Required_: No

_Type_: [FrameCaptureS3Settings](aws-properties-medialive-channel-framecaptures3settings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FollowerChannelSettings

FrameCaptureGroupSettings

All content copied from https://docs.aws.amazon.com/.
