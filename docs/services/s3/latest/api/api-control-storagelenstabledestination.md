# StorageLensTableDestination

A container for configuring your S3 Storage Lens reports to export to read-only S3 table
buckets. This parameter enables you to store your Storage Lens metrics in a structured,
queryable table format in Apache Iceberg.

For more information about S3 Storage Lens, see [Assessing your storage activity and usage with S3 Storage Lens](../userguide/storage-lens.md) in the _Amazon S3 User Guide_.

## Contents

**IsEnabled**

A container that indicates whether the export to read-only S3 table buckets is enabled
for your S3 Storage Lens configuration. When set to true, Storage Lens reports are automatically
exported to tables in addition to other configured destinations.

Type: Boolean

Required: Yes

**Encryption**

A container for the encryption of the S3 Storage Lens metrics exports.

Type: [StorageLensDataExportEncryption](api-control-storagelensdataexportencryption.md) data type

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/s3control-2018-08-20/StorageLensTableDestination)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/s3control-2018-08-20/StorageLensTableDestination)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/s3control-2018-08-20/StorageLensTableDestination)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

StorageLensGroupOrOperator

StorageLensTag
