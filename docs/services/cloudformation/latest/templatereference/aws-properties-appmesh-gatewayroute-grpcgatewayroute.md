This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::GatewayRoute GrpcGatewayRoute

An object that represents a gRPC gateway route.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Action" : GrpcGatewayRouteAction,
  "Match" : GrpcGatewayRouteMatch
}

```

### YAML

```yaml

  Action:
    GrpcGatewayRouteAction
  Match:
    GrpcGatewayRouteMatch

```

## Properties

`Action`

An object that represents the action to take if a match is determined.

_Required_: Yes

_Type_: [GrpcGatewayRouteAction](aws-properties-appmesh-gatewayroute-grpcgatewayrouteaction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Match`

An object that represents the criteria for determining a request match.

_Required_: Yes

_Type_: [GrpcGatewayRouteMatch](aws-properties-appmesh-gatewayroute-grpcgatewayroutematch.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GatewayRouteVirtualService

GrpcGatewayRouteAction

All content copied from https://docs.aws.amazon.com/.
