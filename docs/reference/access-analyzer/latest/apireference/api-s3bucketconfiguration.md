# S3BucketConfiguration

Proposed access control configuration for an Amazon S3 bucket. You can propose a
configuration for a new Amazon S3 bucket or an existing Amazon S3 bucket that you own by specifying
the Amazon S3 bucket policy, bucket ACLs, bucket BPA settings, Amazon S3 access points, and
multi-region access points attached to the bucket. If the configuration is for an existing
Amazon S3 bucket and you do not specify the Amazon S3 bucket policy, the access preview uses the
existing policy attached to the bucket. If the access preview is for a new resource and you
do not specify the Amazon S3 bucket policy, the access preview assumes a bucket without a
policy. To propose deletion of an existing bucket policy, you can specify an empty string.
For more information about bucket policy limits, see [Bucket Policy\
Examples](https://docs.aws.amazon.com/AmazonS3/latest/dev/example-bucket-policies.html).

## Contents

**accessPoints**

The configuration of Amazon S3 access points or multi-region access points for the bucket.
You can propose up to 10 new access points per bucket.

Type: String to [S3AccessPointConfiguration](api-s3accesspointconfiguration.md) object map

Key Pattern: `arn:[^:]*:s3:[^:]*:[^:]*:accesspoint/.*`

Required: No

**bucketAclGrants**

The proposed list of ACL grants for the Amazon S3 bucket. You can propose up to 100 ACL
grants per bucket. If the proposed grant configuration is for an existing bucket, the
access preview uses the proposed list of grant configurations in place of the existing
grants. Otherwise, the access preview uses the existing grants for the bucket.

Type: Array of [S3BucketAclGrantConfiguration](api-s3bucketaclgrantconfiguration.md) objects

Required: No

**bucketPolicy**

The proposed bucket policy for the Amazon S3 bucket.

Type: String

Required: No

**bucketPublicAccessBlock**

The proposed block public access configuration for the Amazon S3 bucket.

Type: [S3PublicAccessBlockConfiguration](api-s3publicaccessblockconfiguration.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/accessanalyzer-2019-11-01/S3BucketConfiguration)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/accessanalyzer-2019-11-01/S3BucketConfiguration)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/accessanalyzer-2019-11-01/S3BucketConfiguration)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

S3BucketAclGrantConfiguration

S3ExpressDirectoryAccessPointConfiguration
