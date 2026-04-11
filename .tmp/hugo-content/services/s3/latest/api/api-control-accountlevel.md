# AccountLevel

A container element for the account-level Amazon S3 Storage Lens configuration.

###### Note

You must enable Storage Lens metrics consistently at both the account level and bucket level, or your request will fail.

For more information about S3 Storage Lens, see [Assessing your storage activity and usage with S3 Storage Lens](../userguide/storage-lens.md) in the _Amazon S3 User Guide_. For a complete list of S3 Storage Lens metrics, see [S3 Storage Lens metrics glossary](../userguide/storage-lens-metrics-glossary.md) in the _Amazon S3 User Guide_.

## Contents

**BucketLevel**

A container element for the S3 Storage Lens bucket-level configuration.

Type: [BucketLevel](api-control-bucketlevel.md) data type

Required: Yes

**ActivityMetrics**

A container element for S3 Storage Lens activity metrics.

Type: [ActivityMetrics](api-control-activitymetrics.md) data type

Required: No

**AdvancedCostOptimizationMetrics**

A container element for S3 Storage Lens advanced cost-optimization metrics.

Type: [AdvancedCostOptimizationMetrics](api-control-advancedcostoptimizationmetrics.md) data type

Required: No

**AdvancedDataProtectionMetrics**

A container element for S3 Storage Lens advanced data-protection metrics.

Type: [AdvancedDataProtectionMetrics](api-control-advanceddataprotectionmetrics.md) data type

Required: No

**AdvancedPerformanceMetrics**

A container element for S3 Storage Lens advanced performance metrics.

Type: [AdvancedPerformanceMetrics](api-control-advancedperformancemetrics.md) data type

Required: No

**DetailedStatusCodesMetrics**

A container element for detailed status code metrics.

Type: [DetailedStatusCodesMetrics](api-control-detailedstatuscodesmetrics.md) data type

Required: No

**StorageLensGroupLevel**

A container element for S3 Storage Lens groups metrics.

Type: [StorageLensGroupLevel](api-control-storagelensgrouplevel.md) data type

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/accountlevel.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/accountlevel.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/accountlevel.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AccessPoint

ActivityMetrics

All content copied from https://docs.aws.amazon.com/.
