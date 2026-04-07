This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Notifications::NotificationHub

Configures a `NotificationHub` for AWS User Notifications. For more
information about notification hub, see the [AWS User Notifications User Guide](https://docs.aws.amazon.com/notifications/latest/userguide/notification-hubs.html).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Notifications::NotificationHub",
  "Properties" : {
      "Region" : String
    }
}

```

### YAML

```yaml

Type: AWS::Notifications::NotificationHub
Properties:
  Region: String

```

## Properties

`Region`

The `NotificationHub` Region.

_Required_: Yes

_Type_: String

_Minimum_: `2`

_Maximum_: `25`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

### Fn::GetAtt

`CreationTime`

The date and time the `NotificationHubOverview` was created.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

NotificationHubStatusSummary
