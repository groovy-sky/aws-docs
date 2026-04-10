This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::CapacityProvider VCpuCountRangeRequest

The minimum and maximum number of vCPUs for instance type selection. This allows you
to specify a range of vCPU counts that meet your workload requirements.

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

The maximum number of vCPUs. Instance types with more vCPUs than this value are
excluded from selection.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Min`

The minimum number of vCPUs. Instance types with fewer vCPUs than this value are
excluded from selection.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TotalLocalStorageGBRequest

AWS::ECS::Cluster

All content copied from https://docs.aws.amazon.com/.
