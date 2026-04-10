This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Timestream::ScheduledQuery SnsConfiguration

Details on SNS that are required to send the notification.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "TopicArn" : String
}

```

### YAML

```yaml

  TopicArn: String

```

## Properties

`TopicArn`

SNS topic ARN that the scheduled query status notifications will be sent to.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ScheduleConfiguration

Tag

All content copied from https://docs.aws.amazon.com/.
