This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::GatewayRoute GrpcGatewayRouteAction

An object that represents the action to take if a match is determined.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Rewrite" : GrpcGatewayRouteRewrite,
  "Target" : GatewayRouteTarget
}

```

### YAML

```yaml

  Rewrite:
    GrpcGatewayRouteRewrite
  Target:
    GatewayRouteTarget

```

## Properties

`Rewrite`

The gateway route action to rewrite.

_Required_: No

_Type_: [GrpcGatewayRouteRewrite](aws-properties-appmesh-gatewayroute-grpcgatewayrouterewrite.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Target`

An object that represents the target that traffic is routed to when a request matches the gateway route.

_Required_: Yes

_Type_: [GatewayRouteTarget](aws-properties-appmesh-gatewayroute-gatewayroutetarget.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GrpcGatewayRoute

GrpcGatewayRouteMatch

All content copied from https://docs.aws.amazon.com/.
