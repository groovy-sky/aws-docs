This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NetworkFirewall::RuleGroup MatchAttributes

Criteria for Network Firewall to use to inspect an individual packet in stateless rule inspection. Each match attributes set can include one or more items such as IP address, CIDR range, port number, protocol, and TCP flags.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DestinationPorts" : [ PortRange, ... ],
  "Destinations" : [ Address, ... ],
  "Protocols" : [ Integer, ... ],
  "SourcePorts" : [ PortRange, ... ],
  "Sources" : [ Address, ... ],
  "TCPFlags" : [ TCPFlagField, ... ]
}

```

### YAML

```yaml

  DestinationPorts:
    - PortRange
  Destinations:
    - Address
  Protocols:
    - Integer
  SourcePorts:
    - PortRange
  Sources:
    - Address
  TCPFlags:
    - TCPFlagField

```

## Properties

`DestinationPorts`

The destination port to inspect for. You can specify an individual port,
for example `1994` and you can specify a port range, for example `1990:1994`.
To match with any port, specify `ANY`.

This setting is only used for protocols 6 (TCP) and 17 (UDP).

_Required_: No

_Type_: Array of [PortRange](aws-properties-networkfirewall-rulegroup-portrange.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Destinations`

The destination IP addresses and address ranges to inspect for, in CIDR notation. If not
specified, this matches with any destination address.

_Required_: No

_Type_: Array of [Address](aws-properties-networkfirewall-rulegroup-address.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Protocols`

The protocols to inspect for, specified using the assigned internet protocol number (IANA)
for each protocol. If not specified, this matches with any protocol.

_Required_: No

_Type_: Array of Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourcePorts`

The source port to inspect for. You can specify an individual port,
for example `1994` and you can specify a port range, for example `1990:1994`.
To match with any port, specify `ANY`.

If not specified, this matches with any source port.

This setting is only used for protocols 6 (TCP) and 17 (UDP).

_Required_: No

_Type_: Array of [PortRange](aws-properties-networkfirewall-rulegroup-portrange.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Sources`

The source IP addresses and address ranges to inspect for, in CIDR notation. If not
specified, this matches with any source address.

_Required_: No

_Type_: Array of [Address](aws-properties-networkfirewall-rulegroup-address.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TCPFlags`

The TCP flags and masks to inspect for. If not specified, this matches with any
settings. This setting is only used for protocol 6 (TCP).

_Required_: No

_Type_: Array of [TCPFlagField](aws-properties-networkfirewall-rulegroup-tcpflagfield.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IPSetReference

PortRange

All content copied from https://docs.aws.amazon.com/.
