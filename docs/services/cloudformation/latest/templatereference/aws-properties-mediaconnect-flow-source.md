This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::Flow Source

The details of the sources of the flow.

If you are creating a flow with a VPC source, you must first create the flow with a
temporary standard source by doing the following:

1. Use CloudFormation to create a flow with a standard source that uses the
    flow’s public IP address.

2. Use CloudFormation to create the VPC interface to add to this flow. This can
    also be done as part of the previous step.

3. After CloudFormation has created the flow and the VPC interface, update the
    source to point to the VPC interface that you created.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Decryption" : Encryption,
  "Description" : String,
  "EntitlementArn" : String,
  "GatewayBridgeSource" : GatewayBridgeSource,
  "IngestIp" : String,
  "IngestPort" : Integer,
  "MaxBitrate" : Integer,
  "MaxLatency" : Integer,
  "MaxSyncBuffer" : Integer,
  "MediaStreamSourceConfigurations" : [ MediaStreamSourceConfiguration, ... ],
  "MinLatency" : Integer,
  "Name" : String,
  "NdiSourceSettings" : NdiSourceSettings,
  "Protocol" : String,
  "RouterIntegrationState" : String,
  "RouterIntegrationTransitDecryption" : FlowTransitEncryption,
  "SenderControlPort" : Integer,
  "SenderIpAddress" : String,
  "SourceArn" : String,
  "SourceIngestPort" : String,
  "SourceListenerAddress" : String,
  "SourceListenerPort" : Integer,
  "StreamId" : String,
  "Tags" : [ Tag, ... ],
  "VpcInterfaceName" : String,
  "WhitelistCidr" : String
}

```

### YAML

```yaml

  Decryption:
    Encryption
  Description: String
  EntitlementArn: String
  GatewayBridgeSource:
    GatewayBridgeSource
  IngestIp: String
  IngestPort: Integer
  MaxBitrate: Integer
  MaxLatency: Integer
  MaxSyncBuffer: Integer
  MediaStreamSourceConfigurations:
    - MediaStreamSourceConfiguration
  MinLatency: Integer
  Name: String
  NdiSourceSettings:
    NdiSourceSettings
  Protocol: String
  RouterIntegrationState: String
  RouterIntegrationTransitDecryption:
    FlowTransitEncryption
  SenderControlPort: Integer
  SenderIpAddress: String
  SourceArn: String
  SourceIngestPort: String
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

The type of encryption that is used on the content ingested from this source.

_Required_: No

_Type_: [Encryption](aws-properties-mediaconnect-flow-encryption.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description for the source. This value is not used or seen outside of the current MediaConnect account.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EntitlementArn`

The ARN of the entitlement that allows you to subscribe to content that comes from another AWS account. The entitlement is set by the content originator and the ARN is generated as part of the originator's flow.

_Required_: No

_Type_: String

_Pattern_: `^arn:(aws[a-zA-Z-]*):mediaconnect:[a-z0-9-]+:[0-9]{12}:entitlement:[a-zA-Z0-9-]+:[a-zA-Z0-9_-]+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GatewayBridgeSource`

The source configuration for cloud flows receiving a stream from a bridge.

_Required_: No

_Type_: [GatewayBridgeSource](aws-properties-mediaconnect-flow-gatewaybridgesource.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IngestIp`

The IP address that the flow will be listening on for incoming content.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IngestPort`

The port that the flow will be listening on for incoming content.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxBitrate`

The maximum bitrate for RIST, RTP, and RTP-FEC streams.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxLatency`

The maximum latency in milliseconds for a RIST or Zixi-based source.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxSyncBuffer`

The size of the buffer (in milliseconds) to use to sync incoming source data.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MediaStreamSourceConfigurations`

The media streams that are associated with the source, and the parameters for those associations.

_Required_: No

_Type_: Array of [MediaStreamSourceConfiguration](aws-properties-mediaconnect-flow-mediastreamsourceconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinLatency`

The minimum latency in milliseconds for SRT-based streams. In streams that use the SRT
protocol, this value that you set on your MediaConnect source or output represents the
minimal potential latency of that connection. The latency of the stream is set to the
highest number between the sender’s minimum latency and the receiver’s minimum
latency.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the source.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NdiSourceSettings`

Property description not available.

_Required_: No

_Type_: [NdiSourceSettings](aws-properties-mediaconnect-flow-ndisourcesettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Protocol`

The protocol that is used by the source. AWS CloudFormation does not currently support CDI or ST 2110 JPEG XS source protocols.

###### Note

AWS Elemental MediaConnect no longer supports the Fujitsu QoS
protocol. This reference is maintained for legacy purposes only.

_Required_: No

_Type_: String

_Allowed values_: `zixi-push | rtp-fec | rtp | rist | srt-listener | srt-caller | st2110-jpegxs | cdi | ndi-speed-hq`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RouterIntegrationState`

Indicates if router integration is enabled or disabled on the flow source.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RouterIntegrationTransitDecryption`

The decryption configuration for the flow source when router integration is enabled.

_Required_: No

_Type_: [FlowTransitEncryption](aws-properties-mediaconnect-flow-flowtransitencryption.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SenderControlPort`

The port that the flow uses to send outbound requests to initiate connection with
the sender.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SenderIpAddress`

The IP address that the flow communicates with to initiate connection with the
sender.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceArn`

The ARN of the source.

_Required_: No

_Type_: String

_Pattern_: `^arn:(aws[a-zA-Z-]*):mediaconnect:[a-z0-9-]+:[0-9]{12}:source:[a-zA-Z0-9-]+:[a-zA-Z0-9_-]+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceIngestPort`

The port that the flow listens on for incoming content. If the protocol of the source is Zixi, the port must
be set to 2088.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

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

The stream ID that you want to use for the transport. This parameter applies only to
Zixi-based streams.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Property description not available.

_Required_: No

_Type_: Array of [Tag](aws-properties-mediaconnect-flow-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcInterfaceName`

The name of the VPC interface that is used for this source.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WhitelistCidr`

The range of IP addresses that should be allowed to contribute content to your source. These IP addresses should be in the form of a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SilentAudio

SourceMonitoringConfig

All content copied from https://docs.aws.amazon.com/.
