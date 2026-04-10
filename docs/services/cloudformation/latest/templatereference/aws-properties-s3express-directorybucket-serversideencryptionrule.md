This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3Express::DirectoryBucket ServerSideEncryptionRule

Specifies the default server-side encryption configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BucketKeyEnabled" : Boolean,
  "ServerSideEncryptionByDefault" : ServerSideEncryptionByDefault
}

```

### YAML

```yaml

  BucketKeyEnabled: Boolean
  ServerSideEncryptionByDefault:
    ServerSideEncryptionByDefault

```

## Properties

`BucketKeyEnabled`

Specifies whether Amazon S3 should use an S3 Bucket Key with server-side encryption
using KMS (SSE-KMS) for new objects in the bucket. S3 Bucket Keys are always enabled for
`GET` and `PUT` operations on a directory bucket and can’t be
disabled. It's only allowed to set the `BucketKeyEnabled` element to
`true`.

S3 Bucket Keys aren't supported, when you copy SSE-KMS encrypted objects from general
purpose buckets to directory buckets, from directory buckets to general purpose buckets, or
between directory buckets, through [CopyObject](../../../s3/latest/api/api-copyobject.md), [UploadPartCopy](../../../s3/latest/api/api-uploadpartcopy.md), [the Copy operation\
in Batch Operations](../../../s3/latest/userguide/directory-buckets-objects-batch-ops.md), or [the import jobs](../../../s3/latest/userguide/create-import-job.md). In this case,
Amazon S3 makes a call to AWS KMS every time a copy request is made for a
KMS-encrypted object.

For more information, see [Amazon S3 Bucket Keys](../../../s3/latest/userguide/s3-express-usingkmsencryption.md#s3-express-sse-kms-bucket-keys) in the _Amazon S3 User Guide_.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServerSideEncryptionByDefault`

Specifies the default server-side encryption to apply to new objects in the bucket. If a PUT Object
request doesn't specify any server-side encryption, this default encryption will be applied.

_Required_: No

_Type_: [ServerSideEncryptionByDefault](aws-properties-s3express-directorybucket-serversideencryptionbydefault.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ServerSideEncryptionByDefault

Tag

All content copied from https://docs.aws.amazon.com/.
