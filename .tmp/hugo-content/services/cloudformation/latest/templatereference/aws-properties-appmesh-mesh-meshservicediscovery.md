This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::Mesh MeshServiceDiscovery

An object that represents the service discovery information for a service mesh.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IpPreference" : String
}

```

### YAML

```yaml

  IpPreference: String

```

## Properties

`IpPreference`

The IP version to use to control traffic within the mesh.

_Required_: No

_Type_: String

_Allowed values_: `IPv6_PREFERRED | IPv4_PREFERRED | IPv4_ONLY | IPv6_ONLY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EgressFilter

MeshSpec

All content copied from https://docs.aws.amazon.com/.
