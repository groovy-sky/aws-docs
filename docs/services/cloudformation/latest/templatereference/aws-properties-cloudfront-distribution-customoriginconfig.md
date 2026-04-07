This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::Distribution CustomOriginConfig

A custom origin. A custom origin is any origin that is _not_ an
Amazon S3 bucket, with one exception. An Amazon S3 bucket that is [configured with\
static website hosting](https://docs.aws.amazon.com/AmazonS3/latest/dev/WebsiteHosting.html) _is_ a custom origin.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "HTTPPort" : Integer,
  "HTTPSPort" : Integer,
  "IpAddressType" : String,
  "OriginKeepaliveTimeout" : Integer,
  "OriginMtlsConfig" : OriginMtlsConfig,
  "OriginProtocolPolicy" : String,
  "OriginReadTimeout" : Integer,
  "OriginSSLProtocols" : [ String, ... ]
}

```

### YAML

```yaml

  HTTPPort: Integer
  HTTPSPort: Integer
  IpAddressType: String
  OriginKeepaliveTimeout: Integer
  OriginMtlsConfig:
    OriginMtlsConfig
  OriginProtocolPolicy: String
  OriginReadTimeout: Integer
  OriginSSLProtocols:
    - String

```

## Properties

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

`IpAddressType`

Specifies which IP protocol CloudFront uses when connecting to your origin. If your origin uses both IPv4 and IPv6 protocols, you can choose `dualstack` to help optimize reliability.

_Required_: No

_Type_: String

_Allowed values_: `ipv4 | ipv6 | dualstack`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OriginKeepaliveTimeout`

Specifies how long, in seconds, CloudFront persists its connection to the origin. The
minimum timeout is 1 second, the maximum is 120 seconds, and the default (if you don't
specify otherwise) is 5 seconds.

For more information, see [Keep-alive timeout (custom origins only)](../../../amazoncloudfront/latest/developerguide/downloaddistvaluesorigin.md#DownloadDistValuesOriginKeepaliveTimeout) in the
_Amazon CloudFront Developer Guide_.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OriginMtlsConfig`

Configures mutual TLS authentication between CloudFront and your origin server.

_Required_: No

_Type_: [OriginMtlsConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudfront-distribution-originmtlsconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OriginProtocolPolicy`

Specifies the protocol (HTTP or HTTPS) that CloudFront uses to connect to the origin. Valid
values are:

- `http-only` – CloudFront always uses HTTP to connect to the
origin.

- `match-viewer` – CloudFront connects to the origin using the same
protocol that the viewer used to connect to CloudFront.

- `https-only` – CloudFront always uses HTTPS to connect to the
origin.

_Required_: Yes

_Type_: String

_Allowed values_: `http-only | match-viewer | https-only`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OriginReadTimeout`

Specifies how long, in seconds, CloudFront waits for a response from the origin. This is
also known as the _origin response timeout_. The minimum timeout is 1
second, the maximum is 120 seconds, and the default (if you don't specify otherwise) is
30 seconds.

For more information, see [Response timeout](../../../amazoncloudfront/latest/developerguide/downloaddistvaluesorigin.md#DownloadDistValuesOriginResponseTimeout) in the
_Amazon CloudFront Developer Guide_.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OriginSSLProtocols`

Specifies the minimum SSL/TLS protocol that CloudFront uses when connecting to your origin
over HTTPS. Valid values include `SSLv3`, `TLSv1`,
`TLSv1.1`, and `TLSv1.2`.

For more information, see [Minimum Origin SSL Protocol](../../../amazoncloudfront/latest/developerguide/downloaddistvaluesorigin.md#DownloadDistValuesOriginSSLProtocols) in the
_Amazon CloudFront Developer Guide_.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [CustomOriginConfig](https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_CustomOriginConfig.html) in the _Amazon CloudFront API Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CustomErrorResponse

DefaultCacheBehavior
