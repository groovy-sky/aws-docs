This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMR::InstanceGroupConfig AutoScalingPolicy

`AutoScalingPolicy` defines how an instance group dynamically adds and terminates EC2 instances in response to the value of a CloudWatch metric. For more information, see [Using Automatic Scaling in Amazon EMR](../../../emr/latest/managementguide/emr-automatic-scaling.md) in the _Amazon EMR Management Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Constraints" : ScalingConstraints,
  "Rules" : [ ScalingRule, ... ]
}

```

### YAML

```yaml

  Constraints:
    ScalingConstraints
  Rules:
    - ScalingRule

```

## Properties

`Constraints`

The upper and lower Amazon EC2 instance limits for an automatic scaling policy.
Automatic scaling activity will not cause an instance group to grow above or below these
limits.

_Required_: Yes

_Type_: [ScalingConstraints](aws-properties-emr-instancegroupconfig-scalingconstraints.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Rules`

The scale-in and scale-out rules that comprise the automatic scaling policy.

_Required_: Yes

_Type_: Array of [ScalingRule](aws-properties-emr-instancegroupconfig-scalingrule.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::EMR::InstanceGroupConfig

CloudWatchAlarmDefinition

All content copied from https://docs.aws.amazon.com/.
