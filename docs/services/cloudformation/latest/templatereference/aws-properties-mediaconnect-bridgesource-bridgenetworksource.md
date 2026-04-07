This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::BridgeSource BridgeNetworkSource

The source of the bridge. A network source originates at your premises.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MulticastIp" : String,
  "MulticastSourceSettings" : MulticastSourceSettings,
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

_Type_: [MulticastSourceSettings](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-mediaconnect-bridgesource-multicastsourcesettings.html)

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

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

BridgeFlowSource

MulticastSourceSettings
