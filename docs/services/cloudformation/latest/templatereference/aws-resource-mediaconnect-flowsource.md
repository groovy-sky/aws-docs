This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::FlowSource

The `AWS::MediaConnect::FlowSource` resource is used to add additional sources to an
existing flow. Adding an additional source requires Failover to be enabled. When you enable
Failover, the additional source must use the same protocol as the existing source. A source
is the external video content that includes configuration information (encryption and
source type) and a network address. Each flow has at least one source. A standard source
comes from a source other than another AWS Elemental MediaConnect flow,
such as an on-premises encoder.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::MediaConnect::FlowSource",
  "Properties" : {
      "Decryption" : Encryption,
      "Description" : String,
      "FlowArn" : String,
      "GatewayBridgeSource" : GatewayBridgeSource,
      "IngestPort" : Integer,
      "MaxBitrate" : Integer,
      "MaxLatency" : Integer,
      "MinLatency" : Integer,
      "Name" : String,
      "Protocol" : String,
      "SourceListenerAddress" : String,
      "SourceListenerPort" : Integer,
      "StreamId" : String,
      "Tags" : [ Tag, ... ],
      "VpcInterfaceName" : String,
      "WhitelistCidr" : String
    }
}

```

### YAML

```yaml

Type: AWS::MediaConnect::FlowSource
Properties:
  Decryption:
    Encryption
  Description: String
  FlowArn: String
  GatewayBridgeSource:
    GatewayBridgeSource
  IngestPort: Integer
  MaxBitrate: Integer
  MaxLatency: Integer
  MinLatency: Integer
  Name: String
  Protocol: String
  SourceListenerAddress: String
  SourceListenerPort: Integer
  StreamId: String
  Tags:
    - Tag
  VpcInterfaceName: String
  WhitelistCidr: String

```

## Properties

`Decryption`

The type of encryption that is used on the content ingested from this source. Allowable encryption types: static-key.

_Required_: No

_Type_: [Encryption](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-mediaconnect-flowsource-encryption.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description for the source. This value is not used or seen outside of the current MediaConnect account.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FlowArn`

The Amazon Resource Name (ARN) of the flow this source is connected to. The flow must have Failover enabled to add an additional source.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:(aws[a-zA-Z-]*):mediaconnect:[a-z0-9-]+:[0-9]{12}:flow:[a-zA-Z0-9-]+:[a-zA-Z0-9_-]+$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`GatewayBridgeSource`

The bridge's source.

_Required_: No

_Type_: [GatewayBridgeSource](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-mediaconnect-flowsource-gatewaybridgesource.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IngestPort`

The port that the flow listens on for incoming content. If the protocol of the
source is Zixi, the port must be set to 2088.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxBitrate`

The smoothing max bitrate (in bps) for RIST, RTP, and RTP-FEC streams.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxLatency`

The maximum latency in milliseconds. This parameter applies only to RIST-based and
Zixi-based streams.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinLatency`

The minimum latency in milliseconds for SRT-based streams. In streams that use the SRT protocol, this value that you set on your MediaConnect source or output represents the minimal potential latency of that connection. The latency of the stream is set to the highest number between the sender’s minimum latency and the receiver’s minimum latency.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the source.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Protocol`

The protocol that the source uses to deliver the content to
MediaConnect. Adding additional sources to an existing flow requires Failover to be
enabled. When you enable Failover, the additional source must use the same protocol as
the existing source. Only the following protocols support failover: Zixi-push, RTP-FEC,
RTP, RIST and SRT protocols.

If you use failover with SRT caller or listener,
the `FailoverMode` property must be set to `FAILOVER`. The
`FailoverMode` property is found in the `FailoverConfig`
resource of the same flow ARN you used for the source's `FlowArn` property.
SRT caller/listener does not support merge mode failover.

_Required_: No

_Type_: String

_Allowed values_: `zixi-push | rtp-fec | rtp | rist | srt-listener | srt-caller`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SourceListenerAddress`

Source IP or domain name for SRT-caller protocol.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceListenerPort`

Source port for SRT-caller protocol.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StreamId`

The stream ID that you want to use for this transport. This parameter applies only to Zixi and SRT caller-based streams.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Property description not available.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-mediaconnect-flowsource-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcInterfaceName`

The name of the VPC interface to use for this source.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WhitelistCidr`

The range of IP addresses that should be allowed to contribute content to your source. These IP addresses should be in the form of a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the source ARN. For example:

`{ "Ref":
            "arn:aws:mediaconnect:us-east-1:111122223333:source:2-3aBC45dEF67hiJ89-c34de5fG678h:AwardsShowSource"
            }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`IngestIp`

The IP address that the flow listens on for incoming content.

`SourceArn`

The ARN of the source.

`SourceIngestPort`

The port that the flow listens on for incoming content. If the protocol of the source is Zixi, the port must be set to 2088.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

VpcInterfaceAttachment

Encryption
