# Resolver

Describes a resolver.

## Contents

**cachingConfig**

The caching configuration for the resolver.

Type: [CachingConfig](api-cachingconfig.md) object

Required: No

**code**

The `resolver` code that contains the request and response functions. When
code is used, the `runtime` is required. The `runtime` value must be
`APPSYNC_JS`.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 32768.

Required: No

**dataSourceName**

The resolver data source name.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 65536.

Pattern: `[_A-Za-z][_0-9A-Za-z]*`

Required: No

**fieldName**

The resolver field name.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 65536.

Pattern: `[_A-Za-z][_0-9A-Za-z]*`

Required: No

**kind**

The resolver type.

- **UNIT**: A UNIT resolver type. A UNIT resolver is
the default resolver type. You can use a UNIT resolver to run a GraphQL query against
a single data source.

- **PIPELINE**: A PIPELINE resolver type. You can
use a PIPELINE resolver to invoke a series of `Function` objects in a
serial manner. You can use a pipeline resolver to run a GraphQL query against
multiple data sources.

Type: String

Valid Values: `UNIT | PIPELINE`

Required: No

**maxBatchSize**

The maximum batching size for a resolver.

Type: Integer

Valid Range: Minimum value of 0. Maximum value of 2000.

Required: No

**metricsConfig**

Enables or disables enhanced resolver metrics for specified resolvers. Note that
`metricsConfig` won't be used unless the
`resolverLevelMetricsBehavior` value is set to
`PER_RESOLVER_METRICS`. If the `resolverLevelMetricsBehavior` is
set to `FULL_REQUEST_RESOLVER_METRICS` instead, `metricsConfig` will
be ignored. However, you can still set its value.

`metricsConfig` can be `ENABLED` or `DISABLED`.

Type: String

Valid Values: `ENABLED | DISABLED`

Required: No

**pipelineConfig**

The `PipelineConfig`.

Type: [PipelineConfig](api-pipelineconfig.md) object

Required: No

**requestMappingTemplate**

The request mapping template.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 65536.

Pattern: `^.*$`

Required: No

**resolverArn**

The resolver Amazon Resource Name (ARN).

Type: String

Required: No

**responseMappingTemplate**

The response mapping template.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 65536.

Pattern: `^.*$`

Required: No

**runtime**

Describes a runtime used by an AWS AppSync pipeline resolver or AWS AppSync function. Specifies the name and version of the runtime to use. Note
that if a runtime is specified, code must also be specified.

Type: [AppSyncRuntime](api-appsyncruntime.md) object

Required: No

**syncConfig**

The `SyncConfig` for a resolver attached to a versioned data source.

Type: [SyncConfig](api-syncconfig.md) object

Required: No

**typeName**

The resolver type name.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 65536.

Pattern: `[_A-Za-z][_0-9A-Za-z]*`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appsync-2017-07-25/resolver.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appsync-2017-07-25/resolver.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appsync-2017-07-25/resolver.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RelationalDatabaseDataSourceConfig

SourceApiAssociation

All content copied from https://docs.aws.amazon.com/.
