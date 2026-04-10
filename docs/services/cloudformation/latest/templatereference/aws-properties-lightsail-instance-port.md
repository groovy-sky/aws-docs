This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lightsail::Instance Port

`Port` is a property of the [Networking](../userguide/aws-properties-lightsail-instance-networking.md) property. It describes information about ports for an
instance.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AccessDirection" : String,
  "AccessFrom" : String,
  "AccessType" : String,
  "CidrListAliases" : [ String, ... ],
  "Cidrs" : [ String, ... ],
  "CommonName" : String,
  "FromPort" : Integer,
  "Ipv6Cidrs" : [ String, ... ],
  "Protocol" : String,
  "ToPort" : Integer
}

```

### YAML

```yaml

  AccessDirection: String
  AccessFrom: String
  AccessType: String
  CidrListAliases:
    - String
  Cidrs:
    - String
  CommonName: String
  FromPort: Integer
  Ipv6Cidrs:
    - String
  Protocol: String
  ToPort: Integer

```

## Properties

`AccessDirection`

The access direction ( `inbound` or `outbound`).

###### Note

Lightsail currently supports only `inbound` access
direction.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AccessFrom`

The location from which access is allowed. For example, `Anywhere
            (0.0.0.0/0)`, or `Custom` if a specific IP address or range of IP
addresses is allowed.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AccessType`

The type of access ( `Public` or `Private`).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CidrListAliases`

An alias that defines access for a preconfigured range of IP addresses.

The only alias currently supported is `lightsail-connect`, which allows IP
addresses of the browser-based RDP/SSH client in the Lightsail console to
connect to your instance.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Cidrs`

The IPv4 address, or range of IPv4 addresses (in CIDR notation) that are allowed to
connect to an instance through the ports, and the protocol.

###### Note

The `ipv6Cidrs` parameter lists the IPv6 addresses that are allowed to
connect to an instance.

Examples:

- To allow the IP address `192.0.2.44`, specify `192.0.2.44`
or `192.0.2.44/32`.

- To allow the IP addresses `192.0.2.0` to `192.0.2.255`,
specify `192.0.2.0/24`.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CommonName`

The common name of the port information.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FromPort`

The first port in a range of open ports on an instance.

Allowed ports:

- TCP and UDP - `0` to `65535`

- ICMP - The ICMP type for IPv4 addresses. For example, specify `8` as
the `fromPort` (ICMP type), and `-1` as the `toPort`
(ICMP code), to enable ICMP Ping.

- ICMPv6 - The ICMP type for IPv6 addresses. For example, specify `128`
as the `fromPort` (ICMPv6 type), and `0` as `toPort`
(ICMPv6 code).

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Ipv6Cidrs`

The IPv6 address, or range of IPv6 addresses (in CIDR notation) that are allowed to
connect to an instance through the ports, and the protocol. Only devices with an IPv6
address can connect to an instance through IPv6; otherwise, IPv4 should be used.

###### Note

The `cidrs` parameter lists the IPv4 addresses that are allowed to connect
to an instance.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Protocol`

The IP protocol name.

The name can be one of the following:

- `tcp` \- Transmission Control Protocol (TCP) provides reliable, ordered,
and error-checked delivery of streamed data between applications running on hosts
communicating by an IP network. If you have an application that doesn't require
reliable data stream service, use UDP instead.

- `all` \- All transport layer protocol types.

- `udp` \- With User Datagram Protocol (UDP), computer applications can
send messages (or datagrams) to other hosts on an Internet Protocol (IP) network.
Prior communications are not required to set up transmission channels or data paths.
Applications that don't require reliable data stream service can use UDP, which
provides a connectionless datagram service that emphasizes reduced latency over
reliability. If you do require reliable data stream service, use TCP instead.

- `icmp` \- Internet Control Message Protocol (ICMP) is used to send error
messages and operational information indicating success or failure when communicating
with an instance. For example, an error is indicated when an instance could not be
reached. When you specify `icmp` as the `protocol`, you must
specify the ICMP type using the `fromPort` parameter, and ICMP code using
the `toPort` parameter.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ToPort`

The last port in a range of open ports on an instance.

Allowed ports:

- TCP and UDP - `0` to `65535`

- ICMP - The ICMP code for IPv4 addresses. For example, specify `8` as
the `fromPort` (ICMP type), and `-1` as the `toPort`
(ICMP code), to enable ICMP Ping.

- ICMPv6 - The ICMP code for IPv6 addresses. For example, specify `128`
as the `fromPort` (ICMPv6 type), and `0` as `toPort`
(ICMPv6 code).

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Networking

State

All content copied from https://docs.aws.amazon.com/.
