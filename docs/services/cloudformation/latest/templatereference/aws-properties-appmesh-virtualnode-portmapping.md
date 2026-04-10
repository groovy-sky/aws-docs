This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::VirtualNode PortMapping

An object representing a virtual node or virtual router listener port mapping.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Port" : Integer,
  "Protocol" : String
}

```

### YAML

```yaml

  Port: Integer
  Protocol: String

```

## Properties

`Port`

The port used for the port mapping.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Maximum_: `65535`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Protocol`

The protocol used for the port mapping. Specify `http`, `http2`,
`grpc`, or `tcp`.

_Required_: Yes

_Type_: String

_Allowed values_: `http | tcp | http2 | grpc`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OutlierDetection

ServiceDiscovery

All content copied from https://docs.aws.amazon.com/.
