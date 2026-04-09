# StorageLensExpandedPrefixesDataExport

A container for your S3 Storage Lens expanded prefix metrics report configuration. Unlike the
default Storage Lens metrics report, the enhanced prefix metrics report includes all
S3 Storage Lens storage and activity data related to the full list of prefixes in your Storage
Lens configuration.

## Contents

**S3BucketDestination**

A container for the bucket where the Amazon S3 Storage Lens metrics export files are
located.

Type: [S3BucketDestination](api-control-s3bucketdestination.md) data type

Required: No

**StorageLensTableDestination**

A container for the bucket where the S3 Storage Lens metric export files are located. At least one export destination must be specified.

Type: [StorageLensTableDestination](api-control-storagelenstabledestination.md) data type

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/storagelensexpandedprefixesdataexport.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/storagelensexpandedprefixesdataexport.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/storagelensexpandedprefixesdataexport.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StorageLensDataExportEncryption

StorageLensGroup

All content copied from https://docs.aws.amazon.com/.
