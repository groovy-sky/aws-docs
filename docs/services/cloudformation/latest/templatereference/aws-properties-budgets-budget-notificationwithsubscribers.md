This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Budgets::Budget NotificationWithSubscribers

A notification with subscribers. A notification can have one SNS subscriber and up to
10 email subscribers, for a total of 11 subscribers.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Notification" : Notification,
  "Subscribers" : [ Subscriber, ... ]
}

```

### YAML

```yaml

  Notification:
    Notification
  Subscribers:
    - Subscriber

```

## Properties

`Notification`

The notification that's associated with a budget.

_Required_: Yes

_Type_: [Notification](aws-properties-budgets-budget-notification.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Subscribers`

A list of subscribers who are subscribed to this notification.

_Required_: Yes

_Type_: Array of [Subscriber](aws-properties-budgets-budget-subscriber.md)

_Minimum_: `1`

_Maximum_: `11`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## See also

- [NotificationWithSubscribers](../../../../reference/aws-cost-management/latest/apireference/api-budgets-notificationwithsubscribers.md)
in the _AWS Cost Explorer Service Cost Management APIs_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Notification

ResourceTag

All content copied from https://docs.aws.amazon.com/.
