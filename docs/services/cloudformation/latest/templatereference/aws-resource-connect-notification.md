This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::Notification

Contains information about a notification, including its content, priority, recipients, and metadata.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Connect::Notification",
  "Properties" : {
      "Content" : NotificationContent,
      "ExpiresAt" : String,
      "InstanceArn" : String,
      "Priority" : String,
      "Recipients" : [ String, ... ],
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Connect::Notification
Properties:
  Content:
    NotificationContent
  ExpiresAt: String
  InstanceArn: String
  Priority: String
  Recipients:
    - String
  Tags:
    - Tag

```

## Properties

`Content`

The localized content of the notification. A map where keys are locale codes and values are the notification text in that locale.

_Required_: Yes

_Type_: [NotificationContent](aws-properties-connect-notification-notificationcontent.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExpiresAt`

The timestamp when the notification expires and is no longer displayed to users.

_Required_: No

_Type_: String

_Pattern_: `^[0-9]+$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`InstanceArn`

The Amazon Resource Name (ARN) of the instance.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Priority`

The priority level of the notification. Valid values are URGENT, HIGH, and LOW.

_Required_: No

_Type_: String

_Allowed values_: `HIGH | LOW`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Recipients`

A list of Amazon Resource Names (ARNs) identifying the recipients of the notification. Maximum of 200 recipients.

_Required_: No

_Type_: Array of String

_Maximum_: `200`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags used to organize, track, or control access for this resource. For example, `{ "Tags": {"key1":"value1", "key2":"value2"} }`.

_Required_: No

_Type_: Array of [Tag](aws-properties-connect-notification-tag.md)

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the notification ARN.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

`Arn`

The Amazon Resource Name (ARN) of the notification.

`CreatedAt`

The timestamp when the notification was created.

`Id`

The unique identifier for the notification.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Connect::IntegrationAssociation

NotificationContent

All content copied from https://docs.aws.amazon.com/.
