This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SMSVOICE::ConfigurationSet KinesisFirehoseDestination

Contains the delivery stream Amazon Resource Name (ARN), and the ARN of the AWS Identity and Access Management (IAM) role associated with a Firehose event
destination.

Event destinations, such as Firehose, are associated with configuration
sets, which enable you to publish message sending events.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DeliveryStreamArn" : String,
  "IamRoleArn" : String
}

```

### YAML

```yaml

  DeliveryStreamArn: String
  IamRoleArn: String

```

## Properties

`DeliveryStreamArn`

The Amazon Resource Name (ARN) of the delivery stream.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:\S+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IamRoleArn`

The ARN of an AWS Identity and Access Management role that is able to write
event data to an Amazon Data Firehose destination.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:\S+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EventDestination

SnsDestination

All content copied from https://docs.aws.amazon.com/.
