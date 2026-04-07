# TableMaintenanceConfigurationValue

The values that define a maintenance configuration for a table.

## Contents

**settings**

Contains details about the settings for the maintenance configuration.

Type: [TableMaintenanceSettings](https://docs.aws.amazon.com/AmazonS3/latest/API/API_s3Buckets_TableMaintenanceSettings.html) object

**Note:** This object is a Union. Only one member of this object can be specified or returned.

Required: No

**status**

The status of the maintenance configuration.

Type: String

Valid Values: `enabled | disabled`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/s3tables-2018-05-10/TableMaintenanceConfigurationValue)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/s3tables-2018-05-10/TableMaintenanceConfigurationValue)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/s3tables-2018-05-10/TableMaintenanceConfigurationValue)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TableBucketSummary

TableMaintenanceJobStatusValue
