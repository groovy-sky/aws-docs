This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IVS::Channel

The `AWS::IVS::Channel` resource specifies an Amazon IVS channel. A channel stores configuration information
related to your live stream. For more information, see [CreateChannel](../../../../reference/ivs/latest/lowlatencyapireference/api-createchannel.md) in the
_Amazon IVS Low-Latency Streaming API Reference_.

###### Note

By default, the IVS API CreateChannel endpoint creates a stream key in addition
to a channel. The Amazon IVS Channel resource
_does not_ create a stream key; to create a stream key, use
the StreamKey resource instead.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IVS::Channel",
  "Properties" : {
      "Authorized" : Boolean,
      "ContainerFormat" : String,
      "InsecureIngest" : Boolean,
      "LatencyMode" : String,
      "MultitrackInputConfiguration" : MultitrackInputConfiguration,
      "Name" : String,
      "Preset" : String,
      "RecordingConfigurationArn" : String,
      "Tags" : [ Tag, ... ],
      "Type" : String
    }
}

```

### YAML

```yaml

Type: AWS::IVS::Channel
Properties:
  Authorized: Boolean
  ContainerFormat: String
  InsecureIngest: Boolean
  LatencyMode: String
  MultitrackInputConfiguration:
    MultitrackInputConfiguration
  Name: String
  Preset: String
  RecordingConfigurationArn: String
  Tags:
    - Tag
  Type: String

```

## Properties

`Authorized`

Whether the channel is authorized.

_Default_: `false`

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ContainerFormat`

Indicates which content-packaging format is used (MPEG-TS or fMP4). If `multitrackInputConfiguration` is specified and `enabled` is `true`, then `containerFormat` is required and must be set to `FRAGMENTED_MP4`. Otherwise, `containerFormat` may be set to `TS` or `FRAGMENTED_MP4`. Default: `TS`.

_Required_: No

_Type_: String

_Allowed values_: `TS | FRAGMENTED_MP4`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InsecureIngest`

Whether the channel allows insecure RTMP ingest.

_Default_: `false`

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LatencyMode`

Channel latency mode. Valid values:

- `NORMAL`: Use NORMAL to broadcast and deliver live video up to Full
HD.

- `LOW`: Use LOW for near real-time interactions with viewers.

###### Note

In the Amazon IVS console, `LOW` and
`NORMAL` correspond to `Ultra-low` and
`Standard`, respectively.

_Default_: `LOW`

_Required_: No

_Type_: String

_Allowed values_: `NORMAL | LOW`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MultitrackInputConfiguration`

Object specifying multitrack input configuration. Default: no multitrack input configuration is specified.

_Required_: No

_Type_: [MultitrackInputConfiguration](aws-properties-ivs-channel-multitrackinputconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

Channel name.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9-_]*$`

_Minimum_: `0`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Preset`

An optional transcode preset for the channel. This is selectable only for `ADVANCED_HD` and `ADVANCED_SD` channel types. For those channel types,
the default preset is `HIGHER_BANDWIDTH_DELIVERY`. For other channel types ( `BASIC` and `STANDARD`), `preset` is the empty string ("").

_Required_: No

_Type_: String

_Allowed values_: ` | HIGHER_BANDWIDTH_DELIVERY | CONSTRAINED_BANDWIDTH_DELIVERY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RecordingConfigurationArn`

The ARN of a RecordingConfiguration resource. An empty string indicates that recording
is disabled for the channel. A RecordingConfiguration ARN indicates that recording is
enabled using the specified recording configuration. See the [RecordingConfiguration](../userguide/aws-resource-ivs-recordingconfiguration.md) resource for more information and an
example.

_Default_: "" (empty string, recording is disabled)

_Required_: No

_Type_: String

_Pattern_: `^$|arn:aws:ivs:[a-z0-9-]+:[0-9]+:recording-configuration/[a-zA-Z0-9-]+$`

_Minimum_: `0`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](../userguide/aws-properties-ivs-channel-tag.md).

_Required_: No

_Type_: Array of [Tag](aws-properties-ivs-channel-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The channel type, which determines the allowable resolution and bitrate. _If_
_you exceed the allowable resolution or bitrate, the stream probably will disconnect_
_immediately._ For details, see [Channel Types](../../../../reference/ivs/latest/lowlatencyapireference/channel-types.md).

_Default_: `STANDARD`

_Required_: No

_Type_: String

_Allowed values_: `STANDARD | BASIC | ADVANCED_SD | ADVANCED_HD`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the channel ARN. For example:

`{ "Ref": "myChannel" }`

For the Amazon IVS channel `myChannel`,
`Ref` returns the channel ARN.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The channel ARN. For example:
`arn:aws:ivs:us-west-2:123456789012:channel/abcdABCDefgh`

`IngestEndpoint`

Channel ingest endpoint, part of the definition of an ingest server, used when you set
up streaming software.

For example: `a1b2c3d4e5f6.global-contribute.live-video.net`

`PlaybackUrl`

Channel playback URL. For example:
`https://a1b2c3d4e5f6.us-west-2.playback.live-video.net/api/video/v1/us-west-2.123456789012.channel.abcdEFGH.m3u8`

## Examples

### Channel and Stream Key Template Examples

The following examples specify an Amazon IVS channel
and stream key.

#### JSON

```json

{
     "AWSTemplateFormatVersion": "2010-09-09",
     "Resources": {
         "Channel": {
             "Type": "AWS::IVS::Channel",
             "Properties": {
                 "Name": "MyChannel",
                 "Tags": [
                     {
                         "Key": "MyKey",
                         "Value": "MyValue"
                     }
                 ],
                 "InsecureIngest": true
             }
         },
         "StreamKey": {
             "Type": "AWS::IVS::StreamKey",
             "Properties": {
                 "ChannelArn": {"Ref": "Channel"},
                 "Tags": [
                     {
                         "Key": "MyKey",
                         "Value": "MyValue"
                     }
                 ]
             }
         }
     },
     "Outputs": {
         "ChannelArn": {
             "Value": {"Ref": "Channel"}
         },
         "ChannelIngestEndpoint": {
             "Value": {
                 "Fn::GetAtt": [
                     "Channel",
                     "IngestEndpoint"
                 ]
             }
         },
         "ChannelPlaybackUrl": {
             "Value": {
                 "Fn::GetAtt": [
                     "Channel",
                     "PlaybackUrl"
                 ]
             }
         },
         "StreamKeyArn": {
             "Value": {"Ref": "StreamKey"}
         }
     }
 }

```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Resources:
  Channel:
    Type: AWS::IVS::Channel
    Properties:
      Name: MyChannel
      Tags:
        - Key: MyKey
          Value: MyValue
      InsecureIngest: true
  StreamKey:
    Type: AWS::IVS::StreamKey
    Properties:
      ChannelArn: !Ref Channel
      Tags:
        - Key: MyKey
          Value: MyValue
Outputs:
  ChannelArn:
    Value: !Ref Channel
  ChannelIngestEndpoint:
    Value: !GetAtt Channel.IngestEndpoint
  ChannelPlaybackUrl:
    Value: !GetAtt Channel.PlaybackUrl
  StreamKeyArn:
    Value: !Ref StreamKey

```

## See also

- [Getting\
Started with IVS Low-Latency Streaming](../../../ivs/latest/lowlatencyuserguide/getting-started.md)

- [Channel](../../../../reference/ivs/latest/lowlatencyapireference/api-channel.md) data type

- [CreateChannel](../../../../reference/ivs/latest/lowlatencyapireference/api-createchannel.md) API endpoint

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon IVS

MultitrackInputConfiguration

All content copied from https://docs.aws.amazon.com/.
