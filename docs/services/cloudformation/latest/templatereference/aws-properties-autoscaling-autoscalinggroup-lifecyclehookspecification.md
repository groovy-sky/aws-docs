This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AutoScaling::AutoScalingGroup LifecycleHookSpecification

`LifecycleHookSpecification` specifies a lifecycle hook for the
`LifecycleHookSpecificationList` property of the [AWS::AutoScaling::AutoScalingGroup](../userguide/aws-resource-autoscaling-autoscalinggroup.md) resource. A lifecycle hook specifies actions to
perform when Amazon EC2 Auto Scaling launches or terminates instances.

For more information, see [Amazon EC2 Auto Scaling lifecycle\
hooks](../../../autoscaling/ec2/userguide/lifecycle-hooks.md) in the _Amazon EC2 Auto Scaling User Guide_. You can find a
sample template snippet in the [Examples](../userguide/aws-resource-as-lifecyclehook.md#aws-resource-as-lifecyclehook--examples) section of the `AWS::AutoScaling::LifecycleHook`
resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DefaultResult" : String,
  "HeartbeatTimeout" : Integer,
  "LifecycleHookName" : String,
  "LifecycleTransition" : String,
  "NotificationMetadata" : String,
  "NotificationTargetARN" : String,
  "RoleARN" : String
}

```

### YAML

```yaml

  DefaultResult: String
  HeartbeatTimeout: Integer
  LifecycleHookName: String
  LifecycleTransition: String
  NotificationMetadata: String
  NotificationTargetARN: String
  RoleARN: String

```

## Properties

`DefaultResult`

The action the Auto Scaling group takes when the lifecycle hook timeout elapses or if an
unexpected failure occurs. The default value is `ABANDON`.

Valid values: `CONTINUE` \| `ABANDON`

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HeartbeatTimeout`

The maximum time, in seconds, that can elapse before the lifecycle hook times out. The
range is from `30` to `7200` seconds. The default value is
`3600` seconds (1 hour).

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LifecycleHookName`

The name of the lifecycle hook.

_Required_: Yes

_Type_: String

_Pattern_: `[A-Za-z0-9\-_\/]+`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LifecycleTransition`

The lifecycle transition. For Auto Scaling groups, there are two major lifecycle
transitions.

- To create a lifecycle hook for scale-out events, specify
`autoscaling:EC2_INSTANCE_LAUNCHING`.

- To create a lifecycle hook for scale-in events, specify
`autoscaling:EC2_INSTANCE_TERMINATING`.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NotificationMetadata`

Additional information that you want to include any time Amazon EC2 Auto Scaling sends a message to
the notification target.

_Required_: No

_Type_: String

_Pattern_: `[\u0009\u000A\u000D\u0020-\u007e]+`

_Minimum_: `1`

_Maximum_: `4000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NotificationTargetARN`

The Amazon Resource Name (ARN) of the notification target that Amazon EC2 Auto Scaling sends
notifications to when an instance is in a wait state for the lifecycle hook. You can
specify an Amazon SNS topic or an Amazon SQS queue.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleARN`

The ARN of the IAM role that allows the Auto Scaling group to publish to the specified
notification target. For information about creating this role, see [Prepare to\
add a lifecycle hook to your Auto Scaling group](../../../autoscaling/ec2/userguide/prepare-for-lifecycle-notifications.md) in the
_Amazon EC2 Auto Scaling User Guide_.

Valid only if the notification target is an Amazon SNS topic or an Amazon SQS queue.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LaunchTemplateSpecification

MemoryGiBPerVCpuRequest

All content copied from https://docs.aws.amazon.com/.
