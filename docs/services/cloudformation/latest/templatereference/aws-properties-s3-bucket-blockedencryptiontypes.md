This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::Bucket BlockedEncryptionTypes

A bucket-level setting for Amazon S3 general purpose buckets used to prevent the upload of new objects encrypted with the specified server-side encryption type. For example, blocking an encryption type will block `PutObject`, `CopyObject`, `PostObject`, multipart upload, and replication requests to the bucket for objects with the specified encryption type. However, you can continue to read and list any pre-existing objects already encrypted with the specified encryption type. For more information, see [Blocking or unblocking SSE-C for a general purpose bucket](../../../s3/latest/userguide/blocking-unblocking-s3-c-encryption-gpb.md).

This data type is used with the following actions:

- [PutBucketEncryption](../../../s3/latest/api/api-putbucketencryption.md)

- [GetBucketEncryption](../../../s3/latest/api/api-getbucketencryption.md)

- [DeleteBucketEncryption](../../../s3/latest/api/api-deletebucketencryption.md)

Permissions

You must have the `s3:PutEncryptionConfiguration` permission to block or unblock an encryption type for a bucket.

You must have the `s3:GetEncryptionConfiguration` permission to view a bucket's encryption type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EncryptionType" : [ String, ... ]
}

```

### YAML

```yaml

  EncryptionType:
    - String

```

## Properties

`EncryptionType`

The object encryption type that you want to block or unblock for an Amazon S3 general purpose bucket.

###### Note

Currently, this parameter only supports blocking or unblocking server side encryption with customer-provided keys (SSE-C). For more information about SSE-C, see [Using server-side encryption with customer-provided keys (SSE-C)](../../../s3/latest/userguide/serversideencryptioncustomerkeys.md).

_Required_: No

_Type_: Array of String

_Allowed values_: `NONE | SSE-C`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AnalyticsConfiguration

BucketEncryption

All content copied from https://docs.aws.amazon.com/.
