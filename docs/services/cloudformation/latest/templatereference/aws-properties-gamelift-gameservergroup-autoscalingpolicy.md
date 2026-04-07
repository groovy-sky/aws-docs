This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GameLift::GameServerGroup AutoScalingPolicy

**This data type is used with the GameLift FleetIQ and game server groups.**

Configuration settings for intelligent automatic scaling that uses target tracking.
After the Auto Scaling group is created, all updates to Auto Scaling policies, including
changing this policy and adding or removing other policies, is done directly on the Auto
Scaling group.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EstimatedInstanceWarmup" : Number,
  "TargetTrackingConfiguration" : TargetTrackingConfiguration
}

```

### YAML

```yaml

  EstimatedInstanceWarmup: Number
  TargetTrackingConfiguration:
    TargetTrackingConfiguration

```

## Properties

`EstimatedInstanceWarmup`

Length of time, in seconds, it takes for a new instance to start new game server
processes and register with Amazon GameLift Servers FleetIQ. Specifying a warm-up time can be useful, particularly
with game servers that take a long time to start up, because it avoids prematurely
starting new instances.

_Required_: No

_Type_: Number

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetTrackingConfiguration`

Settings for a target-based scaling policy applied to Auto Scaling group.
These settings are used to create a target-based policy that tracks the GameLift
FleetIQ metric `PercentUtilizedGameServers` and specifies a target value
for the metric. As player usage changes, the policy triggers to adjust the game server group
capacity so that the metric returns to the target value.

_Required_: Yes

_Type_: [TargetTrackingConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-gamelift-gameservergroup-targettrackingconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::GameLift::GameServerGroup

InstanceDefinition
