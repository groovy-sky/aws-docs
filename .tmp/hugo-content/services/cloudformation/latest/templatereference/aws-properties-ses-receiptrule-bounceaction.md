This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::ReceiptRule BounceAction

The action to send a bounce response for the email. When executed, this action generates a non-delivery report (bounce) back to the sender.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Message" : String,
  "Sender" : String,
  "SmtpReplyCode" : String,
  "StatusCode" : String,
  "TopicArn" : String
}

```

### YAML

```yaml

  Message: String
  Sender: String
  SmtpReplyCode: String
  StatusCode: String
  TopicArn: String

```

## Properties

`Message`

The human-readable text to include in the bounce message.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Sender`

The sender email address of the bounce message.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SmtpReplyCode`

The SMTP reply code for the bounce, as defined by RFC 5321.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StatusCode`

The enhanced status code for the bounce, in the format of x.y.z (e.g. 5.1.1).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TopicArn`

The Amazon Resource Name (ARN) of the Amazon SNS topic to notify when the bounce action is
taken. You can find the ARN of a topic by using the [ListTopics](../../../sns/latest/api/api-listtopics.md) operation in
Amazon SNS.

For more information about Amazon SNS topics, see the [Amazon SNS Developer Guide](../../../sns/latest/dg/createtopic.md).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AddHeaderAction

ConnectAction

All content copied from https://docs.aws.amazon.com/.
