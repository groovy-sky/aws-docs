This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::Bucket MetricsConfiguration

Specifies a metrics configuration for the CloudWatch request metrics (specified by the
metrics configuration ID) from an Amazon S3 bucket. If you're updating an existing metrics
configuration, note that this is a full replacement of the existing metrics configuration. If
you don't include the elements you want to keep, they are erased. For examples, see [AWS::S3::Bucket](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#aws-properties-s3-bucket--examples). For more information, see [PUT Bucket metrics](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUTMetricConfiguration.html)
in the _Amazon S3 API Reference_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AccessPointArn" : String,
  "Id" : String,
  "Prefix" : String,
  "TagFilters" : [ TagFilter, ... ]
}

```

### YAML

```yaml

  AccessPointArn: String
  Id: String
  Prefix: String
  TagFilters:
    - TagFilter

```

## Properties

`AccessPointArn`

The access point that was used while performing operations on the object. The metrics
configuration only includes objects that meet the filter's criteria.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Id`

The ID used to identify the metrics configuration. This can be any value you choose that
helps you identify your metrics configuration.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Prefix`

The prefix that an object must have to be included in the metrics results.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TagFilters`

Specifies a list of tag filters to use as a metrics configuration filter. The metrics
configuration includes only objects that meet the filter's criteria.

_Required_: No

_Type_: Array of [TagFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-s3-bucket-tagfilter.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- AWS::S3::Bucket [Examples](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#aws-properties-s3-bucket--examples)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Metrics

NoncurrentVersionExpiration
