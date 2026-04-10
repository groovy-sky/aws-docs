This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Redshift::ClusterSecurityGroupIngress

Adds an inbound (ingress) rule to an Amazon Redshift security group. Depending on whether
the application accessing your cluster is running on the Internet or an Amazon EC2
instance, you can authorize inbound access to either a Classless Interdomain Routing
(CIDR)/Internet Protocol (IP) range or to an Amazon EC2 security group. You can add as
many as 20 ingress rules to an Amazon Redshift security group.

If you authorize access to an Amazon EC2 security group, specify
_EC2SecurityGroupName_ and
_EC2SecurityGroupOwnerId_. The Amazon EC2 security group and
Amazon Redshift cluster must be in the same AWS Region.

If you authorize access to a CIDR/IP address range, specify
_CIDRIP_. For an overview of CIDR blocks, see the Wikipedia
article on [Classless Inter-Domain Routing](http://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing).

You must also associate the security group with a cluster so that clients running
on these IP addresses or the EC2 instance are authorized to connect to the cluster. For
information about managing security groups, go to [Working with Security\
Groups](../../../redshift/latest/mgmt/working-with-security-groups.md) in the _Amazon Redshift Cluster Management Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Redshift::ClusterSecurityGroupIngress",
  "Properties" : {
      "CIDRIP" : String,
      "ClusterSecurityGroupName" : String,
      "EC2SecurityGroupName" : String,
      "EC2SecurityGroupOwnerId" : String
    }
}

```

### YAML

```yaml

Type: AWS::Redshift::ClusterSecurityGroupIngress
Properties:
  CIDRIP: String
  ClusterSecurityGroupName: String
  EC2SecurityGroupName: String
  EC2SecurityGroupOwnerId: String

```

## Properties

`CIDRIP`

The IP range to be added the Amazon Redshift security group.

_Required_: No

_Type_: String

_Maximum_: `2147483647`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ClusterSecurityGroupName`

The name of the security group to which the ingress rule is added.

_Required_: Yes

_Type_: String

_Maximum_: `2147483647`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EC2SecurityGroupName`

The EC2 security group to be added the Amazon Redshift security group.

_Required_: No

_Type_: String

_Maximum_: `2147483647`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EC2SecurityGroupOwnerId`

The AWS account number of the owner of the security group specified by
the _EC2SecurityGroupName_ parameter. The AWS Access
Key ID is not an acceptable value.

Example: `111122223333`

Conditional. If you specify the `EC2SecurityGroupName` property, you must
specify this property.

_Required_: No

_Type_: String

_Maximum_: `2147483647`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Id`

Specifies an inbound (ingress) rule for an Amazon Redshift security group.

## Examples

### Ingress Rules

The following snippet describes a ingress rules for an Amazon Redshift
cluster security group.

#### JSON

```json

"myClusterSecurityGroupIngressIP" : {
  "Type": "AWS::Redshift::ClusterSecurityGroupIngress",
  "Properties": {
    "ClusterSecurityGroupName" : {"Ref":"myClusterSecurityGroup"},
    "CIDRIP" : "10.0.0.0/16"
  }
}
```

#### YAML

```yaml

myClusterSecurityGroupIngressIP:
  Type: "AWS::Redshift::ClusterSecurityGroupIngress"
  Properties:
    ClusterSecurityGroupName:
    Ref: "myClusterSecurityGroup"
    CIDRIP: "10.0.0.0/16"
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::Redshift::ClusterSubnetGroup

All content copied from https://docs.aws.amazon.com/.
