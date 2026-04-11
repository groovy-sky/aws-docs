This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel M3u8Settings

Settings for the M3U8 container.

The parent of this entity is StandardHlsSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AudioFramesPerPes" : Integer,
  "AudioPids" : String,
  "EcmPid" : String,
  "KlvBehavior" : String,
  "KlvDataPids" : String,
  "NielsenId3Behavior" : String,
  "PatInterval" : Integer,
  "PcrControl" : String,
  "PcrPeriod" : Integer,
  "PcrPid" : String,
  "PmtInterval" : Integer,
  "PmtPid" : String,
  "ProgramNum" : Integer,
  "Scte35Behavior" : String,
  "Scte35Pid" : String,
  "TimedMetadataBehavior" : String,
  "TimedMetadataPid" : String,
  "TransportStreamId" : Integer,
  "VideoPid" : String
}

```

### YAML

```yaml

  AudioFramesPerPes: Integer
  AudioPids: String
  EcmPid: String
  KlvBehavior: String
  KlvDataPids: String
  NielsenId3Behavior: String
  PatInterval: Integer
  PcrControl: String
  PcrPeriod: Integer
  PcrPid: String
  PmtInterval: Integer
  PmtPid: String
  ProgramNum: Integer
  Scte35Behavior: String
  Scte35Pid: String
  TimedMetadataBehavior: String
  TimedMetadataPid: String
  TransportStreamId: Integer
  VideoPid: String

```

## Properties

`AudioFramesPerPes`

The number of audio frames to insert for each PES packet.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AudioPids`

The PID of the elementary audio streams in the transport stream. Multiple values are accepted, and can be
entered in ranges or by comma separation. You can enter the value as a decimal or hexadecimal value.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EcmPid`

This parameter is unused and deprecated.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KlvBehavior`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KlvDataPids`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NielsenId3Behavior`

If set to passthrough, Nielsen inaudible tones for media tracking will be detected in the input audio and an equivalent ID3 tag will be inserted in the output.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PatInterval`

The number of milliseconds between instances of this table in the output transport stream. A value of \\"0\\"
writes out the PMT once per segment file.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PcrControl`

When set to pcrEveryPesPacket, a Program Clock Reference value is inserted for every Packetized Elementary
Stream (PES) header. This parameter is effective only when the PCR PID is the same as the video or audio elementary
stream.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PcrPeriod`

The maximum time, in milliseconds, between Program Clock References (PCRs) inserted into the transport
stream.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PcrPid`

The PID of the Program Clock Reference (PCR) in the transport stream. When no value is given, MediaLive assigns
the same value as the video PID. You can enter the value as a decimal or hexadecimal value.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PmtInterval`

The number of milliseconds between instances of this table in the output transport stream. A value of \\"0\\"
writes out the PMT once per segment file.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PmtPid`

The PID for the Program Map Table (PMT) in the transport stream. You can enter the value as a decimal or
hexadecimal value.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProgramNum`

The value of the program number field in the Program Map Table (PMT).

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Scte35Behavior`

If set to passthrough, passes any SCTE-35 signals from the input source to this output.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Scte35Pid`

The PID of the SCTE-35 stream in the transport stream. You can enter the value as a decimal or hexadecimal
value.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimedMetadataBehavior`

When set to passthrough, timed metadata is passed through from input to output.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimedMetadataPid`

The PID of the timed metadata stream in the transport stream. You can enter the value as a decimal or
hexadecimal value. Valid values are 32 (or 0x20)..8182 (or 0x1ff6).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TransportStreamId`

The value of the transport stream ID field in the Program Map Table (PMT).

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VideoPid`

The PID of the elementary video stream in the transport stream. You can enter the value as a decimal or
hexadecimal value.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

M2tsSettings

MaintenanceCreateSettings

All content copied from https://docs.aws.amazon.com/.
