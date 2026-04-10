This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisFirehose::DeliveryStream CloudWatchLoggingOptions

The `CloudWatchLoggingOptions` property type specifies Amazon CloudWatch
Logs (CloudWatch Logs) logging options that Amazon Kinesis Data Firehose (Kinesis Data
Firehose) uses for the delivery stream.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Enabled" : Boolean,
  "LogGroupName" : String,
  "LogStreamName" : String
}

```

### YAML

```yaml

  Enabled: Boolean
  LogGroupName: String
  LogStreamName: String

```

## Properties

`Enabled`

Indicates whether CloudWatch Logs logging is enabled.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogGroupName`

The name of the CloudWatch Logs log group that contains the log stream that Kinesis
Data Firehose will use.

Conditional. If you enable logging, you must specify this property.

_Required_: Conditional

_Type_: String

_Pattern_: `[\.\-_/#A-Za-z0-9]*`

_Minimum_: `0`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogStreamName`

The name of the CloudWatch Logs log stream that Kinesis Data Firehose uses to send
logs about data delivery.

Conditional. If you enable logging, you must specify this property.

_Required_: Conditional

_Type_: String

_Pattern_: `[^:*]*`

_Minimum_: `0`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CatalogConfiguration

CopyCommand

All content copied from https://docs.aws.amazon.com/.
