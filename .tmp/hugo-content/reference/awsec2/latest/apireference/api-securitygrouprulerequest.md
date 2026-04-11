# SecurityGroupRuleRequest

Describes a security group rule.

You must specify exactly one of the following parameters, based on the rule type:

- CidrIpv4

- CidrIpv6

- PrefixListId

- ReferencedGroupId

###### Note

AWS
[canonicalizes](https://en.wikipedia.org/wiki/Canonicalization) IPv4 and IPv6 CIDRs. For example, if you specify 100.68.0.18/18 for the CIDR block,
AWS canonicalizes the CIDR block to 100.68.0.0/18. Any subsequent DescribeSecurityGroups and DescribeSecurityGroupRules calls will
return the canonicalized form of the CIDR block. Additionally, if you attempt to add another rule with the
non-canonical form of the CIDR (such as 100.68.0.18/18) and there is already a rule for the canonicalized
form of the CIDR block (such as 100.68.0.0/18), the API throws an duplicate rule error.

When you modify a rule, you cannot change the rule type. For example, if the rule
uses an IPv4 address range, you must use `CidrIpv4` to specify a new IPv4
address range.

## Contents

**CidrIpv4**

The IPv4 CIDR range. To specify a single IPv4 address, use the /32 prefix length.

Type: String

Required: No

**CidrIpv6**

The IPv6 CIDR range. To specify a single IPv6 address, use the /128 prefix length.

Type: String

Required: No

**Description**

The description of the security group rule.

Type: String

Required: No

**FromPort**

If the protocol is TCP or UDP, this is the start of the port range.
If the protocol is ICMP or ICMPv6, this is the ICMP type or -1 (all ICMP types).

Type: Integer

Required: No

**IpProtocol**

The IP protocol name ( `tcp`, `udp`, `icmp`,
`icmpv6`) or number (see [Protocol Numbers](http://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml)).

Use `-1` to specify all protocols.

Type: String

Required: No

**PrefixListId**

The ID of the prefix list.

Type: String

Required: No

**ReferencedGroupId**

The ID of the security group that is referenced in the security group rule.

Type: String

Required: No

**ToPort**

If the protocol is TCP or UDP, this is the end of the port range.
If the protocol is ICMP or ICMPv6, this is the ICMP code or -1 (all ICMP codes).
If the start port is -1 (all ICMP types), then the end port must be -1 (all ICMP codes).

Type: Integer

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/securitygrouprulerequest.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/securitygrouprulerequest.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/securitygrouprulerequest.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

SecurityGroupRuleDescription

SecurityGroupRuleUpdate
