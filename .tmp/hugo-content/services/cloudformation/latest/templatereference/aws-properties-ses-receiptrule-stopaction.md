This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::ReceiptRule StopAction

When included in a receipt rule, this action terminates the evaluation of the receipt
rule set and, optionally, publishes a notification to Amazon Simple Notification Service (Amazon SNS).

For information about setting a stop action in a receipt rule, see the [Amazon SES\
Developer Guide](../../../ses/latest/dg/receiving-email-action-stop.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Scope" : String,
  "TopicArn" : String
}

```

### YAML

```yaml

  Scope: String
  TopicArn: String

```

## Properties

`Scope`

The scope of the StopAction. The only acceptable value is `RuleSet`.

_Required_: Yes

_Type_: String

_Allowed values_: `RuleSet`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TopicArn`

The Amazon Resource Name (ARN) of the Amazon SNS topic to notify when the stop action is
taken. You can find the ARN of a topic by using the [ListTopics](../../../sns/latest/api/api-listtopics.md) Amazon SNS operation.

For more information about Amazon SNS topics, see the [Amazon SNS Developer Guide](../../../sns/latest/dg/createtopic.md).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SNSAction

WorkmailAction

All content copied from https://docs.aws.amazon.com/.
