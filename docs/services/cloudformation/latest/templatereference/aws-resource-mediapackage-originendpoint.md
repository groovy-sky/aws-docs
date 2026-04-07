This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaPackage::OriginEndpoint

Create an endpoint on an AWS Elemental MediaPackage channel.

An endpoint represents a single delivery point of a channel, and defines content output handling through various components, such as packaging protocols, DRM and encryption integration, and more.

After it's created, an endpoint provides a fixed public URL. This URL remains the same
throughout the lifetime of the endpoint, regardless of any failures or upgrades that might occur. Integrate the URL with a downstream CDN (such as Amazon CloudFront) or playback
device.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::MediaPackage::OriginEndpoint",
  "Properties" : {
      "Authorization" : Authorization,
      "ChannelId" : String,
      "CmafPackage" : CmafPackage,
      "DashPackage" : DashPackage,
      "Description" : String,
      "HlsPackage" : HlsPackage,
      "Id" : String,
      "ManifestName" : String,
      "MssPackage" : MssPackage,
      "Origination" : String,
      "StartoverWindowSeconds" : Integer,
      "Tags" : [ Tag, ... ],
      "TimeDelaySeconds" : Integer,
      "Whitelist" : [ String, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::MediaPackage::OriginEndpoint
Properties:
  Authorization:
    Authorization
  ChannelId: String
  CmafPackage:
    CmafPackage
  DashPackage:
    DashPackage
  Description: String
  HlsPackage:
    HlsPackage
  Id: String
  ManifestName: String
  MssPackage:
    MssPackage
  Origination: String
  StartoverWindowSeconds: Integer
  Tags:
    - Tag
  TimeDelaySeconds: Integer
  Whitelist:
    - String

```

## Properties

`Authorization`

Parameters for CDN authorization.

_Required_: No

_Type_: [Authorization](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-mediapackage-originendpoint-authorization.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ChannelId`

The ID of the channel associated with this endpoint.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CmafPackage`

Parameters for Common Media Application Format (CMAF) packaging.

_Required_: No

_Type_: [CmafPackage](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-mediapackage-originendpoint-cmafpackage.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DashPackage`

Parameters for DASH packaging.

_Required_: No

_Type_: [DashPackage](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-mediapackage-originendpoint-dashpackage.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

Any descriptive information that you want to add to the endpoint for future identification purposes.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HlsPackage`

Parameters for Apple HLS packaging.

_Required_: No

_Type_: [HlsPackage](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-mediapackage-originendpoint-hlspackage.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Id`

The manifest ID is required and must be unique within the OriginEndpoint. The ID can't be changed after the endpoint is created.

_Required_: Yes

_Type_: String

_Pattern_: `\A[0-9a-zA-Z-_]+\Z`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ManifestName`

A short string that's appended to the end of the endpoint URL to create a unique path to this endpoint.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MssPackage`

Parameters for Microsoft Smooth Streaming packaging.

_Required_: No

_Type_: [MssPackage](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-mediapackage-originendpoint-msspackage.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Origination`

Controls video origination from this endpoint.

Valid values:

- `ALLOW` \- enables this endpoint to serve content to requesting devices.

- `DENY` \- prevents this endpoint from serving content. Denying origination is helpful for harvesting live-to-VOD assets. For more information about harvesting and origination, see
[Live-to-VOD Requirements](https://docs.aws.amazon.com/mediapackage/latest/ug/ltov-reqmts.html).

_Required_: No

_Type_: String

_Allowed values_: `ALLOW | DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StartoverWindowSeconds`

Maximum duration (seconds) of content to retain for startover playback. Omit this attribute or enter `0` to indicate that startover playback is disabled for this endpoint.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags to assign to the endpoint.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-mediapackage-originendpoint-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeDelaySeconds`

Minimum duration (seconds) of delay to enforce on the playback of live content. Omit this attribute or enter `0` to indicate that there is no time delay in effect for this endpoint.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Whitelist`

The IP addresses that can access this endpoint.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the name of the origin endpoint.

For example: `{ "Ref": "myOriginEndpoint" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The endpoint's unique system-generated resource name, based on the AWS record.

`Url`

URL for the key provider’s key retrieval API endpoint. Must start with https://.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Authorization
