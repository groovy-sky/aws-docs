This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::RouteTable

Specifies a route table for the specified VPC. After you create a route table, you can add
routes and associate the table with a subnet.

For more information, see [Route tables](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Route_Tables.html)
in the _Amazon VPC User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::RouteTable",
  "Properties" : {
      "Tags" : [ Tag, ... ],
      "VpcId" : String
    }
}

```

### YAML

```yaml

Type: AWS::EC2::RouteTable
Properties:
  Tags:
    - Tag
  VpcId: String

```

## Properties

`Tags`

Any tags assigned to the route table.

_Required_: No

_Type_: Array of [Tag](aws-properties-ec2-routetable-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcId`

The ID of the VPC.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the route table.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`RouteTableId`

The ID of the route table.

## Examples

### Route table

The following example uses the VPC ID from a VPC named myVPC that was declared
elsewhere in the same template.

#### JSON

```json

"myRouteTable" : {
   "Type" : "AWS::EC2::RouteTable",
   "Properties" : {
      "VpcId" : { "Ref" : "myVPC" },
      "Tags" : [ { "Key" : "stack", "Value" : "production" } ]
   }
}
```

#### YAML

```yaml

  myRouteTable:
    Type: AWS::EC2::RouteTable
    Properties:
      VpcId:
        Ref: myVPC
      Tags:
      - Key: stack
        Value: production
```

## See also

- [AWS::EC2::Route](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-route.html)

- [CreateRouteTable](../../../../reference/awsec2/latest/apireference/api-createroutetable.md) in the _Amazon EC2 API_
_Reference_

- [Route\
tables](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Route_Tables.html) in the _Amazon VPC User Guide_

- [Tag your\
Amazon EC2 resources](../../../ec2/latest/userguide/using-tags.md) in the _Amazon EC2 User_
_Guide_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::EC2::RouteServerPropagation

Tag
