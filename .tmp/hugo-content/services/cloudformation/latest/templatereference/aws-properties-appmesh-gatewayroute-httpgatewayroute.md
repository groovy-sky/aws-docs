This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::GatewayRoute HttpGatewayRoute

An object that represents an HTTP gateway route.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Action" : HttpGatewayRouteAction,
  "Match" : HttpGatewayRouteMatch
}

```

### YAML

```yaml

  Action:
    HttpGatewayRouteAction
  Match:
    HttpGatewayRouteMatch

```

## Properties

`Action`

An object that represents the action to take if a match is determined.

_Required_: Yes

_Type_: [HttpGatewayRouteAction](aws-properties-appmesh-gatewayroute-httpgatewayrouteaction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Match`

An object that represents the criteria for determining a request match.

_Required_: Yes

_Type_: [HttpGatewayRouteMatch](aws-properties-appmesh-gatewayroute-httpgatewayroutematch.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GrpcGatewayRouteRewrite

HttpGatewayRouteAction

All content copied from https://docs.aws.amazon.com/.
