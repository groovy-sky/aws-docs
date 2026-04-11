This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NetworkFirewall::RuleGroup Header

The 5-tuple criteria for AWS Network Firewall to use to inspect packet headers in stateful
traffic flow inspection. Traffic flows that match the criteria are a match for the
corresponding stateful rule.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Destination" : String,
  "DestinationPort" : String,
  "Direction" : String,
  "Protocol" : String,
  "Source" : String,
  "SourcePort" : String
}

```

### YAML

```yaml

  Destination: String
  DestinationPort: String
  Direction: String
  Protocol: String
  Source: String
  SourcePort: String

```

## Properties

`Destination`

The destination IP address or address range to inspect for, in CIDR notation.
To match with any address, specify `ANY`.

Specify an IP address or a block of IP addresses in Classless Inter-Domain Routing (CIDR) notation. Network Firewall supports all address ranges for IPv4 and IPv6.

Examples:

- To configure Network Firewall to inspect for the IP address 192.0.2.44, specify `192.0.2.44/32`.

- To configure Network Firewall to inspect for IP addresses from 192.0.2.0 to 192.0.2.255, specify `192.0.2.0/24`.

- To configure Network Firewall to inspect for the IP address 1111:0000:0000:0000:0000:0000:0000:0111, specify `1111:0000:0000:0000:0000:0000:0000:0111/128`.

- To configure Network Firewall to inspect for IP addresses from 1111:0000:0000:0000:0000:0000:0000:0000 to 1111:0000:0000:0000:ffff:ffff:ffff:ffff, specify `1111:0000:0000:0000:0000:0000:0000:0000/64`.

For more information about CIDR notation, see the Wikipedia entry [Classless\
Inter-Domain Routing](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing).

_Required_: Yes

_Type_: String

_Pattern_: `^.*$`

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DestinationPort`

The destination port to inspect for. You can specify an individual port,
for example `1994` and you can specify a port range, for example `1990:1994`.
To match with any port, specify `ANY`.

_Required_: Yes

_Type_: String

_Pattern_: `^.*$`

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Direction`

The direction of traffic flow to inspect. If set to `ANY`, the inspection
matches bidirectional traffic, both from the source to the destination and from the
destination to the source. If set to `FORWARD`, the inspection only matches
traffic going from the source to the destination.

_Required_: Yes

_Type_: String

_Allowed values_: `FORWARD | ANY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Protocol`

The protocol to inspect for. To specify all, you can use `IP`, because all traffic on AWS and on the internet is IP.

_Required_: Yes

_Type_: String

_Allowed values_: `IP | TCP | UDP | ICMP | HTTP | FTP | TLS | SMB | DNS | DCERPC | SSH | SMTP | IMAP | MSN | KRB5 | IKEV2 | TFTP | NTP | DHCP`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Source`

The source IP address or address range to inspect for, in CIDR notation.
To match with any address, specify `ANY`.

Specify an IP address or a block of IP addresses in Classless Inter-Domain Routing (CIDR) notation. Network Firewall supports all address ranges for IPv4 and IPv6.

Examples:

- To configure Network Firewall to inspect for the IP address 192.0.2.44, specify `192.0.2.44/32`.

- To configure Network Firewall to inspect for IP addresses from 192.0.2.0 to 192.0.2.255, specify `192.0.2.0/24`.

- To configure Network Firewall to inspect for the IP address 1111:0000:0000:0000:0000:0000:0000:0111, specify `1111:0000:0000:0000:0000:0000:0000:0111/128`.

- To configure Network Firewall to inspect for IP addresses from 1111:0000:0000:0000:0000:0000:0000:0000 to 1111:0000:0000:0000:ffff:ffff:ffff:ffff, specify `1111:0000:0000:0000:0000:0000:0000:0000/64`.

For more information about CIDR notation, see the Wikipedia entry [Classless\
Inter-Domain Routing](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing).

_Required_: Yes

_Type_: String

_Pattern_: `^.*$`

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourcePort`

The source port to inspect for. You can specify an individual port,
for example `1994` and you can specify a port range, for example `1990:1994`.
To match with any port, specify `ANY`.

_Required_: Yes

_Type_: String

_Pattern_: `^.*$`

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Dimension

IPSet

All content copied from https://docs.aws.amazon.com/.
