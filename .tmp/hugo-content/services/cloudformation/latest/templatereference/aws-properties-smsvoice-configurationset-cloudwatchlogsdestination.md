This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SMSVOICE::ConfigurationSet CloudWatchLogsDestination

Contains the destination configuration to use when publishing message sending events.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IamRoleArn" : String,
  "LogGroupArn" : String
}

```

### YAML

```yaml

  IamRoleArn: String
  LogGroupArn: String

```

## Properties

`IamRoleArn`

The Amazon Resource Name (ARN) of an AWS Identity and Access Management role
that is able to write event data to an Amazon CloudWatch destination.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:\S+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogGroupArn`

The name of the Amazon CloudWatch log group that you want to record events in.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:\S+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::SMSVOICE::ConfigurationSet

EventDestination

All content copied from https://docs.aws.amazon.com/.
