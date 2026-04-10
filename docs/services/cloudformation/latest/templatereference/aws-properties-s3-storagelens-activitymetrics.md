This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::StorageLens ActivityMetrics

This resource enables Amazon S3 Storage Lens activity metrics. Activity metrics
show details about how your storage is requested, such as requests (for example, All requests,
Get requests, Put requests), bytes uploaded or downloaded, and errors.

For more information, see [Assessing your storage activity and usage\
with S3 Storage Lens](../../../s3/latest/userguide/storage-lens.md) in the _Amazon S3 User Guide_. For a
complete list of metrics, see [S3 Storage Lens metrics\
glossary](../../../s3/latest/userguide/storage-lens-metrics-glossary.md) in the _Amazon S3 User Guide_.

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

A property that indicates whether the activity metrics is enabled.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AccountLevel

AdvancedCostOptimizationMetrics

All content copied from https://docs.aws.amazon.com/.
