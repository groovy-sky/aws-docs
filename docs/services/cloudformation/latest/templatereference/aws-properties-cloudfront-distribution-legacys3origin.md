This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::Distribution LegacyS3Origin

The origin as an Amazon S3 bucket.

###### Note

This property is legacy. We recommend that you use [Origin](../userguide/aws-properties-cloudfront-distribution-origin.md) instead.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DNSName" : String,
  "OriginAccessIdentity" : String
}

```

### YAML

```yaml

  DNSName: String
  OriginAccessIdentity: String

```

## Properties

`DNSName`

The domain name assigned to your CloudFront distribution.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OriginAccessIdentity`

The CloudFront origin access identity to associate with the distribution. Use an origin
access identity to configure the distribution so that end users can only access objects
in an Amazon S3 through CloudFront.

###### Note

This property is legacy. We recommend that you use [OriginAccessControl](../userguide/aws-resource-cloudfront-originaccesscontrol.md) instead.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LegacyCustomOrigin

Logging

All content copied from https://docs.aws.amazon.com/.
