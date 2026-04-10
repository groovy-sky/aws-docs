This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::VirtualNode VirtualNodeGrpcConnectionPool

An object that represents a type of connection pool.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaxRequests" : Integer
}

```

### YAML

```yaml

  MaxRequests: Integer

```

## Properties

`MaxRequests`

Maximum number of inflight requests Envoy can concurrently support across hosts in
upstream cluster.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VirtualNodeConnectionPool

VirtualNodeHttp2ConnectionPool

All content copied from https://docs.aws.amazon.com/.
