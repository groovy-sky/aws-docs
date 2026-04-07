This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AutoScaling::AutoScalingGroup MetricsCollection

`MetricsCollection` is a property of the [AWS::AutoScaling::AutoScalingGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-autoscalinggroup.html) resource that describes the group metrics that
an Amazon EC2 Auto Scaling group sends to Amazon CloudWatch. These metrics describe the group
rather than any of its instances.

For more information, see [Monitor CloudWatch metrics for\
your Auto Scaling groups and instances](https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-instance-monitoring.html) in the _Amazon EC2 Auto Scaling User_
_Guide_. You can find a sample template snippet in the [Examples](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-autoscalinggroup.html#aws-resource-autoscaling-autoscalinggroup--examples) section of the `AWS::AutoScaling::AutoScalingGroup`
resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Granularity" : String,
  "Metrics" : [ String, ... ]
}

```

### YAML

```yaml

  Granularity: String
  Metrics:
    - String

```

## Properties

`Granularity`

The frequency at which Amazon EC2 Auto Scaling sends aggregated data to CloudWatch. The only valid value is
`1Minute`.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Metrics`

Identifies the metrics to enable.

You can specify one or more of the following metrics:

- `GroupMinSize`

- `GroupMaxSize`

- `GroupDesiredCapacity`

- `GroupInServiceInstances`

- `GroupPendingInstances`

- `GroupStandbyInstances`

- `GroupTerminatingInstances`

- `GroupTotalInstances`

- `GroupInServiceCapacity`

- `GroupPendingCapacity`

- `GroupStandbyCapacity`

- `GroupTerminatingCapacity`

- `GroupTotalCapacity`

- `WarmPoolDesiredCapacity`

- `WarmPoolWarmedCapacity`

- `WarmPoolPendingCapacity`

- `WarmPoolTerminatingCapacity`

- `WarmPoolTotalCapacity`

- `GroupAndWarmPoolDesiredCapacity`

- `GroupAndWarmPoolTotalCapacity`

If you specify `Granularity` and don't specify any metrics, all metrics are
enabled.

For more information, see [Amazon CloudWatch metrics for\
Amazon EC2 Auto Scaling](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-metrics.html) in the _Amazon EC2 Auto Scaling User Guide_.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

MemoryMiBRequest

MixedInstancesPolicy
