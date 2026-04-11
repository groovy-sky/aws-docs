This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel FrameCaptureSettings

The frame capture settings.

The parent of this entity is VideoCodecSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CaptureInterval" : Integer,
  "CaptureIntervalUnits" : String,
  "TimecodeBurninSettings" : TimecodeBurninSettings
}

```

### YAML

```yaml

  CaptureInterval: Integer
  CaptureIntervalUnits: String
  TimecodeBurninSettings:
    TimecodeBurninSettings

```

## Properties

`CaptureInterval`

The frequency, in seconds, for capturing frames for inclusion in the output. For example, "10" means capture a
frame every 10 seconds.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CaptureIntervalUnits`

Unit for the frame capture interval.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimecodeBurninSettings`

Property description not available.

_Required_: No

_Type_: [TimecodeBurninSettings](aws-properties-medialive-channel-timecodeburninsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FrameCaptureS3Settings

GlobalConfiguration

All content copied from https://docs.aws.amazon.com/.
