# TableRecordExpirationConfigurationValue

The expiration configuration settings for records in a table, and the status of the configuration. If the status of the configuration is enabled, records expire and are automatically removed after the number of days specified in the record expiration settings for the table.

## Contents

**settings**

The expiration settings for records in the table.

Type: [TableRecordExpirationSettings](https://docs.aws.amazon.com/AmazonS3/latest/API/API_s3Buckets_TableRecordExpirationSettings.html) object

Required: No

**status**

The status of the expiration settings for records in the table.

Type: String

Valid Values: `enabled | disabled`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/s3tables-2018-05-10/TableRecordExpirationConfigurationValue)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/s3tables-2018-05-10/TableRecordExpirationConfigurationValue)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/s3tables-2018-05-10/TableRecordExpirationConfigurationValue)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TableMetadata

TableRecordExpirationJobMetrics
