This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::FlowOutput

The `AWS::MediaConnect::FlowOutput` resource defines the destination address, protocol,
and port that AWS Elemental MediaConnect sends the ingested video to. Each
flow can have up to 50 outputs. An output can have the same protocol or a different
protocol from the source. The following protocols are supported: RIST, RTP, RTP-FEC,
SRT-listener, SRT-caller, Zixi pull, and Zixi push. CDI and ST 2110 JPEG XS protocols are
not currently supported by AWS CloudFormation.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::MediaConnect::FlowOutput",
  "Properties" : {
      "CidrAllowList" : [ String, ... ],
      "Description" : String,
      "Destination" : String,
      "Encryption" : Encryption,
      "FlowArn" : String,
      "MaxLatency" : Integer,
      "MediaStreamOutputConfigurations" : [ MediaStreamOutputConfiguration, ... ],
      "MinLatency" : Integer,
      "Name" : String,
      "NdiProgramName" : String,
      "NdiSpeedHqQuality" : Integer,
      "OutputStatus" : String,
      "Port" : Integer,
      "Protocol" : String,
      "RemoteId" : String,
      "RouterIntegrationState" : String,
      "RouterIntegrationTransitEncryption" : FlowTransitEncryption,
      "SmoothingLatency" : Integer,
      "StreamId" : String,
      "Tags" : [ Tag, ... ],
      "VpcInterfaceAttachment" : VpcInterfaceAttachment
    }
}

```

### YAML

```yaml

Type: AWS::MediaConnect::FlowOutput
Properties:
  CidrAllowList:
    - String
  Description: String
  Destination: String
  Encryption:
    Encryption
  FlowArn: String
  MaxLatency: Integer
  MediaStreamOutputConfigurations:
    - MediaStreamOutputConfiguration
  MinLatency: Integer
  Name: String
  NdiProgramName: String
  NdiSpeedHqQuality: Integer
  OutputStatus: String
  Port: Integer
  Protocol: String
  RemoteId: String
  RouterIntegrationState: String
  RouterIntegrationTransitEncryption:
    FlowTransitEncryption
  SmoothingLatency: Integer
  StreamId: String
  Tags:
    - Tag
  VpcInterfaceAttachment:
    VpcInterfaceAttachment

```

## Properties

`CidrAllowList`

The range of IP addresses that should be allowed to initiate output requests to this flow. These IP addresses should be in the form of a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description of the output. This description appears only on the MediaConnect console and will not be seen by the end user.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Destination`

The IP address where you want to send the output.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Encryption`

The type of key used for the encryption. If no `keyType` is provided, the
service will use the default setting (static-key). Allowable encryption types:
static-key.

_Required_: No

_Type_: [Encryption](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-mediaconnect-flowoutput-encryption.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FlowArn`

The Amazon Resource Name (ARN) of the flow this output is attached to.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:(aws[a-zA-Z-]*):mediaconnect:[a-z0-9-]+:[0-9]{12}:flow:[a-zA-Z0-9-]+:[a-zA-Z0-9_-]+$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MaxLatency`

The maximum latency in milliseconds. This parameter applies only to RIST-based and
Zixi-based streams.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MediaStreamOutputConfigurations`

The media streams that are associated with the output, and the parameters for those associations.

_Required_: No

_Type_: Array of [MediaStreamOutputConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-mediaconnect-flowoutput-mediastreamoutputconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinLatency`

The minimum latency in milliseconds for SRT-based streams. In streams that use the SRT protocol, this value that you set on your MediaConnect source or output represents the minimal potential latency of that connection. The latency of the stream is set to the highest number between the sender’s minimum latency and the receiver’s minimum latency.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the bridge's output.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NdiProgramName`

A suffix for the name of the NDI® sender that the flow creates. If a custom name isn't specified, MediaConnect uses the output name.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NdiSpeedHqQuality`

A quality setting for the NDI Speed HQ encoder.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OutputStatus`

An indication of whether the output should transmit data or not. If you don't specify the
`outputStatus` field in your request, MediaConnect leaves the value
unchanged.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Port`

The port to use when content is distributed to this output.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Protocol`

The protocol to use for the output.

###### Note

AWS Elemental MediaConnect no longer supports the Fujitsu QoS
protocol. This reference is maintained for legacy purposes only.

_Required_: No

_Type_: String

_Allowed values_: `zixi-push | rtp-fec | rtp | zixi-pull | rist | srt-listener | srt-caller | st2110-jpegxs | cdi | ndi-speed-hq`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RemoteId`

The remote ID for the Zixi-pull stream.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RouterIntegrationState`

Indicates if router integration is enabled or disabled on the flow output.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RouterIntegrationTransitEncryption`

The encryption configuration for the output when router integration is enabled.

_Required_: No

_Type_: [FlowTransitEncryption](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-mediaconnect-flowoutput-flowtransitencryption.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SmoothingLatency`

The smoothing latency in milliseconds for RIST, RTP, and RTP-FEC streams.

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

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-mediaconnect-flowoutput-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcInterfaceAttachment`

The name of the VPC interface attachment to use for this output.

_Required_: No

_Type_: [VpcInterfaceAttachment](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-mediaconnect-flowoutput-vpcinterfaceattachment.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the output ARN. For example:

`{ "Ref":
            "arn:aws:mediaconnect:us-east-1:111122223333:output:2-3aBC45dEF67hiJ89-c34de5fG678h:NYCOutput"
            }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`OutputArn`

The ARN of the output.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

DestinationConfiguration
