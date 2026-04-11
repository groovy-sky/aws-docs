This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GameLift::GameServerGroup TargetTrackingConfiguration

**This data type is used with the Amazon GameLift FleetIQ and game server groups.**

Settings for a target-based scaling policy as part of a `GameServerGroupAutoScalingPolicy`.
These settings are used to
create a target-based policy that tracks the GameLift FleetIQ metric
`"PercentUtilizedGameServers"` and specifies a target value for the
metric. As player usage changes, the policy triggers to adjust the game server group
capacity so that the metric returns to the target value.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "TargetValue" : Number
}

```

### YAML

```yaml

  TargetValue: Number

```

## Properties

`TargetValue`

Desired value to use with a game server group target-based scaling policy.

_Required_: Yes

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::GameLift::GameSessionQueue

All content copied from https://docs.aws.amazon.com/.
