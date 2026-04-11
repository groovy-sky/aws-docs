This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::ReceiptRule SNSAction

The action to publish the email content to an Amazon SNS topic. When executed, this action
will send the email as a notification to the specified SNS topic.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Encoding" : String,
  "TopicArn" : String
}

```

### YAML

```yaml

  Encoding: String
  TopicArn: String

```

## Properties

`Encoding`

The encoding to use for the email within the Amazon SNS notification. The default value is `UTF-8`.
Use `BASE64` if you need to preserve all special characters, especially when the original message uses
a different encoding format.

_Required_: No

_Type_: String

_Allowed values_: `UTF-8 | Base64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TopicArn`

The Amazon Resource Name (ARN) of the Amazon SNS Topic to which notification for the email received will be published.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

S3Action

StopAction

All content copied from https://docs.aws.amazon.com/.
