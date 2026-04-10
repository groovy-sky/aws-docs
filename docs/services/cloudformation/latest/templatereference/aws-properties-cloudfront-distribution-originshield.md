This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::Distribution OriginShield

CloudFront Origin Shield.

Using Origin Shield can help reduce the load on your origin. For more information, see
[Using Origin Shield](../../../amazoncloudfront/latest/developerguide/origin-shield.md) in the
_Amazon CloudFront Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Enabled" : Boolean,
  "OriginShieldRegion" : String
}

```

### YAML

```yaml

  Enabled: Boolean
  OriginShieldRegion: String

```

## Properties

`Enabled`

A flag that specifies whether Origin Shield is enabled.

When it's enabled, CloudFront routes all requests through Origin Shield, which can help
protect your origin. When it's disabled, CloudFront might send requests directly to your
origin from multiple edge locations or regional edge caches.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OriginShieldRegion`

The AWS Region for Origin Shield.

Specify the AWS Region that has the lowest latency to your origin. To specify a
region, use the region code, not the region name. For example, specify the US East
(Ohio) region as `us-east-2`.

When you enable CloudFront Origin Shield, you must specify the AWS Region for Origin
Shield. For the list of AWS Regions that you can specify, and for help choosing the
best Region for your origin, see [Choosing the AWS Region for Origin Shield](../../../amazoncloudfront/latest/developerguide/origin-shield.md#choose-origin-shield-region) in the
_Amazon CloudFront Developer Guide_.

_Required_: No

_Type_: String

_Pattern_: `[a-z]{2}-[a-z]+-\d`

_Minimum_: `1`

_Maximum_: `32`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OriginMtlsConfig

ParameterDefinition

All content copied from https://docs.aws.amazon.com/.
