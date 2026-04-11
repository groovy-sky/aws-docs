This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::CapacityProvider NetworkInterfaceCountRequest

The minimum and maximum number of network interfaces for instance type selection. This
is useful for workloads that require multiple network interfaces.

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

The maximum number of network interfaces. Instance types that support more network
interfaces are excluded from selection.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Min`

The minimum number of network interfaces. Instance types that support fewer network
interfaces are excluded from selection.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NetworkBandwidthGbpsRequest

Tag

All content copied from https://docs.aws.amazon.com/.
