This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel OutputLocationRef

A reference to an OutputDestination ID that is defined in the channel.

This entity is used by ArchiveGroupSettings, FrameCaptureGroupSettings, HlsGroupSettings,
MediaPackageGroupSettings, MSSmoothGroupSettings, RtmpOutputSettings, and UdpOutputSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DestinationRefId" : String
}

```

### YAML

```yaml

  DestinationRefId: String

```

## Properties

`DestinationRefId`

A reference ID for this destination.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OutputGroupSettings

OutputLockingSettings

All content copied from https://docs.aws.amazon.com/.
