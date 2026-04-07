# StorageLensDataExport

A container to specify the properties of your S3 Storage Lens metrics export, including the
destination, schema, and format.

## Contents

**CloudWatchMetrics**

A container for enabling Amazon CloudWatch publishing for S3 Storage Lens metrics.

Type: [CloudWatchMetrics](https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CloudWatchMetrics.html) data type

Required: No

**S3BucketDestination**

A container for the bucket where the S3 Storage Lens metrics export will be located.

###### Note

This bucket must be located in the same Region as the storage lens configuration.

Type: [S3BucketDestination](https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_S3BucketDestination.html) data type

Required: No

**StorageLensTableDestination**

A container for configuring S3 Storage Lens data exports to read-only S3 table buckets.

Type: [StorageLensTableDestination](https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_StorageLensTableDestination.html) data type

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/s3control-2018-08-20/StorageLensDataExport)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/s3control-2018-08-20/StorageLensDataExport)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/s3control-2018-08-20/StorageLensDataExport)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

StorageLensConfiguration

StorageLensDataExportEncryption
