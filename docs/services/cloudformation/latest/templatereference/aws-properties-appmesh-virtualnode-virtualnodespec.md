This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::VirtualNode VirtualNodeSpec

An object that represents the specification of a virtual node.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BackendDefaults" : BackendDefaults,
  "Backends" : [ Backend, ... ],
  "Listeners" : [ Listener, ... ],
  "Logging" : Logging,
  "ServiceDiscovery" : ServiceDiscovery
}

```

### YAML

```yaml

  BackendDefaults:
    BackendDefaults
  Backends:
    - Backend
  Listeners:
    - Listener
  Logging:
    Logging
  ServiceDiscovery:
    ServiceDiscovery

```

## Properties

`BackendDefaults`

A reference to an object that represents the defaults for backends.

_Required_: No

_Type_: [BackendDefaults](aws-properties-appmesh-virtualnode-backenddefaults.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Backends`

The backends that the virtual node is expected to send outbound traffic to.

###### Important

App Mesh doesn't validate the existence of those virtual services specified in
backends. This is to prevent a cyclic dependency between virtual nodes and virtual
services creation. Make sure the virtual service name is correct. The virtual service
can be created afterwards if it doesn't already exist.

_Required_: No

_Type_: Array of [Backend](aws-properties-appmesh-virtualnode-backend.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Listeners`

The listener that the virtual node is expected to receive inbound traffic from. You can
specify one listener.

_Required_: No

_Type_: Array of [Listener](aws-properties-appmesh-virtualnode-listener.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Logging`

The inbound and outbound access logging information for the virtual node.

_Required_: No

_Type_: [Logging](aws-properties-appmesh-virtualnode-logging.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceDiscovery`

The service discovery information for the virtual node. If your virtual node does not
expect ingress traffic, you can omit this parameter. If you specify a
`listener`, then you must specify service discovery information.

_Required_: No

_Type_: [ServiceDiscovery](aws-properties-appmesh-virtualnode-servicediscovery.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VirtualNodeHttpConnectionPool

VirtualNodeTcpConnectionPool

All content copied from https://docs.aws.amazon.com/.
