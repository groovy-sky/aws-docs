This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::VirtualNode VirtualNodeHttpConnectionPool

An object that represents a type of connection pool.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaxConnections" : Integer,
  "MaxPendingRequests" : Integer
}

```

### YAML

```yaml

  MaxConnections: Integer
  MaxPendingRequests: Integer

```

## Properties

`MaxConnections`

Maximum number of outbound TCP connections Envoy can establish concurrently with all
hosts in upstream cluster.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxPendingRequests`

Number of overflowing requests after `max_connections` Envoy will queue to
upstream cluster.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VirtualNodeHttp2ConnectionPool

VirtualNodeSpec

All content copied from https://docs.aws.amazon.com/.
