This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::StorageLens CloudWatchMetrics

This resource enables the Amazon CloudWatch publishing option for Amazon S3
Storage Lens metrics.

For more information, see [Monitor S3\
Storage Lens metrics in CloudWatch](../../../s3/latest/userguide/storage-lens-view-metrics-cloudwatch.md) in the _Amazon S3 User_
_Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IsEnabled" : Boolean
}

```

### YAML

```yaml

  IsEnabled: Boolean

```

## Properties

`IsEnabled`

This property identifies whether the CloudWatch publishing option for S3 Storage
Lens is enabled.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BucketsAndRegions

DataExport

All content copied from https://docs.aws.amazon.com/.
