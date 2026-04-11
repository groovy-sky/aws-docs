This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::VirtualNode VirtualServiceBackend

An object that represents a virtual service backend for a virtual node.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ClientPolicy" : ClientPolicy,
  "VirtualServiceName" : String
}

```

### YAML

```yaml

  ClientPolicy:
    ClientPolicy
  VirtualServiceName: String

```

## Properties

`ClientPolicy`

A reference to an object that represents the client policy for a backend.

_Required_: No

_Type_: [ClientPolicy](aws-properties-appmesh-virtualnode-clientpolicy.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VirtualServiceName`

The name of the virtual service that is acting as a virtual node backend.

###### Important

App Mesh doesn't validate the existence of those virtual services specified in
backends. This is to prevent a cyclic dependency between virtual nodes and virtual
services creation. Make sure the virtual service name is correct. The virtual service
can be created afterwards if it doesn't already exist.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VirtualNodeTcpConnectionPool

AWS::AppMesh::VirtualRouter

All content copied from https://docs.aws.amazon.com/.
