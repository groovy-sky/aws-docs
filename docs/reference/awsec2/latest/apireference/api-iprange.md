# IpRange

Describes an IPv4 address range.

## Contents

**CidrIp** (request), **cidrIp** (response)

The IPv4 address range. You can either specify a CIDR block or a source security group,
not both. To specify a single IPv4 address, use the /32 prefix length.

###### Note

AWS
[canonicalizes](https://en.wikipedia.org/wiki/Canonicalization) IPv4 and IPv6 CIDRs. For example, if you specify 100.68.0.18/18 for the CIDR block,
AWS canonicalizes the CIDR block to 100.68.0.0/18. Any subsequent DescribeSecurityGroups and DescribeSecurityGroupRules calls will
return the canonicalized form of the CIDR block. Additionally, if you attempt to add another rule with the
non-canonical form of the CIDR (such as 100.68.0.18/18) and there is already a rule for the canonicalized
form of the CIDR block (such as 100.68.0.0/18), the API throws an duplicate rule error.

Type: String

Required: No

**Description** (request), **description** (response)

A description for the security group rule that references this IPv4 address range.

Constraints: Up to 255 characters in length. Allowed characters are a-z, A-Z, 0-9,
spaces, and .\_-:/()#,@\[\]+=&;{}!$\*

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/iprange.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/iprange.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/iprange.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

IpPermission

Ipv4PrefixSpecification
