This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::VerifiedAccessInstance KinesisDataFirehose

Options for Kinesis as a logging destination.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DeliveryStream" : String,
  "Enabled" : Boolean
}

```

### YAML

```yaml

  DeliveryStream: String
  Enabled: Boolean

```

## Properties

`DeliveryStream`

The ID of the delivery stream.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Enabled`

Indicates whether logging is enabled.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CloudWatchLogs

S3

All content copied from https://docs.aws.amazon.com/.
