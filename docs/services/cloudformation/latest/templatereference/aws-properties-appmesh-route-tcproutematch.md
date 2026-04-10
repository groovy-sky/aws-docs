This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::Route TcpRouteMatch

An object representing the TCP route to match.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Port" : Integer
}

```

### YAML

```yaml

  Port: Integer

```

## Properties

`Port`

The port number to match on.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `65535`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TcpRouteAction

TcpTimeout

All content copied from https://docs.aws.amazon.com/.
