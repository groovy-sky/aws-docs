This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::CapacityProvider MemoryGiBPerVCpuRequest

The minimum and maximum amount of memory per vCPU in gibibytes (GiB). This helps
ensure that instance types have the appropriate memory-to-CPU ratio for your
workloads.

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

The maximum amount of memory per vCPU in GiB. Instance types with a higher
memory-to-vCPU ratio are excluded from selection.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Min`

The minimum amount of memory per vCPU in GiB. Instance types with a lower
memory-to-vCPU ratio are excluded from selection.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ManagedScaling

MemoryMiBRequest

All content copied from https://docs.aws.amazon.com/.
