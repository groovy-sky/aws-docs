This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::VirtualNode

Creates a virtual node within a service mesh.

A virtual node acts as a logical pointer to a particular task group, such as an Amazon ECS service or a Kubernetes deployment. When you create a virtual node, you can
specify the service discovery information for your task group, and whether the proxy
running in a task group will communicate with other proxies using Transport Layer Security
(TLS).

You define a `listener` for any inbound traffic that your virtual node
expects. Any virtual service that your virtual node expects to communicate to is specified
as a `backend`.

The response metadata for your new virtual node contains the `arn` that is
associated with the virtual node. Set this value to the full ARN; for example,
`arn:aws:appmesh:us-west-2:123456789012:myMesh/default/virtualNode/myApp`)
as the `APPMESH_RESOURCE_ARN` environment variable for your task group's Envoy
proxy container in your task definition or pod spec. This is then mapped to the
`node.id` and `node.cluster` Envoy parameters.

###### Note

By default, App Mesh uses the name of the resource you specified in
`APPMESH_RESOURCE_ARN` when Envoy is referring to itself in metrics and
traces. You can override this behavior by setting the
`APPMESH_RESOURCE_CLUSTER` environment variable with your own name.

For more information about virtual nodes, see [Virtual nodes](https://docs.aws.amazon.com/app-mesh/latest/userguide/virtual_nodes.html). You must be using `1.15.0` or later of the Envoy image when
setting these variables. For more information aboutApp Mesh Envoy variables, see
[Envoy\
image](https://docs.aws.amazon.com/app-mesh/latest/userguide/envoy.html) in the AWS App Mesh User Guide.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::AppMesh::VirtualNode",
  "Properties" : {
      "MeshName" : String,
      "MeshOwner" : String,
      "Spec" : VirtualNodeSpec,
      "Tags" : [ Tag, ... ],
      "VirtualNodeName" : String
    }
}

```

### YAML

```yaml

Type: AWS::AppMesh::VirtualNode
Properties:
  MeshName: String
  MeshOwner: String
  Spec:
    VirtualNodeSpec
  Tags:
    - Tag
  VirtualNodeName: String

```

## Properties

`MeshName`

The name of the service mesh to create the virtual node in.

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

The virtual node specification to apply.

_Required_: Yes

_Type_: [VirtualNodeSpec](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appmesh-virtualnode-virtualnodespec.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Optional metadata that you can apply to the virtual node to assist with categorization
and organization. Each tag consists of a key and an optional value, both of which you
define. Tag keys can have a maximum character length of 128 characters, and tag values can have
a maximum length of 256 characters.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appmesh-virtualnode-tag.html)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VirtualNodeName`

The name to use for the virtual node.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource ARN. For example:

`{ "Ref": "myVirtualNode" }`

When you pass the logical ID of an `AWS::AppMesh::VirtualNode` resource to
the intrinsic Ref function, the function returns the virtual node ARN, such as
`arn:aws:appmesh:us-east-1:555555555555:virtualNode/myVirtualNode`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The full Amazon Resource Name (ARN) for the virtual node.

`MeshName`

The name of the service mesh that the virtual node resides in.

`MeshOwner`

The AWS IAM account ID of the service mesh owner. If the account ID is
not your own, then it's the ID of the account that shared the mesh with your account. For
more information about mesh sharing, see [Working with Shared Meshes](https://docs.aws.amazon.com/app-mesh/latest/userguide/sharing.html).

`ResourceOwner`

The AWS IAM account ID of the resource owner. If the account ID is not
your own, then it's the ID of the mesh owner or of another account that the mesh is shared
with. For more information about mesh sharing, see [Working with Shared Meshes](https://docs.aws.amazon.com/app-mesh/latest/userguide/sharing.html).

`Uid`

The unique identifier for the virtual node.

`VirtualNodeName`

The name of the virtual node.

## Examples

### Create a Virtual Node

This example creates a virtual node with two backends and a listener with a health
check policy. It also sends access logs to a file path and uses DNS service
discovery.

#### JSON

```json

{
   "Description": "Basic Test Virtual Node",
   "Resources": {
      "BasicVirtualNode": {
         "Type": "AWS::AppMesh::VirtualNode",
         "Properties": {
            "VirtualNodeName": "TestVirtualNode",
            "MeshName": null,
            "Spec": {
               "Backends": [
                  {
                     "VirtualService": {
                        "VirtualServiceName": "Backend_1"
                     }
                  },
                  {
                     "VirtualService": {
                        "VirtualServiceName": "Backend_2"
                     }
                  }
               ],
               "Listeners": [
                  {
                     "HealthCheck": {
                        "HealthyThreshold": 2,
                        "IntervalMillis": 5000,
                        "Path": "Path",
                        "Port": 8080,
                        "Protocol": "http",
                        "TimeoutMillis": 2000,
                        "UnhealthyThreshold": 2
                     },
                     "PortMapping": {
                        "Port": 8080,
                        "Protocol": "http"
                     }
                  }
               ],
               "ServiceDiscovery": {
                  "DNS": {
                     "Hostname": "Hostname"
                  }
               },
               "Logging": {
                  "AccessLog": {
                     "File": {
                        "Path": "Path"
                     }
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
      "VirtualNodeName": {
         "Description": "Name of the VirtualNode",
         "Value": {
            "Fn::GetAtt": [
               "BasicVirtualNode",
               "VirtualNodeName"
            ]
         }
      },
      "MeshName": {
         "Description": "Name of the Mesh",
         "Value": {
            "Fn::GetAtt": [
               "BasicVirtualNode",
               "MeshName"
            ]
         }
      },
      "Arn": {
         "Description": "Arn of the VirtualNode created",
         "Value": {
            "Fn::GetAtt": [
               "BasicVirtualNode",
               "Arn"
            ]
         }
      },
      "Uid": {
         "Description": "Uid of the VirtualNode created",
         "Value": {
            "Fn::GetAtt": [
               "BasicVirtualNode",
               "Uid"
            ]
         }
      }
   }
}
```

#### YAML

```yaml

Description: "Basic Test Virtual Node"
Resources:
  BasicVirtualNode:
    Type: "AWS::AppMesh::VirtualNode"
    Properties:
      VirtualNodeName: "TestVirtualNode"
      MeshName: !ImportValue TestMeshName
      Spec:
        Backends:
        - VirtualService:
            VirtualServiceName: "Backend_1"
        - VirtualService:
            VirtualServiceName: "Backend_2"
        Listeners:
        - HealthCheck:
            HealthyThreshold: 2
            IntervalMillis: 5000
            Path: "Path"
            Port: 8080
            Protocol: "http"
            TimeoutMillis: 2000
            UnhealthyThreshold: 2
          PortMapping:
            Port: 8080
            Protocol: "http"
        ServiceDiscovery:
          DNS:
            Hostname: "Hostname"
        Logging:
          AccessLog:
            File:
              Path: "Path"
      Tags:
      - Key: "Key1"
        Value: "Value1"
      - Key: "Key2"
        Value: "Value2"

Outputs:
  VirtualNodeName:
    Description: Name of the VirtualNode
    Value:
      Fn::GetAtt:
      - BasicVirtualNode
      - VirtualNodeName
  MeshName:
    Description: Name of the Mesh
    Value:
      Fn::GetAtt:
      - BasicVirtualNode
      - MeshName
  Arn:
    Description: Arn of the VirtualNode created
    Value:
      Fn::GetAtt:
      - BasicVirtualNode
      - Arn
  Uid:
    Description: Uid of the VirtualNode created
    Value:
      Fn::GetAtt:
      - BasicVirtualNode
      - Uid
```

## See also

- [Virtual Nodes](https://docs.aws.amazon.com/app-mesh/latest/userguide/virtual_nodes.html) in the _AWS App Mesh User Guide_.

- [CreateVirtualNode](https://docs.aws.amazon.com/app-mesh/latest/APIReference/API_CreateVirtualNode.html) in the _AWS App Mesh API Reference_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

VirtualGatewayTlsValidationContextTrust

AccessLog
