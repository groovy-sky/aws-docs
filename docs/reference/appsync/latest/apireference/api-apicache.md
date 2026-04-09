# ApiCache

The `ApiCache` object.

## Contents

**apiCachingBehavior**

Caching behavior.

- **FULL\_REQUEST\_CACHING**: All requests from the
same user are cached. Individual resolvers are automatically cached. All API calls
will try to return responses from the cache.

- **PER\_RESOLVER\_CACHING**: Individual resolvers
that you specify are cached.

- **OPERATION\_LEVEL\_CACHING**: Full requests are cached together and returned without executing resolvers.

Type: String

Valid Values: `FULL_REQUEST_CACHING | PER_RESOLVER_CACHING | OPERATION_LEVEL_CACHING`

Required: No

**atRestEncryptionEnabled**

At-rest encryption flag for cache. You cannot update this setting after creation.

Type: Boolean

Required: No

**healthMetricsConfig**

Controls how cache health metrics will be emitted to CloudWatch. Cache health metrics
include:

- NetworkBandwidthOutAllowanceExceeded: The network packets dropped because the
throughput exceeded the aggregated bandwidth limit. This is useful for diagnosing
bottlenecks in a cache configuration.

- EngineCPUUtilization: The CPU utilization (percentage) allocated to the Redis
process. This is useful for diagnosing bottlenecks in a cache
configuration.

Metrics will be recorded by API ID. You can set the value to `ENABLED` or
`DISABLED`.

Type: String

Valid Values: `ENABLED | DISABLED`

Required: No

**status**

The cache instance status.

- **AVAILABLE**: The instance is available for
use.

- **CREATING**: The instance is currently
creating.

- **DELETING**: The instance is currently
deleting.

- **MODIFYING**: The instance is currently
modifying.

- **FAILED**: The instance has failed
creation.

Type: String

Valid Values: `AVAILABLE | CREATING | DELETING | MODIFYING | FAILED`

Required: No

**transitEncryptionEnabled**

Transit encryption flag when connecting to cache. You cannot update this setting after
creation.

Type: Boolean

Required: No

**ttl**

TTL in seconds for cache entries.

Valid values are 1–3,600 seconds.

Type: Long

Required: No

**type**

The cache instance type. Valid values are

- `SMALL`

- `MEDIUM`

- `LARGE`

- `XLARGE`

- `LARGE_2X`

- `LARGE_4X`

- `LARGE_8X` (not available in all regions)

- `LARGE_12X`

Historically, instance types were identified by an EC2-style value. As of July 2020, this is deprecated, and the generic identifiers above should be used.

The following legacy instance types are available, but their use is discouraged:

- **T2\_SMALL**: A t2.small instance type.

- **T2\_MEDIUM**: A t2.medium instance type.

- **R4\_LARGE**: A r4.large instance type.

- **R4\_XLARGE**: A r4.xlarge instance type.

- **R4\_2XLARGE**: A r4.2xlarge instance type.

- **R4\_4XLARGE**: A r4.4xlarge instance type.

- **R4\_8XLARGE**: A r4.8xlarge instance type.

Type: String

Valid Values: `T2_SMALL | T2_MEDIUM | R4_LARGE | R4_XLARGE | R4_2XLARGE | R4_4XLARGE | R4_8XLARGE | SMALL | MEDIUM | LARGE | XLARGE | LARGE_2X | LARGE_4X | LARGE_8X | LARGE_12X`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appsync-2017-07-25/apicache.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appsync-2017-07-25/apicache.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appsync-2017-07-25/apicache.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ApiAssociation

ApiKey

All content copied from https://docs.aws.amazon.com/.
