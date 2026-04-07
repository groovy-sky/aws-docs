# S3DeliveryConfiguration

This structure contains delivery configurations that apply only when the delivery
destination resource is an S3 bucket.

## Contents

**enableHiveCompatiblePath**

This parameter causes the S3 objects that contain delivered logs to use a prefix structure
that allows for integration with Apache Hive.

Type: Boolean

Required: No

**suffixPath**

This string allows re-configuring the S3 object prefix to contain either static or
variable sections. The valid variables to use in the suffix path will vary by each log source.
To find the values supported for the suffix path for each log source, use the [DescribeConfigurationTemplates](https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_DescribeConfigurationTemplates.html) operation and check the
`allowedSuffixPathFields` field in the response.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/logs-2014-03-28/S3DeliveryConfiguration)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/logs-2014-03-28/S3DeliveryConfiguration)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/logs-2014-03-28/S3DeliveryConfiguration)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

S3Configuration

S3TableIntegrationSource
