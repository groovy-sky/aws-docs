This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CE::AnomalyMonitor

The `AWS::CE::AnomalyMonitor` resource is a Cost Explorer resource type that
continuously inspects your account's cost data for anomalies, based on
`MonitorType` and `MonitorSpecification`. The content consists of
detailed metadata and the current status of the monitor object.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CE::AnomalyMonitor",
  "Properties" : {
      "MonitorDimension" : String,
      "MonitorName" : String,
      "MonitorSpecification" : String,
      "MonitorType" : String,
      "ResourceTags" : [ ResourceTag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::CE::AnomalyMonitor
Properties:
  MonitorDimension: String
  MonitorName: String
  MonitorSpecification: String
  MonitorType: String
  ResourceTags:
    - ResourceTag

```

## Properties

`MonitorDimension`

For customer managed monitors, do not specify this field.

For AWS managed monitors, this field controls which cost dimension is
automatically analyzed by the monitor. For `TAG` and `COST_CATEGORY
            ` dimensions, you must also specify MonitorSpecification to configure the specific
tag or cost category key to analyze.

_Required_: No

_Type_: String

_Allowed values_: `SERVICE | LINKED_ACCOUNT | TAG | COST_CATEGORY`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MonitorName`

The name of the monitor.

_Required_: Yes

_Type_: String

_Pattern_: `[\S\s]*`

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MonitorSpecification`

The array of `MonitorSpecification` in JSON array format. For instance, you can
use `MonitorSpecification` to specify a tag, Cost Category, or linked account for
your custom anomaly monitor. For further information, see the [Examples](../userguide/aws-resource-ce-anomalymonitor.md#aws-resource-ce-anomalymonitor--examples) section of this page.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MonitorType`

The type of the monitor.

Set this to `DIMENSIONAL` for an AWS managed monitor.
AWS managed monitors automatically track up to the top 5,000 values by
cost within a dimension of your choosing. Each dimension value is evaluated
independently. If you start incurring cost in a new value of your chosen dimension, it
will automatically be analyzed by an AWS managed monitor.

Set this to `CUSTOM` for a customer managed monitor. Customer managed
monitors let you select specific dimension values that get monitored in aggregate.

For more information about monitor types, see [Monitor\
types](../../../cost-management/latest/userguide/getting-started-ad.md#monitor-type-def) in the _Billing and Cost Management User Guide_.

_Required_: Yes

_Type_: String

_Allowed values_: `DIMENSIONAL | CUSTOM`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ResourceTags`

Property description not available.

_Required_: No

_Type_: Array of [ResourceTag](aws-properties-ce-anomalymonitor-resourcetag.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns `MonitorArn`. For example:

`{ "Ref": "myAnomalySubscriptionLogicalName" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreationDate`

The date when the monitor was created.

`DimensionalValueCount`

The value for evaluated dimensions.

`LastEvaluatedDate`

The date when the monitor last evaluated for anomalies.

`LastUpdatedDate`

The date when the monitor was last updated.

`MonitorArn`

The Amazon Resource Name (ARN) value for the monitor.

## Examples

- [Service monitor](#aws-resource-ce-anomalymonitor--examples--Service_monitor)

- [Anomaly monitor with tags](#aws-resource-ce-anomalymonitor--examples--Anomaly_monitor_with_tags)

- [Anomaly monitor with Cost Category](#aws-resource-ce-anomalymonitor--examples--Anomaly_monitor_with_Cost_Category)

- [Anomaly monitor with linked account](#aws-resource-ce-anomalymonitor--examples--Anomaly_monitor_with_linked_account)

- [Attaching subscriptions to monitors](#aws-resource-ce-anomalymonitor--examples--Attaching_subscriptions_to_monitors)

### Service monitor

The following example creates a service anomaly monitor.

#### JSON

```json

{
  "Resources": {
    "AnomalyServiceMonitor": {
      "Type": "AWS::CE::AnomalyMonitor",
      "Properties": {
        "MonitorName": "MonitorName",
        "MonitorType": "DIMENSIONAL",
        "MonitorDimension": "SERVICE"
      }
    }
  }
}

```

#### YAML

```yaml

Resources:
  AnomalyServiceMonitor:
    Type: 'AWS::CE::AnomalyMonitor'
    Properties:
      MonitorName: 'MonitorName'
      MonitorType: 'DIMENSIONAL'
      MonitorDimension: 'SERVICE'

```

### Anomaly monitor with tags

The following example creates a custom anomaly monitor with tags.

#### JSON

```json

{
  "Resources": {
    "CustomAnomalyMonitorWithTags": {
      "Type": "AWS::CE::AnomalyMonitor",
      "Properties": {
        "MonitorName": "MonitorName",
        "MonitorType": "CUSTOM",
        "MonitorSpecification": " { \"Tags\" : { \"Key\" : \"Tag Key\", \"Values\" : [ \"TagValue1\", \"TagValue2\" ] } }"
      }
    }
  }
}

```

#### YAML

```yaml

Resources:
  CustomAnomalyMonitorWithTags:
    Type: 'AWS::CE::AnomalyMonitor'
    Properties:
      MonitorName: "MonitorName"
      MonitorType: "CUSTOM"
      MonitorSpecification: '
            {
              "Tags" : {
                "Key" : "Tag Key",
                "Values" : [ "TagValue1", "TagValue2" ]
              }
            }'

```

### Anomaly monitor with Cost Category

The following example creates a custom anomaly monitor with a Cost Category.

#### JSON

```json

{
  "Resources": {
    "CustomAnomalyMonitorWithCC": {
      "Type": "AWS::CE::AnomalyMonitor",
      "Properties": {
        "MonitorName": "MonitorName",
        "MonitorType": "CUSTOM",
        "MonitorSpecification": " { \"CostCategories\" : { \"Key\" : \"CostCategoryKey\", \"Values\" : [ \"CostCategoryValue\" ] } }"
      }
    }
  }
}

```

#### YAML

```yaml

Resources:
   CustomAnomalyMonitorWithCC:
    Type: 'AWS::CE::AnomalyMonitor'
    Properties:
      MonitorName: "MonitorName"
      MonitorType: "CUSTOM"
      MonitorSpecification: '
            {
              "CostCategories" : {
                "Key" : "CostCategoryKey",
                "Values" : [ "CostCategoryValue" ]
              }
            }'

```

### Anomaly monitor with linked account

The following example creates a custom anomaly monitor with a linked account.

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

```

### Attaching subscriptions to monitors

The following example shows two anomaly monitors attached to an anomaly
subscription.

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
         { “Ref”: "CustomAnomalyMonitorWithLinkedAccount"},
         { "Ref": "AnomalyServiceMonitor"}
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

## See also

- [AnomalyMonitor](../../../../reference/aws-cost-management/latest/apireference/api-anomalymonitor.md) in the _AWS Billing and Cost Management API_
_Reference_.

- [CreateAnomalyMonitor](../../../../reference/aws-cost-management/latest/apireference/api-createanomalymonitor.md) in the _AWS Billing and Cost Management API_
_Reference_.

- [DeleteAnomalyMonitor](../../../../reference/aws-cost-management/latest/apireference/api-deleteanomalymonitor.md) in the _AWS Billing and Cost Management API_
_Reference_.

- [GetAnomalyMonitors](../../../../reference/aws-cost-management/latest/apireference/api-getanomalymonitors.md) in the _AWS Billing and Cost Management API_
_Reference_.

- [UpdateAnomalyMonitor](../../../../reference/aws-cost-management/latest/apireference/api-updateanomalymonitor.md) in the _AWS Billing and Cost Management API_
_Reference_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS Cost Explorer

ResourceTag

All content copied from https://docs.aws.amazon.com/.
