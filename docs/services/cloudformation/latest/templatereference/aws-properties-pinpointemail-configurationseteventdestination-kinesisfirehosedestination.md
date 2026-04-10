This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::PinpointEmail::ConfigurationSetEventDestination KinesisFirehoseDestination

An object that defines an Amazon Kinesis Data Firehose destination for email events. You can use Amazon Kinesis Data Firehose to
stream data to other services, such as Amazon S3 and Amazon Redshift.

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

The Amazon Resource Name (ARN) of the Amazon Kinesis Data Firehose stream that Amazon Pinpoint sends email events
to.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IamRoleArn`

The Amazon Resource Name (ARN) of the IAM role that Amazon Pinpoint uses when sending email
events to the Amazon Kinesis Data Firehose stream.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EventDestination

PinpointDestination

All content copied from https://docs.aws.amazon.com/.
