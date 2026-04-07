This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::SecurityGroupVpcAssociation

A security group association with a VPC.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::SecurityGroupVpcAssociation",
  "Properties" : {
      "GroupId" : String,
      "VpcId" : String
    }
}

```

### YAML

```yaml

Type: AWS::EC2::SecurityGroupVpcAssociation
Properties:
  GroupId: String
  VpcId: String

```

## Properties

`GroupId`

The association's security group ID.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VpcId`

The association's VPC ID.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns a concatenation of the security group ID and VPC ID in the format `sg-id|vpc-id`. Example: `sg-a1b2c3d4e5a6b7d8e|vpc-a1b2c3d4e5a6b7d8e`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`State`

The association's state.

`StateReason`

The association's state reason.

`VpcOwnerId`

The AWS account ID of the owner of the VPC.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::EC2::SecurityGroupIngress

AWS::EC2::SnapshotBlockPublicAccess
