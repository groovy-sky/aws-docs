This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::CapacityProvider MemoryMiBRequest

The minimum and maximum amount of memory in mebibytes (MiB) for instance type
selection. This ensures that selected instance types have adequate memory for your
workloads.

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

The maximum amount of memory in MiB. Instance types with more memory than this value
are excluded from selection.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Min`

The minimum amount of memory in MiB. Instance types with less memory than this value
are excluded from selection.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MemoryGiBPerVCpuRequest

NetworkBandwidthGbpsRequest

All content copied from https://docs.aws.amazon.com/.
