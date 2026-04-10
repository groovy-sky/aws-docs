This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::VirtualService VirtualServiceProvider

An object that represents the provider for a virtual service.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "VirtualNode" : VirtualNodeServiceProvider,
  "VirtualRouter" : VirtualRouterServiceProvider
}

```

### YAML

```yaml

  VirtualNode:
    VirtualNodeServiceProvider
  VirtualRouter:
    VirtualRouterServiceProvider

```

## Properties

`VirtualNode`

The virtual node associated with a virtual service.

_Required_: No

_Type_: [VirtualNodeServiceProvider](aws-properties-appmesh-virtualservice-virtualnodeserviceprovider.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VirtualRouter`

The virtual router associated with a virtual service.

_Required_: No

_Type_: [VirtualRouterServiceProvider](aws-properties-appmesh-virtualservice-virtualrouterserviceprovider.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VirtualRouterServiceProvider

VirtualServiceSpec

All content copied from https://docs.aws.amazon.com/.
