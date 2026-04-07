This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::SecurityGroup

Specifies a security group.

You must specify ingress rules to allow inbound traffic. By default, no inbound
traffic is allowed.

When you create a security group, if you do not add egress rules, we add egress
rules that allow all outbound IPv4 and IPv6 traffic. Otherwise, we do not add them.
After the security group is created, if you remove all egress rules that you added,
we do not add egress rules, so no outbound traffic is allowed.

If you modify a rule, CloudFormation removes the existing rule and then adds a new rule.
There is a brief period when neither the original rule or the new rule exists, so the
corresponding traffic is dropped.

This type supports updates. For more information about updating stacks, see [AWS CloudFormation Stacks Updates](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks.html).

###### Important

To cross-reference two security groups in the ingress and egress rules of those
security groups, use the [AWS::EC2::SecurityGroupEgress](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-security-group-egress.html) and [AWS::EC2::SecurityGroupIngress](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-security-group-ingress.html) resources to define your rules. Do not use
the embedded ingress and egress rules in the `AWS::EC2::SecurityGroup`. Doing
so creates a circular dependency, which CloudFormation doesn't allow.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::SecurityGroup",
  "Properties" : {
      "GroupDescription" : String,
      "GroupName" : String,
      "SecurityGroupEgress" : [ Egress, ... ],
      "SecurityGroupIngress" : [ Ingress, ... ],
      "Tags" : [ Tag, ... ],
      "VpcId" : String
    }
}

```

### YAML

```yaml

Type: AWS::EC2::SecurityGroup
Properties:
  GroupDescription: String
  GroupName: String
  SecurityGroupEgress:
    - Egress
  SecurityGroupIngress:
    - Ingress
  Tags:
    - Tag
  VpcId: String

```

## Properties

`GroupDescription`

A description for the security group.

Constraints: Up to 255 characters in length

Valid characters: a-z, A-Z, 0-9, spaces, and .\_-:/()#,@\[\]+=&;{}!$\*

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`GroupName`

The name of the security group. Names are case-insensitive and must be unique within the VPC.

Constraints: Up to 255 characters in length. Can't start with `sg-`.

Valid characters: a-z, A-Z, 0-9, spaces, and .\_-:/()#,@\[\]+=&;{}!$\*

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SecurityGroupEgress`

The outbound rules associated with the security group.

_Required_: No

_Type_: Array of [Egress](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-securitygroup-egress.html)

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`SecurityGroupIngress`

The inbound rules associated with the security group.

_Required_: No

_Type_: Array of [Ingress](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-securitygroup-ingress.html)

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`Tags`

Any tags assigned to the security group.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-securitygroup-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcId`

The ID of the VPC for the security group. If you do not specify a VPC, the default is
to use the default VPC for the Region. If there's no specified VPC and no default VPC,
security group creation fails.

_Required_: Conditional

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the security group if you specified the `VpcId` property.
Otherwise, it returns the name of the security group. If you omit the `VpcId` property
and need the ID of the security group, use `Fn::GetAtt` instead.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`GroupId`

The ID of the security group, such as `sg-94b3a1f6`.

`VpcId`

The ID of the VPC, such as `vpc-0669f8f9`.

## Examples

- [Define basic ingress and egress rules](#aws-resource-ec2-securitygroup--examples--Define_basic_ingress_and_egress_rules)

- [Remove the default rule](#aws-resource-ec2-securitygroup--examples--Remove_the_default_rule)

- [Allow ping requests](#aws-resource-ec2-securitygroup--examples--Allow_ping_requests)

### Define basic ingress and egress rules

The following example specifies a security group with an ingress and egress rule.

#### JSON

```json

"InstanceSecurityGroup" : {
    "Type" : "AWS::EC2::SecurityGroup",
    "Properties" : {
        "GroupDescription" : "Allow http to client host",
        "VpcId" : {"Ref" : "myVPC"},
        "SecurityGroupIngress" : [{
            "IpProtocol" : "tcp",
            "FromPort" : 80,
            "ToPort" : 80,
            "CidrIp" : "0.0.0.0/0"
        }],
        "SecurityGroupEgress" : [{
            "IpProtocol" : "tcp",
            "FromPort" : 80,
            "ToPort" : 80,
            "CidrIp" : "0.0.0.0/0"
        }]
    }
}
```

#### YAML

```yaml

InstanceSecurityGroup:
  Type: AWS::EC2::SecurityGroup
  Properties:
    GroupDescription: Allow http to client host
    VpcId: !Ref myVPC
    SecurityGroupIngress:
      - IpProtocol: tcp
        FromPort: 80
        ToPort: 80
        CidrIp: 0.0.0.0/0
    SecurityGroupEgress:
      - IpProtocol: tcp
        FromPort: 80
        ToPort: 80
        CidrIp: 0.0.0.0/0
```

### Remove the default rule

When you specify a VPC security group, Amazon EC2 creates a default egress rule
that allows egress traffic on all ports and IP protocols to any location. The default
rule is removed only when you specify one or more egress rules. If you want to remove
the default rule and limit egress traffic to just the localhost (127.0.0.1/32), use
the following example.

#### JSON

```json

"sgwithoutegress": {
    "Type": "AWS::EC2::SecurityGroup",
    "Properties": {
        "GroupDescription": "Limits security group egress traffic",
        "SecurityGroupEgress": [{
            "CidrIp": "127.0.0.1/32",
            "IpProtocol": "-1"
        }],
        "VpcId": { "Ref": "myVPC"}
    }
}
```

#### YAML

```yaml

sgwithoutegress:
  Type: AWS::EC2::SecurityGroup
  Properties:
    GroupDescription: Limits security group egress traffic
    SecurityGroupEgress:
      - CidrIp: 127.0.0.1/32
        IpProtocol: "-1"
    VpcId: !Ref myVPC
```

### Allow ping requests

To allow ping requests, add the ICMP protocol type and specify 8 (echo request)
for the ICMP type and either 0 or -1 (all) for the ICMP code.

#### JSON

```json

"SGPing" : {
    "Type" : "AWS::EC2::SecurityGroup",
    "DependsOn": "VPC",
    "Properties" : {
        "GroupDescription" : "SG to test ping",
        "VpcId" : {"Ref" : "VPC"},
        "SecurityGroupIngress" : [
        {
            "IpProtocol" : "tcp",
            "FromPort" : 22,
            "ToPort" : 22,
            "CidrIp" : "10.0.0.0/24"
        },
        {
            "IpProtocol" : "icmp",
            "FromPort" : 8,
            "ToPort" : -1,
            "CidrIp" : "10.0.0.0/24"
        }]
    }
}
```

#### YAML

```yaml

SGPing:
  Type: AWS::EC2::SecurityGroup
  DependsOn: VPC
  Properties:
    GroupDescription: SG to test ping
    VpcId: !Ref VPC
    SecurityGroupIngress:
      - IpProtocol: tcp
        FromPort: 22
        ToPort: 22
        CidrIp: 10.0.0.0/24
      - IpProtocol: icmp
        FromPort: 8
        ToPort: -1
        CidrIp: 10.0.0.0/24
```

## See also

- [Security groups for your VPC](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_SecurityGroups.html) in the _Amazon VPC User_
_Guide_

- [Amazon EC2 security groups](../../../ec2/latest/userguide/ec2-security-groups.md) in the _Amazon EC2 User Guide_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Egress
