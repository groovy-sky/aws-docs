This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::Route TcpRoute

An object that represents a TCP route type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Action" : TcpRouteAction,
  "Match" : TcpRouteMatch,
  "Timeout" : TcpTimeout
}

```

### YAML

```yaml

  Action:
    TcpRouteAction
  Match:
    TcpRouteMatch
  Timeout:
    TcpTimeout

```

## Properties

`Action`

The action to take if a match is determined.

_Required_: Yes

_Type_: [TcpRouteAction](aws-properties-appmesh-route-tcprouteaction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Match`

An object that represents the criteria for determining a request match.

_Required_: No

_Type_: [TcpRouteMatch](aws-properties-appmesh-route-tcproutematch.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Timeout`

An object that represents types of timeouts.

_Required_: No

_Type_: [TcpTimeout](aws-properties-appmesh-route-tcptimeout.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

TcpRouteAction

All content copied from https://docs.aws.amazon.com/.
