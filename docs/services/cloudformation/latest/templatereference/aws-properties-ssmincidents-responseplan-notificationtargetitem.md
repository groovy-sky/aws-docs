This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSMIncidents::ResponsePlan NotificationTargetItem

The Amazon SNS topic that's used by Amazon Q Developer in chat applications to notify the incidents chat
channel.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SnsTopicArn" : String
}

```

### YAML

```yaml

  SnsTopicArn: String

```

## Properties

`SnsTopicArn`

The Amazon Resource Name (ARN) of the Amazon SNS topic.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws(-(cn|us-gov))?:sns:(([a-z]+-)+[0-9])?:([0-9]{12})?:[^.]+$`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Integration

PagerDutyConfiguration

All content copied from https://docs.aws.amazon.com/.
