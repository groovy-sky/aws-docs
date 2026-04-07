This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppSync::ApiCache

The `AWS::AppSync::ApiCache` resource represents the input of a
`CreateApiCache` operation.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::AppSync::ApiCache",
  "Properties" : {
      "ApiCachingBehavior" : String,
      "ApiId" : String,
      "AtRestEncryptionEnabled" : Boolean,
      "HealthMetricsConfig" : String,
      "TransitEncryptionEnabled" : Boolean,
      "Ttl" : Number,
      "Type" : String
    }
}

```

### YAML

```yaml

Type: AWS::AppSync::ApiCache
Properties:
  ApiCachingBehavior: String
  ApiId: String
  AtRestEncryptionEnabled: Boolean
  HealthMetricsConfig: String
  TransitEncryptionEnabled: Boolean
  Ttl: Number
  Type: String

```

## Properties

`ApiCachingBehavior`

Caching behavior.

- **FULL\_REQUEST\_CACHING**: All requests from
the same user are cached. Individual resolvers are automatically cached. All API
calls will try to return responses from the cache.

- **PER\_RESOLVER\_CACHING**: Individual resolvers
that you specify are cached.

- **OPERATION\_LEVEL\_CACHING**: Full requests are
cached together and returned without executing resolvers.

_Required_: Yes

_Type_: String

_Allowed values_: `FULL_REQUEST_CACHING | PER_RESOLVER_CACHING | OPERATION_LEVEL_CACHING`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ApiId`

The GraphQL API ID.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`AtRestEncryptionEnabled`

_This parameter has been deprecated_.

At-rest encryption flag for cache. You cannot update this setting after creation.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HealthMetricsConfig`

Controls how cache health metrics will be emitted to CloudWatch. Cache health metrics
include:

- **NetworkBandwidthOutAllowanceExceeded**: The
network packets dropped because the throughput exceeded the aggregated bandwidth
limit. This is useful for diagnosing bottlenecks in a cache
configuration.

- **EngineCPUUtilization**: The CPU utilization
(percentage) allocated to the Redis process. This is useful for diagnosing
bottlenecks in a cache configuration.

Metrics will be recorded by API ID. You can set the value to `ENABLED` or
`DISABLED`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TransitEncryptionEnabled`

_This parameter has been deprecated_.

Transit encryption flag when connecting to cache. You cannot update this setting after creation.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Ttl`

TTL in seconds for cache entries.

Valid values are 1–3,600 seconds.

_Required_: Yes

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

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

_Required_: Yes

_Type_: String

_Allowed values_: `T2_SMALL | T2_MEDIUM | R4_LARGE | R4_XLARGE | R4_2XLARGE | R4_4XLARGE | R4_8XLARGE | SMALL | MEDIUM | LARGE | XLARGE | LARGE_2X | LARGE_4X | LARGE_8X | LARGE_12X`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

### ApiCache Creation Example

The following example creates an ApiCache for your GraphQL API.

#### YAML

```yaml

Parameters:
  graphQlApiId:
    Type: String
Resources:
  ApiCache:
    Type:  AWS::AppSync::ApiCache
    Properties:
      ApiId: !Ref graphQlApiId
      Type: SMALL
      ApiCachingBehavior: FULL_REQUEST_CACHING
      Ttl: 1200
      TransitEncryptionEnabled: true
      AtRestEncryptionEnabled: true
```

#### JSON

```json

{
  "Parameters": {
    "graphQlApiId": {
      "Type": "String"
    }
  },
  "Resources": {
    "ApiCache": {
      "Type": "AWS::AppSync::ApiCache",
      "Properties": {
        "ApiId": { "Ref": "graphQlApiId" },
        "Type": "SMALL",
        "ApiCachingBehavior": "FULL_REQUEST_CACHING",
        "Ttl": 1200,
        "TransitEncryptionEnabled": true,
        "AtRestEncryptionEnabled": true
      }
    }
  }
}
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AWS::AppSync::ApiKey
