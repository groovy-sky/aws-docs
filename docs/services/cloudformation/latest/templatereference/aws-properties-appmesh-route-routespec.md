This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::Route RouteSpec

An object that represents a route specification. Specify one route type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "GrpcRoute" : GrpcRoute,
  "Http2Route" : HttpRoute,
  "HttpRoute" : HttpRoute,
  "Priority" : Integer,
  "TcpRoute" : TcpRoute
}

```

### YAML

```yaml

  GrpcRoute:
    GrpcRoute
  Http2Route:
    HttpRoute
  HttpRoute:
    HttpRoute
  Priority: Integer
  TcpRoute:
    TcpRoute

```

## Properties

`GrpcRoute`

An object that represents the specification of a gRPC route.

_Required_: No

_Type_: [GrpcRoute](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appmesh-route-grpcroute.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Http2Route`

An object that represents the specification of an HTTP/2 route.

_Required_: No

_Type_: [HttpRoute](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appmesh-route-httproute.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HttpRoute`

An object that represents the specification of an HTTP route.

_Required_: No

_Type_: [HttpRoute](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appmesh-route-httproute.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Priority`

The priority for the route. Routes are matched based on the specified value, where 0 is
the highest priority.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TcpRoute`

An object that represents the specification of a TCP route.

_Required_: No

_Type_: [TcpRoute](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appmesh-route-tcproute.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

QueryParameter

Tag
