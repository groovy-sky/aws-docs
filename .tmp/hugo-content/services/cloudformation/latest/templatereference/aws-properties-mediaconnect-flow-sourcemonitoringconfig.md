This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::Flow SourceMonitoringConfig

The `SourceMonitoringConfig` property type specifies the source monitoring settings for an `AWS::MediaConnect::Flow`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AudioMonitoringSettings" : [ AudioMonitoringSetting, ... ],
  "ContentQualityAnalysisState" : String,
  "ThumbnailState" : String,
  "VideoMonitoringSettings" : [ VideoMonitoringSetting, ... ]
}

```

### YAML

```yaml

  AudioMonitoringSettings:
    - AudioMonitoringSetting
  ContentQualityAnalysisState: String
  ThumbnailState: String
  VideoMonitoringSettings:
    - VideoMonitoringSetting

```

## Properties

`AudioMonitoringSettings`

Contains the settings for audio stream metrics monitoring.

_Required_: No

_Type_: Array of [AudioMonitoringSetting](aws-properties-mediaconnect-flow-audiomonitoringsetting.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ContentQualityAnalysisState`

Indicates whether content quality analysis is enabled or disabled.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ThumbnailState`

The current state of the thumbnail monitoring.

- If you don't explicitly specify a value when creating a flow, no thumbnail state will be set.

- If you update an existing flow and remove a previously set thumbnail state, the value will change to `DISABLED`.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VideoMonitoringSettings`

Contains the settings for video stream metrics monitoring.

_Required_: No

_Type_: Array of [VideoMonitoringSetting](aws-properties-mediaconnect-flow-videomonitoringsetting.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Source

SourcePriority

All content copied from https://docs.aws.amazon.com/.
