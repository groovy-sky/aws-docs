This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::Daemon DaemonDeploymentConfiguration

Optional deployment parameters that control how a daemon rolls out updates across
container instances.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Alarms" : DaemonAlarmConfiguration,
  "BakeTimeInMinutes" : Integer,
  "DrainPercent" : Number
}

```

### YAML

```yaml

  Alarms:
    DaemonAlarmConfiguration
  BakeTimeInMinutes: Integer
  DrainPercent: Number

```

## Properties

`Alarms`

The CloudWatch alarm configuration for the daemon deployment. When alarms are
triggered during a deployment, the deployment can be automatically rolled back.

_Required_: No

_Type_: [DaemonAlarmConfiguration](aws-properties-ecs-daemon-daemonalarmconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BakeTimeInMinutes`

The amount of time (in minutes) to wait after a successful deployment step before
proceeding. This allows time to monitor for issues before continuing. The default value
is 0.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `1440`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DrainPercent`

The percentage of container instances to drain simultaneously during a daemon
deployment. Valid values are between 0.0 and 100.0.

_Required_: No

_Type_: Number

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DaemonAlarmConfiguration

Tag

All content copied from https://docs.aws.amazon.com/.
