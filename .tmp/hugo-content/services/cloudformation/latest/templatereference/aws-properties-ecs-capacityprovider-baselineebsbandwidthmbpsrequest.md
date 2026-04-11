This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::CapacityProvider BaselineEbsBandwidthMbpsRequest

The minimum and maximum baseline Amazon EBS bandwidth in megabits per second (Mbps)
for instance type selection. This is important for workloads with high storage I/O
requirements.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Max" : Integer,
  "Min" : Integer
}

```

### YAML

```yaml

  Max: Integer
  Min: Integer

```

## Properties

`Max`

The maximum baseline Amazon EBS bandwidth in Mbps. Instance types with higher Amazon
EBS bandwidth are excluded from selection.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Min`

The minimum baseline Amazon EBS bandwidth in Mbps. Instance types with lower Amazon
EBS bandwidth are excluded from selection.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AutoScalingGroupProvider

CapacityReservationRequest

All content copied from https://docs.aws.amazon.com/.
