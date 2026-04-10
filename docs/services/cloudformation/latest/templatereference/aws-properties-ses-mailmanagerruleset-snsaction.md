This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::MailManagerRuleSet SnsAction

The action to publish the email content to an Amazon SNS topic. When executed, this action
will send the email as a notification to the specified SNS topic.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ActionFailurePolicy" : String,
  "Encoding" : String,
  "PayloadType" : String,
  "RoleArn" : String,
  "TopicArn" : String
}

```

### YAML

```yaml

  ActionFailurePolicy: String
  Encoding: String
  PayloadType: String
  RoleArn: String
  TopicArn: String

```

## Properties

`ActionFailurePolicy`

A policy that states what to do in the case of failure. The action will fail if there are configuration errors.
For example, specified SNS topic has been deleted or the role lacks necessary permissions to call the `sns:Publish` API.

_Required_: No

_Type_: String

_Allowed values_: `CONTINUE | DROP`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Encoding`

The encoding to use for the email within the Amazon SNS notification. The default value is `UTF-8`.
Use `BASE64` if you need to preserve all special characters, especially when the original message uses
a different encoding format.

_Required_: No

_Type_: String

_Allowed values_: `UTF-8 | BASE64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PayloadType`

The expected payload type within the Amazon SNS notification. `CONTENT` attempts to publish the full email
content with 20KB of headers content. `HEADERS` extracts up to 100KB of header content to include in the notification,
email content will not be included to the notification. The default value is `CONTENT`.

_Required_: No

_Type_: String

_Allowed values_: `CONTENT | HEADERS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The Amazon Resource Name (ARN) of the IAM Role to use while writing to Amazon SNS. This role must have access to
the `sns:Publish` API for the given topic.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9:_/+=,@.#-]+$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TopicArn`

The Amazon Resource Name (ARN) of the Amazon SNS Topic to which notification for the email received will be published.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:(aws|aws-cn|aws-us-gov|aws-eusc):sns:[a-z]{2}-[a-z]+-\d{1}:\d{12}:[\w\-]{1,256}$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SendAction

Tag

All content copied from https://docs.aws.amazon.com/.
