This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisAnalyticsV2::Application KinesisFirehoseInput

For a SQL-based Kinesis Data Analytics application, identifies a Kinesis Data
Firehose delivery stream as the streaming source. You provide the delivery stream's Amazon
Resource Name (ARN).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ResourceARN" : String
}

```

### YAML

```yaml

  ResourceARN: String

```

## Properties

`ResourceARN`

The Amazon Resource Name (ARN) of the delivery stream.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:.*$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [KinesisFirehoseInput](../../../managed-flink/latest/apiv2/api-kinesisfirehoseinput.md) in the _Amazon Kinesis Data_
_Analytics API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

JSONMappingParameters

KinesisStreamsInput

All content copied from https://docs.aws.amazon.com/.
