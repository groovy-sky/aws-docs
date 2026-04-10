This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::StorageLens BucketsAndRegions

This resource contains the details of the buckets and Regions for the Amazon S3 Storage
Lens configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Buckets" : [ String, ... ],
  "Regions" : [ String, ... ]
}

```

### YAML

```yaml

  Buckets:
    - String
  Regions:
    - String

```

## Properties

`Buckets`

This property contains the details of the buckets for the Amazon S3 Storage Lens
configuration. This should be the bucket Amazon Resource Name(ARN). For valid values, see
[Buckets ARN\
format here](../../../s3/latest/api/api-control-include.md#API_control_Include_Contents) in the _Amazon S3 API Reference_.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Regions`

This property contains the details of the Regions for the S3 Storage Lens
configuration.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BucketLevel

CloudWatchMetrics

All content copied from https://docs.aws.amazon.com/.
