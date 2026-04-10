This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::Route GrpcRetryPolicy

An object that represents a retry policy. Specify at least one value for at least one of the types of `RetryEvents`, a value for `maxRetries`, and a value for `perRetryTimeout`.
Both `server-error` and `gateway-error` under `httpRetryEvents` include the Envoy `reset` policy. For more information on the
`reset` policy, see the [Envoy documentation](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/router_filter).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "GrpcRetryEvents" : [ String, ... ],
  "HttpRetryEvents" : [ String, ... ],
  "MaxRetries" : Integer,
  "PerRetryTimeout" : Duration,
  "TcpRetryEvents" : [ String, ... ]
}

```

### YAML

```yaml

  GrpcRetryEvents:
    - String
  HttpRetryEvents:
    - String
  MaxRetries: Integer
  PerRetryTimeout:
    Duration
  TcpRetryEvents:
    - String

```

## Properties

`GrpcRetryEvents`

Specify at least one of the valid values.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HttpRetryEvents`

Specify at least one of the following values.

- **server-error** – HTTP status codes 500, 501,
502, 503, 504, 505, 506, 507, 508, 510, and 511

- **gateway-error** – HTTP status codes 502,
503, and 504

- **client-error** – HTTP status code 409

- **stream-error** – Retry on refused
stream

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `25`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxRetries`

The maximum number of retry attempts.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PerRetryTimeout`

The timeout for each retry attempt.

_Required_: Yes

_Type_: [Duration](aws-properties-appmesh-route-duration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TcpRetryEvents`

Specify a valid value. The event occurs before any processing of a request has started and is encountered when the upstream is temporarily or permanently unavailable.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Duration

GrpcRoute

All content copied from https://docs.aws.amazon.com/.
