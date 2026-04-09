# Scope

You can use the access point scope to restrict access to specific prefixes, API operations, or a combination of both.

For more information, see [Manage the scope of your access points for directory buckets](../userguide/access-points-directory-buckets-manage-scope.md).

## Contents

**Permissions**

You can include one or more API operations as permissions.

Type: Array of strings

Valid Values: `GetObject | GetObjectAttributes | ListMultipartUploadParts | ListBucket | ListBucketMultipartUploads | PutObject | DeleteObject | AbortMultipartUpload`

Required: No

**Prefixes**

You can specify any amount of prefixes, but the total length of characters of all prefixes must be less than 256 bytes in size.

Type: Array of strings

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/scope.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/scope.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/scope.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

S3UpdateObjectEncryptionSSEKMS

SelectionCriteria

All content copied from https://docs.aws.amazon.com/.
