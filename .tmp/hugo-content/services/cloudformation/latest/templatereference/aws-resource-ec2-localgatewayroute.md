This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::LocalGatewayRoute

Creates a static route for the specified local gateway route table. You must specify one of the
following targets:

- `LocalGatewayVirtualInterfaceGroupId`

- `NetworkInterfaceId`

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::LocalGatewayRoute",
  "Properties" : {
      "DestinationCidrBlock" : String,
      "LocalGatewayRouteTableId" : String,
      "LocalGatewayVirtualInterfaceGroupId" : String,
      "NetworkInterfaceId" : String
    }
}

```

### YAML

```yaml

Type: AWS::EC2::LocalGatewayRoute
Properties:
  DestinationCidrBlock: String
  LocalGatewayRouteTableId: String
  LocalGatewayVirtualInterfaceGroupId: String
  NetworkInterfaceId: String

```

## Properties

`DestinationCidrBlock`

The CIDR block used for destination matches.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LocalGatewayRouteTableId`

The ID of the local gateway route table.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LocalGatewayVirtualInterfaceGroupId`

The ID of the virtual interface group.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NetworkInterfaceId`

The ID of the network interface.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the local gateway route table and its destination address range. For example:

`{ "Ref": "lgw-rtb-12346789abcdef|10.0.0.0/24" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`State`

The state of the local gateway route table.

`Type`

The type of local gateway route.

## See also

- [Local\
gateway](../../../outposts/latest/userguide/outposts-local-gateways.md) in _AWS Outposts User_
_Guide_

- [CreateLocalGatewayRouteTable](../../../../reference/awsec2/latest/apireference/api-createlocalgatewayroute.md) in the _Amazon EC2 API_
_Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VCpuCount

AWS::EC2::LocalGatewayRouteTable

All content copied from https://docs.aws.amazon.com/.
