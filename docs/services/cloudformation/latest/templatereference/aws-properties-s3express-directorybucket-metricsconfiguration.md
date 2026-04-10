This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3Express::DirectoryBucket MetricsConfiguration

Specifies a metrics configuration for the CloudWatch request metrics (specified by the metrics
configuration ID) from an Amazon S3 bucket. If you're updating an existing metrics configuration, note that
this is a full replacement of the existing metrics configuration. If you don't include the elements you
want to keep, they are erased. For more information, see [PutBucketMetricsConfiguration](../../../s3/latest/api/restbucketputmetricconfiguration.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AccessPointArn" : String,
  "Id" : String,
  "Prefix" : String
}

```

### YAML

```yaml

  AccessPointArn: String
  Id: String
  Prefix: String

```

## Properties

`AccessPointArn`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Id`

The ID used to identify the metrics configuration. The ID has a 64 character limit and can only
contain letters, numbers, periods, dashes, and underscores.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Prefix`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LifecycleConfiguration

Rule

All content copied from https://docs.aws.amazon.com/.
