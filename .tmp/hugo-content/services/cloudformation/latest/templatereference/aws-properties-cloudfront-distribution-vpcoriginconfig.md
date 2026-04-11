This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::Distribution VpcOriginConfig

An Amazon CloudFront VPC origin configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "OriginKeepaliveTimeout" : Integer,
  "OriginReadTimeout" : Integer,
  "OwnerAccountId" : String,
  "VpcOriginId" : String
}

```

### YAML

```yaml

  OriginKeepaliveTimeout: Integer
  OriginReadTimeout: Integer
  OwnerAccountId: String
  VpcOriginId: String

```

## Properties

`OriginKeepaliveTimeout`

Specifies how long, in seconds, CloudFront persists its connection to the origin. The
minimum timeout is 1 second, the maximum is 120 seconds, and the default (if you don't
specify otherwise) is 5 seconds.

For more information, see [Keep-alive timeout (custom origins only)](../../../amazoncloudfront/latest/developerguide/downloaddistvaluesorigin.md#DownloadDistValuesOriginKeepaliveTimeout) in the
_Amazon CloudFront Developer Guide_.

_Required_: No

_Type_: Integer

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

`OwnerAccountId`

The account ID of the AWS account that owns the VPC origin.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcOriginId`

The VPC origin ID.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ViewerMtlsConfig

AWS::CloudFront::DistributionTenant

All content copied from https://docs.aws.amazon.com/.
