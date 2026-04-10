This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IVSChat::LoggingConfiguration FirehoseDestinationConfiguration

The FirehoseDestinationConfiguration property type specifies a Kinesis Firehose location where chat logs will be stored.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DeliveryStreamName" : String
}

```

### YAML

```yaml

  DeliveryStreamName: String

```

## Properties

`DeliveryStreamName`

Name of the Amazon Kinesis Firehose delivery stream where chat activity will be
logged.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_.-]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DestinationConfiguration

S3DestinationConfiguration

All content copied from https://docs.aws.amazon.com/.
