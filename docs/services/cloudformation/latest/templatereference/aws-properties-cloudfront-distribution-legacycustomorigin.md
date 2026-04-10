This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::Distribution LegacyCustomOrigin

A custom origin. A custom origin is any origin that is _not_ an
Amazon S3 bucket, with one exception. An Amazon S3 bucket that is [configured with\
static website hosting](../../../s3/latest/dev/websitehosting.md) _is_ a custom origin.

###### Note

This property is legacy. We recommend that you use [Origin](../userguide/aws-properties-cloudfront-distribution-origin.md) instead.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DNSName" : String,
  "HTTPPort" : Integer,
  "HTTPSPort" : Integer,
  "OriginProtocolPolicy" : String,
  "OriginSSLProtocols" : [ String, ... ]
}

```

### YAML

```yaml

  DNSName: String
  HTTPPort: Integer
  HTTPSPort: Integer
  OriginProtocolPolicy: String
  OriginSSLProtocols:
    - String

```

## Properties

`DNSName`

The domain name assigned to your CloudFront distribution.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HTTPPort`

The HTTP port that CloudFront uses to connect to the origin. Specify the HTTP port that the
origin listens on.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HTTPSPort`

The HTTPS port that CloudFront uses to connect to the origin. Specify the HTTPS port that
the origin listens on.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OriginProtocolPolicy`

Specifies the protocol (HTTP or HTTPS) that CloudFront uses to connect to the origin.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OriginSSLProtocols`

The minimum SSL/TLS protocol version that CloudFront uses when communicating with your origin server over HTTPs.

For more information, see [Minimum Origin SSL Protocol](../../../amazoncloudfront/latest/developerguide/distribution-web-values-specify.md#DownloadDistValuesOriginSSLProtocols) in the
_Amazon CloudFront Developer Guide_.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LambdaFunctionAssociation

LegacyS3Origin

All content copied from https://docs.aws.amazon.com/.
