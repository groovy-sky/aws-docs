This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AutoScaling::AutoScalingGroup LaunchTemplate

Use this structure to specify the launch templates and instance types (overrides) for a
mixed instances policy.

`LaunchTemplate` is a property of the [AWS::AutoScaling::AutoScalingGroup MixedInstancesPolicy](../userguide/aws-properties-autoscaling-autoscalinggroup-mixedinstancespolicy.md) property type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LaunchTemplateSpecification" : LaunchTemplateSpecification,
  "Overrides" : [ LaunchTemplateOverrides, ... ]
}

```

### YAML

```yaml

  LaunchTemplateSpecification:
    LaunchTemplateSpecification
  Overrides:
    - LaunchTemplateOverrides

```

## Properties

`LaunchTemplateSpecification`

The launch template.

_Required_: Yes

_Type_: [LaunchTemplateSpecification](aws-properties-autoscaling-autoscalinggroup-launchtemplatespecification.md)

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`Overrides`

Any properties that you specify override the same properties in the launch
template.

_Required_: No

_Type_: Array of [LaunchTemplateOverrides](aws-properties-autoscaling-autoscalinggroup-launchtemplateoverrides.md)

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InstancesDistribution

LaunchTemplateOverrides

All content copied from https://docs.aws.amazon.com/.
