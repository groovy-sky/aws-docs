This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMR::InstanceGroupConfig ScalingRule

`ScalingRule` is a subproperty of the `AutoScalingPolicy` property type. `ScalingRule` defines the scale-in or scale-out rules for scaling activity, including the CloudWatch metric alarm that triggers activity, how EC2 instances are added or removed, and the periodicity of adjustments. The automatic scaling policy for an instance group can comprise one or more automatic scaling rules.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Action" : ScalingAction,
  "Description" : String,
  "Name" : String,
  "Trigger" : ScalingTrigger
}

```

### YAML

```yaml

  Action:
    ScalingAction
  Description: String
  Name: String
  Trigger:
    ScalingTrigger

```

## Properties

`Action`

The conditions that trigger an automatic scaling activity.

_Required_: Yes

_Type_: [ScalingAction](aws-properties-emr-instancegroupconfig-scalingaction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A friendly, more verbose description of the automatic scaling rule.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name used to identify an automatic scaling rule. Rule names must be unique within a
scaling policy.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Trigger`

The CloudWatch alarm definition that determines when automatic scaling activity is
triggered.

_Required_: Yes

_Type_: [ScalingTrigger](aws-properties-emr-instancegroupconfig-scalingtrigger.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ScalingConstraints

ScalingTrigger

All content copied from https://docs.aws.amazon.com/.
