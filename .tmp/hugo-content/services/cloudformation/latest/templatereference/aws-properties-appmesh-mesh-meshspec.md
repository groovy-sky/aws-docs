This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::Mesh MeshSpec

An object that represents the specification of a service mesh.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EgressFilter" : EgressFilter,
  "ServiceDiscovery" : MeshServiceDiscovery
}

```

### YAML

```yaml

  EgressFilter:
    EgressFilter
  ServiceDiscovery:
    MeshServiceDiscovery

```

## Properties

`EgressFilter`

The egress filter rules for the service mesh.

_Required_: No

_Type_: [EgressFilter](aws-properties-appmesh-mesh-egressfilter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceDiscovery`

An object that represents the service discovery information for a service mesh.

_Required_: No

_Type_: [MeshServiceDiscovery](aws-properties-appmesh-mesh-meshservicediscovery.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MeshServiceDiscovery

Tag

All content copied from https://docs.aws.amazon.com/.
