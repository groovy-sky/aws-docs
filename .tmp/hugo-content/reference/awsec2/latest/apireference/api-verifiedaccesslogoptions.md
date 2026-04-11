# VerifiedAccessLogOptions

Options for Verified Access logs.

## Contents

**CloudWatchLogs**

Sends Verified Access logs to CloudWatch Logs.

Type: [VerifiedAccessLogCloudWatchLogsDestinationOptions](api-verifiedaccesslogcloudwatchlogsdestinationoptions.md) object

Required: No

**IncludeTrustContext**

Indicates whether to include trust data sent by trust providers in the logs.

Type: Boolean

Required: No

**KinesisDataFirehose**

Sends Verified Access logs to Kinesis.

Type: [VerifiedAccessLogKinesisDataFirehoseDestinationOptions](api-verifiedaccesslogkinesisdatafirehosedestinationoptions.md) object

Required: No

**LogVersion**

The logging version.

Valid values: `ocsf-0.1` \| `ocsf-1.0.0-rc.2`

Type: String

Required: No

**S3**

Sends Verified Access logs to Amazon S3.

Type: [VerifiedAccessLogS3DestinationOptions](api-verifiedaccesslogs3destinationoptions.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/verifiedaccesslogoptions.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/verifiedaccesslogoptions.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/verifiedaccesslogoptions.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

VerifiedAccessLogKinesisDataFirehoseDestinationOptions

VerifiedAccessLogs
