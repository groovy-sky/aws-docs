# RegionReport

A combination of a bucket and Region that's part of a Multi-Region Access Point.

## Contents

**Bucket**

The name of the bucket.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 255.

Required: No

**BucketAccountId**

The AWS account ID that owns the Amazon S3 bucket that's associated with this
Multi-Region Access Point.

Type: String

Length Constraints: Maximum length of 64.

Pattern: `^\d{12}$`

Required: No

**Region**

The name of the Region.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/regionreport.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/regionreport.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/regionreport.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RegionalBucket

ReplicaModifications

All content copied from https://docs.aws.amazon.com/.
