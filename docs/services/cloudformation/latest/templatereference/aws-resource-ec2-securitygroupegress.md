This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::SecurityGroupEgress

Adds the specified outbound (egress) rule to a security group.

An outbound rule permits instances to send traffic to the specified IPv4 or IPv6
address range, the IP addresses that are specified by a prefix list, or the instances
that are associated with a destination security group. For more information, see [Security group rules](../../../vpc/latest/userguide/security-group-rules.md).

You must specify exactly one of the following destinations: an IPv4 address range, an IPv6 address range,
a prefix list, or a security group.

You must specify a protocol for each rule (for example, TCP). If the protocol is TCP or UDP,
you must also specify a port or port range. If the protocol is ICMP or ICMPv6, you must also
specify the ICMP/ICMPv6 type and code. To specify all types or all codes, use -1.

Rule changes are propagated to instances associated with the security group as quickly
as possible. However, a small delay might occur.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::SecurityGroupEgress",
  "Properties" : {
      "CidrIp" : String,
      "CidrIpv6" : String,
      "Description" : String,
      "DestinationPrefixListId" : String,
      "DestinationSecurityGroupId" : String,
      "FromPort" : Integer,
      "GroupId" : String,
      "IpProtocol" : String,
      "ToPort" : Integer
    }
}

```

### YAML

```yaml

Type: AWS::EC2::SecurityGroupEgress
Properties:
  CidrIp: String
  CidrIpv6: String
  Description: String
  DestinationPrefixListId: String
  DestinationSecurityGroupId: String
  FromPort: Integer
  GroupId: String
  IpProtocol: String
  ToPort: Integer

```

## Properties

`CidrIp`

The IPv4 address range, in CIDR format.

You must specify exactly one of the following:
`CidrIp`, `CidrIpv6`, `DestinationPrefixListId`, or `DestinationSecurityGroupId`.

For examples of rules that you can add to
security groups for specific access scenarios, see [Security group rules\
for different use cases](../../../ec2/latest/userguide/security-group-rules-reference.md) in the _Amazon EC2 User_
_Guide_.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CidrIpv6`

The IPv6 address range, in CIDR format.

You must specify exactly one of the following:
`CidrIp`, `CidrIpv6`, `DestinationPrefixListId`, or `DestinationSecurityGroupId`.

For examples of rules that you can add to security groups for specific access scenarios, see [Security group rules for different use cases](../../../ec2/latest/userguide/security-group-rules-reference.md) in the
_Amazon EC2 User Guide_.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

The description of an egress (outbound) security group rule.

Constraints: Up to 255 characters in length. Allowed characters are a-z, A-Z, 0-9,
spaces, and .\_-:/()#,@\[\]+=;{}!$\*

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DestinationPrefixListId`

The prefix list IDs for an AWS service. This is the AWS service to access through a
VPC endpoint from instances associated with the security group.

You must specify exactly one of the following:
`CidrIp`, `CidrIpv6`, `DestinationPrefixListId`, or `DestinationSecurityGroupId`.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DestinationSecurityGroupId`

The ID of the security group.

You must specify exactly one of the following:
`CidrIp`, `CidrIpv6`, `DestinationPrefixListId`, or `DestinationSecurityGroupId`.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FromPort`

If the protocol is TCP or UDP, this is the start of the port range.
If the protocol is ICMP or ICMPv6, this is the ICMP type or -1 (all ICMP types).

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`GroupId`

The ID of the security group. You must specify either the security group ID or the
security group name in the request. For security groups in a nondefault VPC, you must
specify the security group ID.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

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

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ToPort`

If the protocol is TCP or UDP, this is the end of the port range.
If the protocol is ICMP or ICMPv6, this is the ICMP code or -1 (all ICMP codes).
If the start port is -1 (all ICMP types), then the end port must be -1 (all ICMP codes).

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the name of the security egress rule.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## Examples

### VPC security groups with egress and ingress rules

In some cases, you might have an originating (source) security group to which you
want to add an outbound rule that allows traffic to a destination (target) security
group. The target security group also needs an inbound rule that allows traffic from
the source security group. Note that you cannot use the `Ref` function to
specify the outbound and inbound rules for each security group. Doing so creates a
circular dependency; you cannot have two resources that depend on each other.
Instead, use the egress and ingress resources to declare these outbound and inbound
rules, as shown in the following template example.

#### JSON

```json

"SourceSG": {
   "Type": "AWS::EC2::SecurityGroup",
   "Properties": {
      "VpcId" : "vpc-1a2b3c4d",
      "GroupDescription": "Sample source security group"
   }
},
"TargetSG": {
   "Type": "AWS::EC2::SecurityGroup",
   "Properties": {
      "VpcId" : "vpc-1a2b3c4d",
      "GroupDescription": "Sample target security group"
   }
},
"OutboundRule": {
   "Type": "AWS::EC2::SecurityGroupEgress",
   "Properties":{
      "IpProtocol": "tcp",
      "FromPort": 0,
      "ToPort": 65535,
      "DestinationSecurityGroupId": {
         "Fn::GetAtt": [
            "TargetSG",
            "GroupId"
         ]
      },
      "GroupId": {
         "Fn::GetAtt": [
            "SourceSG",
            "GroupId"
         ]
      }
   }
},
"InboundRule": {
   "Type": "AWS::EC2::SecurityGroupIngress",
   "Properties":{
      "IpProtocol": "tcp",
      "FromPort": 0,
      "ToPort": 65535,
      "SourceSecurityGroupId": {
         "Fn::GetAtt": [
            "SourceSG",
            "GroupId"
         ]
      },
      "GroupId": {
         "Fn::GetAtt": [
            "TargetSG",
            "GroupId"
         ]
      }
   }
}
```

#### YAML

```yaml

SourceSG:
  Type: AWS::EC2::SecurityGroup
  Properties:
    VpcId: vpc-1a2b3c4d
    GroupDescription: Sample source security group
TargetSG:
  Type: AWS::EC2::SecurityGroup
  Properties:
    VpcId: vpc-1a2b3c4d
    GroupDescription: Sample target security group
OutboundRule:
  Type: AWS::EC2::SecurityGroupEgress
  Properties:
    IpProtocol: tcp
    FromPort: 0
    ToPort: 65535
    DestinationSecurityGroupId:
      Fn::GetAtt:
        - TargetSG
        - GroupId
    GroupId:
      Fn::GetAtt:
        - SourceSG
        - GroupId
InboundRule:
  Type: AWS::EC2::SecurityGroupIngress
  Properties:
    IpProtocol: tcp
    FromPort: 0
    ToPort: 65535
    SourceSecurityGroupId:
      Fn::GetAtt:
        - SourceSG
        - GroupId
    GroupId:
      Fn::GetAtt:
        - TargetSG
        - GroupId
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AWS::EC2::SecurityGroupIngress
