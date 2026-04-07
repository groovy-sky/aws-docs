This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IVS::StreamKey

The `AWS::IVS::StreamKey` resource specifies an Amazon IVS stream key associated with the referenced channel. Use
a stream key to initiate a live stream.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IVS::StreamKey",
  "Properties" : {
      "ChannelArn" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::IVS::StreamKey
Properties:
  ChannelArn: String
  Tags:
    - Tag

```

## Properties

`ChannelArn`

Channel ARN for the stream.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws:ivs:[a-z0-9-]+:[0-9]+:channel/[a-zA-Z0-9-]+$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-streamkey-tag.html).

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ivs-streamkey-tag.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource ARN. For example:

`{ "Ref": "myStreamKey" }`

For the Amazon IVS stream key `myStreamKey`,
`Ref` returns the stream key ARN.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The stream-key ARN. For example:
`arn:aws:ivs:us-west-2:123456789012:stream-key/g1H2I3j4k5L6`

`Value`

The stream-key value. For example:
`sk_us-west-2_abcdABCDefgh_567890abcdef`

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
Started with IVS Low-Latency Streaming](https://docs.aws.amazon.com/ivs/latest/LowLatencyUserGuide/getting-started.html)

- [StreamKey](https://docs.aws.amazon.com/ivs/latest/LowLatencyAPIReference/API_StreamKey.html) data
type

- [CreateStreamKey](https://docs.aws.amazon.com/ivs/latest/LowLatencyAPIReference/API_CreateStreamKey.html) API endpoint

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Tag
