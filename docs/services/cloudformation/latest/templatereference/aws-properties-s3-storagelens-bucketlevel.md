This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::StorageLens BucketLevel

A property for the bucket-level storage metrics for Amazon S3 Storage Lens.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ActivityMetrics" : ActivityMetrics,
  "AdvancedCostOptimizationMetrics" : AdvancedCostOptimizationMetrics,
  "AdvancedDataProtectionMetrics" : AdvancedDataProtectionMetrics,
  "AdvancedPerformanceMetrics" : AdvancedPerformanceMetrics,
  "DetailedStatusCodesMetrics" : DetailedStatusCodesMetrics,
  "PrefixLevel" : PrefixLevel
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
  DetailedStatusCodesMetrics:
    DetailedStatusCodesMetrics
  PrefixLevel:
    PrefixLevel

```

## Properties

`ActivityMetrics`

A property for bucket-level activity metrics for S3 Storage Lens.

_Required_: No

_Type_: [ActivityMetrics](aws-properties-s3-storagelens-activitymetrics.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AdvancedCostOptimizationMetrics`

A property for bucket-level advanced cost optimization metrics for S3 Storage Lens.

_Required_: No

_Type_: [AdvancedCostOptimizationMetrics](aws-properties-s3-storagelens-advancedcostoptimizationmetrics.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AdvancedDataProtectionMetrics`

A property for bucket-level advanced data protection metrics for S3 Storage Lens.

_Required_: No

_Type_: [AdvancedDataProtectionMetrics](aws-properties-s3-storagelens-advanceddataprotectionmetrics.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AdvancedPerformanceMetrics`

A property for bucket-level advanced performance metrics for S3 Storage Lens.

_Required_: No

_Type_: [AdvancedPerformanceMetrics](aws-properties-s3-storagelens-advancedperformancemetrics.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DetailedStatusCodesMetrics`

A property for bucket-level detailed status code metrics for S3 Storage Lens.

_Required_: No

_Type_: [DetailedStatusCodesMetrics](aws-properties-s3-storagelens-detailedstatuscodesmetrics.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrefixLevel`

A property for bucket-level prefix-level storage metrics for S3 Storage Lens.

_Required_: No

_Type_: [PrefixLevel](aws-properties-s3-storagelens-prefixlevel.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AwsOrg

BucketsAndRegions

All content copied from https://docs.aws.amazon.com/.
