This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaPackageV2::OriginEndpoint

Specifies the configuration parameters for a MediaPackage V2 origin endpoint.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::MediaPackageV2::OriginEndpoint",
  "Properties" : {
      "ChannelGroupName" : String,
      "ChannelName" : String,
      "ContainerType" : String,
      "DashManifests" : [ DashManifestConfiguration, ... ],
      "Description" : String,
      "ForceEndpointErrorConfiguration" : ForceEndpointErrorConfiguration,
      "HlsManifests" : [ HlsManifestConfiguration, ... ],
      "LowLatencyHlsManifests" : [ LowLatencyHlsManifestConfiguration, ... ],
      "MssManifests" : [ MssManifestConfiguration, ... ],
      "OriginEndpointName" : String,
      "Segment" : Segment,
      "StartoverWindowSeconds" : Integer,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::MediaPackageV2::OriginEndpoint
Properties:
  ChannelGroupName: String
  ChannelName: String
  ContainerType: String
  DashManifests:
    - DashManifestConfiguration
  Description: String
  ForceEndpointErrorConfiguration:
    ForceEndpointErrorConfiguration
  HlsManifests:
    - HlsManifestConfiguration
  LowLatencyHlsManifests:
    - LowLatencyHlsManifestConfiguration
  MssManifests:
    - MssManifestConfiguration
  OriginEndpointName: String
  Segment:
    Segment
  StartoverWindowSeconds: Integer
  Tags:
    - Tag

```

## Properties

`ChannelGroupName`

The name of the channel group associated with the origin endpoint configuration.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ChannelName`

The channel name associated with the origin endpoint.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ContainerType`

The container type associated with the origin endpoint configuration.

_Required_: Yes

_Type_: String

_Allowed values_: `TS | CMAF | ISM`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DashManifests`

A DASH manifest configuration.

_Required_: No

_Type_: Array of [DashManifestConfiguration](aws-properties-mediapackagev2-originendpoint-dashmanifestconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description associated with the origin endpoint.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ForceEndpointErrorConfiguration`

The failover settings for the endpoint.

_Required_: No

_Type_: [ForceEndpointErrorConfiguration](aws-properties-mediapackagev2-originendpoint-forceendpointerrorconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HlsManifests`

The HLS manifests associated with the origin endpoint configuration.

_Required_: No

_Type_: Array of [HlsManifestConfiguration](aws-properties-mediapackagev2-originendpoint-hlsmanifestconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LowLatencyHlsManifests`

The low-latency HLS (LL-HLS) manifests associated with the origin endpoint.

_Required_: No

_Type_: Array of [LowLatencyHlsManifestConfiguration](aws-properties-mediapackagev2-originendpoint-lowlatencyhlsmanifestconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MssManifests`

A list of Microsoft Smooth Streaming (MSS) manifest configurations associated with the origin endpoint. Each configuration represents a different MSS streaming option available from this endpoint.

_Required_: No

_Type_: Array of [MssManifestConfiguration](aws-properties-mediapackagev2-originendpoint-mssmanifestconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OriginEndpointName`

The name of the origin endpoint associated with the origin endpoint configuration.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Segment`

The segment associated with the origin endpoint.

_Required_: No

_Type_: [Segment](aws-properties-mediapackagev2-originendpoint-segment.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StartoverWindowSeconds`

The size of the window (in seconds) to specify a window of the live stream that's available for on-demand viewing. Viewers can start-over or catch-up on content that falls within the window.

_Required_: No

_Type_: Integer

_Minimum_: `60`

_Maximum_: `1209600`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags associated with the origin endpoint.

_Required_: No

_Type_: Array of [Tag](aws-properties-mediapackagev2-originendpoint-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns `arn:aws:mediapackagev2:region:AccountId:ChannelGroup/ChannelGroupName/Channel/ChannelName/OriginEndpoint/OriginEndpointName`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The attributes of the origin endpoint. You can only use the `GetAtt` function for `readOnly` properties. For example, you can use the `GetAtt` function to retrieve the read-only `CreatedAt` property.

`Arn`

The Amazon Resource Name (ARN) of the origin endpoint.

`CreatedAt`

The timestamp of the creation of the origin endpoint.

`DashManifestUrls`

The egress domain URL for stream delivery from MediaPackage.

`HlsManifestUrls`

The egress domain URL for stream delivery from MediaPackage.

`LowLatencyHlsManifestUrls`

The egress domain URL for stream delivery from MediaPackage.

`ModifiedAt`

The timestamp of the modification of the origin endpoint.

`MssManifestUrls`

Property description not available.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::MediaPackageV2::ChannelPolicy

DashBaseUrl

All content copied from https://docs.aws.amazon.com/.
