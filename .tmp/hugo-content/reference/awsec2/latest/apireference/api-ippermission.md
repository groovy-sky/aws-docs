# IpPermission

Describes the permissions for a security group rule.

## Contents

**FromPort** (request), **fromPort** (response)

If the protocol is TCP or UDP, this is the start of the port range.
If the protocol is ICMP or ICMPv6, this is the ICMP type or -1 (all ICMP types).

Type: Integer

Required: No

**Groups.N**

The security group and AWS account ID pairs.

Type: Array of [UserIdGroupPair](api-useridgrouppair.md) objects

Required: No

**IpProtocol** (request), **ipProtocol** (response)

The IP protocol name ( `tcp`, `udp`, `icmp`, `icmpv6`)
or number (see [Protocol Numbers](http://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml)).

Use `-1` to specify all protocols. When authorizing
security group rules, specifying `-1` or a protocol number other than
`tcp`, `udp`, `icmp`, or `icmpv6` allows
traffic on all ports, regardless of any port range you specify. For `tcp`,
`udp`, and `icmp`, you must specify a port range. For `icmpv6`,
the port range is optional; if you omit the port range, traffic for all types and codes is allowed.

Type: String

Required: No

**IpRanges.N**

The IPv4 address ranges.

Type: Array of [IpRange](api-iprange.md) objects

Required: No

**Ipv6Ranges.N**

The IPv6 address ranges.

Type: Array of [Ipv6Range](api-ipv6range.md) objects

Required: No

**PrefixListIds.N**

The prefix list IDs.

Type: Array of [PrefixListId](api-prefixlistid.md) objects

Required: No

**ToPort** (request), **toPort** (response)

If the protocol is TCP or UDP, this is the end of the port range.
If the protocol is ICMP or ICMPv6, this is the ICMP code or -1 (all ICMP codes).
If the start port is -1 (all ICMP types), then the end port must be -1 (all ICMP codes).

Type: Integer

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/ippermission.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/ippermission.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/ippermission.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

IpamScopeExternalAuthorityConfiguration

IpRange
