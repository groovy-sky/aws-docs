This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::GatewayRoute

Creates a gateway route.

A gateway route is attached to a virtual gateway and routes traffic to an existing
virtual service. If a route matches a request, it can distribute traffic to a target
virtual service.

For more information about gateway routes, see [Gateway routes](https://docs.aws.amazon.com/app-mesh/latest/userguide/gateway-routes.html).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::AppMesh::GatewayRoute",
  "Properties" : {
      "GatewayRouteName" : String,
      "MeshName" : String,
      "MeshOwner" : String,
      "Spec" : GatewayRouteSpec,
      "Tags" : [ Tag, ... ],
      "VirtualGatewayName" : String
    }
}

```

### YAML

```yaml

Type: AWS::AppMesh::GatewayRoute
Properties:
  GatewayRouteName: String
  MeshName: String
  MeshOwner: String
  Spec:
    GatewayRouteSpec
  Tags:
    - Tag
  VirtualGatewayName: String

```

## Properties

`GatewayRouteName`

The name of the gateway route.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MeshName`

The name of the service mesh that the resource resides in.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MeshOwner`

The AWS IAM account ID of the service mesh owner. If the account ID is not your own, then it's
the ID of the account that shared the mesh with your account. For more information about mesh sharing, see [Working with shared meshes](https://docs.aws.amazon.com/app-mesh/latest/userguide/sharing.html).

_Required_: No

_Type_: String

_Minimum_: `12`

_Maximum_: `12`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Spec`

The specifications of the gateway route.

_Required_: Yes

_Type_: [GatewayRouteSpec](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appmesh-gatewayroute-gatewayroutespec.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Optional metadata that you can apply to the gateway route to assist with categorization
and organization. Each tag consists of a key and an optional value, both of which you
define. Tag keys can have a maximum character length of 128 characters, and tag values can have
a maximum length of 256 characters.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appmesh-gatewayroute-tag.html)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VirtualGatewayName`

The virtual gateway that the gateway route is associated with.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource ARN. For example:

`{ "Ref": "myGatewayRoute" }`

When you pass the logical ID of an `AWS::AppMesh::GatewayRoute` resource to
the intrinsic Ref function, the function returns the gateway route ARN, such as
`arn:aws:appmesh:us-east-1:555555555555:gatewayRoute/myGatewayRoute`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The full Amazon Resource Name (ARN) for the gateway route.

`GatewayRouteName`

The name of the gateway route.

`MeshName`

The name of the service mesh that the gateway route resides in.

`MeshOwner`

The AWS IAM account ID of the service mesh owner. If the account ID is
not your own, then it's the ID of the account that shared the mesh with your account. For
more information about mesh sharing, see [Working with Shared Meshes](https://docs.aws.amazon.com/app-mesh/latest/userguide/sharing.html).

`ResourceOwner`

The IAM account ID of the resource owner. If the account ID is not your
own, then it's the ID of the mesh owner or of another account that the mesh is shared with.
For more information about mesh sharing, see [Working with Shared Meshes](https://docs.aws.amazon.com/app-mesh/latest/userguide/sharing.html).

`Uid`

The unique identifier for the gateway route.

`VirtualGatewayName`

The name of the virtual gateway that the gateway route is associated with.

## See also

- [Gateway routes](https://docs.aws.amazon.com/app-mesh/latest/userguide/gateway-routes.html) in the _AWS App Mesh User Guide_.

- [CreateGatewayRoute](https://docs.aws.amazon.com/app-mesh/latest/APIReference/API_CreateGatewayRoute.html) in the _AWS App Mesh API Reference_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS App Mesh

GatewayRouteHostnameMatch
