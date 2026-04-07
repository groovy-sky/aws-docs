This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::VirtualService

Creates a virtual service within a service mesh.

A virtual service is an abstraction of a real service that is provided by a virtual node
directly or indirectly by means of a virtual router. Dependent services call your virtual
service by its `virtualServiceName`, and those requests are routed to the
virtual node or virtual router that is specified as the provider for the virtual
service.

For more information about virtual services, see [Virtual services](https://docs.aws.amazon.com/app-mesh/latest/userguide/virtual_services.html).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::AppMesh::VirtualService",
  "Properties" : {
      "MeshName" : String,
      "MeshOwner" : String,
      "Spec" : VirtualServiceSpec,
      "Tags" : [ Tag, ... ],
      "VirtualServiceName" : String
    }
}

```

### YAML

```yaml

Type: AWS::AppMesh::VirtualService
Properties:
  MeshName: String
  MeshOwner: String
  Spec:
    VirtualServiceSpec
  Tags:
    - Tag
  VirtualServiceName: String

```

## Properties

`MeshName`

The name of the service mesh to create the virtual service in.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MeshOwner`

The AWS IAM account ID of the service mesh owner. If the account ID is not your own, then
the account that you specify must share the mesh with your account before you can create
the resource in the service mesh. For more information about mesh sharing, see [Working with shared meshes](https://docs.aws.amazon.com/app-mesh/latest/userguide/sharing.html).

_Required_: No

_Type_: String

_Minimum_: `12`

_Maximum_: `12`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Spec`

The virtual service specification to apply.

_Required_: Yes

_Type_: [VirtualServiceSpec](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appmesh-virtualservice-virtualservicespec.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Optional metadata that you can apply to the virtual service to assist with
categorization and organization. Each tag consists of a key and an optional value, both of
which you define. Tag keys can have a maximum character length of 128 characters, and tag values can have
a maximum length of 256 characters.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appmesh-virtualservice-tag.html)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VirtualServiceName`

The name to use for the virtual service.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource ARN. For example:

`{ "Ref": "myVirtualService" }`

When you pass the logical ID of an `AWS::AppMesh::VirtualService` resource to
the intrinsic Ref function, the function returns the virtual service ARN, such as
`arn:aws:appmesh:us-east-1:555555555555:virtualService/myVirtualService`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The full Amazon Resource Name (ARN) for the virtual service.

`MeshName`

The name of the service mesh that the virtual service resides in.

`MeshOwner`

The AWS IAM account ID of the service mesh owner. If the account ID is
not your own, then it's the ID of the account that shared the mesh with your account. For
more information about mesh sharing, see [Working with Shared Meshes](https://docs.aws.amazon.com/app-mesh/latest/userguide/sharing.html).

`ResourceOwner`

The AWS IAM account ID of the resource owner. If the account ID is not
your own, then it's the ID of the mesh owner or of another account that the mesh is shared
with. For more information about mesh sharing, see [Working with Shared Meshes](https://docs.aws.amazon.com/app-mesh/latest/userguide/sharing.html).

`Uid`

The unique identifier for the virtual service.

`VirtualServiceName`

The name of the virtual service.

## See also

- [Virtual Services](https://docs.aws.amazon.com/app-mesh/latest/userguide/virtual_services.html) in the _AWS App Mesh User Guide_.

- [CreateVirtualService](https://docs.aws.amazon.com/app-mesh/latest/APIReference/API_CreateVirtualService.html) in the _AWS App Mesh API Reference_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

VirtualRouterSpec

Tag
