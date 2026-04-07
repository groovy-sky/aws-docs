This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Notifications::ChannelAssociation

The `AWS::Notifications::ChannelAssociation` resource associates a
`Channel` with a `NotificationConfiguration` for AWS User Notifications. For more information about AWS User Notifications, see the [AWS User Notifications User Guide](https://docs.aws.amazon.com/notifications/latest/userguide/what-is-service.html).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Notifications::ChannelAssociation",
  "Properties" : {
      "Arn" : String,
      "NotificationConfigurationArn" : String
    }
}

```

### YAML

```yaml

Type: AWS::Notifications::ChannelAssociation
Properties:
  Arn: String
  NotificationConfigurationArn: String

```

## Properties

`Arn`

The Amazon Resource Name (ARN) of the `Channel`.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:[a-z-]{3,10}:(chatbot|consoleapp|notifications-contacts):[a-zA-Z0-9-]*:[0-9]{12}:[a-zA-Z0-9-_.@]+/[a-zA-Z0-9/_.@:-]+$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NotificationConfigurationArn`

The ARN of the `NotificationConfiguration` associated with the `Channel`.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:[a-z-]{3,10}:notifications::[0-9]{12}:configuration/[a-z0-9]{27}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic Ref function, Ref returns the ARN of the configuration created.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Notifications

AWS::Notifications::EventRule
