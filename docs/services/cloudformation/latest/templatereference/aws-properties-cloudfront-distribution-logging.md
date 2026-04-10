This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::Distribution Logging

A complex type that specifies whether access logs are written for the distribution.

###### Note

If you already enabled standard logging (legacy) and you want to enable standard logging
(v2) to send your access logs to Amazon S3, we recommend that you specify a
_different_ Amazon S3 bucket or use a _separate_
_path_ in the same bucket (for example, use a log prefix or
partitioning). This helps you keep track of which log files are associated with
which logging subscription and prevents log files from overwriting each other. For
more information, see [Standard logging (access logs)](../../../amazoncloudfront/latest/developerguide/accesslogs.md) in the
_Amazon CloudFront Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Bucket" : String,
  "IncludeCookies" : Boolean,
  "Prefix" : String
}

```

### YAML

```yaml

  Bucket: String
  IncludeCookies: Boolean
  Prefix: String

```

## Properties

`Bucket`

The Amazon S3 bucket to store the access logs in, for example,
`amzn-s3-demo-bucket.s3.amazonaws.com`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncludeCookies`

Specifies whether you want CloudFront to include cookies in access logs, specify
`true` for `IncludeCookies`. If you choose to include cookies
in logs, CloudFront logs all cookies regardless of how you configure the cache behaviors for
this distribution. If you don't want to include cookies when you create a distribution
or if you want to disable include cookies for an existing distribution, specify
`false` for `IncludeCookies`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Prefix`

An optional string that you want CloudFront to prefix to the access log
`filenames` for this distribution, for example, `myprefix/`.
If you want to enable logging, but you don't want to specify a prefix, you still must
include an empty `Prefix` element in the `Logging` element.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [LoggingConfig](../../../../reference/cloudfront/latest/apireference/api-loggingconfig.md) in the _Amazon CloudFront API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LegacyS3Origin

Origin

All content copied from https://docs.aws.amazon.com/.
