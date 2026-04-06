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

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/athena-2017-05-18/CloudWatchLoggingConfiguration)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/athena-2017-05-18/CloudWatchLoggingConfiguration)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/athena-2017-05-18/CloudWatchLoggingConfiguration)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Classification

Column
