This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::Distribution Origin

An origin.

An origin is the location where content is stored, and from which CloudFront gets content to
serve to viewers. To specify an origin:

- Use `S3OriginConfig` to specify an Amazon S3 bucket that is not
configured with static website hosting.

- Use `VpcOriginConfig` to specify a VPC origin.

- Use `CustomOriginConfig` to specify all other kinds of origins,
including:

- An Amazon S3 bucket that is configured with static website hosting

- An Elastic Load Balancing load balancer

- An AWS Elemental MediaPackage endpoint

- An AWS Elemental MediaStore container

- Any other HTTP server, running on an Amazon EC2 instance or any other kind
of host

For the current maximum number of origins that you can specify per distribution, see
[General Quotas on Web Distributions](../../../amazoncloudfront/latest/developerguide/cloudfront-limits.md#limits-web-distributions) in the
_Amazon CloudFront Developer Guide_ (quotas were formerly referred to as
limits).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConnectionAttempts" : Integer,
  "ConnectionTimeout" : Integer,
  "CustomOriginConfig" : CustomOriginConfig,
  "DomainName" : String,
  "Id" : String,
  "OriginAccessControlId" : String,
  "OriginCustomHeaders" : [ OriginCustomHeader, ... ],
  "OriginPath" : String,
  "OriginShield" : OriginShield,
  "ResponseCompletionTimeout" : Integer,
  "S3OriginConfig" : S3OriginConfig,
  "VpcOriginConfig" : VpcOriginConfig
}

```

### YAML

```yaml

  ConnectionAttempts: Integer
  ConnectionTimeout: Integer
  CustomOriginConfig:
    CustomOriginConfig
  DomainName: String
  Id: String
  OriginAccessControlId: String
  OriginCustomHeaders:
    - OriginCustomHeader
  OriginPath: String
  OriginShield:
    OriginShield
  ResponseCompletionTimeout: Integer
  S3OriginConfig:
    S3OriginConfig
  VpcOriginConfig:
    VpcOriginConfig

```

## Properties

`ConnectionAttempts`

The number of times that CloudFront attempts to connect to the origin. The minimum number is
1, the maximum is 3, and the default (if you don't specify otherwise) is 3.

For a custom origin (including an Amazon S3 bucket that's configured with static website
hosting), this value also specifies the number of times that CloudFront attempts to get a
response from the origin, in the case of an [Origin Response Timeout](../../../amazoncloudfront/latest/developerguide/distribution-web-values-specify.md#DownloadDistValuesOriginResponseTimeout).

For more information, see [Origin Connection Attempts](../../../amazoncloudfront/latest/developerguide/distribution-web-values-specify.md#origin-connection-attempts) in the
_Amazon CloudFront Developer Guide_.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConnectionTimeout`

The number of seconds that CloudFront waits when trying to establish a connection to the
origin. The minimum timeout is 1 second, the maximum is 10 seconds, and the default (if
you don't specify otherwise) is 10 seconds.

For more information, see [Origin Connection Timeout](../../../amazoncloudfront/latest/developerguide/distribution-web-values-specify.md#origin-connection-timeout) in the
_Amazon CloudFront Developer Guide_.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomOriginConfig`

Use this type to specify an origin that is not an Amazon S3 bucket, with one exception. If
the Amazon S3 bucket is configured with static website hosting, use this type. If the Amazon S3
bucket is not configured with static website hosting, use the
`S3OriginConfig` type instead.

_Required_: Conditional

_Type_: [CustomOriginConfig](aws-properties-cloudfront-distribution-customoriginconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DomainName`

The domain name for the origin.

For more information, see [Origin Domain Name](../../../amazoncloudfront/latest/developerguide/distribution-web-values-specify.md#DownloadDistValuesDomainName) in the _Amazon CloudFront Developer Guide_.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Id`

A unique identifier for the origin. This value must be unique within the
distribution.

Use this value to specify the `TargetOriginId` in a
`CacheBehavior` or `DefaultCacheBehavior`.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OriginAccessControlId`

The unique identifier of an origin access control for this origin.

For more information, see [Restricting access to an Amazon S3 origin](../../../amazoncloudfront/latest/developerguide/private-content-restricting-access-to-s3.md) in the
_Amazon CloudFront Developer Guide_.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OriginCustomHeaders`

A list of HTTP header names and values that CloudFront adds to the requests that it sends to
the origin.

For more information, see [Adding Custom Headers to Origin Requests](../../../amazoncloudfront/latest/developerguide/add-origin-custom-headers.md) in the
_Amazon CloudFront Developer Guide_.

_Required_: No

_Type_: Array of [OriginCustomHeader](aws-properties-cloudfront-distribution-origincustomheader.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OriginPath`

An optional path that CloudFront appends to the origin domain name when CloudFront requests
content from the origin.

For more information, see [Origin Path](../../../amazoncloudfront/latest/developerguide/distribution-web-values-specify.md#DownloadDistValuesOriginPath) in the
_Amazon CloudFront Developer Guide_.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OriginShield`

CloudFront Origin Shield. Using Origin Shield can help reduce the load on your
origin.

For more information, see [Using Origin Shield](../../../amazoncloudfront/latest/developerguide/origin-shield.md) in the _Amazon CloudFront Developer Guide_.

_Required_: No

_Type_: [OriginShield](aws-properties-cloudfront-distribution-originshield.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResponseCompletionTimeout`

The time (in seconds) that a request from CloudFront to the origin can stay open and wait
for a response. If the complete response isn't received from the origin by this time,
CloudFront ends the connection.

The value for `ResponseCompletionTimeout` must be equal to or greater than
the value for `OriginReadTimeout`. If you don't set a value for
`ResponseCompletionTimeout`, CloudFront doesn't enforce a maximum value.

For more information, see [Response completion timeout](../../../amazoncloudfront/latest/developerguide/downloaddistvaluesorigin.md#response-completion-timeout) in the
_Amazon CloudFront Developer Guide_.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3OriginConfig`

Use this type to specify an origin that is an Amazon S3 bucket that is not configured with
static website hosting. To specify any other type of origin, including an Amazon S3 bucket
that is configured with static website hosting, use the `CustomOriginConfig`
type instead.

_Required_: Conditional

_Type_: [S3OriginConfig](aws-properties-cloudfront-distribution-s3originconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcOriginConfig`

The VPC origin configuration.

_Required_: No

_Type_: [VpcOriginConfig](aws-properties-cloudfront-distribution-vpcoriginconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [Origin](../../../../reference/cloudfront/latest/apireference/api-origin.md) in the _Amazon CloudFront API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Logging

OriginCustomHeader

All content copied from https://docs.aws.amazon.com/.
