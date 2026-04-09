# BucketInfo

Specifies the information about the bucket that will be created. For more information about
directory buckets, see [Directory buckets](../userguide/directory-buckets-overview.md) in the
_Amazon S3 User Guide_.

###### Note

This functionality is only supported by directory buckets.

## Contents

**DataRedundancy**

The number of Zone (Availability Zone or Local Zone) that's used for redundancy for the bucket.

Type: String

Valid Values: `SingleAvailabilityZone | SingleLocalZone`

Required: No

**Type**

The type of bucket.

Type: String

Valid Values: `Directory`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3-2006-03-01/bucketinfo.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3-2006-03-01/bucketinfo.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3-2006-03-01/bucketinfo.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Bucket

BucketLifecycleConfiguration

All content copied from https://docs.aws.amazon.com/.
