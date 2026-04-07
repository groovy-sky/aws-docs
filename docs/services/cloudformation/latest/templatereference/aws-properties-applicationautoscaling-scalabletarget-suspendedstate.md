This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApplicationAutoScaling::ScalableTarget SuspendedState

`SuspendedState` is a property of the [AWS::ApplicationAutoScaling::ScalableTarget](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalabletarget.html) resource that specifies whether the
scaling activities for a scalable target are in a suspended state.

For more information, see [Suspending and resuming scaling](https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-suspend-resume-scaling.html) in the _Application Auto Scaling User_
_Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DynamicScalingInSuspended" : Boolean,
  "DynamicScalingOutSuspended" : Boolean,
  "ScheduledScalingSuspended" : Boolean
}

```

### YAML

```yaml

  DynamicScalingInSuspended: Boolean
  DynamicScalingOutSuspended: Boolean
  ScheduledScalingSuspended: Boolean

```

## Properties

`DynamicScalingInSuspended`

Whether scale in by a target tracking scaling policy or a step scaling policy is
suspended. Set the value to `true` if you don't want Application Auto Scaling to
remove capacity when a scaling policy is triggered. The default is `false`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DynamicScalingOutSuspended`

Whether scale out by a target tracking scaling policy or a step scaling policy is
suspended. Set the value to `true` if you don't want Application Auto Scaling to
add capacity when a scaling policy is triggered. The default is `false`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScheduledScalingSuspended`

Whether scheduled scaling is suspended. Set the value to `true` if you don't
want Application Auto Scaling to add or remove capacity by initiating scheduled actions. The
default is `false`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ScheduledAction

AWS::ApplicationAutoScaling::ScalingPolicy
