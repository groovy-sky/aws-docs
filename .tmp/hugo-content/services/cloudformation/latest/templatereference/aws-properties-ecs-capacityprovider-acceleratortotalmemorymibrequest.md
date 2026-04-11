This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::CapacityProvider AcceleratorTotalMemoryMiBRequest

The minimum and maximum total accelerator memory in mebibytes (MiB) for instance type
selection. This is important for GPU workloads that require specific amounts of video
memory.

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

The maximum total accelerator memory in MiB. Instance types with more accelerator
memory are excluded from selection.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Min`

The minimum total accelerator memory in MiB. Instance types with less accelerator
memory are excluded from selection.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AcceleratorCountRequest

AutoScalingGroupProvider

All content copied from https://docs.aws.amazon.com/.
