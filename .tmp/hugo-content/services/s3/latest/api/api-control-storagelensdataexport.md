# StorageLensDataExport

A container to specify the properties of your S3 Storage Lens metrics export, including the
destination, schema, and format.

## Contents

**CloudWatchMetrics**

A container for enabling Amazon CloudWatch publishing for S3 Storage Lens metrics.

Type: [CloudWatchMetrics](api-control-cloudwatchmetrics.md) data type

Required: No

**S3BucketDestination**

A container for the bucket where the S3 Storage Lens metrics export will be located.

###### Note

This bucket must be located in the same Region as the storage lens configuration.

Type: [S3BucketDestination](api-control-s3bucketdestination.md) data type

Required: No

**StorageLensTableDestination**

A container for configuring S3 Storage Lens data exports to read-only S3 table buckets.

Type: [StorageLensTableDestination](api-control-storagelenstabledestination.md) data type

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/storagelensdataexport.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/storagelensdataexport.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/storagelensdataexport.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StorageLensConfiguration

StorageLensDataExportEncryption

All content copied from https://docs.aws.amazon.com/.
