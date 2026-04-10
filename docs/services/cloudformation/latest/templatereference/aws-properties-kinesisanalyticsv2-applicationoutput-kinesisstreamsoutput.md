This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisAnalyticsV2::ApplicationOutput KinesisStreamsOutput

When you configure a SQL-based Kinesis Data Analytics application's output,
identifies a Kinesis data stream as the destination. You provide the stream Amazon Resource
Name (ARN).

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

The ARN of the destination Kinesis data stream to write to.

_Required_: Yes

_Type_: String

_Pattern_: `arn:.*`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [KinesisStreamsOutput](../../../managed-flink/latest/apiv2/api-kinesisstreamsoutput.md) in the _Amazon Kinesis Data_
_Analytics API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

KinesisFirehoseOutput

LambdaOutput

All content copied from https://docs.aws.amazon.com/.
