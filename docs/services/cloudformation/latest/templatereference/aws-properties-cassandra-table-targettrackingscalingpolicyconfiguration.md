This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cassandra::Table TargetTrackingScalingPolicyConfiguration

Amazon Keyspaces supports the `target tracking` auto scaling policy for a provisioned table. This policy
scales a table based on the ratio of consumed to provisioned capacity.
The auto scaling target is a percentage of the provisioned capacity of the table.

- `targetTrackingScalingPolicyConfiguration`: To define the target tracking policy, you must define the target value.

- `targetValue`: The target utilization rate of the table. Amazon Keyspaces auto scaling ensures that the ratio of
consumed capacity to provisioned capacity stays at or near this value. You
define `targetValue` as a percentage. A `double` between 20 and 90. (Required)

- `disableScaleIn`: A `boolean` that specifies if `scale-in` is
disabled or enabled for the table. This parameter is disabled by default.
To turn on `scale-in`, set the `boolean` value to
`FALSE`. This means that capacity for a table can be
automatically scaled down on your behalf. (Optional)

- `scaleInCooldown`: A cooldown period in seconds between scaling activities that lets the table stabilize
before another scale in activity starts. If no value is provided, the default is 0. (Optional)

- `scaleOutCooldown`: A cooldown period in seconds between scaling activities that lets the table stabilize
before another scale out activity starts. If no value is provided, the default is 0. (Optional)

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DisableScaleIn" : Boolean,
  "ScaleInCooldown" : Integer,
  "ScaleOutCooldown" : Integer,
  "TargetValue" : Integer
}

```

### YAML

```yaml

  DisableScaleIn: Boolean
  ScaleInCooldown: Integer
  ScaleOutCooldown: Integer
  TargetValue: Integer

```

## Properties

`DisableScaleIn`

Specifies if `scale-in` is enabled.

When auto scaling automatically decreases capacity for a table,
the table _scales in_. When scaling policies are set, they can't
scale in the table lower than its minimum capacity.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScaleInCooldown`

Specifies a `scale-in` cool down period.

A cooldown period in seconds between scaling activities that lets the table stabilize before another scaling activity starts.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScaleOutCooldown`

Specifies a scale out cool down period.

A cooldown period in seconds between scaling activities that lets the table stabilize before another scaling activity starts.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetValue`

Specifies the target value for the target tracking auto scaling policy.

Amazon Keyspaces auto scaling scales up capacity automatically when traffic exceeds this target utilization
rate, and then back down when it falls below the target. This ensures that the ratio of
consumed capacity to provisioned capacity stays at or near this value. You
define `targetValue` as a percentage. An `integer` between 20 and 90.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

WarmThroughput

All content copied from https://docs.aws.amazon.com/.
