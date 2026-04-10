This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisAnalyticsV2::Application KinesisStreamsInput

Identifies a Kinesis data stream as the streaming source. You provide the
stream's Amazon Resource Name (ARN).

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

The ARN of the input Kinesis data stream to read.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:.*$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [KinesisStreamsInput](../../../managed-flink/latest/apiv2/api-kinesisstreamsinput.md) in the _Amazon Kinesis Data_
_Analytics API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

KinesisFirehoseInput

MappingParameters

All content copied from https://docs.aws.amazon.com/.
