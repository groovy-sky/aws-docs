# IcebergCompactionSettings

Contains details about the compaction settings for an Iceberg table.

## Contents

**strategy**

The compaction strategy to use for the table. This determines how files are selected and combined during compaction operations.

Type: String

Valid Values: `auto | binpack | sort | z-order`

Required: No

**targetFileSizeMB**

The target file size for the table in MB.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 2147483647.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3tables-2018-05-10/icebergcompactionsettings.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3tables-2018-05-10/icebergcompactionsettings.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3tables-2018-05-10/icebergcompactionsettings.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EncryptionConfiguration

IcebergMetadata

All content copied from https://docs.aws.amazon.com/.
