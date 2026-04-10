This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Budgets::Budget

The `AWS::Budgets::Budget` resource allows customers to take pre-defined actions that will trigger once a budget threshold has been exceeded. creates, replaces, or deletes budgets for Billing and Cost Management.
For more information, see
[Managing Your Costs with Budgets](../../../awsaccountbilling/latest/aboutv2/budgets-managing-costs.md)
in the _AWS Billing and Cost Management User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Budgets::Budget",
  "Properties" : {
      "Budget" : BudgetData,
      "NotificationsWithSubscribers" : [ NotificationWithSubscribers, ... ],
      "ResourceTags" : [ ResourceTag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Budgets::Budget
Properties:
  Budget:
    BudgetData
  NotificationsWithSubscribers:
    - NotificationWithSubscribers
  ResourceTags:
    - ResourceTag

```

## Properties

`Budget`

The budget object that you want to create.

_Required_: Yes

_Type_: [BudgetData](aws-properties-budgets-budget-budgetdata.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NotificationsWithSubscribers`

A notification that you want to associate with a budget. A budget can have up to five notifications, and each notification can have one SNS subscriber and up to 10 email subscribers. If you include notifications and subscribers in your `CreateBudget` call, AWS creates the notifications and subscribers for you.

_Required_: No

_Type_: Array of [NotificationWithSubscribers](aws-properties-budgets-budget-notificationwithsubscribers.md)

_Maximum_: `10`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ResourceTags`

An optional list of tags to associate with the specified budget. Each tag consists of a
key and a value, and each key must be unique for the resource.

_Required_: No

_Type_: Array of [ResourceTag](aws-properties-budgets-budget-resourcetag.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the name of the budget that is created by the template.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

## Examples

### Budget for 100 USD with two notifications

The following example creates a budget for 100 USD amount of costs, with notifications for when you have spent over
80 USD or over 99 USD. The notifications are sent to the subscribers `email@example.com` and `email2@example.com`.

#### JSON

```json

{
  "Description": "Basic Budget test",
  "Resources": {
    "Budget": {
      "Type": "AWS::Budgets::Budget",
      "Properties": {
        "Budget": {
          "BudgetLimit": {
            "Amount": "100",
            "Unit": "USD"
          },
          "TimeUnit": "MONTHLY",
          "TimePeriod": {
            "Start": "1225864800",
            "End": "1926864800"
          },
          "BudgetType": "COST",
          "CostFilters": {
            "AZ": [
              "us-east-1",
              "us-west-1",
              "us-east-2"
            ]
          }
        },
        "NotificationsWithSubscribers": [
          {
            "Notification": {
              "NotificationType": "ACTUAL",
              "ComparisonOperator": "GREATER_THAN",
              "Threshold": 99
            },
            "Subscribers": [
              {
                "SubscriptionType": "EMAIL",
                "Address": "email@example.com"
              },
              {
                "SubscriptionType": "EMAIL",
                "Address": "email2@example.com"
              }
            ]
          },
          {
            "Notification": {
              "NotificationType": "ACTUAL",
              "ComparisonOperator": "GREATER_THAN",
              "Threshold": 80
            },
            "Subscribers": [
              {
                "SubscriptionType": "EMAIL",
                "Address": "email@example.com"
              }
            ]
          }
        ]
      }
    }
  },
  "Outputs": {
    "BudgetId": {
      "Value": "BudgetExample"
    }
  }
}
```

#### YAML

```yaml

---
Description: "Basic Budget test"
Resources:
  BudgetExample:
    Type: "AWS::Budgets::Budget"
    Properties:
      Budget:
        BudgetLimit:
          Amount: 100
          Unit: USD
        TimeUnit: MONTHLY
        TimePeriod:
          Start: 1225864800
          End: 1926864800
        BudgetType: COST
        CostFilters:
          AZ:
            - us-east-1
            - us-west-1
            - us-east-2
      NotificationsWithSubscribers:
        - Notification:
            NotificationType: ACTUAL
            ComparisonOperator: GREATER_THAN
            Threshold: 99
          Subscribers:
            - SubscriptionType: EMAIL
              Address: email@example.com
            - SubscriptionType: EMAIL
              Address: email2@example.com
        - Notification:
            NotificationType: ACTUAL
            ComparisonOperator: GREATER_THAN
            Threshold: 80
          Subscribers:
          - SubscriptionType: EMAIL
            Address: email@example.com
Outputs:
  BudgetId:
    Value: !Ref BudgetExample
```

## See also

- [CreateBudget](../../../../reference/aws-cost-management/latest/apireference/api-budgets-createbudget.md)
in the _AWS Cost Explorer Service Cost Management APIs_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS Budgets

AutoAdjustData

All content copied from https://docs.aws.amazon.com/.
