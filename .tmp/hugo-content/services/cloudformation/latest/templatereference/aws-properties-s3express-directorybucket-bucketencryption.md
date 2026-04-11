This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3Express::DirectoryBucket BucketEncryption

Specifies default encryption for a bucket using server-side encryption with Amazon S3
managed keys (SSE-S3) or AWS KMS keys (SSE-KMS). For information about
default encryption for directory buckets, see [Setting and monitoring\
default encryption for directory buckets](../../../s3/latest/userguide/s3-express-bucket-encryption.md) in the _Amazon S3 User_
_Guide_.

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

_Type_: Array of [ServerSideEncryptionRule](aws-properties-s3express-directorybucket-serversideencryptionrule.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AbortIncompleteMultipartUpload

LifecycleConfiguration

All content copied from https://docs.aws.amazon.com/.
