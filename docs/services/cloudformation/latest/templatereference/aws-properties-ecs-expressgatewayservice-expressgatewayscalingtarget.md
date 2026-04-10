This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::ExpressGatewayService ExpressGatewayScalingTarget

Defines the auto-scaling configuration for an Express service. This determines how the
service automatically adjusts the number of running tasks based on demand metrics such as
CPU utilization, memory utilization, or request count per target.

Auto-scaling helps ensure your application can handle varying levels of traffic while
optimizing costs by scaling down during low-demand periods. You can specify the minimum and
maximum number of tasks, the scaling metric, and the target value for that metric.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AutoScalingMetric" : String,
  "AutoScalingTargetValue" : Integer,
  "MaxTaskCount" : Integer,
  "MinTaskCount" : Integer
}

```

### YAML

```yaml

  AutoScalingMetric: String
  AutoScalingTargetValue: Integer
  MaxTaskCount: Integer
  MinTaskCount: Integer

```

## Properties

`AutoScalingMetric`

The metric used for auto-scaling decisions. The default metric used for an Express service is `CPUUtilization`.

_Required_: No

_Type_: String

_Allowed values_: `AVERAGE_CPU | AVERAGE_MEMORY | REQUEST_COUNT_PER_TARGET`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AutoScalingTargetValue`

The target value for the auto-scaling metric. The default value for an Express service is 60.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxTaskCount`

The maximum number of tasks to run in the Express service.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinTaskCount`

The minimum number of tasks to run in the Express service.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ExpressGatewayRepositoryCredentials

ExpressGatewayServiceAwsLogsConfiguration

All content copied from https://docs.aws.amazon.com/.
