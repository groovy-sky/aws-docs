This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::GatewayRoute GatewayRouteSpec

An object that represents a gateway route specification. Specify one gateway route
type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "GrpcRoute" : GrpcGatewayRoute,
  "Http2Route" : HttpGatewayRoute,
  "HttpRoute" : HttpGatewayRoute,
  "Priority" : Integer
}

```

### YAML

```yaml

  GrpcRoute:
    GrpcGatewayRoute
  Http2Route:
    HttpGatewayRoute
  HttpRoute:
    HttpGatewayRoute
  Priority: Integer

```

## Properties

`GrpcRoute`

An object that represents the specification of a gRPC gateway route.

_Required_: No

_Type_: [GrpcGatewayRoute](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appmesh-gatewayroute-grpcgatewayroute.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Http2Route`

An object that represents the specification of an HTTP/2 gateway route.

_Required_: No

_Type_: [HttpGatewayRoute](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appmesh-gatewayroute-httpgatewayroute.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HttpRoute`

An object that represents the specification of an HTTP gateway route.

_Required_: No

_Type_: [HttpGatewayRoute](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appmesh-gatewayroute-httpgatewayroute.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Priority`

The ordering of the gateway routes spec.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

GatewayRouteRangeMatch

GatewayRouteTarget
