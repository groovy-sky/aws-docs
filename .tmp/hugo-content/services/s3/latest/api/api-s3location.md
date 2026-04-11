# S3Location

Describes an Amazon S3 location that will receive the results of the restore request.

## Contents

**BucketName**

The name of the bucket where the restore results will be placed.

Type: String

Required: Yes

**Prefix**

The prefix that is prepended to the restore results for this request.

Type: String

Required: Yes

**AccessControlList**

A list of grants that control access to the staged results.

Type: Array of [Grant](api-grant.md) data types

Required: No

**CannedACL**

The canned ACL to apply to the restore results.

Type: String

Valid Values: `private | public-read | public-read-write | authenticated-read | aws-exec-read | bucket-owner-read | bucket-owner-full-control`

Required: No

**Encryption**

Contains the type of server-side encryption used.

Type: [Encryption](api-encryption.md) data type

Required: No

**StorageClass**

The class of storage used to store the restore results.

Type: String

Valid Values: `STANDARD | REDUCED_REDUNDANCY | STANDARD_IA | ONEZONE_IA | INTELLIGENT_TIERING | GLACIER | DEEP_ARCHIVE | OUTPOSTS | GLACIER_IR | SNOW | EXPRESS_ONEZONE | FSX_OPENZFS | FSX_ONTAP`

Required: No

**Tagging**

The tag-set that is applied to the restore results.

Type: [Tagging](api-tagging.md) data type

Required: No

**UserMetadata**

A list of metadata to store with the restore results in S3.

Type: Array of [MetadataEntry](api-metadataentry.md) data types

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3-2006-03-01/s3location.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3-2006-03-01/s3location.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3-2006-03-01/s3location.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

S3KeyFilter

S3TablesDestination

All content copied from https://docs.aws.amazon.com/.
