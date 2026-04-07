This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lightsail::Distribution

The `AWS::Lightsail::Distribution` resource specifies a content delivery
network (CDN) distribution. You can create distributions only in the `us-east-1` AWS Region.

A distribution is a globally distributed network of caching servers that improve the
performance of your website or web application hosted on a Lightsail
instance, static content hosted on a Lightsail bucket, or through a Lightsail load balancer.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Lightsail::Distribution",
  "Properties" : {
      "BundleId" : String,
      "CacheBehaviors" : [ CacheBehaviorPerPath, ... ],
      "CacheBehaviorSettings" : CacheSettings,
      "CertificateName" : String,
      "DefaultCacheBehavior" : CacheBehavior,
      "DistributionName" : String,
      "IpAddressType" : String,
      "IsEnabled" : Boolean,
      "Origin" : InputOrigin,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Lightsail::Distribution
Properties:
  BundleId: String
  CacheBehaviors:
    - CacheBehaviorPerPath
  CacheBehaviorSettings:
    CacheSettings
  CertificateName: String
  DefaultCacheBehavior:
    CacheBehavior
  DistributionName: String
  IpAddressType: String
  IsEnabled: Boolean
  Origin:
    InputOrigin
  Tags:
    - Tag

```

## Properties

`BundleId`

The ID of the bundle applied to the distribution.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CacheBehaviors`

An array of objects that describe the per-path cache behavior of the
distribution.

_Required_: No

_Type_: Array of [CacheBehaviorPerPath](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-lightsail-distribution-cachebehaviorperpath.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CacheBehaviorSettings`

An object that describes the cache behavior settings of the distribution.

_Required_: No

_Type_: [CacheSettings](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-lightsail-distribution-cachesettings.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CertificateName`

The name of the SSL/TLS certificate attached to the distribution.

_Required_: No

_Type_: String

_Pattern_: `\w[\w\-]*\w`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultCacheBehavior`

An object that describes the default cache behavior of the distribution.

_Required_: Yes

_Type_: [CacheBehavior](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-lightsail-distribution-cachebehavior.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DistributionName`

The name of the distribution

_Required_: Yes

_Type_: String

_Pattern_: `\w[\w\-]*\w`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`IpAddressType`

The IP address type of the distribution.

The possible values are `ipv4` for IPv4 only, and `dualstack` for
IPv4 and IPv6.

_Required_: No

_Type_: String

_Allowed values_: `dualstack | ipv4 | ipv6`

_Update requires_: Updates are not supported.

`IsEnabled`

A Boolean value indicating whether the distribution is enabled.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Origin`

An object that describes the origin resource of the distribution, such as a Lightsail instance, bucket, or load balancer.

The distribution pulls, caches, and serves content from the origin.

_Required_: Yes

_Type_: [InputOrigin](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-lightsail-distribution-inputorigin.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html)
in the _AWS CloudFormation User Guide_.

###### Note

The `Value` of `Tags` is optional for Lightsail resources.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-lightsail-distribution-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns a unique identifier for this resource.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AbleToUpdateBundle`

Indicates whether you can update the distribution’s current bundle to another bundle.

`DistributionArn`

The Amazon Resource Name (ARN) of the distribution.

`Status`

The status of the distribution.

## Remarks

_Configuring cache behavior settings_

The `CacheBehaviorSettings` parameter can be set only if the `DefaultCacheBehavior` parameter is
set to `cache`, or if the `CacheBehaviors` parameter has a path with a `cache` behavior.
If neither of those conditions are true, the `CacheBehaviorSettings` will not be set for the distribution
and the stack will drift.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

CacheBehavior
