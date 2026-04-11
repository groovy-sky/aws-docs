This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::CapacityProvider InfrastructureOptimization

The configuration that controls how Amazon ECS optimizes your infrastructure.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ScaleInAfter" : Integer
}

```

### YAML

```yaml

  ScaleInAfter: Integer

```

## Properties

`ScaleInAfter`

This parameter defines the number of seconds Amazon ECS Managed Instances waits before optimizing EC2 instances that have become idle or underutilized.
A longer delay increases the likelihood of placing new tasks on idle or underutilized instances instances, reducing startup time.
A shorter delay helps reduce infrastructure costs by optimizing idle or underutilized instances,instances more quickly.

Valid values are:

- `null` \- Uses the default optimization behavior.

- `-1` \- Disables automatic infrastructure optimization.

- A value between `0` and `3600` (inclusive) - Specifies the number of seconds to wait before optimizing instances.

_Required_: No

_Type_: Integer

_Minimum_: `-1`

_Maximum_: `3600`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CapacityReservationRequest

InstanceLaunchTemplate

All content copied from https://docs.aws.amazon.com/.
