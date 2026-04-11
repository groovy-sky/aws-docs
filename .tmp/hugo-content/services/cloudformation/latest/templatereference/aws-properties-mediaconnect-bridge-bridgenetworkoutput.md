This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::Bridge BridgeNetworkOutput

The output of the bridge. A network output is delivered to your premises.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IpAddress" : String,
  "Name" : String,
  "NetworkName" : String,
  "Port" : Integer,
  "Protocol" : String,
  "Ttl" : Integer
}

```

### YAML

```yaml

  IpAddress: String
  Name: String
  NetworkName: String
  Port: Integer
  Protocol: String
  Ttl: Integer

```

## Properties

`IpAddress`

The network output IP address.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The network output name.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NetworkName`

The network output's gateway network name.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Port`

The network output's port.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Protocol`

The network output protocol.

###### Note

AWS Elemental MediaConnect no longer supports the Fujitsu QoS
protocol. This reference is maintained for legacy purposes only.

_Required_: Yes

_Type_: String

_Allowed values_: `rtp-fec | rtp | udp`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Ttl`

The network output TTL.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BridgeFlowSource

BridgeNetworkSource

All content copied from https://docs.aws.amazon.com/.
