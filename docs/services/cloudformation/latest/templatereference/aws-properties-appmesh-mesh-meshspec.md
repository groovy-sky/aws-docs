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

_Type_: [EgressFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appmesh-mesh-egressfilter.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceDiscovery`

An object that represents the service discovery information for a service mesh.

_Required_: No

_Type_: [MeshServiceDiscovery](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appmesh-mesh-meshservicediscovery.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

MeshServiceDiscovery

Tag
