This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Workteam NotificationConfiguration

Configures Amazon SNS notifications of available or expiring work items for work
teams.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "NotificationTopicArn" : String
}

```

### YAML

```yaml

  NotificationTopicArn: String

```

## Properties

`NotificationTopicArn`

The ARN for the Amazon SNS topic to which notifications should be published.

_Required_: Yes

_Type_: String

_Pattern_: `arn:aws[a-z\-]*:sns:[a-z0-9\-]*:[0-9]{12}:[a-zA-Z0-9_.-]*`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MemberDefinition

OidcMemberDefinition

All content copied from https://docs.aws.amazon.com/.
