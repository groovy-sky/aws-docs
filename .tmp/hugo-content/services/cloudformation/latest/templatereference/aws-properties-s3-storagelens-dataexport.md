This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::StorageLens DataExport

This resource contains the details of the Amazon S3 Storage Lens metrics export.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CloudWatchMetrics" : CloudWatchMetrics,
  "S3BucketDestination" : S3BucketDestination,
  "StorageLensTableDestination" : StorageLensTableDestination
}

```

### YAML

```yaml

  CloudWatchMetrics:
    CloudWatchMetrics
  S3BucketDestination:
    S3BucketDestination
  StorageLensTableDestination:
    StorageLensTableDestination

```

## Properties

`CloudWatchMetrics`

This property enables the Amazon CloudWatch publishing option for S3 Storage Lens
metrics.

_Required_: No

_Type_: [CloudWatchMetrics](aws-properties-s3-storagelens-cloudwatchmetrics.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3BucketDestination`

This property contains the details of the bucket where the S3 Storage Lens metrics export
will be placed.

_Required_: No

_Type_: [S3BucketDestination](aws-properties-s3-storagelens-s3bucketdestination.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StorageLensTableDestination`

This property contains the details of the S3 table bucket where the S3 Storage Lens default metrics report will be placed. This property enables you to store your Storage Lens metrics in read-only S3 Tables.

_Required_: No

_Type_: [StorageLensTableDestination](aws-properties-s3-storagelens-storagelenstabledestination.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CloudWatchMetrics

DetailedStatusCodesMetrics

All content copied from https://docs.aws.amazon.com/.
