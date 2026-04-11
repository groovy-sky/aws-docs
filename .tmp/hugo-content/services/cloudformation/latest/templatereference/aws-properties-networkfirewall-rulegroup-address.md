This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NetworkFirewall::RuleGroup Address

A single IP address specification. This is used in the match attributes
source and destination specifications.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AddressDefinition" : String
}

```

### YAML

```yaml

  AddressDefinition: String

```

## Properties

`AddressDefinition`

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

_Pattern_: `^([a-fA-F\d:\.]+/\d{1,3})$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ActionDefinition

CustomAction

All content copied from https://docs.aws.amazon.com/.
