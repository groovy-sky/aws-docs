# VerifiedAccessLogOptions

Options for Verified Access logs.

## Contents

**CloudWatchLogs**

Sends Verified Access logs to CloudWatch Logs.

Type: [VerifiedAccessLogCloudWatchLogsDestinationOptions](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_VerifiedAccessLogCloudWatchLogsDestinationOptions.html) object

Required: No

**IncludeTrustContext**

Indicates whether to include trust data sent by trust providers in the logs.

Type: Boolean

Required: No

**KinesisDataFirehose**

Sends Verified Access logs to Kinesis.

Type: [VerifiedAccessLogKinesisDataFirehoseDestinationOptions](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_VerifiedAccessLogKinesisDataFirehoseDestinationOptions.html) object

Required: No

**LogVersion**

The logging version.

Valid values: `ocsf-0.1` \| `ocsf-1.0.0-rc.2`

Type: String

Required: No

**S3**

Sends Verified Access logs to Amazon S3.

Type: [VerifiedAccessLogS3DestinationOptions](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_VerifiedAccessLogS3DestinationOptions.html) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/VerifiedAccessLogOptions)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/VerifiedAccessLogOptions)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/VerifiedAccessLogOptions)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

VerifiedAccessLogKinesisDataFirehoseDestinationOptions

VerifiedAccessLogs
