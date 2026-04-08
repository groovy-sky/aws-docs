# S3ExpressDirectoryBucketConfiguration

Proposed access control configuration for an Amazon S3 directory bucket. You can propose a
configuration for a new Amazon S3 directory bucket or an existing Amazon S3 directory bucket that you
own by specifying the Amazon S3 bucket policy. If the configuration is for an existing Amazon S3
directory bucket and you do not specify the Amazon S3 bucket policy, the access preview uses the
existing policy attached to the directory bucket. If the access preview is for a new
resource and you do not specify the Amazon S3 bucket policy, the access preview assumes an
directory bucket without a policy. To propose deletion of an existing bucket policy, you
can specify an empty string. For more information about Amazon S3 directory bucket policies, see
[Example bucket policies for directory buckets](../../../../services/s3/latest/userguide/s3-express-security-iam-example-bucket-policies.md) in the Amazon Simple Storage Service User
Guide.

## Contents

**accessPoints**

The proposed access points for the Amazon S3 directory bucket.

Type: String to [S3ExpressDirectoryAccessPointConfiguration](api-s3expressdirectoryaccesspointconfiguration.md) object map

Key Pattern: `arn:[^:]*:s3express:[^:]*:[^:]*:accesspoint/.*`

Required: No

**bucketPolicy**

The proposed bucket policy for the Amazon S3 directory bucket.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/accessanalyzer-2019-11-01/s3expressdirectorybucketconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/accessanalyzer-2019-11-01/s3expressdirectorybucketconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/accessanalyzer-2019-11-01/s3expressdirectorybucketconfiguration.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

S3ExpressDirectoryAccessPointConfiguration

S3PublicAccessBlockConfiguration
