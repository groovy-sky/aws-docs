This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApplicationAutoScaling::ScalableTarget ScalableTargetAction

`ScalableTargetAction` specifies the minimum and maximum capacity for the
`ScalableTargetAction` property of the [AWS::ApplicationAutoScaling::ScalableTarget ScheduledAction](../userguide/aws-properties-applicationautoscaling-scalabletarget-scheduledaction.md) property type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaxCapacity" : Integer,
  "MinCapacity" : Integer
}

```

### YAML

```yaml

  MaxCapacity: Integer
  MinCapacity: Integer

```

## Properties

`MaxCapacity`

The maximum capacity.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinCapacity`

The minimum capacity.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [Configure Application Auto Scaling resources](../userguide/quickref-application-auto-scaling.md)

- [Getting started](../../../autoscaling/application/userguide/getting-started.md)
in the _Application Auto Scaling User Guide_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::ApplicationAutoScaling::ScalableTarget

ScheduledAction

All content copied from https://docs.aws.amazon.com/.
