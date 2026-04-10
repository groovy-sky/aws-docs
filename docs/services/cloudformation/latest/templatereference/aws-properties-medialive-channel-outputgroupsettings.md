This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel OutputGroupSettings

The configuration of the output group.

The parent of this entity is OutputGroup.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ArchiveGroupSettings" : ArchiveGroupSettings,
  "CmafIngestGroupSettings" : CmafIngestGroupSettings,
  "FrameCaptureGroupSettings" : FrameCaptureGroupSettings,
  "HlsGroupSettings" : HlsGroupSettings,
  "MediaPackageGroupSettings" : MediaPackageGroupSettings,
  "MsSmoothGroupSettings" : MsSmoothGroupSettings,
  "MultiplexGroupSettings" : Json,
  "RtmpGroupSettings" : RtmpGroupSettings,
  "SrtGroupSettings" : SrtGroupSettings,
  "UdpGroupSettings" : UdpGroupSettings
}

```

### YAML

```yaml

  ArchiveGroupSettings:
    ArchiveGroupSettings
  CmafIngestGroupSettings:
    CmafIngestGroupSettings
  FrameCaptureGroupSettings:
    FrameCaptureGroupSettings
  HlsGroupSettings:
    HlsGroupSettings
  MediaPackageGroupSettings:
    MediaPackageGroupSettings
  MsSmoothGroupSettings:
    MsSmoothGroupSettings
  MultiplexGroupSettings: Json
  RtmpGroupSettings:
    RtmpGroupSettings
  SrtGroupSettings:
    SrtGroupSettings
  UdpGroupSettings:
    UdpGroupSettings

```

## Properties

`ArchiveGroupSettings`

The configuration of an archive output group.

The parent of this entity is OutputGroupSettings.

_Required_: No

_Type_: [ArchiveGroupSettings](aws-properties-medialive-channel-archivegroupsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CmafIngestGroupSettings`

Property description not available.

_Required_: No

_Type_: [CmafIngestGroupSettings](aws-properties-medialive-channel-cmafingestgroupsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FrameCaptureGroupSettings`

The configuration of a frame capture output group.

_Required_: No

_Type_: [FrameCaptureGroupSettings](aws-properties-medialive-channel-framecapturegroupsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HlsGroupSettings`

The configuration of an HLS output group.

_Required_: No

_Type_: [HlsGroupSettings](aws-properties-medialive-channel-hlsgroupsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MediaPackageGroupSettings`

The configuration of a MediaPackage output group.

_Required_: No

_Type_: [MediaPackageGroupSettings](aws-properties-medialive-channel-mediapackagegroupsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MsSmoothGroupSettings`

The configuration of a Microsoft Smooth output group.

_Required_: No

_Type_: [MsSmoothGroupSettings](aws-properties-medialive-channel-mssmoothgroupsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MultiplexGroupSettings`

The settings for a Multiplex output group.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RtmpGroupSettings`

The configuration of an RTMP output group.

_Required_: No

_Type_: [RtmpGroupSettings](aws-properties-medialive-channel-rtmpgroupsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SrtGroupSettings`

Property description not available.

_Required_: No

_Type_: [SrtGroupSettings](aws-properties-medialive-channel-srtgroupsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UdpGroupSettings`

The configuration of a UDP output group.

_Required_: No

_Type_: [UdpGroupSettings](aws-properties-medialive-channel-udpgroupsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OutputGroup

OutputLocationRef

All content copied from https://docs.aws.amazon.com/.
