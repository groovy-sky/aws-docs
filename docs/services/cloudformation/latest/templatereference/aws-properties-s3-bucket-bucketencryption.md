This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::Bucket BucketEncryption

Specifies default encryption for a bucket using server-side encryption with Amazon
S3-managed keys (SSE-S3), AWS KMS-managed keys (SSE-KMS), or dual-layer server-side encryption with KMS-managed keys (DSSE-KMS). For
information about the Amazon S3 default encryption feature, see [Amazon S3 Default Encryption for S3\
Buckets](https://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-encryption.html) in the _Amazon S3 User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ServerSideEncryptionConfiguration" : [ ServerSideEncryptionRule, ... ]
}

```

### YAML

```yaml

  ServerSideEncryptionConfiguration:
    - ServerSideEncryptionRule

```

## Properties

`ServerSideEncryptionConfiguration`

Specifies the default server-side-encryption configuration.

_Required_: Yes

_Type_: Array of [ServerSideEncryptionRule](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-s3-bucket-serversideencryptionrule.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- AWS::S3::Bucket [Examples](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#aws-properties-s3-bucket--examples)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

BlockedEncryptionTypes

CorsConfiguration
