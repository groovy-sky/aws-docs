This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Notifications::NotificationConfiguration

Configures a `NotificationConfiguration` for AWS User Notifications.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Notifications::NotificationConfiguration",
  "Properties" : {
      "AggregationDuration" : String,
      "Description" : String,
      "Name" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Notifications::NotificationConfiguration
Properties:
  AggregationDuration: String
  Description: String
  Name: String
  Tags:
    - Tag

```

## Properties

`AggregationDuration`

The aggregation preference of the `NotificationConfiguration`.

- Values:

- `LONG`

- Aggregate notifications for long periods of time (12 hours).

- `SHORT`

- Aggregate notifications for short periods of time (5 minutes).

- `NONE`

- Don't aggregate notifications.

_Required_: No

_Type_: String

_Allowed values_: `LONG | SHORT | NONE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the `NotificationConfiguration`.

_Required_: Yes

_Type_: String

_Pattern_: `^[^\u0001-\u001F\u007F-\u009F]*$`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the `NotificationConfiguration`. Supports RFC 3986's
unreserved characters.

_Required_: Yes

_Type_: String

_Pattern_: `^[A-Za-z0-9_\-]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A map of tags assigned to a `NotificationConfiguration`.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-notifications-notificationconfiguration-tag.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic Ref function, Ref returns the ARN of the configuration created.

### Fn::GetAtt

`Arn`

The Amazon Resource Name (ARN) of the `NotificationConfiguration`
resource.

`CreationTime`

The creation time of the `NotificationConfiguration`.

`Status`

The current status of the `NotificationConfiguration`.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Notifications::ManagedNotificationAdditionalChannelAssociation

Tag
