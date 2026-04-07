# TableMaintenanceSettings

Contains details about maintenance settings for the table.

## Contents

###### Important

This data type is a UNION, so only one of the following members can be specified when used or returned.

**icebergCompaction**

Contains details about the Iceberg compaction settings for the table.

Type: [IcebergCompactionSettings](https://docs.aws.amazon.com/AmazonS3/latest/API/API_s3Buckets_IcebergCompactionSettings.html) object

Required: No

**icebergSnapshotManagement**

Contains details about the Iceberg snapshot management settings for the table.

Type: [IcebergSnapshotManagementSettings](https://docs.aws.amazon.com/AmazonS3/latest/API/API_s3Buckets_IcebergSnapshotManagementSettings.html) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/s3tables-2018-05-10/TableMaintenanceSettings)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/s3tables-2018-05-10/TableMaintenanceSettings)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/s3tables-2018-05-10/TableMaintenanceSettings)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TableMaintenanceJobStatusValue

TableMetadata
