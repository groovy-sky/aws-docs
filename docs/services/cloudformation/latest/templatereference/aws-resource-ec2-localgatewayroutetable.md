This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::LocalGatewayRouteTable

Describes a local gateway route table.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::LocalGatewayRouteTable",
  "Properties" : {
      "LocalGatewayId" : String,
      "Mode" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::EC2::LocalGatewayRouteTable
Properties:
  LocalGatewayId: String
  Mode: String
  Tags:
    - Tag

```

## Properties

`LocalGatewayId`

The ID of the local gateway.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Mode`

The mode of the local gateway route table.

_Required_: No

_Type_: String

_Allowed values_: `direct-vpc-routing | coip`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags assigned to the local gateway route table.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-localgatewayroutetable-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the local gateway route table. For example:

`{ "Ref": "lgw-rtb-059615ef7deEXAMPLE" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`LocalGatewayRouteTableArn`

The Amazon Resource Name (ARN) of the local gateway route table.

`LocalGatewayRouteTableId`

The ID of the local gateway route table.

`OutpostArn`

The Amazon Resource Name (ARN) of the Outpost.

`OwnerId`

The ID of the AWS account that owns the local gateway route table.

`State`

The state of the local gateway route table.

## See also

- [Local\
gateway](https://docs.aws.amazon.com/outposts/latest/userguide/outposts-local-gateways.html) in _AWS Outposts User_
_Guide_

- [CreateLocalGatewayRouteTable](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateLocalGatewayRoute.html) in the _Amazon EC2 API_
_Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::EC2::LocalGatewayRoute

Tag
