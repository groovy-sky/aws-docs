This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GameLift::Fleet TargetConfiguration

Settings for a target-based scaling policy. A target-based policy tracks a particular
fleet metric specifies a target value for the metric. As player usage changes, the
policy triggers Amazon GameLift Servers to adjust capacity so that the metric returns to the target
value. The target configuration specifies settings as needed for the target based
policy, including the target value.

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

Desired value to use with a target-based scaling policy. The value must be relevant
for whatever metric the scaling policy is using. For example, in a policy using the
metric PercentAvailableGameSessions, the target value should be the preferred size of
the fleet's buffer (the percent of capacity that should be idle and ready for new game
sessions).

_Required_: Yes

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::GameLift::GameServerGroup

All content copied from https://docs.aws.amazon.com/.
