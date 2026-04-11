This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel UdpOutputSettings

The settings for one UDP output.

The parent of this entity is OutputSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BufferMsec" : Integer,
  "ContainerSettings" : UdpContainerSettings,
  "Destination" : OutputLocationRef,
  "FecOutputSettings" : FecOutputSettings
}

```

### YAML

```yaml

  BufferMsec: Integer
  ContainerSettings:
    UdpContainerSettings
  Destination:
    OutputLocationRef
  FecOutputSettings:
    FecOutputSettings

```

## Properties

`BufferMsec`

The UDP output buffering in milliseconds. Larger values increase latency through the transcoder but
simultaneously assist the transcoder in maintaining a constant, low-jitter UDP/RTP output while accommodating clock
recovery, input switching, input disruptions, picture reordering, and so on.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ContainerSettings`

The settings for the UDP output.

_Required_: No

_Type_: [UdpContainerSettings](aws-properties-medialive-channel-udpcontainersettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Destination`

The destination address and port number for RTP or UDP packets. These can be unicast or multicast RTP or UDP
(for example, rtp://239.10.10.10:5001 or udp://10.100.100.100:5002).

_Required_: No

_Type_: [OutputLocationRef](aws-properties-medialive-channel-outputlocationref.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FecOutputSettings`

The settings for enabling and adjusting Forward Error Correction on UDP outputs.

_Required_: No

_Type_: [FecOutputSettings](aws-properties-medialive-channel-fecoutputsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UdpGroupSettings

VideoBlackFailoverSettings

All content copied from https://docs.aws.amazon.com/.
