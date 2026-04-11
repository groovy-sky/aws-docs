This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::Route HttpRouteMatch

An object that represents the requirements for a route to match HTTP requests for a
virtual router.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Headers" : [ HttpRouteHeader, ... ],
  "Method" : String,
  "Path" : HttpPathMatch,
  "Port" : Integer,
  "Prefix" : String,
  "QueryParameters" : [ QueryParameter, ... ],
  "Scheme" : String
}

```

### YAML

```yaml

  Headers:
    - HttpRouteHeader
  Method: String
  Path:
    HttpPathMatch
  Port: Integer
  Prefix: String
  QueryParameters:
    - QueryParameter
  Scheme: String

```

## Properties

`Headers`

The client request headers to match on.

_Required_: No

_Type_: Array of [HttpRouteHeader](aws-properties-appmesh-route-httprouteheader.md)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Method`

The client request method to match on. Specify only one.

_Required_: No

_Type_: String

_Allowed values_: `GET | HEAD | POST | PUT | DELETE | CONNECT | OPTIONS | TRACE | PATCH`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Path`

The client request path to match on.

_Required_: No

_Type_: [HttpPathMatch](aws-properties-appmesh-route-httppathmatch.md)

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

The client request query parameters to match on.

_Required_: No

_Type_: Array of [QueryParameter](aws-properties-appmesh-route-queryparameter.md)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Scheme`

The client request scheme to match on. Specify only one. Applicable only for HTTP2
routes.

_Required_: No

_Type_: String

_Allowed values_: `http | https`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HttpRouteHeader

HttpTimeout

All content copied from https://docs.aws.amazon.com/.
