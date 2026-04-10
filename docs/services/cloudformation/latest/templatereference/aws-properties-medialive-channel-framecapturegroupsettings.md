This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel FrameCaptureGroupSettings

The settings for a frame capture output group.

The parent of this entity is OutputGroupSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Destination" : OutputLocationRef,
  "FrameCaptureCdnSettings" : FrameCaptureCdnSettings
}

```

### YAML

```yaml

  Destination:
    OutputLocationRef
  FrameCaptureCdnSettings:
    FrameCaptureCdnSettings

```

## Properties

`Destination`

The destination for the frame capture files. The destination is either the URI for an Amazon S3 bucket and
object, plus a file name prefix (for example, s3ssl://sportsDelivery/highlights/20180820/curling\_) or the URI for a
MediaStore container, plus a file name prefix (for example, mediastoressl://sportsDelivery/20180820/curling\_). The
final file names consist of the prefix from the destination field (for example, "curling\_") + name modifier + the
counter (5 digits, starting from 00001) + extension (which is always .jpg). For example, curlingLow.00001.jpg.

_Required_: No

_Type_: [OutputLocationRef](aws-properties-medialive-channel-outputlocationref.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FrameCaptureCdnSettings`

Settings to configure the destination of a Frame Capture output.

_Required_: No

_Type_: [FrameCaptureCdnSettings](aws-properties-medialive-channel-framecapturecdnsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FrameCaptureCdnSettings

FrameCaptureOutputSettings

All content copied from https://docs.aws.amazon.com/.
