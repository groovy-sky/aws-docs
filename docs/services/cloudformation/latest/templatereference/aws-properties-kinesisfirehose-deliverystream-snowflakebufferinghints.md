This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisFirehose::DeliveryStream SnowflakeBufferingHints

Describes the buffering to perform before delivering data to the Snowflake destination. If you do not specify any value, Firehose uses the default values.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IntervalInSeconds" : Integer,
  "SizeInMBs" : Integer
}

```

### YAML

```yaml

  IntervalInSeconds: Integer
  SizeInMBs: Integer

```

## Properties

`IntervalInSeconds`

Buffer incoming data for the specified period of time, in seconds, before delivering it to the destination. The default value is 0.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `900`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SizeInMBs`

Buffer incoming data to the specified size, in MBs, before delivering it to the
destination. The default value is 128.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Serializer

SnowflakeDestinationConfiguration

All content copied from https://docs.aws.amazon.com/.
