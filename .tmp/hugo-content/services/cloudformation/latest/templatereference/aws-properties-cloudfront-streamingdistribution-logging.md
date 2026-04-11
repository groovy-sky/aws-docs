This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::StreamingDistribution Logging

A complex type that controls whether access logs are written for the streaming
distribution.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Bucket" : String,
  "Enabled" : Boolean,
  "Prefix" : String
}

```

### YAML

```yaml

  Bucket: String
  Enabled: Boolean
  Prefix: String

```

## Properties

`Bucket`

The Amazon S3 bucket to store the access logs in, for example,
`amzn-s3-demo-bucket.s3.amazonaws.com`.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Enabled`

Specifies whether you want CloudFront to save access logs to an Amazon S3 bucket. If you don't
want to enable logging when you create a streaming distribution or if you want to
disable logging for an existing streaming distribution, specify `false` for
`Enabled`, and specify `empty Bucket` and `Prefix`
elements. If you specify `false` for `Enabled` but you specify
values for `Bucket` and `Prefix`, the values are automatically
deleted.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Prefix`

An optional string that you want CloudFront to prefix to the access log filenames for this
streaming distribution, for example, `myprefix/`. If you want to enable
logging, but you don't want to specify a prefix, you still must include an empty
`Prefix` element in the `Logging` element.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [StreamingLoggingConfig](../../../../reference/cloudfront/latest/apireference/api-streamingloggingconfig.md) in the _Amazon CloudFront API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::CloudFront::StreamingDistribution

S3Origin

All content copied from https://docs.aws.amazon.com/.
