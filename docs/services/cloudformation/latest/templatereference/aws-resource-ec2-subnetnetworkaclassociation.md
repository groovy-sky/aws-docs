This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::SubnetNetworkAclAssociation

Associates a subnet with a network ACL. For more information, see [ReplaceNetworkAclAssociation](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-ReplaceNetworkAclAssociation.html) in the _Amazon EC2 API_
_Reference_.

When `AWS::EC2::SubnetNetworkAclAssociation` resources are created during
create or update operations, AWS CloudFormation adopts existing resources that share
the same key properties (the properties that contribute to uniquely identify the resource).
However, if the operation fails and rolls back, AWS CloudFormation deletes the
previously out-of-band resources. You can protect against this behavior by using
`Retain` deletion policies. For more information, see [DeletionPolicy Attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::SubnetNetworkAclAssociation",
  "Properties" : {
      "NetworkAclId" : String,
      "SubnetId" : String
    }
}

```

### YAML

```yaml

Type: AWS::EC2::SubnetNetworkAclAssociation
Properties:
  NetworkAclId: String
  SubnetId: String

```

## Properties

`NetworkAclId`

The ID of the network ACL.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SubnetId`

The ID of the subnet.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the subnet network ACL association.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

`AssociationId`

Returns the value of this object's AssociationId property.

## Examples

### Subnet network ACL association

The following example associates subnet mySubnet with the myNetworkAcl network
ACL.

#### JSON

```json

"mySubnetNetworkAclAssociation" : {
   "Type" : "AWS::EC2::SubnetNetworkAclAssociation",
   "Properties" : {
      "SubnetId" : { "Ref" : "mySubnet" },
      "NetworkAclId" : { "Ref" : "myNetworkAcl" }
   }
}
```

#### YAML

```yaml

   mySubnetNetworkAclAssociation:
     Type: AWS::EC2::SubnetNetworkAclAssociation
     Properties:
       SubnetId:
         Ref: mySubnet
       NetworkAclId:
         Ref: myNetworkAcl
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::EC2::SubnetCidrBlock

AWS::EC2::SubnetRouteTableAssociation
