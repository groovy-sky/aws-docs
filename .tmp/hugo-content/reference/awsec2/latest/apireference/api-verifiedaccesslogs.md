# VerifiedAccessLogs

Describes the options for Verified Access logs.

## Contents

**cloudWatchLogs**

CloudWatch Logs logging destination.

Type: [VerifiedAccessLogCloudWatchLogsDestination](api-verifiedaccesslogcloudwatchlogsdestination.md) object

Required: No

**includeTrustContext**

Indicates whether trust data is included in the logs.

Type: Boolean

Required: No

**kinesisDataFirehose**

Kinesis logging destination.

Type: [VerifiedAccessLogKinesisDataFirehoseDestination](api-verifiedaccesslogkinesisdatafirehosedestination.md) object

Required: No

**logVersion**

The log version.

Type: String

Required: No

**s3**

Amazon S3 logging options.

Type: [VerifiedAccessLogS3Destination](api-verifiedaccesslogs3destination.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/verifiedaccesslogs.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/verifiedaccesslogs.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/verifiedaccesslogs.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

VerifiedAccessLogOptions

VerifiedAccessLogS3Destination
