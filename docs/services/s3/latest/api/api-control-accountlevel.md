# AccountLevel

A container element for the account-level Amazon S3 Storage Lens configuration.

###### Note

You must enable Storage Lens metrics consistently at both the account level and bucket level, or your request will fail.

For more information about S3 Storage Lens, see [Assessing your storage activity and usage with S3 Storage Lens](../userguide/storage-lens.md) in the _Amazon S3 User Guide_. For a complete list of S3 Storage Lens metrics, see [S3 Storage Lens metrics glossary](../userguide/storage-lens-metrics-glossary.md) in the _Amazon S3 User Guide_.

## Contents

**BucketLevel**

A container element for the S3 Storage Lens bucket-level configuration.

Type: [BucketLevel](https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_BucketLevel.html) data type

Required: Yes

**ActivityMetrics**

A container element for S3 Storage Lens activity metrics.

Type: [ActivityMetrics](https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_ActivityMetrics.html) data type

Required: No

**AdvancedCostOptimizationMetrics**

A container element for S3 Storage Lens advanced cost-optimization metrics.

Type: [AdvancedCostOptimizationMetrics](https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_AdvancedCostOptimizationMetrics.html) data type

Required: No

**AdvancedDataProtectionMetrics**

A container element for S3 Storage Lens advanced data-protection metrics.

Type: [AdvancedDataProtectionMetrics](https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_AdvancedDataProtectionMetrics.html) data type

Required: No

**AdvancedPerformanceMetrics**

A container element for S3 Storage Lens advanced performance metrics.

Type: [AdvancedPerformanceMetrics](https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_AdvancedPerformanceMetrics.html) data type

Required: No

**DetailedStatusCodesMetrics**

A container element for detailed status code metrics.

Type: [DetailedStatusCodesMetrics](https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DetailedStatusCodesMetrics.html) data type

Required: No

**StorageLensGroupLevel**

A container element for S3 Storage Lens groups metrics.

Type: [StorageLensGroupLevel](https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_StorageLensGroupLevel.html) data type

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/s3control-2018-08-20/AccountLevel)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/s3control-2018-08-20/AccountLevel)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/s3control-2018-08-20/AccountLevel)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AccessPoint

ActivityMetrics
