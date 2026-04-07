This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::VirtualGateway

Creates a virtual gateway.

A virtual gateway allows resources outside your mesh to communicate to resources that
are inside your mesh. The virtual gateway represents an Envoy proxy running in an Amazon ECS task, in a Kubernetes service, or on an Amazon EC2 instance. Unlike a
virtual node, which represents an Envoy running with an application, a virtual gateway
represents Envoy deployed by itself.

For more information about virtual gateways, see [Virtual gateways](https://docs.aws.amazon.com/app-mesh/latest/userguide/virtual_gateways.html).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::AppMesh::VirtualGateway",
  "Properties" : {
      "MeshName" : String,
      "MeshOwner" : String,
      "Spec" : VirtualGatewaySpec,
      "Tags" : [ Tag, ... ],
      "VirtualGatewayName" : String
    }
}

```

### YAML

```yaml

Type: AWS::AppMesh::VirtualGateway
Properties:
  MeshName: String
  MeshOwner: String
  Spec:
    VirtualGatewaySpec
  Tags:
    - Tag
  VirtualGatewayName: String

```

## Properties

`MeshName`

The name of the service mesh that the virtual gateway resides in.

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

The specifications of the virtual gateway.

_Required_: Yes

_Type_: [VirtualGatewaySpec](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appmesh-virtualgateway-virtualgatewayspec.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Optional metadata that you can apply to the virtual gateway to assist with
categorization and organization. Each tag consists of a key and an optional value, both of
which you define. Tag keys can have a maximum character length of 128 characters, and tag values can have
a maximum length of 256 characters.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appmesh-virtualgateway-tag.html)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VirtualGatewayName`

The name of the virtual gateway.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource ARN. For example:

`{ "Ref": "myVirtualGateway" }`

When you pass the logical ID of an `AWS::AppMesh::VirtualGateway` resource to
the intrinsic Ref function, the function returns the gateway route ARN, such as
`arn:aws:appmesh:us-east-1:555555555555:virtualGateway/myVirtualGateway`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The full Amazon Resource Name (ARN) for the virtual gateway.

`MeshName`

The name of the service mesh that the virtual gateway resides in.

`MeshOwner`

The AWS IAM account ID of the service mesh owner. If the account ID is
not your own, then it's the ID of the account that shared the mesh with your account. For
more information about mesh sharing, see [Working with Shared Meshes](https://docs.aws.amazon.com/app-mesh/latest/userguide/sharing.html).

`ResourceOwner`

The AWS IAM account ID of the resource owner. If the account ID is not
your own, then it's the ID of the mesh owner or of another account that the mesh is shared
with. For more information about mesh sharing, see [Working with Shared Meshes](https://docs.aws.amazon.com/app-mesh/latest/userguide/sharing.html).

`Uid`

The unique identifier for the virtual gateway.

`VirtualGatewayName`

The name of the virtual gateway.

## See also

- [Virtual gateways](https://docs.aws.amazon.com/app-mesh/latest/userguide/virtual_gateways.html) in the _AWS App Mesh User Guide_.

- [CreateVirtualGateway](https://docs.aws.amazon.com/app-mesh/latest/APIReference/API_CreateVirtualGateway.html) in the _AWS App Mesh API Reference_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

WeightedTarget

JsonFormatRef
