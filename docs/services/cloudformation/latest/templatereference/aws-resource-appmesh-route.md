This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::Route

Creates a route that is associated with a virtual router.

You can route several different protocols and define a retry policy for a route.
Traffic can be routed to one or more virtual nodes.

For more information about routes, see [Routes](https://docs.aws.amazon.com/app-mesh/latest/userguide/routes.html).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::AppMesh::Route",
  "Properties" : {
      "MeshName" : String,
      "MeshOwner" : String,
      "RouteName" : String,
      "Spec" : RouteSpec,
      "Tags" : [ Tag, ... ],
      "VirtualRouterName" : String
    }
}

```

### YAML

```yaml

Type: AWS::AppMesh::Route
Properties:
  MeshName: String
  MeshOwner: String
  RouteName: String
  Spec:
    RouteSpec
  Tags:
    - Tag
  VirtualRouterName: String

```

## Properties

`MeshName`

The name of the service mesh to create the route in.

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

`RouteName`

The name to use for the route.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Spec`

The route specification to apply.

_Required_: Yes

_Type_: [RouteSpec](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appmesh-route-routespec.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Optional metadata that you can apply to the route to assist with categorization and
organization. Each tag consists of a key and an optional value, both of which you define.
Tag keys can have a maximum character length of 128 characters, and tag values can have
a maximum length of 256 characters.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appmesh-route-tag.html)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VirtualRouterName`

The name of the virtual router in which to create the route. If the virtual router is in
a shared mesh, then you must be the owner of the virtual router resource.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource ARN. For example:

`{ "Ref": "myRoute" }`

When you pass the logical ID of an `AWS::AppMesh::Route` resource to the
intrinsic Ref function, the function returns the route ARN, such as
`arn:aws:appmesh:us-east-1:555555555555:route/myRoute`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The full Amazon Resource Name (ARN) for the route.

`MeshName`

The name of the service mesh that the route resides in.

`MeshOwner`

The AWS IAM account ID of the service mesh owner. If the account ID is
not your own, then it's the ID of the account that shared the mesh with your account. For
more information about mesh sharing, see [Working with Shared Meshes](https://docs.aws.amazon.com/app-mesh/latest/userguide/sharing.html).

`ResourceOwner`

The AWS IAM account ID of the resource owner. If the account ID is not
your own, then it's the ID of the mesh owner or of another account that the mesh is shared
with. For more information about mesh sharing, see [Working with Shared Meshes](https://docs.aws.amazon.com/app-mesh/latest/userguide/sharing.html).

`RouteName`

The name of the route.

`Uid`

The unique identifier for the route.

`VirtualRouterName`

The name of the virtual router that the route is associated with.

## Examples

### Create a Route

This example creates a route with two weighted targets.

#### JSON

```json

{
   "Description": "Basic Test Route",
   "Resources": {
      "BasicRoute": {
         "Type": "AWS::AppMesh::Route",
         "Properties": {
            "RouteName": "TestRoute",
            "MeshName": null,
            "VirtualRouterName": null,
            "Spec": {
               "HttpRoute": {
                  "Match": {
                     "Prefix": "routePrefix"
                  },
                  "Action": {
                     "WeightedTargets": [
                        {
                           "VirtualNode": null,
                           "Weight": 10
                        },
                        {
                           "VirtualNode": null,
                           "Weight": 20
                        }
                     ]
                  }
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
      "RouteName": {
         "Description": "Name of the Route",
         "Value": {
            "Fn::GetAtt": [
               "BasicRoute",
               "RouteName"
            ]
         }
      },
      "MeshName": {
         "Description": "Name of the Mesh",
         "Value": {
            "Fn::GetAtt": [
               "BasicRoute",
               "MeshName"
            ]
         }
      },
      "VirtualRouterName": {
         "Description": "Name of the VirtualRouter",
         "Value": {
            "Fn::GetAtt": [
               "BasicRoute",
               "VirtualRouterName"
            ]
         }
      },
      "Arn": {
         "Description": "Arn of the Route created",
         "Value": {
            "Fn::GetAtt": [
               "BasicRoute",
               "Arn"
            ]
         }
      },
      "Uid": {
         "Description": "Uid of the Route created",
         "Value": {
            "Fn::GetAtt": [
               "BasicRoute",
               "Uid"
            ]
         }
      }
   }
}
```

#### YAML

```yaml

Description: "Basic Test Route"
Resources:
  BasicRoute:
    Type: "AWS::AppMesh::Route"
    Properties:
      RouteName: "TestRoute"
      MeshName: !ImportValue TestMeshName
      VirtualRouterName: !ImportValue TestVirtualRouterName1
      Spec:
        HttpRoute:
          Match:
            Prefix: "routePrefix"
          Action:
            WeightedTargets:
            - VirtualNode: !ImportValue TestVirtualNodeName1
              Weight: 10
            - VirtualNode: !ImportValue TestVirtualNodeName2
              Weight: 20
      Tags:
      - Key: "Key1"
        Value: "Value1"
      - Key: "Key2"
        Value: "Value2"

Outputs:
  RouteName:
    Description: Name of the Route
    Value:
      Fn::GetAtt:
      - BasicRoute
      - RouteName
  MeshName:
    Description: Name of the Mesh
    Value:
      Fn::GetAtt:
      - BasicRoute
      - MeshName
  VirtualRouterName:
    Description: Name of the VirtualRouter
    Value:
      Fn::GetAtt:
      - BasicRoute
      - VirtualRouterName
  Arn:
    Description: Arn of the Route created
    Value:
      Fn::GetAtt:
      - BasicRoute
      - Arn
  Uid:
    Description: Uid of the Route created
    Value:
      Fn::GetAtt:
      - BasicRoute
      - Uid
```

## See also

- [Routes](https://docs.aws.amazon.com/app-mesh/latest/userguide/routes.html) in the _AWS App Mesh User Guide_.

- [CreateRoute](https://docs.aws.amazon.com/app-mesh/latest/APIReference/API_CreateRoute.html) in
the _AWS App Mesh API Reference_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Duration
