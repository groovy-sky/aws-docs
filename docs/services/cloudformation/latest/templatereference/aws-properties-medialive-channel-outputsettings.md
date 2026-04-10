This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel OutputSettings

The output settings.

The parent of this entity is Output.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ArchiveOutputSettings" : ArchiveOutputSettings,
  "CmafIngestOutputSettings" : CmafIngestOutputSettings,
  "FrameCaptureOutputSettings" : FrameCaptureOutputSettings,
  "HlsOutputSettings" : HlsOutputSettings,
  "MediaPackageOutputSettings" : MediaPackageOutputSettings,
  "MsSmoothOutputSettings" : MsSmoothOutputSettings,
  "MultiplexOutputSettings" : MultiplexOutputSettings,
  "RtmpOutputSettings" : RtmpOutputSettings,
  "SrtOutputSettings" : SrtOutputSettings,
  "UdpOutputSettings" : UdpOutputSettings
}

```

### YAML

```yaml

  ArchiveOutputSettings:
    ArchiveOutputSettings
  CmafIngestOutputSettings:
    CmafIngestOutputSettings
  FrameCaptureOutputSettings:
    FrameCaptureOutputSettings
  HlsOutputSettings:
    HlsOutputSettings
  MediaPackageOutputSettings:
    MediaPackageOutputSettings
  MsSmoothOutputSettings:
    MsSmoothOutputSettings
  MultiplexOutputSettings:
    MultiplexOutputSettings
  RtmpOutputSettings:
    RtmpOutputSettings
  SrtOutputSettings:
    SrtOutputSettings
  UdpOutputSettings:
    UdpOutputSettings

```

## Properties

`ArchiveOutputSettings`

The settings for an archive output.

_Required_: No

_Type_: [ArchiveOutputSettings](aws-properties-medialive-channel-archiveoutputsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CmafIngestOutputSettings`

Property description not available.

_Required_: No

_Type_: [CmafIngestOutputSettings](aws-properties-medialive-channel-cmafingestoutputsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FrameCaptureOutputSettings`

The settings for a frame capture output.

The parent of this entity is OutputGroupSettings.

_Required_: No

_Type_: [FrameCaptureOutputSettings](aws-properties-medialive-channel-framecaptureoutputsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HlsOutputSettings`

The settings for an HLS output.

The parent of this entity is OutputGroupSettings.

_Required_: No

_Type_: [HlsOutputSettings](aws-properties-medialive-channel-hlsoutputsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MediaPackageOutputSettings`

The settings for a MediaPackage output.

The parent of this entity is OutputGroupSettings.

_Required_: No

_Type_: [MediaPackageOutputSettings](aws-properties-medialive-channel-mediapackageoutputsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MsSmoothOutputSettings`

The settings for a Microsoft Smooth output.

_Required_: No

_Type_: [MsSmoothOutputSettings](aws-properties-medialive-channel-mssmoothoutputsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MultiplexOutputSettings`

Configuration of a Multiplex output.

_Required_: No

_Type_: [MultiplexOutputSettings](aws-properties-medialive-channel-multiplexoutputsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RtmpOutputSettings`

The settings for an RTMP output.

The parent of this entity is OutputGroupSettings.

_Required_: No

_Type_: [RtmpOutputSettings](aws-properties-medialive-channel-rtmpoutputsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SrtOutputSettings`

Property description not available.

_Required_: No

_Type_: [SrtOutputSettings](aws-properties-medialive-channel-srtoutputsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UdpOutputSettings`

The settings for a UDP output.

The parent of this entity is OutputGroupSettings.

_Required_: No

_Type_: [UdpOutputSettings](aws-properties-medialive-channel-udpoutputsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OutputLockingSettings

PipelineLockingSettings

All content copied from https://docs.aws.amazon.com/.
