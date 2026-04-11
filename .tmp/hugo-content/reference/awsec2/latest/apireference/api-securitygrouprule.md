# SecurityGroupRule

Describes a security group rule.

## Contents

**cidrIpv4**

The IPv4 CIDR range.

Type: String

Required: No

**cidrIpv6**

The IPv6 CIDR range.

Type: String

Required: No

**description**

The security group rule description.

Type: String

Required: No

**fromPort**

If the protocol is TCP or UDP, this is the start of the port range.
If the protocol is ICMP or ICMPv6, this is the ICMP type or -1 (all ICMP types).

Type: Integer

Required: No

**groupId**

The ID of the security group.

Type: String

Required: No

**groupOwnerId**

The ID of the AWS account that owns the security group.

Type: String

Required: No

**ipProtocol**

The IP protocol name ( `tcp`, `udp`, `icmp`,
`icmpv6`) or number (see [Protocol Numbers](http://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml)).

Use `-1` to specify all protocols.

Type: String

Required: No

**isEgress**

Indicates whether the security group rule is an outbound rule.

Type: Boolean

Required: No

**prefixListId**

The ID of the prefix list.

Type: String

Required: No

**referencedGroupInfo**

Describes the security group that is referenced in the rule.

Type: [ReferencedSecurityGroup](api-referencedsecuritygroup.md) object

Required: No

**securityGroupRuleArn**

The ARN of the security group rule.

Type: String

Required: No

**securityGroupRuleId**

The ID of the security group rule.

Type: String

Required: No

**TagSet.N**

The tags applied to the security group rule.

Type: Array of [Tag](api-tag.md) objects

Required: No

**toPort**

If the protocol is TCP or UDP, this is the end of the port range.
If the protocol is ICMP or ICMPv6, this is the ICMP code or -1 (all ICMP codes).
If the start port is -1 (all ICMP types), then the end port must be -1 (all ICMP codes).

Type: Integer

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/securitygrouprule.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/securitygrouprule.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/securitygrouprule.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

SecurityGroupReference

SecurityGroupRuleDescription
