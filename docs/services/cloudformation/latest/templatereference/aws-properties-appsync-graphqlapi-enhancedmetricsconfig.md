This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppSync::GraphQLApi EnhancedMetricsConfig

Describes an enhanced metrics configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DataSourceLevelMetricsBehavior" : String,
  "OperationLevelMetricsConfig" : String,
  "ResolverLevelMetricsBehavior" : String
}

```

### YAML

```yaml

  DataSourceLevelMetricsBehavior: String
  OperationLevelMetricsConfig: String
  ResolverLevelMetricsBehavior: String

```

## Properties

`DataSourceLevelMetricsBehavior`

Controls how data source metrics will be emitted to CloudWatch. Data source metrics
include:

- **Requests**: The number of invocations that
occured during a request.

- **Latency**: The time to complete a data source
invocation.

- **Errors**: The number of errors that occurred
during a data source invocation.

These metrics can be emitted to CloudWatch per data source or for all data sources in
the request. Metrics will be recorded by API ID and data source name.
`dataSourceLevelMetricsBehavior` accepts one of these values at a
time:

- `FULL_REQUEST_DATA_SOURCE_METRICS`: Records and emits metric data
for all data sources in the request.

- `PER_DATA_SOURCE_METRICS`: Records and emits metric data for data
sources that have the `MetricsConfig` value set to
`ENABLED`.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OperationLevelMetricsConfig`

Controls how operation metrics will be emitted to CloudWatch. Operation metrics
include:

- **Requests**: The number of times a specified
GraphQL operation was called.

- **GraphQL errors**: The number of GraphQL errors
that occurred during a specified GraphQL operation.

Metrics will be recorded by API ID and operation name. You can set the value to
`ENABLED` or `DISABLED`.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResolverLevelMetricsBehavior`

Controls how resolver metrics will be emitted to CloudWatch. Resolver metrics
include:

- **GraphQL errors**: The number of GraphQL errors
that occurred.

- **Requests**: The number of invocations that
occurred during a request.

- **Latency**: The time to complete a resolver
invocation.

- **Cache hits**: The number of cache hits during a
request.

- **Cache misses**: The number of cache misses
during a request.

These metrics can be emitted to CloudWatch per resolver or for all resolvers in the
request. Metrics will be recorded by API ID and resolver name.
`resolverLevelMetricsBehavior` accepts one of these values at a
time:

- `FULL_REQUEST_RESOLVER_METRICS`: Records and emits metric data for
all resolvers in the request.

- `PER_RESOLVER_METRICS`: Records and emits metric data for resolvers
that have the `MetricsConfig` value set to
`ENABLED`.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CognitoUserPoolConfig

LambdaAuthorizerConfig
