This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::VirtualRouter

Creates a virtual router within a service mesh.

Specify a `listener` for any inbound traffic that your virtual router
receives. Create a virtual router for each protocol and port that you need to route.
Virtual routers handle traffic for one or more virtual services within your mesh. After you
create your virtual router, create and associate routes for your virtual router that direct
incoming requests to different virtual nodes.

For more information about virtual routers, see [Virtual routers](../../../app-mesh/latest/userguide/virtual-routers.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::AppMesh::VirtualRouter",
  "Properties" : {
      "MeshName" : String,
      "MeshOwner" : String,
      "Spec" : VirtualRouterSpec,
      "Tags" : [ Tag, ... ],
      "VirtualRouterName" : String
    }
}

```

### YAML

```yaml

Type: AWS::AppMesh::VirtualRouter
Properties:
  MeshName: String
  MeshOwner: String
  Spec:
    VirtualRouterSpec
  Tags:
    - Tag
  VirtualRouterName: String

```

## Properties

`MeshName`

The name of the service mesh to create the virtual router in.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MeshOwner`

The AWS IAM account ID of the service mesh owner. If the account ID is not your own, then
the account that you specify must share the mesh with your account before you can create
the resource in the service mesh. For more information about mesh sharing, see [Working with shared meshes](../../../app-mesh/latest/userguide/sharing.md).

_Required_: No

_Type_: String

_Minimum_: `12`

_Maximum_: `12`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Spec`

The virtual router specification to apply.

_Required_: Yes

_Type_: [VirtualRouterSpec](aws-properties-appmesh-virtualrouter-virtualrouterspec.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Optional metadata that you can apply to the virtual router to assist with categorization
and organization. Each tag consists of a key and an optional value, both of which you
define. Tag keys can have a maximum character length of 128 characters, and tag values can have
a maximum length of 256 characters.

_Required_: No

_Type_: Array of [Tag](aws-properties-appmesh-virtualrouter-tag.md)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VirtualRouterName`

The name to use for the virtual router.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource ARN. For example:

`{ "Ref": "myVirtualRouter" }`

When you pass the logical ID of an `AWS::AppMesh::VirtualRouter` resource to
the intrinsic Ref function, the function returns the virtual router ARN, such as
`arn:aws:appmesh:us-east-1:555555555555:virtualRouter/myVirtualRouter`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The full Amazon Resource Name (ARN) for the virtual router.

`MeshName`

The name of the service mesh that the virtual router resides in.

`MeshOwner`

The AWS IAM account ID of the service mesh owner. If the account ID is
not your own, then it's the ID of the account that shared the mesh with your account. For
more information about mesh sharing, see [Working with Shared Meshes](../../../app-mesh/latest/userguide/sharing.md).

`ResourceOwner`

The AWS IAM account ID of the resource owner. If the account ID is not
your own, then it's the ID of the mesh owner or of another account that the mesh is shared
with. For more information about mesh sharing, see [Working with Shared Meshes](../../../app-mesh/latest/userguide/sharing.md).

`Uid`

The unique identifier for the virtual router.

`VirtualRouterName`

The name of the virtual router.

## Examples

### Create a Virtual Router

This example creates a basic virtual router with an HTTP port mapping and two
tags.

#### JSON

```json

{
  "Description": "Basic Test Virtual Router",
  "Resources": {
    "BasicVirtualRouter": {
      "Type": "AWS::AppMesh::VirtualRouter",
      "Properties": {
        "VirtualRouterName": "TestVirtualRouter",
        "MeshName": null,
        "Spec": {
          "Listeners": [
            {
              "PortMapping": {
                "Port": 8080,
                "Protocol": "http"
              }
            }
          ]
        },
        "Tags": [
          {
            "Key": "Key1",
            "Value": "Value1"
          },
          {
            "Key": "Key2",
            "Value": "Value2"
          }
        ]
      }
    }
  },
  "Outputs": {
    "VirtualRouterName": {
      "Description": "Name of the VirtualRouter",
      "Value": {
        "Fn::GetAtt": [
          "BasicVirtualRouter",
          "VirtualRouterName"
        ]
      }
    },
    "MeshName": {
      "Description": "Name of the Mesh",
      "Value": {
        "Fn::GetAtt": [
          "BasicVirtualRouter",
          "MeshName"
        ]
      }
    },
    "Arn": {
      "Description": "Arn of the VirtualRouter created",
      "Value": {
        "Fn::GetAtt": [
          "BasicVirtualRouter",
          "Arn"
        ]
      }
    },
    "Uid": {
      "Description": "Uid of the VirtualRouter created",
      "Value": {
        "Fn::GetAtt": [
          "BasicVirtualRouter",
          "Uid"
        ]
      }
    }
  }
}
```

#### YAML

```yaml

Description: Basic Test Virtual Router
Resources:
  BasicVirtualRouter:
    Type: AWS::AppMesh::VirtualRouter
    Properties:
      VirtualRouterName: TestVirtualRouter
      MeshName:
      Spec:
        Listeners:
        - PortMapping:
            Port: 8080
            Protocol: http
      Tags:
      - Key: Key1
        Value: Value1
      - Key: Key2
        Value: Value2
Outputs:
  VirtualRouterName:
    Description: Name of the VirtualRouter
    Value:
      Fn::GetAtt:
      - BasicVirtualRouter
      - VirtualRouterName
  MeshName:
    Description: Name of the Mesh
    Value:
      Fn::GetAtt:
      - BasicVirtualRouter
      - MeshName
  Arn:
    Description: Arn of the VirtualRouter created
    Value:
      Fn::GetAtt:
      - BasicVirtualRouter
      - Arn
  Uid:
    Description: Uid of the VirtualRouter created
    Value:
      Fn::GetAtt:
      - BasicVirtualRouter
      - Uid

```

## See also

- [Virtual Routers](../../../app-mesh/latest/userguide/virtual-routers.md) in the _AWS App Mesh User Guide_.

- [CreateVirtualRouter](../../../../reference/app-mesh/latest/apireference/api-createvirtualrouter.md) in the _AWS App Mesh API Reference_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VirtualServiceBackend

PortMapping

All content copied from https://docs.aws.amazon.com/.
