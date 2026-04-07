This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::SecurityGroupIngress

Adds an inbound (ingress) rule to a security group.

An inbound rule permits instances to receive traffic from the specified IPv4 or IPv6
address range, the IP addresses that are specified by a prefix list, or the instances
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
  "Type" : "AWS::EC2::SecurityGroupIngress",
  "Properties" : {
      "CidrIp" : String,
      "CidrIpv6" : String,
      "Description" : String,
      "FromPort" : Integer,
      "GroupId" : String,
      "GroupName" : String,
      "IpProtocol" : String,
      "SourcePrefixListId" : String,
      "SourceSecurityGroupId" : String,
      "SourceSecurityGroupName" : String,
      "SourceSecurityGroupOwnerId" : String,
      "ToPort" : Integer
    }
}

```

### YAML

```yaml

Type: AWS::EC2::SecurityGroupIngress
Properties:
  CidrIp: String
  CidrIpv6: String
  Description: String
  FromPort: Integer
  GroupId: String
  GroupName: String
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

For examples of rules that you can add to security groups for specific access scenarios,
see [Security group rules\
for different use cases](../../../ec2/latest/userguide/security-group-rules-reference.md) in the _Amazon EC2 User_
_Guide_.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CidrIpv6`

The IPv6 address range, in CIDR format.

You must specify exactly one of the following:
`CidrIp`, `CidrIpv6`, `SourcePrefixListId`, or `SourceSecurityGroupId`.

For examples of rules that you can add to security groups for specific access scenarios,
see [Security group rules\
for different use cases](../../../ec2/latest/userguide/security-group-rules-reference.md) in the _Amazon EC2 User_
_Guide_.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

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

The start of port range for the TCP and UDP protocols, or an ICMP/ICMPv6 type number. A
value of `-1` indicates all ICMP/ICMPv6 types. If you specify all ICMP/ICMPv6
types, you must specify all codes.

Use this for ICMP and any protocol that uses ports.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`GroupId`

The ID of the security group.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`GroupName`

\[Default VPC\] The name of the security group. For security groups for a default VPC
you can specify either the ID or the name of the security group. For security groups for
a nondefault VPC, you must specify the ID of the security group.

_Required_: No

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

`SourcePrefixListId`

The ID of a prefix list.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SourceSecurityGroupId`

The ID of the security group. You must specify either the security group ID or the
security group name. For security groups in a nondefault VPC, you must specify the security
group ID.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SourceSecurityGroupName`

\[Default VPC\] The name of the source security group. You must specify either the security group ID
or the security group name. You can't specify the group name in combination with an IP address range.
Creates rules that grant full ICMP, UDP, and TCP access.

For security groups in a nondefault VPC, you must specify the group ID.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SourceSecurityGroupOwnerId`

\[nondefault VPC\] The AWS account ID for the source security group, if
the source security group is in a different account. You can't specify this property with
an IP address range. Creates rules that grant full ICMP, UDP, and TCP access.

If you specify `SourceSecurityGroupName` or
`SourceSecurityGroupId` and that security group is owned by a different
account than the account creating the stack, you must specify
`SourceSecurityGroupOwnerId`; otherwise, this property is optional.

_Required_: Conditional

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ToPort`

The end of port range for the TCP and UDP protocols, or an ICMP/ICMPv6 code. A value of
`-1` indicates all ICMP/ICMPv6 codes for the specified ICMP type. If you
specify all ICMP/ICMPv6 types, you must specify all codes.

Use this for ICMP and any protocol that uses ports.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Examples

- [VPC security groups with egress and ingress rules](#aws-resource-ec2-securitygroupingress--examples--VPC_security_groups_with_egress_and_ingress_rules)

- [Allow traffic from a security group in a peered VPC](#aws-resource-ec2-securitygroupingress--examples--Allow_traffic_from_a_security_group_in_a_peered_VPC)

### VPC security groups with egress and ingress rules

In some cases, you might have an originating (source) security group to which you
want to add an outbound rule that allows traffic to a destination (target) security
group. The target security group also needs an inbound rule that allows traffic from
the source security group. Note that you cannot use the Ref function to specify the
outbound and inbound rules for each security group. Doing so creates a circular
dependency; you cannot have two resources that depend on each other. Instead, use the
egress and ingress resources to declare these outbound and inbound rules, as shown in
the following template example.

#### JSON

```json

{
    "Resources": {
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
    }
}
```

#### YAML

```yaml

Resources:
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

### Allow traffic from a security group in a peered VPC

The following example allows one-way traffic from an
originating (source) security group to a destination (target) security group.
However, in this example the security groups are in peered VPCs across AWS accounts. You might want to allow cross-account traffic if, for
example, you create a security scanning resource in one AWS account
that you'll use to run diagnostics in another account. This example adds an ingress
rule to a target VPC security group that allows incoming traffic from a source
security group in a different AWS account. Note that the source
security group also needs an egress rule that allows outgoing traffic to the target
security group. Because the source security group is in a different account, the
following example doesn't use the Ref function to reference the source security group
ID but instead directly specifies the security group ID `sg-12345678`.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Resources": {
        "TargetSG": {
            "Type": "AWS::EC2::SecurityGroup",
            "Properties": {
                "VpcId": "vpc-1a2b3c4d",
                "GroupDescription": "Security group allowing ingress for security scanners"
            }
        },
        "InboundRule": {
            "Type": "AWS::EC2::SecurityGroupIngress",
            "Properties": {
                "GroupId": {
                    "Fn::GetAtt": [
                        "TargetSG",
                        "GroupId"
                    ]
                },
                "IpProtocol": "tcp",
                "FromPort": 80,
                "ToPort": 80,
                "SourceSecurityGroupId": "sg-12345678",
                "SourceSecurityGroupOwnerId": "123456789012"
            }
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Resources:
  TargetSG:
    Type: 'AWS::EC2::SecurityGroup'
    Properties:
      VpcId: vpc-1a2b3c4d
      GroupDescription: Security group allowing ingress for security scanners
  InboundRule:
    Type: 'AWS::EC2::SecurityGroupIngress'
    Properties:
      GroupId: !GetAtt TargetSG.GroupId
      IpProtocol: tcp
      FromPort: 80
      ToPort: 80
      SourceSecurityGroupId: sg-12345678
      SourceSecurityGroupOwnerId: '123456789012'

```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::EC2::SecurityGroupEgress

AWS::EC2::SecurityGroupVpcAssociation
