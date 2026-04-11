This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::Mesh

Creates a service mesh.

A service mesh is a logical boundary for network traffic between services that are
represented by resources within the mesh. After you create your service mesh, you can
create virtual services, virtual nodes, virtual routers, and routes to distribute traffic
between the applications in your mesh.

For more information about service meshes, see [Service meshes](../../../app-mesh/latest/userguide/meshes.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::AppMesh::Mesh",
  "Properties" : {
      "MeshName" : String,
      "Spec" : MeshSpec,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::AppMesh::Mesh
Properties:
  MeshName: String
  Spec:
    MeshSpec
  Tags:
    - Tag

```

## Properties

`MeshName`

The name to use for the service mesh.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Spec`

The service mesh specification to apply.

_Required_: No

_Type_: [MeshSpec](aws-properties-appmesh-mesh-meshspec.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Optional metadata that you can apply to the service mesh to assist with categorization
and organization. Each tag consists of a key and an optional value, both of which you
define. Tag keys can have a maximum character length of 128 characters, and tag values can have
a maximum length of 256 characters.

_Required_: No

_Type_: Array of [Tag](aws-properties-appmesh-mesh-tag.md)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource ARN. For example:

`{ "Ref": "myMesh" }`

When you pass the logical ID of an `AWS::AppMesh::Mesh` resource to the
intrinsic Ref function, the function returns the mesh ARN, such as
`arn:aws:appmesh:us-east-1:555555555555:mesh/myMesh`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The full Amazon Resource Name (ARN) for the mesh.

`MeshName`

The name of the service mesh.

`MeshOwner`

The IAM account ID of the service mesh owner. If the account ID is not
your own, then it's the ID of the account that shared the mesh with your account. For more
information about mesh sharing, see [Working with Shared Meshes](../../../app-mesh/latest/userguide/sharing.md).

`ResourceOwner`

The IAM account ID of the resource owner. If the account ID is not your
own, then it's the ID of the mesh owner or of another account that the mesh is shared with.
For more information about mesh sharing, see [Working with Shared Meshes](../../../app-mesh/latest/userguide/sharing.md).

`Uid`

The unique identifier for the mesh.

## Examples

### Create a Service Mesh

This example creates a service mesh that allows all egress traffic.

#### JSON

```json

{
   "Description": "Basic Test Mesh",
   "Resources": {
      "BasicMesh": {
         "Type": "AWS::AppMesh::Mesh",
         "Properties": {
            "MeshName": "BasicMesh1",
            "Spec": {
               "EgressFilter": {
                  "Type": "ALLOW_ALL"
               }
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
      "MeshName": {
         "Description": "Name of the Mesh",
         "Value": {
            "Fn::GetAtt": [
               "BasicMesh",
               "MeshName"
            ]
         }
      },
      "Arn": {
         "Description": "Arn of the Mesh created",
         "Value": {
            "Fn::GetAtt": [
               "BasicMesh",
               "Arn"
            ]
         }
      },
      "Uid": {
         "Description": "Uid of the Mesh created",
         "Value": {
            "Fn::GetAtt": [
               "BasicMesh",
               "Uid"
            ]
         }
      }
   }
}
```

#### YAML

```yaml

Description: "Basic Test Mesh"
Resources:
  BasicMesh:
    Type: "AWS::AppMesh::Mesh"
    Properties:
      MeshName: "BasicMesh1"
      Spec:
        EgressFilter:
          Type: "ALLOW_ALL"
      Tags:
      - Key: "Key1"
        Value: "Value1"
      - Key: "Key2"
        Value: "Value2"
Outputs:
  MeshName:
    Description: Name of the Mesh
    Value:
      Fn::GetAtt:
      - BasicMesh
      - MeshName
  Arn:
    Description: Arn of the Mesh created
    Value:
      Fn::GetAtt:
      - BasicMesh
      - Arn
  Uid:
    Description: Uid of the Mesh created
    Value:
      Fn::GetAtt:
      - BasicMesh
      - Uid
```

## See also

- [Service\
Meshes](../../../app-mesh/latest/userguide/meshes.md) in the _AWS App Mesh User Guide_.

- [CreateMesh](../../../../reference/app-mesh/latest/apireference/api-createmesh.md) in the _AWS App Mesh API Reference_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

EgressFilter

All content copied from https://docs.aws.amazon.com/.
