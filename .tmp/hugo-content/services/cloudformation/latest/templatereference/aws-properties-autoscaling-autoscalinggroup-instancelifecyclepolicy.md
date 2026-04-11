This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AutoScaling::AutoScalingGroup InstanceLifecyclePolicy

The instance lifecycle policy for the Auto Scaling group. This policy controls instance behavior when an instance
transitions through its lifecycle states. Configure retention triggers to specify when instances should
move to a `Retained` state instead of automatic termination.

For more information, see
[Control instance retention with instance lifecycle policies](../../../autoscaling/ec2/userguide/instance-lifecycle-policy.md)
in the _Amazon EC2 Auto Scaling User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RetentionTriggers" : RetentionTriggers
}

```

### YAML

```yaml

  RetentionTriggers:
    RetentionTriggers

```

## Properties

`RetentionTriggers`

Specifies the conditions that trigger instance retention behavior. These triggers determine when instances
should move to a `Retained` state instead of automatic termination. This allows you to maintain
control over instance management when lifecycles transition and operations fail.

_Required_: No

_Type_: [RetentionTriggers](aws-properties-autoscaling-autoscalinggroup-retentiontriggers.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CpuPerformanceFactorRequest

InstanceMaintenancePolicy

All content copied from https://docs.aws.amazon.com/.
