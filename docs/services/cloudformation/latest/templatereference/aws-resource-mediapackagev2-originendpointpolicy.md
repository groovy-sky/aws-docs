This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaPackageV2::OriginEndpointPolicy

Specifies the configuration parameters of a policy associated with a MediaPackage V2 origin endpoint.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::MediaPackageV2::OriginEndpointPolicy",
  "Properties" : {
      "CdnAuthConfiguration" : CdnAuthConfiguration,
      "ChannelGroupName" : String,
      "ChannelName" : String,
      "OriginEndpointName" : String,
      "Policy" : Json
    }
}

```

### YAML

```yaml

Type: AWS::MediaPackageV2::OriginEndpointPolicy
Properties:
  CdnAuthConfiguration:
    CdnAuthConfiguration
  ChannelGroupName: String
  ChannelName: String
  OriginEndpointName: String
  Policy: Json

```

## Properties

`CdnAuthConfiguration`

The settings to enable CDN authorization headers in MediaPackage.

_Required_: No

_Type_: [CdnAuthConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-mediapackagev2-originendpointpolicy-cdnauthconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ChannelGroupName`

The name of the channel group associated with the origin endpoint policy.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ChannelName`

The channel name associated with the origin endpoint policy.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`OriginEndpointName`

The name of the origin endpoint associated with the origin endpoint policy.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Policy`

The policy associated with the origin endpoint.

_Required_: Yes

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns `ChannelGroupName/ChannelName/OriginEndpointName`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

CdnAuthConfiguration
