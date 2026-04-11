# ServerSideEncryptionRule

Specifies the default server-side encryption configuration.

###### Note

- **General purpose buckets** \- If you're specifying a customer
managed KMS key, we recommend using a fully qualified KMS key ARN. If you use a KMS key
alias instead, then AWS KMS resolves the key within the requester’s account. This behavior can
result in data that's encrypted with a KMS key that belongs to the requester, and not the bucket
owner.

- **Directory buckets** -
When you specify an [AWS KMS customer managed key](../../../kms/latest/developerguide/concepts.md#customer-cmk) for encryption in your directory bucket, only use the key ID or key ARN. The key alias format of the KMS key isn't supported.

## Contents

**ApplyServerSideEncryptionByDefault**

Specifies the default server-side encryption to apply to new objects in the bucket. If a PUT Object
request doesn't specify any server-side encryption, this default encryption will be applied.

Type: [ServerSideEncryptionByDefault](api-serversideencryptionbydefault.md) data type

Required: No

**BlockedEncryptionTypes**

A bucket-level setting for Amazon S3 general purpose buckets used to prevent the upload of new objects encrypted with the specified server-side encryption type. For example, blocking an encryption type will block `PutObject`, `CopyObject`, `PostObject`, multipart upload, and replication requests to the bucket for objects with the specified encryption type. However, you can continue to read and list any pre-existing objects already encrypted with the specified encryption type. For more information, see [Blocking or unblocking SSE-C for a general purpose bucket](../userguide/blocking-unblocking-s3-c-encryption-gpb.md).

###### Note

Currently, this parameter only supports blocking or unblocking server-side encryption with customer-provided keys (SSE-C). For more information about SSE-C, see [Using server-side encryption with customer-provided keys (SSE-C)](../userguide/serversideencryptioncustomerkeys.md).

Type: [BlockedEncryptionTypes](api-blockedencryptiontypes.md) data type

Required: No

**BucketKeyEnabled**

Specifies whether Amazon S3 should use an S3 Bucket Key with server-side encryption using KMS (SSE-KMS)
for new objects in the bucket. Existing objects are not affected. Setting the
`BucketKeyEnabled` element to `true` causes Amazon S3 to use an S3 Bucket Key.

###### Note

- **General purpose buckets** \- By default, S3 Bucket Key is not
enabled. For more information, see [Amazon S3 Bucket Keys](../dev/bucket-key.md) in the
_Amazon S3 User Guide_.

- **Directory buckets** -
S3 Bucket Keys are always enabled for `GET` and `PUT` operations in a directory bucket and can’t be disabled. S3 Bucket Keys aren't supported, when you copy SSE-KMS encrypted objects from general purpose buckets
to directory buckets, from directory buckets to general purpose buckets, or between directory buckets, through [CopyObject](api-copyobject.md), [UploadPartCopy](api-uploadpartcopy.md), [the Copy operation in Batch Operations](../userguide/directory-buckets-objects-batch-ops.md), or
[the import jobs](../userguide/create-import-job.md). In this case, Amazon S3 makes a call to AWS KMS every time a copy request is made for a KMS-encrypted object.

Type: Boolean

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3-2006-03-01/serversideencryptionrule.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3-2006-03-01/serversideencryptionrule.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3-2006-03-01/serversideencryptionrule.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ServerSideEncryptionConfiguration

SessionCredentials

All content copied from https://docs.aws.amazon.com/.
