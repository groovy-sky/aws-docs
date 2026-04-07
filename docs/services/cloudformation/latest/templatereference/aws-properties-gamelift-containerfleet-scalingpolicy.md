This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GameLift::ContainerFleet ScalingPolicy

Rule that controls how a fleet is scaled. Scaling policies are uniquely identified by
the combination of name and fleet ID.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ComparisonOperator" : String,
  "EvaluationPeriods" : Integer,
  "MetricName" : String,
  "Name" : String,
  "PolicyType" : String,
  "ScalingAdjustment" : Integer,
  "ScalingAdjustmentType" : String,
  "TargetConfiguration" : TargetConfiguration,
  "Threshold" : Number
}

```

### YAML

```yaml

  ComparisonOperator: String
  EvaluationPeriods: Integer
  MetricName: String
  Name: String
  PolicyType: String
  ScalingAdjustment: Integer
  ScalingAdjustmentType: String
  TargetConfiguration:
    TargetConfiguration
  Threshold: Number

```

## Properties

`ComparisonOperator`

Comparison operator to use when measuring a metric against the threshold value.

_Required_: No

_Type_: String

_Allowed values_: `GreaterThanOrEqualToThreshold | GreaterThanThreshold | LessThanThreshold | LessThanOrEqualToThreshold`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EvaluationPeriods`

Length of time (in minutes) the metric must be at or beyond the threshold before a
scaling event is triggered.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MetricName`

Name of the Amazon GameLift Servers-defined metric that is used to trigger a scaling adjustment. For
detailed descriptions of fleet metrics, see [Monitor Amazon GameLift Servers\
with Amazon CloudWatch](https://docs.aws.amazon.com/gamelift/latest/developerguide/monitoring-cloudwatch.html).

- **ActivatingGameSessions** \-\- Game sessions in
the process of being created.

- **ActiveGameSessions** \-\- Game sessions that
are currently running.

- **ActiveInstances** \-\- Fleet instances that
are currently running at least one game session.

- **AvailableGameSessions** \-\- Additional game
sessions that fleet could host simultaneously, given current capacity.

- **AvailablePlayerSessions** \-\- Empty player
slots in currently active game sessions. This includes game sessions that are
not currently accepting players. Reserved player slots are not
included.

- **CurrentPlayerSessions** \-\- Player slots in
active game sessions that are being used by a player or are reserved for a
player.

- **IdleInstances** \-\- Active instances that are
currently hosting zero game sessions.

- **PercentAvailableGameSessions** \-\- Unused
percentage of the total number of game sessions that a fleet could host
simultaneously, given current capacity. Use this metric for a target-based
scaling policy.

- **PercentIdleInstances** \-\- Percentage of the
total number of active instances that are hosting zero game sessions.

- **QueueDepth** \-\- Pending game session
placement requests, in any queue, where the current fleet is the top-priority
destination.

- **WaitTime** \-\- Current wait time for pending
game session placement requests, in any queue, where the current fleet is the
top-priority destination.

_Required_: Yes

_Type_: String

_Allowed values_: `ActivatingGameSessions | ActiveGameSessions | ActiveInstances | AvailableGameSessions | AvailablePlayerSessions | CurrentPlayerSessions | IdleInstances | PercentAvailableGameSessions | PercentIdleInstances | QueueDepth | WaitTime | ConcurrentActivatableGameSessions`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

A descriptive label that is associated with a fleet's scaling policy. Policy names do not need to be unique.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PolicyType`

The type of scaling policy to create. For a target-based policy, set the parameter
_MetricName_ to 'PercentAvailableGameSessions' and specify a
_TargetConfiguration_. For a rule-based policy set the following
parameters: _MetricName_, _ComparisonOperator_,
_Threshold_, _EvaluationPeriods_,
_ScalingAdjustmentType_, and
_ScalingAdjustment_.

_Required_: No

_Type_: String

_Allowed values_: `RuleBased | TargetBased`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScalingAdjustment`

Amount of adjustment to make, based on the scaling adjustment type.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScalingAdjustmentType`

The type of adjustment to make to a fleet's instance count.

- **ChangeInCapacity** \-\- add (or subtract) the
scaling adjustment value from the current instance count. Positive values scale
up while negative values scale down.

- **ExactCapacity** \-\- set the instance count to the
scaling adjustment value.

- **PercentChangeInCapacity** \-\- increase or reduce
the current instance count by the scaling adjustment, read as a percentage.
Positive values scale up while negative values scale down.

_Required_: No

_Type_: String

_Allowed values_: `ChangeInCapacity | ExactCapacity | PercentChangeInCapacity`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetConfiguration`

An object that contains settings for a target-based scaling policy.

_Required_: No

_Type_: [TargetConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-gamelift-containerfleet-targetconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Threshold`

Metric value used to trigger a scaling event.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ManagedCapacityConfiguration

Tag
