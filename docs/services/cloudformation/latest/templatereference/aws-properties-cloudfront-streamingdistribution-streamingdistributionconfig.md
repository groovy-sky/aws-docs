This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::StreamingDistribution StreamingDistributionConfig

The RTMP distribution's configuration information.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Aliases" : [ String, ... ],
  "Comment" : String,
  "Enabled" : Boolean,
  "Logging" : Logging,
  "PriceClass" : String,
  "S3Origin" : S3Origin,
  "TrustedSigners" : TrustedSigners
}

```

### YAML

```yaml

  Aliases:
    - String
  Comment: String
  Enabled: Boolean
  Logging:
    Logging
  PriceClass: String
  S3Origin:
    S3Origin
  TrustedSigners:
    TrustedSigners

```

## Properties

`Aliases`

A complex type that contains information about CNAMEs (alternate domain names), if
any, for this streaming distribution.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Comment`

Any comments you want to include about the streaming distribution.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Enabled`

Whether the streaming distribution is enabled to accept user requests for
content.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Logging`

A complex type that controls whether access logs are written for the streaming
distribution.

_Required_: No

_Type_: [Logging](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudfront-streamingdistribution-logging.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PriceClass`

A complex type that contains information about price class for this streaming
distribution.

_Required_: No

_Type_: String

_Allowed values_: `PriceClass_100 | PriceClass_200 | PriceClass_All | None`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3Origin`

A complex type that contains information about the Amazon S3 bucket from which you want
CloudFront to get your media files for distribution.

_Required_: Yes

_Type_: [S3Origin](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudfront-streamingdistribution-s3origin.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TrustedSigners`

A complex type that specifies any AWS accounts that you want to permit to create
signed URLs for private content. If you want the distribution to use signed URLs,
include this element; if you want the distribution to use public URLs, remove this
element. For more information, see [Serving Private\
Content through CloudFront](../../../amazoncloudfront/latest/developerguide/privatecontent.md) in the _Amazon CloudFront Developer Guide_.

_Required_: Yes

_Type_: [TrustedSigners](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudfront-streamingdistribution-trustedsigners.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [StreamingDistributionConfig](https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_StreamingDistributionConfig.html) in the _Amazon CloudFront API Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

S3Origin

Tag
