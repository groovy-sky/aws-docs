This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel RtmpGroupSettings

The configuration of an RTMP output group.

The parent of this entity is OutputGroupSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AdMarkers" : [ String, ... ],
  "AuthenticationScheme" : String,
  "CacheFullBehavior" : String,
  "CacheLength" : Integer,
  "CaptionData" : String,
  "IncludeFillerNalUnits" : String,
  "InputLossAction" : String,
  "RestartDelay" : Integer
}

```

### YAML

```yaml

  AdMarkers:
    - String
  AuthenticationScheme: String
  CacheFullBehavior: String
  CacheLength: Integer
  CaptionData: String
  IncludeFillerNalUnits: String
  InputLossAction: String
  RestartDelay: Integer

```

## Properties

`AdMarkers`

Choose the ad marker type for this output group. MediaLive will create a message based on the content of each SCTE-35 message, format it for that marker type, and insert it in the datastream.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AuthenticationScheme`

An authentication scheme to use when connecting with a CDN.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CacheFullBehavior`

Controls behavior when the content cache fills up. If a remote origin server stalls the RTMP connection and
doesn't accept content fast enough, the media cache fills up. When the cache reaches the duration specified by
cacheLength, the cache stops accepting new content. If set to disconnectImmediately, the RTMP output forces a
disconnect. Clear the media cache, and reconnect after restartDelay seconds. If set to waitForServer, the RTMP output
waits up to 5 minutes to allow the origin server to begin accepting data again.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CacheLength`

The cache length, in seconds, that is used to calculate buffer size.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CaptionData`

Controls the types of data that pass to onCaptionInfo outputs. If set to all, 608 and 708 carried DTVCC data is
passed. If set to field1AndField2608, DTVCC data is stripped out, but 608 data from both fields is passed. If set to
field1608, only the data carried in 608 from field 1 video is passed.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncludeFillerNalUnits`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InputLossAction`

Controls the behavior of this RTMP group if the input becomes unavailable. emitOutput: Emit a slate until the
input returns. pauseOutput: Stop transmitting data until the input returns. This does not close the underlying RTMP
connection.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RestartDelay`

If a streaming output fails, the number of seconds to wait until a restart is initiated. A value of 0 means
never restart.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RemixSettings

RtmpOutputSettings

All content copied from https://docs.aws.amazon.com/.
