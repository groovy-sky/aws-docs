This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::VPCPeeringConnection

Requests a VPC peering connection between two VPCs: a requester VPC that you own and an
accepter VPC with which to create the connection. The accepter VPC can belong to a different
AWS account and can be in a different Region than the requester VPC.

The requester VPC and accepter VPC cannot have overlapping CIDR blocks. If you create a
VPC peering connection request between VPCs with overlapping CIDR blocks, the VPC peering
connection has a status of `failed`.

If the VPCs belong to different accounts, the acceptor account must have a role that
allows the requester account to accept the VPC peering connection. For an example,
see [Walkthrough: Peer with a VPC in another AWS account](../userguide/peer-with-vpc-in-another-account.md).

If the requester and acceptor VPCs are in the same account, the peering request is
accepted without a peering role.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::VPCPeeringConnection",
  "Properties" : {
      "AssumeRoleRegion" : String,
      "PeerOwnerId" : String,
      "PeerRegion" : String,
      "PeerRoleArn" : String,
      "PeerVpcId" : String,
      "Tags" : [ Tag, ... ],
      "VpcId" : String
    }
}

```

### YAML

```yaml

Type: AWS::EC2::VPCPeeringConnection
Properties:
  AssumeRoleRegion: String
  PeerOwnerId: String
  PeerRegion: String
  PeerRoleArn: String
  PeerVpcId: String
  Tags:
    - Tag
  VpcId: String

```

## Properties

`AssumeRoleRegion`

The Region code to use when calling Security Token Service (STS) to assume the
PeerRoleArn, if provided.

_Required_: Conditional

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PeerOwnerId`

The AWS account ID of the owner of the accepter VPC.

Default: Your AWS account ID

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PeerRegion`

The Region code for the accepter VPC, if the accepter VPC is located in a Region
other than the Region in which you make the request.

Default: The Region in which you make the request.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PeerRoleArn`

The Amazon Resource Name (ARN) of the VPC peer role for the peering connection in
another AWS account.

This is required when you are peering a VPC in a different AWS
account.

_Required_: Conditional

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PeerVpcId`

The ID of the VPC with which you are creating the VPC peering connection. You must
specify this parameter in the request.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Any tags assigned to the resource.

_Required_: No

_Type_: Array of [Tag](aws-properties-ec2-vpcpeeringconnection-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcId`

The ID of the VPC.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the VPC peering connection.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Id`

The ID of the peering connection.

## Examples

### Peer VPCs in the same account

This example shows how to peer two VPCs in the same account. It uses an existing
VPC as the requester VPC and creates the accepter VPC.

#### JSON

```json

"Resources": {
    "myVpc": {
        "Type": "AWS::EC2::VPC",
        "Properties": {
            "CidrBlock": "10.0.0.0/16",
            "EnableDnsSupport": true,
            "EnableDnsHostnames": true,
            "Tags": [
                {
                    "Key": "Name",
                    "Value": "accepter-vpc"
                }
            ]
        }
    },
    "vpcPeeringConnection": {
        "Type": "AWS::EC2::VPCPeeringConnection",
        "Properties": {
            "VpcId": "vpc-e03dd489",
            "PeerVpcId": {
                "Ref": "myVpc"
            },
            "Tags": [
                {
                    "Key": "Name",
                    "Value": "cfn-peering-example"
                }
            ]
        }
    }
}
```

#### YAML

```yaml

Resources:
  myVpc:
    Type: AWS::EC2::VPC
    Properties:
      CidrBlock: 10.0.0.0/16
      EnableDnsSupport: true
      EnableDnsHostnames: true
      Tags:
      - Key: Name
        Value: accepter-vpc
  vpcPeeringConnection:
    Type: AWS::EC2::VPCPeeringConnection
    Properties:
      VpcId: vpc-e03dd489
      PeerVpcId: !Ref myVpc
      Tags:
      - Key: Name
        Value: cfn-peering-example
```

## See also

- [Walkthrough: Peer with a VPC in another AWS account](../userguide/peer-with-vpc-in-another-account.md)

- [What is VPC peering](../../../vpc/latest/peering/what-is-vpc-peering.md) in the _VPC Peering Guide_

- [CreateVpcPeeringConnection](../../../../reference/awsec2/latest/apireference/api-createvpcpeeringconnection.md) in the _Amazon EC2 API_
_Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::EC2::VPCGatewayAttachment

Tag

All content copied from https://docs.aws.amazon.com/.
