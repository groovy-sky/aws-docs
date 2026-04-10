This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GroundStation::DataflowEndpointGroup ConnectionDetails

Egress address of AgentEndpoint with an optional mtu.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Mtu" : Integer,
  "SocketAddress" : SocketAddress
}

```

### YAML

```yaml

  Mtu: Integer
  SocketAddress:
    SocketAddress

```

## Properties

`Mtu`

Maximum transmission unit (MTU) size in bytes of a dataflow endpoint.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SocketAddress`

A socket address.

_Required_: No

_Type_: [SocketAddress](aws-properties-groundstation-dataflowendpointgroup-socketaddress.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AwsGroundStationAgentEndpoint

DataflowEndpoint

All content copied from https://docs.aws.amazon.com/.
