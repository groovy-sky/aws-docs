This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NetworkFirewall::TLSInspectionConfiguration ServerCertificateScope

Settings that define the Secure Sockets Layer/Transport Layer Security (SSL/TLS) traffic that Network Firewall should decrypt for inspection by the stateful rule engine.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DestinationPorts" : [ PortRange, ... ],
  "Destinations" : [ Address, ... ],
  "Protocols" : [ Integer, ... ],
  "SourcePorts" : [ PortRange, ... ],
  "Sources" : [ Address, ... ]
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

```

## Properties

`DestinationPorts`

The destination ports to decrypt for inspection, in Transmission Control Protocol (TCP) format. If not specified, this matches with any destination port.

You can specify individual ports, for example `1994`, and you can specify port ranges, such as `1990:1994`.

_Required_: No

_Type_: Array of [PortRange](aws-properties-networkfirewall-tlsinspectionconfiguration-portrange.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Destinations`

The destination IP addresses and address ranges to decrypt for inspection, in CIDR notation. If not specified, this
matches with any destination address.

_Required_: No

_Type_: Array of [Address](aws-properties-networkfirewall-tlsinspectionconfiguration-address.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Protocols`

The protocols to inspect for, specified using the assigned internet protocol number (IANA)
for each protocol. If not specified, this matches with any protocol.

Network Firewall currently supports only TCP.

_Required_: No

_Type_: Array of Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourcePorts`

The source ports to decrypt for inspection, in Transmission Control Protocol (TCP) format. If not specified, this matches with any source port.

You can specify individual ports, for example `1994`, and you can specify port ranges, such as `1990:1994`.

_Required_: No

_Type_: Array of [PortRange](aws-properties-networkfirewall-tlsinspectionconfiguration-portrange.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Sources`

The source IP addresses and address ranges to decrypt for inspection, in CIDR notation. If not specified, this
matches with any source address.

_Required_: No

_Type_: Array of [Address](aws-properties-networkfirewall-tlsinspectionconfiguration-address.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ServerCertificateConfiguration

Tag

All content copied from https://docs.aws.amazon.com/.
