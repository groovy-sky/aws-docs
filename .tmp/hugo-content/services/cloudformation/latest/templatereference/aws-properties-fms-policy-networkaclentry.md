This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FMS::Policy NetworkAclEntry

Describes a rule in a network ACL.

Each network ACL has a set of numbered ingress rules and a separate set of numbered egress rules. When determining
whether a packet should be allowed in or out of a subnet associated with the network ACL, AWS processes the
entries in the network ACL according to the rule numbers, in ascending order.

When you manage an individual network ACL, you explicitly specify the rule numbers. When you specify the network ACL rules in a Firewall Manager policy,
you provide the rules to run first, in the order that you want them to run, and the rules to run last, in the order
that you want them to run. Firewall Manager assigns the rule numbers for you when you save the network ACL policy specification.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CidrBlock" : String,
  "Egress" : Boolean,
  "IcmpTypeCode" : IcmpTypeCode,
  "Ipv6CidrBlock" : String,
  "PortRange" : PortRange,
  "Protocol" : String,
  "RuleAction" : String
}

```

### YAML

```yaml

  CidrBlock: String
  Egress: Boolean
  IcmpTypeCode:
    IcmpTypeCode
  Ipv6CidrBlock: String
  PortRange:
    PortRange
  Protocol: String
  RuleAction: String

```

## Properties

`CidrBlock`

The IPv4 network range to allow or deny, in CIDR notation.

_Required_: No

_Type_: String

_Pattern_: `^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(\/([0-9]|[1-2][0-9]|3[0-2]))$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Egress`

Indicates whether the rule is an egress, or outbound, rule (applied to traffic leaving the subnet). If it's not
an egress rule, then it's an ingress, or inbound, rule.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IcmpTypeCode`

ICMP protocol: The ICMP type and code.

_Required_: No

_Type_: [IcmpTypeCode](aws-properties-fms-policy-icmptypecode.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Ipv6CidrBlock`

The IPv6 network range to allow or deny, in CIDR notation.

_Required_: No

_Type_: String

_Pattern_: `^(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))(/(1[0-2]|[0-9]))?$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PortRange`

TCP or UDP protocols: The range of ports the rule applies to.

_Required_: No

_Type_: [PortRange](aws-properties-fms-policy-portrange.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Protocol`

The protocol number. A value of "-1" means all protocols.

_Required_: Yes

_Type_: String

_Pattern_: `^(tcp|udp|icmp|-1|([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]))$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RuleAction`

Indicates whether to allow or deny the traffic that matches the rule.

_Required_: Yes

_Type_: String

_Allowed values_: `allow | deny`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NetworkAclCommonPolicy

NetworkAclEntrySet

All content copied from https://docs.aws.amazon.com/.
