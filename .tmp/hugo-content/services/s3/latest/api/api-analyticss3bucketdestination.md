# AnalyticsS3BucketDestination

Contains information about where to publish the analytics results.

## Contents

**Bucket**

The Amazon Resource Name (ARN) of the bucket to which data is exported.

Type: String

Required: Yes

**Format**

Specifies the file format used when exporting data to Amazon S3.

Type: String

Valid Values: `CSV`

Required: Yes

**BucketAccountId**

The account ID that owns the destination S3 bucket. If no account ID is provided, the owner is not
validated before exporting data.

###### Note

Although this value is optional, we strongly recommend that you set it to help prevent problems
if the destination bucket ownership changes.

Type: String

Required: No

**Prefix**

The prefix to use when exporting data. The prefix is prepended to all results.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3-2006-03-01/analyticss3bucketdestination.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3-2006-03-01/analyticss3bucketdestination.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3-2006-03-01/analyticss3bucketdestination.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AnalyticsFilter

BlockedEncryptionTypes

All content copied from https://docs.aws.amazon.com/.
