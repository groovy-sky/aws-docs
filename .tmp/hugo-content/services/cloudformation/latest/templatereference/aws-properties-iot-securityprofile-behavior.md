This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::SecurityProfile Behavior

A Device Defender security profile behavior.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Criteria" : BehaviorCriteria,
  "ExportMetric" : Boolean,
  "Metric" : String,
  "MetricDimension" : MetricDimension,
  "Name" : String,
  "SuppressAlerts" : Boolean
}

```

### YAML

```yaml

  Criteria:
    BehaviorCriteria
  ExportMetric: Boolean
  Metric: String
  MetricDimension:
    MetricDimension
  Name: String
  SuppressAlerts: Boolean

```

## Properties

`Criteria`

The criteria that determine if a device is behaving normally in regard to the `metric`.

###### Note

In the AWS IoT console, you can choose to be sent an alert through Amazon SNS when AWS IoT Device Defender detects that a device is behaving anomalously.

_Required_: No

_Type_: [BehaviorCriteria](aws-properties-iot-securityprofile-behaviorcriteria.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExportMetric`

Value indicates exporting metrics related to the behavior when it is true.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Metric`

What is measured by the behavior.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9:_-]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MetricDimension`

The dimension of the metric.

_Required_: No

_Type_: [MetricDimension](aws-properties-iot-securityprofile-metricdimension.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name
you've given to the behavior.

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z0-9:_-]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SuppressAlerts`

The alert status. If you set the value to `true`, alerts will be
suppressed.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AlertTarget

BehaviorCriteria

All content copied from https://docs.aws.amazon.com/.
