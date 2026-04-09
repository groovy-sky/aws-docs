# TableRecordExpirationConfigurationValue

The expiration configuration settings for records in a table, and the status of the configuration. If the status of the configuration is enabled, records expire and are automatically removed after the number of days specified in the record expiration settings for the table.

## Contents

**settings**

The expiration settings for records in the table.

Type: [TableRecordExpirationSettings](api-s3buckets-tablerecordexpirationsettings.md) object

Required: No

**status**

The status of the expiration settings for records in the table.

Type: String

Valid Values: `enabled | disabled`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3tables-2018-05-10/tablerecordexpirationconfigurationvalue.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3tables-2018-05-10/tablerecordexpirationconfigurationvalue.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3tables-2018-05-10/tablerecordexpirationconfigurationvalue.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TableMetadata

TableRecordExpirationJobMetrics

All content copied from https://docs.aws.amazon.com/.
