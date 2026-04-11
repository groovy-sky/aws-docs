This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Budgets::Budget Notification

A notification that's associated with a budget. A budget can have up to ten
notifications.

Each notification must have at least one subscriber. A notification can have one SNS
subscriber and up to 10 email subscribers, for a total of 11 subscribers.

For example, if you have a budget for 200 dollars and you want to be notified when you
go over 160 dollars, create a notification with the following parameters:

- A notificationType of `ACTUAL`

- A `thresholdType` of `PERCENTAGE`

- A `comparisonOperator` of `GREATER_THAN`

- A notification `threshold` of `80`

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ComparisonOperator" : String,
  "NotificationType" : String,
  "Threshold" : Number,
  "ThresholdType" : String
}

```

### YAML

```yaml

  ComparisonOperator: String
  NotificationType: String
  Threshold: Number
  ThresholdType: String

```

## Properties

`ComparisonOperator`

The comparison that's used for this notification.

_Required_: Yes

_Type_: String

_Allowed values_: `GREATER_THAN | LESS_THAN | EQUAL_TO`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NotificationType`

Specifies whether the notification is for how much you have spent
( `ACTUAL`) or for how much that you're forecasted to spend
( `FORECASTED`).

_Required_: Yes

_Type_: String

_Allowed values_: `ACTUAL | FORECASTED`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Threshold`

The threshold that's associated with a notification. Thresholds are always a
percentage, and many customers find value being alerted between 50% - 200% of the
budgeted amount. The maximum limit for your threshold is 1,000,000% above the budgeted
amount.

_Required_: Yes

_Type_: Number

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ThresholdType`

The type of threshold for a notification. For `ABSOLUTE_VALUE` thresholds,
AWS notifies you when you go over or are forecasted to go over your
total cost threshold. For
`PERCENTAGE` thresholds, AWS notifies you when you go over
or are forecasted to go over a certain percentage of your forecasted spend. For example,
if you have a budget for 200 dollars and you have a `PERCENTAGE` threshold of
80%, AWS notifies you when you go over 160 dollars.

_Required_: No

_Type_: String

_Allowed values_: `PERCENTAGE | ABSOLUTE_VALUE`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## See also

- [Notification](../../../../reference/aws-cost-management/latest/apireference/api-budgets-notification.md)
in the _AWS Cost Explorer Service Cost Management APIs_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HistoricalOptions

NotificationWithSubscribers

All content copied from https://docs.aws.amazon.com/.
