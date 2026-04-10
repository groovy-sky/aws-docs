This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::Mesh EgressFilter

An object that represents the egress filter rules for a service mesh.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : String
}

```

### YAML

```yaml

  Type: String

```

## Properties

`Type`

The egress filter type. By default, the type is `DROP_ALL`, which allows
egress only from virtual nodes to other defined resources in the service mesh (and any
traffic to `*.amazonaws.com` for AWS API calls). You can set the
egress filter type to `ALLOW_ALL` to allow egress to any endpoint inside or
outside of the service mesh.

###### Note

If you specify any backends on a virtual node when using `ALLOW_ALL`, you
must specifiy all egress for that virtual node as backends. Otherwise,
`ALLOW_ALL` will no longer work for that virtual node.

_Required_: Yes

_Type_: String

_Allowed values_: `ALLOW_ALL | DROP_ALL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::AppMesh::Mesh

MeshServiceDiscovery

All content copied from https://docs.aws.amazon.com/.
