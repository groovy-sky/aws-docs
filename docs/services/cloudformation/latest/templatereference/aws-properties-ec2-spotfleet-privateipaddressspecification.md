This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::SpotFleet PrivateIpAddressSpecification

Describes a secondary private IPv4 address for a network interface.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Primary" : Boolean,
  "PrivateIpAddress" : String
}

```

### YAML

```yaml

  Primary: Boolean
  PrivateIpAddress: String

```

## Properties

`Primary`

Indicates whether the private IPv4 address is the primary private IPv4 address. Only
one IPv4 address can be designated as primary.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PrivateIpAddress`

The private IPv4 address.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PerformanceFactorReferenceRequest

SpotCapacityRebalance

All content copied from https://docs.aws.amazon.com/.
