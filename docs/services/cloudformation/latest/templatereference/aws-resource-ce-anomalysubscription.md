This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CE::AnomalySubscription

The `AWS::CE::AnomalySubscription` resource (also referred to as an alert
subscription) is a Cost Explorer resource type that sends notifications about specific
anomalies that meet an alerting criteria defined by you.

You can specify the frequency of the alerts and the subscribers to notify.

Anomaly subscriptions can be associated with one or more [`AWS::CE::AnomalyMonitor`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ce-anomalymonitor.html) resources, and they only send notifications
about anomalies detected by those associated monitors. You can also configure a threshold to
further control which anomalies are included in the notifications.

Anomalies that don’t exceed the chosen threshold and therefore don’t trigger notifications
from an anomaly subscription will still be available on the console and from the [`GetAnomalies`](../../../../reference/aws-cost-management/latest/apireference/api-getanomalies.md) API.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CE::AnomalySubscription",
  "Properties" : {
      "Frequency" : String,
      "MonitorArnList" : [ String, ... ],
      "ResourceTags" : [ ResourceTag, ... ],
      "Subscribers" : [ Subscriber, ... ],
      "SubscriptionName" : String,
      "Threshold" : Number,
      "ThresholdExpression" : String
    }
}

```

### YAML

```yaml

Type: AWS::CE::AnomalySubscription
Properties:
  Frequency: String
  MonitorArnList:
    - String
  ResourceTags:
    - ResourceTag
  Subscribers:
    - Subscriber
  SubscriptionName: String
  Threshold: Number
  ThresholdExpression: String

```

## Properties

`Frequency`

The frequency that anomaly notifications are sent. Notifications are sent either over
email (for DAILY and WEEKLY frequencies) or SNS (for IMMEDIATE frequency). For more
information, see [Creating an Amazon SNS topic for\
anomaly notifications](https://docs.aws.amazon.com/cost-management/latest/userguide/ad-SNS.html).

_Required_: Yes

_Type_: String

_Allowed values_: `DAILY | IMMEDIATE | WEEKLY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MonitorArnList`

A list of cost anomaly monitors.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceTags`

Property description not available.

_Required_: No

_Type_: Array of [ResourceTag](aws-properties-ce-anomalysubscription-resourcetag.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Subscribers`

A list of subscribers to notify.

_Required_: Yes

_Type_: Array of [Subscriber](aws-properties-ce-anomalysubscription-subscriber.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubscriptionName`

The name for the subscription.

_Required_: Yes

_Type_: String

_Pattern_: `[\S\s]*`

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Threshold`

(deprecated)

An absolute dollar value that must be exceeded by the anomaly's total impact (see [Impact](../../../../reference/aws-cost-management/latest/apireference/api-impact.md) for more details) for an anomaly notification to be generated.

This field has been deprecated. To specify a threshold, use ThresholdExpression. Continued
use of Threshold will be treated as shorthand syntax for a ThresholdExpression.

One of Threshold or ThresholdExpression is required for
`AWS::CE::AnomalySubscription`. You cannot specify both.

_Required_: No

_Type_: Number

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ThresholdExpression`

An [Expression](../../../../reference/aws-cost-management/latest/apireference/api-expression.md) object
in JSON string format used to specify the anomalies that you want to generate alerts for. This
supports dimensions and nested expressions. The supported dimensions are
`ANOMALY_TOTAL_IMPACT_ABSOLUTE` and `ANOMALY_TOTAL_IMPACT_PERCENTAGE`,
corresponding to an anomaly’s TotalImpact and TotalImpactPercentage, respectively (see [Impact](../../../../reference/aws-cost-management/latest/apireference/api-impact.md) for more details). The supported nested expression types are
`AND` and `OR`. The match option `GREATER_THAN_OR_EQUAL` is
required. Values must be numbers between 0 and 10,000,000,000 in string format.

One of Threshold or ThresholdExpression is required for
`AWS::CE::AnomalySubscription`. You cannot specify both.

For further information, see the [Examples](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ce-anomalysubscription.html#aws-resource-ce-anomalysubscription--examples) section of this page.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns `SubscriptionArn`. For example:

`{ "Ref": "myAnomalyMonitorLogicalName" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AccountId`

Your unique account identifier.

`SubscriptionArn`

The `AnomalySubscription` Amazon Resource Name (ARN).

## Examples

- [Attaching subscriptions to monitors](#aws-resource-ce-anomalysubscription--examples--Attaching_subscriptions_to_monitors)

- [Using a percentage-based threshold](#aws-resource-ce-anomalysubscription--examples--Using_a_percentage-based_threshold)

- [Using a percentage-based threshold and absolute threshold combined with AND](#aws-resource-ce-anomalysubscription--examples--Using_a_percentage-based_threshold_and_absolute_threshold_combined_with_AND)

- [Using a percentage-based threshold and absolute threshold combined with OR](#aws-resource-ce-anomalysubscription--examples--Using_a_percentage-based_threshold_and_absolute_threshold_combined_with_OR)

### Attaching subscriptions to monitors

The following example shows two monitors attached to a subscription.

#### JSON

```json

{
  "Resources": {
    "CustomAnomalyMonitorWithLinkedAccount": {
      "Type": "AWS::CE::AnomalyMonitor",
      "Properties": {
        "MonitorName": "MonitorName",
        "MonitorType": "CUSTOM",
        "MonitorSpecification": " { \"Dimensions\" : { \"Key\" : \"LINKED_ACCOUNT\", \"Values\" : [ \"123456789012\", \"123456789013\" ] } }"
      }
    },
    "AnomalyServiceMonitor": {
      "Type": "AWS::CE::AnomalyMonitor",
      "Properties": {
        "MonitorName": "MonitorName",
        "MonitorType": "DIMENSIONAL",
        "MonitorDimension": "SERVICE"
      }
    },
    "AnomalySubscription": {
      "Type": "AWS::CE::AnomalySubscription",
      "Properties": {
        "SubscriptionName": "SubscriptionName",
        "Threshold": 100,
        "Frequency": "DAILY",
        "MonitorArnList": [
          {"Ref": "CustomAnomalyMonitorWithLinkedAccount"},
          {"Ref": "AnomalyServiceMonitor"}
        ],
        "Subscribers": [
          {
            "Type": "EMAIL",
            "Address": "abc@def.com"
          }
        ]
      }
    }
  }
}
```

#### YAML

```yaml

Resources:
   CustomAnomalyMonitorWithLinkedAccount:
    Type: 'AWS::CE::AnomalyMonitor'
    Properties:
      MonitorName: "MonitorName"
      MonitorType: "CUSTOM"
      MonitorSpecification: '
            {
              "Dimensions" : {
                "Key" : "LINKED_ACCOUNT",
                "Values" : [ "123456789012", "123456789013" ]
              }
            }'

   AnomalyServiceMonitor:
    Type: 'AWS::CE::AnomalyMonitor'
    Properties:
      MonitorName: 'MonitorName'
      MonitorType: 'DIMENSIONAL'
      MonitorDimension: 'SERVICE'

   AnomalySubscription:
    Type: 'AWS::CE::AnomalySubscription'
    Properties:
      SubscriptionName: "SubscriptionName"
      Threshold: 100
      Frequency: "DAILY"
      MonitorArnList: [
         !Ref CustomAnomalyMonitorWithLinkedAccount,
         !Ref AnomalyServiceMonitor
      ]
      Subscribers: [
        {
          "Type": "EMAIL",
          "Address": "abc@def.com"
        }
      ]
```

### Using a percentage-based threshold

The following example creates a subscription using a percentage-based
threshold.

#### JSON

```json

{
  "Resources": {
    "AnomalySubscription": {
      "Type": "AWS::CE::AnomalySubscription",
      "Properties": {
        "SubscriptionName": "SubscriptionName",
        "ThresholdExpression": "{ \"Dimensions\": { \"Key\": \"ANOMALY_TOTAL_IMPACT_PERCENTAGE\", \"Values\": [ \"100\" ], \"MatchOptions\": [ \"GREATER_THAN_OR_EQUAL\" ] } }",
        "MonitorArnList": [],
        "Subscribers": [
          {
            "Address": "abc@def.com",
            "Type": "EMAIL"
          }
        ],
        "Frequency": "DAILY"
      }
    }
  }
}
```

#### YAML

```yaml

Resources:
   AnomalySubscription:
    Type: 'AWS::CE::AnomalySubscription'
    Properties:
      SubscriptionName: "Subscription 1"
      ThresholdExpression: '{
        "Dimensions": {
          "Key": "ANOMALY_TOTAL_IMPACT_PERCENTAGE",
          "MatchOptions": [ "GREATER_THAN_OR_EQUAL" ],
          "Values": [ "100" ]
        }
      }'
      Frequency: "DAILY"
      MonitorArnList: [ ]
      Subscribers: [
        {
          "Type": "EMAIL",
          "Address": "abc@def.com"
        }
      ]
```

### Using a percentage-based threshold and absolute threshold combined with AND

The following example creates a subscription using a percentage-based threshold and
absolute threshold combined with AND.

#### JSON

```json

{
  "Resources": {
    "AnomalySubscription": {
      "Type": "AWS::CE::AnomalySubscription",
      "Properties": {
        "SubscriptionName": "SubscriptionName",
        "ThresholdExpression": "{ \"And\": [ { \"Dimensions\": { \"Key\": \"ANOMALY_TOTAL_IMPACT_PERCENTAGE\", \"MatchOptions\": [ \"GREATER_THAN_OR_EQUAL\" ], \"Values\": [ \"100\" ] } }, { \"Dimensions\": { \"Key\": \"ANOMALY_TOTAL_IMPACT_ABSOLUTE\", \"MatchOptions\": [ \"GREATER_THAN_OR_EQUAL\" ], \"Values\": [ \"200\" ] } } ] }",
        "MonitorArnList": [],
        "Subscribers": [
          {
            "Address": "abc@def.com",
            "Type": "EMAIL"
          }
        ],
        "Frequency": "DAILY"
      }
    }
  }
}
```

#### YAML

```yaml

Resources:
  AnomalySubscription:
    Type: 'AWS::CE::AnomalySubscription'
    Properties:
      SubscriptionName: "SubscriptionName"
      ThresholdExpression: '{
        "And": [
          {
            "Dimensions": {
              "Key": "ANOMALY_TOTAL_IMPACT_PERCENTAGE",
              "MatchOptions": [ "GREATER_THAN_OR_EQUAL" ],
              "Values": [ "100" ]
            }
          },
          {
            "Dimensions": {
              "Key": "ANOMALY_TOTAL_IMPACT_ABSOLUTE",
              "MatchOptions": [ "GREATER_THAN_OR_EQUAL" ],
              "Values": [ "200" ]
            }
          }
        ]
      }'
      Frequency: "DAILY"
      MonitorArnList: [ ]
      Subscribers: [
        {
          "Type": "EMAIL",
          "Address": "abc@def.com"
        }
      ]
```

### Using a percentage-based threshold and absolute threshold combined with OR

The following example creates a subscription using a percentage-based threshold and
absolute threshold combined with OR.

#### JSON

```json

{
  "Resources": {
    "AnomalySubscription": {
      "Type": "AWS::CE::AnomalySubscription",
      "Properties": {
        "SubscriptionName": "SubscriptionName",
        "ThresholdExpression": "{ \"Or\": [ { \"Dimensions\": { \"Key\": \"ANOMALY_TOTAL_IMPACT_PERCENTAGE\", \"MatchOptions\": [ \"GREATER_THAN_OR_EQUAL\" ], \"Values\": [ \"100\" ] } }, { \"Dimensions\": { \"Key\": \"ANOMALY_TOTAL_IMPACT_ABSOLUTE\", \"MatchOptions\": [ \"GREATER_THAN_OR_EQUAL\" ], \"Values\": [ \"200\" ] } } ] }",
        "MonitorArnList": [],
        "Subscribers": [
          {
            "Address": "abc@def.com",
            "Type": "EMAIL"
          }
        ],
        "Frequency": "DAILY"
      }
    }
  }
}
```

#### YAML

```yaml

Resources:
  AnomalySubscription:
    Type: 'AWS::CE::AnomalySubscription'
    Properties:
      SubscriptionName: "SubscriptionName"
      ThresholdExpression: '{
        "Or": [
          {
            "Dimensions": {
              "Key": "ANOMALY_TOTAL_IMPACT_PERCENTAGE",
              "MatchOptions": [ "GREATER_THAN_OR_EQUAL" ],
              "Values": [ "100" ]
            }
          },
          {
            "Dimensions": {
              "Key": "ANOMALY_TOTAL_IMPACT_ABSOLUTE",
              "MatchOptions": [ "GREATER_THAN_OR_EQUAL" ],
              "Values": [ "200" ]
            }
          }
        ]
      }'
      Frequency: "DAILY"
      MonitorArnList: [ ]
      Subscribers: [
        {
          "Type": "EMAIL",
          "Address": "abc@def.com"
        }
      ]
```

## See also

- [AnomalySubscription](https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_AnomalySubscription.html) in the _AWS Billing and Cost Management API_
_Reference_.

- [CreateAnomalySubscription](https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_CreateAnomalySubscription.html) in the _AWS Billing and Cost Management API_
_Reference_.

- [DeleteAnomalySubscription](https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_DeleteAnomalySubscription.html) in the _AWS Billing and Cost Management API_
_Reference_.

- [GetAnomalySubscriptions](https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_GetAnomalySubscriptions.html) in the _AWS Billing and Cost Management API_
_Reference_.

- [UpdateAnomalySubscription](https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_UpdateAnomalySubscription.html) in the _AWS Billing and Cost Management API_
_Reference_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ResourceTag

ResourceTag
