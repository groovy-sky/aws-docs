This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::InternetMonitor::Monitor S3Config

The configuration for publishing Amazon CloudWatch Internet Monitor internet measurements to Amazon S3. The configuration
includes the bucket name and (optionally) bucket prefix for the S3 bucket to store the measurements, and the delivery status.
The delivery status is `ENABLED` if you choose to deliver internet measurements to S3 logs, and `DISABLED` otherwise.

The measurements are also published to Amazon CloudWatch Logs.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BucketName" : String,
  "BucketPrefix" : String,
  "LogDeliveryStatus" : String
}

```

### YAML

```yaml

  BucketName: String
  BucketPrefix: String
  LogDeliveryStatus: String

```

## Properties

`BucketName`

The Amazon S3 bucket name for internet measurements publishing.

_Required_: No

_Type_: String

_Minimum_: `3`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BucketPrefix`

An optional Amazon S3 bucket prefix for internet measurements publishing.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogDeliveryStatus`

The status of publishing Internet Monitor internet measurements to an Amazon S3 bucket. The delivery status is `ENABLED`
if you choose to deliver internet measurements to an S3 bucket, and `DISABLED` otherwise.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LocalHealthEventsConfig

Tag

All content copied from https://docs.aws.amazon.com/.
