This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::Bridge BridgeNetworkSource

The source of the bridge. A network source originates at your premises.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MulticastIp" : String,
  "MulticastSourceSettings" : MulticastSourceSettings,
  "Name" : String,
  "NetworkName" : String,
  "Port" : Integer,
  "Protocol" : String
}

```

### YAML

```yaml

  MulticastIp: String
  MulticastSourceSettings:
    MulticastSourceSettings
  Name: String
  NetworkName: String
  Port: Integer
  Protocol: String

```

## Properties

`MulticastIp`

The network source multicast IP.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MulticastSourceSettings`

The settings related to the multicast source.

_Required_: No

_Type_: [MulticastSourceSettings](aws-properties-mediaconnect-bridge-multicastsourcesettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the network source.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NetworkName`

The network source's gateway network name.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Port`

The network source port.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Protocol`

The network source protocol.

###### Note

AWS Elemental MediaConnect no longer supports the Fujitsu QoS
protocol. This reference is maintained for legacy purposes only.

_Required_: Yes

_Type_: String

_Allowed values_: `rtp-fec | rtp | udp`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BridgeNetworkOutput

BridgeOutput

All content copied from https://docs.aws.amazon.com/.
