This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::VirtualRouter VirtualRouterSpec

An object that represents the specification of a virtual router.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Listeners" : [ VirtualRouterListener, ... ]
}

```

### YAML

```yaml

  Listeners:
    - VirtualRouterListener

```

## Properties

`Listeners`

The listeners that the virtual router is expected to receive inbound traffic
from.

_Required_: Yes

_Type_: Array of [VirtualRouterListener](aws-properties-appmesh-virtualrouter-virtualrouterlistener.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VirtualRouterListener

AWS::AppMesh::VirtualService

All content copied from https://docs.aws.amazon.com/.
