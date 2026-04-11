This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::StorageLens AccountLevel

This resource contains the details of the account-level metrics for Amazon S3 Storage
Lens.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ActivityMetrics" : ActivityMetrics,
  "AdvancedCostOptimizationMetrics" : AdvancedCostOptimizationMetrics,
  "AdvancedDataProtectionMetrics" : AdvancedDataProtectionMetrics,
  "AdvancedPerformanceMetrics" : AdvancedPerformanceMetrics,
  "BucketLevel" : BucketLevel,
  "DetailedStatusCodesMetrics" : DetailedStatusCodesMetrics,
  "StorageLensGroupLevel" : StorageLensGroupLevel
}

```

### YAML

```yaml

  ActivityMetrics:
    ActivityMetrics
  AdvancedCostOptimizationMetrics:
    AdvancedCostOptimizationMetrics
  AdvancedDataProtectionMetrics:
    AdvancedDataProtectionMetrics
  AdvancedPerformanceMetrics:
    AdvancedPerformanceMetrics
  BucketLevel:
    BucketLevel
  DetailedStatusCodesMetrics:
    DetailedStatusCodesMetrics
  StorageLensGroupLevel:
    StorageLensGroupLevel

```

## Properties

`ActivityMetrics`

This property contains the details of account-level activity metrics for S3 Storage
Lens.

_Required_: No

_Type_: [ActivityMetrics](aws-properties-s3-storagelens-activitymetrics.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AdvancedCostOptimizationMetrics`

This property contains the details of account-level advanced cost optimization metrics for
S3 Storage Lens.

_Required_: No

_Type_: [AdvancedCostOptimizationMetrics](aws-properties-s3-storagelens-advancedcostoptimizationmetrics.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AdvancedDataProtectionMetrics`

This property contains the details of account-level advanced data protection metrics for
S3 Storage Lens.

_Required_: No

_Type_: [AdvancedDataProtectionMetrics](aws-properties-s3-storagelens-advanceddataprotectionmetrics.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AdvancedPerformanceMetrics`

This property contains the account-level details for S3 Storage Lens advanced
performance metrics.

_Required_: No

_Type_: [AdvancedPerformanceMetrics](aws-properties-s3-storagelens-advancedperformancemetrics.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BucketLevel`

This property contains the details of the account-level bucket-level configurations for
Amazon S3 Storage Lens. To enable bucket-level configurations, make sure to also set the same metrics at the account level.

_Required_: Yes

_Type_: [BucketLevel](aws-properties-s3-storagelens-bucketlevel.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DetailedStatusCodesMetrics`

This property contains the details of account-level detailed status code metrics for S3
Storage Lens.

_Required_: No

_Type_: [DetailedStatusCodesMetrics](aws-properties-s3-storagelens-detailedstatuscodesmetrics.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StorageLensGroupLevel`

This property determines the scope of Storage Lens group data that is displayed in the
Storage Lens dashboard.

_Required_: No

_Type_: [StorageLensGroupLevel](aws-properties-s3-storagelens-storagelensgrouplevel.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::S3::StorageLens

ActivityMetrics

All content copied from https://docs.aws.amazon.com/.
