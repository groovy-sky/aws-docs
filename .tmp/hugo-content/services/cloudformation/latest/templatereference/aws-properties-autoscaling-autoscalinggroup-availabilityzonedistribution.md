This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AutoScaling::AutoScalingGroup AvailabilityZoneDistribution

`AvailabilityZoneDistribution` is a property of the [AWS::AutoScaling::AutoScalingGroup](../userguide/aws-resource-autoscaling-autoscalinggroup.md)
resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CapacityDistributionStrategy" : String
}

```

### YAML

```yaml

  CapacityDistributionStrategy: String

```

## Properties

`CapacityDistributionStrategy`

If launches fail in an Availability Zone, the following strategies are available. The default is `balanced-best-effort`.

- `balanced-only` \- If launches fail in an Availability Zone, Auto Scaling will continue to attempt to launch in the unhealthy zone to preserve a balanced distribution.

- `balanced-best-effort` \- If launches fail in an Availability Zone, Auto Scaling will attempt to launch in another healthy Availability Zone instead.

_Required_: No

_Type_: String

_Allowed values_: `balanced-best-effort | balanced-only`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AcceleratorTotalMemoryMiBRequest

AvailabilityZoneImpairmentPolicy

All content copied from https://docs.aws.amazon.com/.
