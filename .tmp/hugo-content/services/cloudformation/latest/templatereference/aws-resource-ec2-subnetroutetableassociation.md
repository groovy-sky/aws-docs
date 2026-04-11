This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::SubnetRouteTableAssociation

Associates a subnet with a route table. The subnet and route table must be in the same
VPC. This association causes traffic originating from the subnet to be routed according to
the routes in the route table. A route table can be associated with multiple subnets. To
create a route table, see [AWS::EC2::RouteTable](../userguide/aws-resource-ec2-routetable.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::SubnetRouteTableAssociation",
  "Properties" : {
      "RouteTableId" : String,
      "SubnetId" : String
    }
}

```

### YAML

```yaml

Type: AWS::EC2::SubnetRouteTableAssociation
Properties:
  RouteTableId: String
  SubnetId: String

```

## Properties

`RouteTableId`

The ID of the route table.

The physical ID changes when the route table ID is changed.

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

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the subnet route table association.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Id`

The ID of the subnet route table association.

## Examples

### Subnet route table association

The following example associates a subnet with a route table.

#### JSON

```json

"mySubnetRouteTableAssociation" : {
   "Type" : "AWS::EC2::SubnetRouteTableAssociation",
   "Properties" : {
      "SubnetId" : { "Ref" : "mySubnet" },
      "RouteTableId" : { "Ref" : "myRouteTable" }
   }
}
```

#### YAML

```yaml

  mySubnetRouteTableAssociation:
    Type: AWS::EC2::SubnetRouteTableAssociation
    Properties:
      SubnetId:
        Ref: mySubnet
      RouteTableId:
        Ref: myRouteTable
```

## See also

- [AssociateRouteTable](../../../../reference/awsec2/latest/apireference/api-associateroutetable.md) in the _Amazon EC2 API_
_Reference_

- [Route tables](../../../vpc/latest/userguide/vpc-route-tables.md) in the _Amazon VPC User Guide_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::EC2::SubnetNetworkAclAssociation

AWS::EC2::TrafficMirrorFilter

All content copied from https://docs.aws.amazon.com/.
