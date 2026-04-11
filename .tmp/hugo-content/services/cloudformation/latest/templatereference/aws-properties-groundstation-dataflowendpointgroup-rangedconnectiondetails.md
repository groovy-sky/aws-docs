This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GroundStation::DataflowEndpointGroup RangedConnectionDetails

Ingress address of AgentEndpoint with a port range and an optional mtu.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Mtu" : Integer,
  "SocketAddress" : RangedSocketAddress
}

```

### YAML

```yaml

  Mtu: Integer
  SocketAddress:
    RangedSocketAddress

```

## Properties

`Mtu`

Maximum transmission unit (MTU) size in bytes of a dataflow endpoint.

_Required_: No

_Type_: Integer

_Minimum_: `1400`

_Maximum_: `1500`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SocketAddress`

A ranged socket address.

_Required_: No

_Type_: [RangedSocketAddress](aws-properties-groundstation-dataflowendpointgroup-rangedsocketaddress.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IntegerRange

RangedSocketAddress

All content copied from https://docs.aws.amazon.com/.
