This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::CapacityProvider NetworkBandwidthGbpsRequest

The minimum and maximum network bandwidth in gigabits per second (Gbps) for instance
type selection. This is important for network-intensive workloads.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Max" : Number,
  "Min" : Number
}

```

### YAML

```yaml

  Max: Number
  Min: Number

```

## Properties

`Max`

The maximum network bandwidth in Gbps. Instance types with higher network bandwidth
are excluded from selection.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Min`

The minimum network bandwidth in Gbps. Instance types with lower network bandwidth are
excluded from selection.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MemoryMiBRequest

NetworkInterfaceCountRequest

All content copied from https://docs.aws.amazon.com/.
