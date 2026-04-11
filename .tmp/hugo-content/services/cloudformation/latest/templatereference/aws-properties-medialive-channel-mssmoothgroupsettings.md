This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel MsSmoothGroupSettings

The settings for a Microsoft Smooth output group.

The parent of this entity is OutputGroupSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AcquisitionPointId" : String,
  "AudioOnlyTimecodeControl" : String,
  "CertificateMode" : String,
  "ConnectionRetryInterval" : Integer,
  "Destination" : OutputLocationRef,
  "EventId" : String,
  "EventIdMode" : String,
  "EventStopBehavior" : String,
  "FilecacheDuration" : Integer,
  "FragmentLength" : Integer,
  "InputLossAction" : String,
  "NumRetries" : Integer,
  "RestartDelay" : Integer,
  "SegmentationMode" : String,
  "SendDelayMs" : Integer,
  "SparseTrackType" : String,
  "StreamManifestBehavior" : String,
  "TimestampOffset" : String,
  "TimestampOffsetMode" : String
}

```

### YAML

```yaml

  AcquisitionPointId: String
  AudioOnlyTimecodeControl: String
  CertificateMode: String
  ConnectionRetryInterval: Integer
  Destination:
    OutputLocationRef
  EventId: String
  EventIdMode: String
  EventStopBehavior: String
  FilecacheDuration: Integer
  FragmentLength: Integer
  InputLossAction: String
  NumRetries: Integer
  RestartDelay: Integer
  SegmentationMode: String
  SendDelayMs: Integer
  SparseTrackType: String
  StreamManifestBehavior: String
  TimestampOffset: String
  TimestampOffsetMode: String

```

## Properties

`AcquisitionPointId`

The value of the Acquisition Point Identity element that is used in each message placed in the sparse track.
Enabled only if sparseTrackType is not "none."

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AudioOnlyTimecodeControl`

If set to passthrough for an audio-only Microsoft Smooth output, the fragment absolute time is set to the
current timecode. This option does not write timecodes to the audio elementary stream.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CertificateMode`

If set to verifyAuthenticity, verifies the HTTPS certificate chain to a trusted certificate authority (CA). This
causes HTTPS outputs to self-signed certificates to fail.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConnectionRetryInterval`

The number of seconds to wait before retrying the connection to the IIS server if the connection is lost.
Content is cached during this time, and the cache is delivered to the IIS server after the connection is
re-established.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Destination`

The Smooth Streaming publish point on an IIS server. MediaLive acts as a "Push" encoder to IIS.

_Required_: No

_Type_: [OutputLocationRef](aws-properties-medialive-channel-outputlocationref.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EventId`

The Microsoft Smooth channel ID that is sent to the IIS server. Specify the ID only if eventIdMode is set to
useConfigured.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EventIdMode`

Specifies whether to send a channel ID to the IIS server. If no channel ID is sent and the same channel is used
without changing the publishing point, clients might see cached video from the previous run. Options: -
"useConfigured" - use the value provided in eventId - "useTimestamp" - generate and send a channel ID based on the
current timestamp - "noEventId" - do not send a channel ID to the IIS server.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EventStopBehavior`

When set to sendEos, sends an EOS signal to an IIS server when stopping the channel.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FilecacheDuration`

The size, in seconds, of the file cache for streaming outputs.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FragmentLength`

The length, in seconds, of mp4 fragments to generate. The fragment length must be compatible with GOP size and
frame rate.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InputLossAction`

A parameter that controls output group behavior on an input loss.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NumRetries`

The number of retry attempts.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RestartDelay`

The number of seconds before initiating a restart due to output failure, due to exhausting the numRetries on one
segment, or exceeding filecacheDuration.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SegmentationMode`

useInputSegmentation has been deprecated. The configured segment size is always used.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SendDelayMs`

The number of milliseconds to delay the output from the second pipeline.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SparseTrackType`

If set to scte35, uses incoming SCTE-35 messages to generate a sparse track in this group of Microsoft Smooth
outputs.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StreamManifestBehavior`

When set to send, sends a stream manifest so that the publishing point doesn't start until all streams
start.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimestampOffset`

The timestamp offset for the channel. Used only if timestampOffsetMode is set to useConfiguredOffset.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimestampOffsetMode`

The type of timestamp date offset to use. - useEventStartDate: Use the date the channel was started as the
offset - useConfiguredOffset: Use an explicitly configured date as the offset.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Mpeg2Settings

MsSmoothOutputSettings

All content copied from https://docs.aws.amazon.com/.
