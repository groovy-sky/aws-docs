This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::ConfigurationSetEventDestination SnsDestination

Contains the topic ARN associated with an Amazon Simple Notification Service (Amazon SNS) event destination.

Event destinations, such as Amazon SNS, are associated with configuration sets, which
enable you to publish email sending events. For information about using configuration
sets, see the [Amazon SES Developer\
Guide](../../../ses/latest/dg/monitor-sending-activity.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "TopicARN" : String
}

```

### YAML

```yaml

  TopicARN: String

```

## Properties

`TopicARN`

The ARN of the Amazon SNS topic for email sending events. You can find the ARN of a topic
by using the [ListTopics](../../../sns/latest/api/api-listtopics.md) Amazon SNS operation.

For more information about Amazon SNS topics, see the [Amazon SNS Developer Guide](../../../sns/latest/dg/createtopic.md).

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[a-z0-9-]*:sns:[a-z0-9-]+:\d{12}:[^:]+$`

_Minimum_: `36`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

KinesisFirehoseDestination

AWS::SES::ContactList

All content copied from https://docs.aws.amazon.com/.
