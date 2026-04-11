This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::VirtualService VirtualNodeServiceProvider

An object that represents a virtual node service provider.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "VirtualNodeName" : String
}

```

### YAML

```yaml

  VirtualNodeName: String

```

## Properties

`VirtualNodeName`

The name of the virtual node that is acting as a service provider.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

VirtualRouterServiceProvider

All content copied from https://docs.aws.amazon.com/.
