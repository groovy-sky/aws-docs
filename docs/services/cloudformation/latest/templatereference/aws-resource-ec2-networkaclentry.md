This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::NetworkAclEntry

Specifies an entry, known as a rule, in a network ACL with a rule number you specify.
Each network ACL has a set of numbered ingress rules and a separate set of numbered egress
rules.

To create the network ACL, see [AWS::EC2::NetworkAcl](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkacl.html).

For information about the protocol value, see [Protocol\
Numbers](https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml) on the Internet Assigned Numbers Authority (IANA) website.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::NetworkAclEntry",
  "Properties" : {
      "CidrBlock" : String,
      "Egress" : Boolean,
      "Icmp" : Icmp,
      "Ipv6CidrBlock" : String,
      "NetworkAclId" : String,
      "PortRange" : PortRange,
      "Protocol" : Integer,
      "RuleAction" : String,
      "RuleNumber" : Integer
    }
}

```

### YAML

```yaml

Type: AWS::EC2::NetworkAclEntry
Properties:
  CidrBlock: String
  Egress: Boolean
  Icmp:
    Icmp
  Ipv6CidrBlock: String
  NetworkAclId: String
  PortRange:
    PortRange
  Protocol: Integer
  RuleAction: String
  RuleNumber: Integer

```

## Properties

`CidrBlock`

The IPv4 CIDR range to allow or deny, in CIDR notation (for example, 172.16.0.0/24).
You must specify an IPv4 CIDR block or an IPv6 CIDR block.

_Required_: Conditional

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Egress`

Whether this rule applies to egress traffic from the subnet ( `true`) or
ingress traffic to the subnet ( `false`). By default, AWS CloudFormation
specifies `false`.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Icmp`

The Internet Control Message Protocol (ICMP) code and type.
Required if specifying 1 (ICMP) for the protocol parameter.

_Required_: Conditional

_Type_: [Icmp](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-networkaclentry-icmp.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Ipv6CidrBlock`

The IPv6 network range to allow or deny, in CIDR notation.
You must specify an IPv4 CIDR block or an IPv6 CIDR block.

_Required_: Conditional

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NetworkAclId`

The ID of the ACL for the entry.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PortRange`

The range of port numbers for the UDP/TCP protocol. Required if specifying 6
(TCP) or 17 (UDP) for the protocol parameter.

_Required_: Conditional

_Type_: [PortRange](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-networkaclentry-portrange.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Protocol`

The IP protocol that the rule applies to. You must specify -1 or a protocol number. You
can specify -1 for all protocols.

###### Note

If you specify -1, all ports are opened and the `PortRange` property is
ignored.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RuleAction`

Whether to allow or deny traffic that matches the rule; valid values are "allow" or
"deny".

_Required_: Yes

_Type_: String

_Allowed values_: `allow | deny`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RuleNumber`

Rule number to assign to the entry, such as 100. ACL entries are processed in ascending
order by rule number. Entries can't use the same rule number unless one is an egress rule
and the other is an ingress rule.

_Required_: Yes

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the network ACL entry.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Id`

The ID of the network ACL entry.

## Examples

### Network ACL entries for inbound and outbound traffic

The following example creates a network ACL, and creates two entries in the NACL.
The first entry allows inbound SSH traffic from the specified network. The second
entry allows all outbound IPv4 traffic.

#### JSON

```json

{
    "Resources": {
        "MyNACL": {
            "Type": "AWS::EC2::NetworkAcl",
            "Properties": {
                "VpcId": "vpc-1122334455aabbccd",
                "Tags": [
                    {
                        "Key": "Name",
                        "Value": "NACLforSSHTraffic"
                    }
                ]
            }
        },
        "InboundRule": {
            "Type": "AWS::EC2::NetworkAclEntry",
            "Properties": {
                "NetworkAclId": {
                    "Ref": "MyNACL"
                },
                "RuleNumber": 100,
                "Protocol": 6,
                "RuleAction": "allow",
                "CidrBlock": "172.16.0.0/24",
                "PortRange": {
                    "From": 22,
                    "To": 22
                }
            }
        },
        "OutboundRule": {
            "Type": "AWS::EC2::NetworkAclEntry",
            "Properties": {
                "NetworkAclId": {
                    "Ref": "MyNACL"
                },
                "RuleNumber": 100,
                "Protocol": -1,
                "Egress": true,
                "RuleAction": "allow",
                "CidrBlock": "0.0.0.0/0"
            }
        }
    }
}
```

#### YAML

```yaml

Resources:
  MyNACL:
    Type: AWS::EC2::NetworkAcl
    Properties:
       VpcId: vpc-1122334455aabbccd
       Tags:
       - Key: Name
         Value: NACLforSSHTraffic
  InboundRule:
    Type: AWS::EC2::NetworkAclEntry
    Properties:
       NetworkAclId:
         Ref: MyNACL
       RuleNumber: 100
       Protocol: 6
       RuleAction: allow
       CidrBlock: 172.16.0.0/24
       PortRange:
         From: 22
         To: 22
  OutboundRule:
    Type: AWS::EC2::NetworkAclEntry
    Properties:
       NetworkAclId:
         Ref: MyNACL
       RuleNumber: 100
       Protocol: -1
       Egress: true
       RuleAction: allow
       CidrBlock: 0.0.0.0/0
```

## See also

- [NetworkAclEntry](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateNetworkAclEntry.html) in the _Amazon EC2 API_
_Reference_

- [Network ACLs](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-network-acls.html) in the _Amazon VPC User Guide_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Icmp
