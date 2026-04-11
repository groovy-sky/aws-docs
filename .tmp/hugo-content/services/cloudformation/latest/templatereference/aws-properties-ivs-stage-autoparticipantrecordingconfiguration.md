This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IVS::Stage AutoParticipantRecordingConfiguration

The `AWS::IVS::AutoParticipantRecordingConfiguration` property type describes a configuration for individual participant recording.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "HlsConfiguration" : HlsConfiguration,
  "MediaTypes" : [ String, ... ],
  "RecordingReconnectWindowSeconds" : Integer,
  "StorageConfigurationArn" : String,
  "ThumbnailConfiguration" : ThumbnailConfiguration
}

```

### YAML

```yaml

  HlsConfiguration:
    HlsConfiguration
  MediaTypes:
    - String
  RecordingReconnectWindowSeconds: Integer
  StorageConfigurationArn: String
  ThumbnailConfiguration:
    ThumbnailConfiguration

```

## Properties

`HlsConfiguration`

HLS configuration object for individual participant recording.

_Required_: No

_Type_: [HlsConfiguration](aws-properties-ivs-stage-hlsconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MediaTypes`

Types of media to be recorded. Default: `AUDIO_VIDEO`.

_Required_: No

_Type_: Array of String

_Allowed values_: `AUDIO_VIDEO | AUDIO_ONLY`

_Minimum_: `0`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RecordingReconnectWindowSeconds`

If a stage publisher disconnects and then reconnects within the specified interval, the multiple recordings will be considered a single recording and merged together.

The default value is 0, which disables merging.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `300`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StorageConfigurationArn`

ARN of the [StorageConfiguration](aws-resource-ivs-storageconfiguration.md)
resource to use for individual participant recording.
Default: "" (empty string, no storage configuration is specified).
Individual participant recording cannot be started unless a storage configuration is specified,
when a [Stage](aws-resource-ivs-stage.md) is created or updated.

_Required_: Yes

_Type_: String

_Pattern_: `^$|^arn:aws:ivs:[a-z0-9-]+:[0-9]+:storage-configuration/[a-zA-Z0-9-]+$`

_Minimum_: `0`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ThumbnailConfiguration`

A complex type that allows you to enable/disable the recording of thumbnails for individual participant recording and modify the interval at which thumbnails are
generated for the live session.

_Required_: No

_Type_: [ThumbnailConfiguration](aws-properties-ivs-stage-thumbnailconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::IVS::Stage

HlsConfiguration

All content copied from https://docs.aws.amazon.com/.
