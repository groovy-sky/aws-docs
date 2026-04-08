# Listing Amazon RDS event notification subscriptions

You can list your current Amazon RDS event notification subscriptions.

###### To list your current Amazon RDS event notification subscriptions

1. Sign in to the AWS Management Console and open the Amazon RDS console at
    [https://console.aws.amazon.com/rds/](https://console.aws.amazon.com/rds).

2. In the navigation pane, choose **Event subscriptions**. The **Event**
**subscriptions** pane shows all your event notification subscriptions.

![List DB event notification subscriptions](https://docs.aws.amazon.com/images/AmazonRDS/latest/UserGuide/images/EventNotification-ListSubs.png)

To list your current Amazon RDS event notification subscriptions, use the AWS CLI [`describe-event-subscriptions`](../../../cli/latest/reference/rds/describe-event-subscriptions.md) command.

###### Example

The following example describes all event subscriptions.

```nohighlight

aws rds describe-event-subscriptions
```

The following example describes the `myfirsteventsubscription`.

```nohighlight

aws rds describe-event-subscriptions --subscription-name myfirsteventsubscription
```

To list your current Amazon RDS event notification subscriptions, call the Amazon RDS API [`DescribeEventSubscriptions`](../../../../reference/amazonrds/latest/apireference/api-describeeventsubscriptions.md) action.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon RDS event notification tags and attributes

Modifying an Amazon RDS event notification subscription

All content copied from https://docs.aws.amazon.com/.
