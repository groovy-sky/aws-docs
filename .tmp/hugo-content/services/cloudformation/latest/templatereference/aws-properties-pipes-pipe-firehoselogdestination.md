This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pipes::Pipe FirehoseLogDestination

Represents the Amazon Data Firehose logging configuration settings for the pipe.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DeliveryStreamArn" : String
}

```

### YAML

```yaml

  DeliveryStreamArn: String

```

## Properties

`DeliveryStreamArn`

The Amazon Resource Name (ARN) of the Firehose delivery stream to which EventBridge delivers the pipe log records.

_Required_: No

_Type_: String

_Pattern_: `^(^arn:aws([a-z]|\-)*:firehose:([a-z]{2}((-gov)|(-iso([a-z]?)))?-[a-z]+-\d{1}):(\d{12}):deliverystream/.+)$`

_Minimum_: `1`

_Maximum_: `1600`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FilterCriteria

MQBrokerAccessCredentials

All content copied from https://docs.aws.amazon.com/.
