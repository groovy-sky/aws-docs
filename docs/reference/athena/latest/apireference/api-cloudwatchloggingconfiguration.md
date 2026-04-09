# CloudWatchLoggingConfiguration

Configuration settings for delivering logs to Amazon CloudWatch log groups.

## Contents

**Enabled**

Enables CloudWatch logging.

Type: Boolean

Required: Yes

**LogGroup**

The name of the log group in Amazon CloudWatch Logs where you want to publish
your logs.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 512.

Pattern: `^[a-zA-Z0-9._/-]+$`

Required: No

**LogStreamNamePrefix**

Prefix for the CloudWatch log stream name.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 512.

Pattern: `^[^:*]*$`

Required: No

**LogTypes**

The types of logs that you want to publish to CloudWatch.

Type: String to array of strings map

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/athena-2017-05-18/cloudwatchloggingconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/athena-2017-05-18/cloudwatchloggingconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/athena-2017-05-18/cloudwatchloggingconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Classification

Column

All content copied from https://docs.aws.amazon.com/.
