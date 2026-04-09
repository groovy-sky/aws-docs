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
To find the values supported for the suffix path for each log source, use the [DescribeConfigurationTemplates](api-describeconfigurationtemplates.md) operation and check the
`allowedSuffixPathFields` field in the response.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/s3deliveryconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/s3deliveryconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/s3deliveryconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

S3Configuration

S3TableIntegrationSource

All content copied from https://docs.aws.amazon.com/.
