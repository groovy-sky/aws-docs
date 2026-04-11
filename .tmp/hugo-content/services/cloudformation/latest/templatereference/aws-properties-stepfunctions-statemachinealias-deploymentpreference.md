This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::StepFunctions::StateMachineAlias DeploymentPreference

Enables gradual state machine deployments.
CloudFormation
automatically shifts traffic from the version the alias currently points to, to a new state machine version that you specify.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Alarms" : [ String, ... ],
  "Interval" : Integer,
  "Percentage" : Integer,
  "StateMachineVersionArn" : String,
  "Type" : String
}

```

### YAML

```yaml

  Alarms:
    - String
  Interval: Integer
  Percentage: Integer
  StateMachineVersionArn: String
  Type: String

```

## Properties

`Alarms`

A list of
Amazon CloudWatch
alarm names to be monitored during the deployment. The deployment fails and rolls back if any of these alarms go into the `ALARM` state.

###### Important

Amazon CloudWatch
considers nonexistent alarms to have an `OK` state. If you provide an invalid alarm name or provide the ARN of an alarm instead of its name, your deployment may not roll back correctly.

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `256 | 100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Interval`

The time in minutes between each traffic shifting increment.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `2100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Percentage`

The percentage of traffic to shift to the new version in each increment.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `99`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StateMachineVersionArn`

The Amazon Resource Name (ARN) of the [`AWS::StepFunctions::StateMachineVersion`](../userguide/aws-resource-stepfunctions-statemachineversion.md) resource that will be the final version to which the alias points to when the traffic shifting is complete.

While performing gradual deployments, you can only provide a single state machine version ARN. To explicitly set version weights in a
CloudFormation
template, use `RoutingConfiguration` instead.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of deployment you want to perform. You can specify one of the following types:

- `LINEAR` \- Shifts traffic to the new version in equal increments with an equal number of minutes between each increment.

For example, if you specify the increment percent as `20` with an interval of `600` minutes, this deployment increases traffic by 20 percent every 600 minutes until the new version receives 100 percent of the traffic.

This deployment immediately rolls back the new version if any
CloudWatch
alarms are triggered.

- `ALL_AT_ONCE` \- Shifts 100 percent of traffic to the new version immediately.
CloudFormation
monitors the new version and rolls it back automatically to the previous version if any
CloudWatch
alarms are triggered.

- `CANARY` \- Shifts traffic in two increments.

In the first increment, a small percentage of traffic, for example, 10 percent is shifted to the new version. In the second increment, before a specified time interval in seconds gets over, the remaining traffic is shifted to the new version. The shift to the new version for the remaining traffic takes place only if no
CloudWatch
alarms are triggered during the specified time interval.

_Required_: Yes

_Type_: String

_Allowed values_: `LINEAR | ALL_AT_ONCE | CANARY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::StepFunctions::StateMachineAlias

RoutingConfigurationVersion

All content copied from https://docs.aws.amazon.com/.
