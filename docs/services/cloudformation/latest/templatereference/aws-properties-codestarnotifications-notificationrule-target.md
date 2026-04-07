This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeStarNotifications::NotificationRule Target

Information about the Amazon Q Developer in chat applications topics or Amazon Q Developer in chat applications clients associated with a notification rule.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "TargetAddress" : String,
  "TargetType" : String
}

```

### YAML

```yaml

  TargetAddress: String
  TargetType: String

```

## Properties

`TargetAddress`

The Amazon Resource Name (ARN) of the Amazon Q Developer in chat applications topic or Amazon Q Developer in chat applications client.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `320`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetType`

The target type. Can be an Amazon Simple Notification Service topic or Amazon Q Developer in chat applications client.

- Amazon Simple Notification Service topics are specified as `SNS`.

- Amazon Q Developer in chat applications clients are specified as `AWSChatbotSlack`.

- Amazon Q Developer in chat applications clients for Microsoft Teams are specified as `AWSChatbotMicrosoftTeams`.

_Required_: Yes

_Type_: String

_Pattern_: `^[A-Za-z]+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::CodeStarNotifications::NotificationRule

Next
