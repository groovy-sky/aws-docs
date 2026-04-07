This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Notifications::ManagedNotificationAdditionalChannelAssociation

Associates a `Channel` with a `ManagedNotificationAdditionalChannelAssociation` for AWS User Notifications.
For more information about AWS User Notifications, see the [AWS User Notifications User Guide](https://docs.aws.amazon.com/notifications/latest/userguide/what-is-service.html).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Notifications::ManagedNotificationAdditionalChannelAssociation",
  "Properties" : {
      "ChannelArn" : String,
      "ManagedNotificationConfigurationArn" : String
    }
}

```

### YAML

```yaml

Type: AWS::Notifications::ManagedNotificationAdditionalChannelAssociation
Properties:
  ChannelArn: String
  ManagedNotificationConfigurationArn: String

```

## Properties

`ChannelArn`

The ARN of the `Channel`.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:[a-z-]{3,10}:(chatbot|consoleapp|notifications-contacts):[a-zA-Z0-9-]*:[0-9]{12}:[a-zA-Z0-9-_.@]+/[a-zA-Z0-9/_.@:-]+$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ManagedNotificationConfigurationArn`

The ARN of the `ManagedNotificationAdditionalChannelAssociation` associated with the `Channel`.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:[a-z-]{3,10}:notifications::([0-9]{12}|):managed-notification-configuration/category/[a-zA-Z0-9\-]{3,64}/sub-category/[a-zA-Z0-9\-]{3,64}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic Ref function, Ref returns the ARN of the configuration created.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Notifications::ManagedNotificationAccountContactAssociation

AWS::Notifications::NotificationConfiguration
