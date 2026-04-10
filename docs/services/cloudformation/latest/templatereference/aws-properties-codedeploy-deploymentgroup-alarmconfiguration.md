This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeDeploy::DeploymentGroup AlarmConfiguration

The `AlarmConfiguration` property type configures CloudWatch alarms
for an AWS CodeDeploy deployment group. `AlarmConfiguration` is a
property of the [DeploymentGroup](../userguide/aws-resource-codedeploy-deploymentgroup.md) resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Alarms" : [ Alarm, ... ],
  "Enabled" : Boolean,
  "IgnorePollAlarmFailure" : Boolean
}

```

### YAML

```yaml

  Alarms:
    - Alarm
  Enabled: Boolean
  IgnorePollAlarmFailure: Boolean

```

## Properties

`Alarms`

A list of alarms configured for the deployment or deployment group. A maximum of 10
alarms can be added.

_Required_: No

_Type_: Array of [Alarm](aws-properties-codedeploy-deploymentgroup-alarm.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Enabled`

Indicates whether the alarm configuration is enabled.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IgnorePollAlarmFailure`

Indicates whether a deployment should continue if information about the current state of
alarms cannot be retrieved from Amazon CloudWatch. The default value is
`false`.

- `true`: The deployment proceeds even if alarm status information can't be
retrieved from CloudWatch.

- `false`: The deployment stops if alarm status information can't be retrieved
from CloudWatch.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Alarm

AutoRollbackConfiguration

All content copied from https://docs.aws.amazon.com/.
