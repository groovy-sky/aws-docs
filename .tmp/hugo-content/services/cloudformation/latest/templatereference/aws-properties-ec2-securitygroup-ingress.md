This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::SecurityGroup Ingress

Adds an inbound (ingress) rule to a security group.

An inbound rule permits instances to receive traffic from the specified IPv4 or IPv6
address range, the IP address ranges that are specified by a prefix list, or the instances
that are associated with a source security group. For more information, see [Security group rules](../../../vpc/latest/userguide/security-group-rules.md).

You must specify exactly one of the following sources: an IPv4 address range, an IPv6 address range,
a prefix list, or a security group.

You must specify a protocol for each rule (for example, TCP). If the protocol is TCP or UDP,
you must also specify a port or port range. If the protocol is ICMP or ICMPv6, you must also
specify the ICMP/ICMPv6 type and code.

Rule changes are propagated to instances associated with the security group as quickly
as possible. However, a small delay might occur.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CidrIp" : String,
  "CidrIpv6" : String,
  "Description" : String,
  "FromPort" : Integer,
  "IpProtocol" : String,
  "SourcePrefixListId" : String,
  "SourceSecurityGroupId" : String,
  "SourceSecurityGroupName" : String,
  "SourceSecurityGroupOwnerId" : String,
  "ToPort" : Integer
}

```

### YAML

```yaml

  CidrIp: String
  CidrIpv6: String
  Description: String
  FromPort: Integer
  IpProtocol: String
  SourcePrefixListId: String
  SourceSecurityGroupId: String
  SourceSecurityGroupName: String
  SourceSecurityGroupOwnerId: String
  ToPort: Integer

```

## Properties

`CidrIp`

The IPv4 address range, in CIDR format.

You must specify exactly one of the following:
`CidrIp`, `CidrIpv6`, `SourcePrefixListId`, or `SourceSecurityGroupId`.

For examples of rules that you can add to
security groups for specific access scenarios, see [Security group rules\
for different use cases](../../../ec2/latest/userguide/security-group-rules-reference.md) in the _Amazon EC2 User_
_Guide_.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CidrIpv6`

The IPv6 address range, in CIDR format.

You must specify exactly one of the following:
`CidrIp`, `CidrIpv6`, `SourcePrefixListId`, or `SourceSecurityGroupId`.

For examples of rules that you can add to
security groups for specific access scenarios, see [Security group rules\
for different use cases](../../../ec2/latest/userguide/security-group-rules-reference.md) in the _Amazon EC2 User_
_Guide_.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

Updates the description of an ingress (inbound) security group rule. You can replace an
existing description, or add a description to a rule that did not have one
previously.

Constraints: Up to 255 characters in length. Allowed characters are a-z, A-Z, 0-9,
spaces, and .\_-:/()#,@\[\]+=;{}!$\*

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FromPort`

If the protocol is TCP or UDP, this is the start of the port range.
If the protocol is ICMP or ICMPv6, this is the ICMP type or -1 (all ICMP types).

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IpProtocol`

The IP protocol name ( `tcp`, `udp`, `icmp`, `icmpv6`)
or number (see [Protocol Numbers](http://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml)).

Use `-1` to specify all protocols. When authorizing
security group rules, specifying `-1` or a protocol number other than
`tcp`, `udp`, `icmp`, or `icmpv6` allows
traffic on all ports, regardless of any port range you specify. For `tcp`,
`udp`, and `icmp`, you must specify a port range. For `icmpv6`,
the port range is optional; if you omit the port range, traffic for all types and codes is allowed.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourcePrefixListId`

The ID of a prefix list.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceSecurityGroupId`

The ID of the security group.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceSecurityGroupName`

\[Default VPC\] The name of the source security group. You must specify either the security group ID
or the security group name. You can't specify the group name in combination with an IP address range.
Creates rules that grant full ICMP, UDP, and TCP access.

For security groups in a nondefault VPC, you must specify the group ID.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceSecurityGroupOwnerId`

\[nondefault VPC\] The AWS account ID for the source security group, if
the source security group is in a different account. You can't specify this property with
an IP address range. Creates rules that grant full ICMP, UDP, and TCP access.

If you specify `SourceSecurityGroupName` or
`SourceSecurityGroupId` and that security group is owned by a different
account than the account creating the stack, you must specify the
`SourceSecurityGroupOwnerId`; otherwise, this property is optional.

_Required_: Conditional

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ToPort`

If the protocol is TCP or UDP, this is the end of the port range.
If the protocol is ICMP or ICMPv6, this is the ICMP code or -1 (all ICMP codes).
If the start port is -1 (all ICMP types), then the end port must be -1 (all ICMP codes).

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Egress

Tag

All content copied from https://docs.aws.amazon.com/.
