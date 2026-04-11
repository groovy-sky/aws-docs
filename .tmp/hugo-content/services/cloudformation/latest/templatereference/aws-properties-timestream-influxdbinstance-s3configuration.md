This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Timestream::InfluxDBInstance S3Configuration

Configuration for S3 bucket log delivery.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BucketName" : String,
  "Enabled" : Boolean
}

```

### YAML

```yaml

  BucketName: String
  Enabled: Boolean

```

## Properties

`BucketName`

The bucket name of the customer S3 bucket.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9a-z]+[0-9a-z\.\-]*[0-9a-z]+$`

_Minimum_: `3`

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Enabled`

Indicates whether log delivery to the S3 bucket is enabled.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LogDeliveryConfiguration

Tag

All content copied from https://docs.aws.amazon.com/.
