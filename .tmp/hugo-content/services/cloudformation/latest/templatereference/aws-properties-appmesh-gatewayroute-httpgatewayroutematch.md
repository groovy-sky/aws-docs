This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::GatewayRoute HttpGatewayRouteMatch

An object that represents the criteria for determining a request match.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Headers" : [ HttpGatewayRouteHeader, ... ],
  "Hostname" : GatewayRouteHostnameMatch,
  "Method" : String,
  "Path" : HttpPathMatch,
  "Port" : Integer,
  "Prefix" : String,
  "QueryParameters" : [ QueryParameter, ... ]
}

```

### YAML

```yaml

  Headers:
    - HttpGatewayRouteHeader
  Hostname:
    GatewayRouteHostnameMatch
  Method: String
  Path:
    HttpPathMatch
  Port: Integer
  Prefix: String
  QueryParameters:
    - QueryParameter

```

## Properties

`Headers`

The client request headers to match on.

_Required_: No

_Type_: Array of [HttpGatewayRouteHeader](aws-properties-appmesh-gatewayroute-httpgatewayrouteheader.md)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Hostname`

The host name to match on.

_Required_: No

_Type_: [GatewayRouteHostnameMatch](aws-properties-appmesh-gatewayroute-gatewayroutehostnamematch.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Method`

The method to match on.

_Required_: No

_Type_: String

_Allowed values_: `GET | HEAD | POST | PUT | DELETE | CONNECT | OPTIONS | TRACE | PATCH`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Path`

The path to match on.

_Required_: No

_Type_: [HttpPathMatch](aws-properties-appmesh-gatewayroute-httppathmatch.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Port`

The port number to match on.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `65535`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Prefix`

Specifies the path to match requests with. This parameter must always start with
`/`, which by itself matches all requests to the virtual service name. You
can also match for path-based routing of requests. For example, if your virtual service
name is `my-service.local` and you want the route to match requests to
`my-service.local/metrics`, your prefix should be
`/metrics`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`QueryParameters`

The query parameter to match on.

_Required_: No

_Type_: Array of [QueryParameter](aws-properties-appmesh-gatewayroute-queryparameter.md)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HttpGatewayRouteHeaderMatch

HttpGatewayRoutePathRewrite

All content copied from https://docs.aws.amazon.com/.
