This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DynamoDB::GlobalTable TargetTrackingScalingPolicyConfiguration

Defines a target tracking scaling policy.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DisableScaleIn" : Boolean,
  "ScaleInCooldown" : Integer,
  "ScaleOutCooldown" : Integer,
  "TargetValue" : Number
}

```

### YAML

```yaml

  DisableScaleIn: Boolean
  ScaleInCooldown: Integer
  ScaleOutCooldown: Integer
  TargetValue: Number

```

## Properties

`DisableScaleIn`

Indicates whether scale in by the target tracking scaling policy is disabled. The
default value is `false`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScaleInCooldown`

The amount of time, in seconds, after a scale-in activity completes before another
scale-in activity can start.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScaleOutCooldown`

The amount of time, in seconds, after a scale-out activity completes before another
scale-out activity can start.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetValue`

Defines a target value for the scaling policy.

_Required_: Yes

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

TimeToLiveSpecification

All content copied from https://docs.aws.amazon.com/.
