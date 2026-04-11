This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApiGateway::Stage MethodSetting

The `MethodSetting` property type configures settings for all methods in a stage.

The `MethodSettings` property of the `AWS::ApiGateway::Stage` resource contains a list of `MethodSetting` property types.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CacheDataEncrypted" : Boolean,
  "CacheTtlInSeconds" : Integer,
  "CachingEnabled" : Boolean,
  "DataTraceEnabled" : Boolean,
  "HttpMethod" : String,
  "LoggingLevel" : String,
  "MetricsEnabled" : Boolean,
  "ResourcePath" : String,
  "ThrottlingBurstLimit" : Integer,
  "ThrottlingRateLimit" : Number
}

```

### YAML

```yaml

  CacheDataEncrypted: Boolean
  CacheTtlInSeconds: Integer
  CachingEnabled: Boolean
  DataTraceEnabled: Boolean
  HttpMethod: String
  LoggingLevel: String
  MetricsEnabled: Boolean
  ResourcePath: String
  ThrottlingBurstLimit: Integer
  ThrottlingRateLimit: Number

```

## Properties

`CacheDataEncrypted`

Specifies whether the cached responses are encrypted.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CacheTtlInSeconds`

Specifies the time to live (TTL), in seconds, for cached responses. The higher the TTL, the longer the response will be cached.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CachingEnabled`

Specifies whether responses should be cached and returned for requests. A cache cluster must be enabled on the stage for responses to be cached.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataTraceEnabled`

Specifies whether data trace logging is enabled for this method, which affects the log entries pushed to Amazon CloudWatch Logs. This can be useful to troubleshoot APIs, but can result in logging sensitive data. We recommend that you don't enable this option for production APIs.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HttpMethod`

The HTTP method. To apply settings to multiple resources and methods, specify an asterisk ( `*`) for the `HttpMethod` and `/*` for the `ResourcePath`. This parameter is required when you specify a `MethodSetting`.

_Required_: Conditional

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LoggingLevel`

Specifies the logging level for this method, which affects the log entries pushed to Amazon CloudWatch Logs. Valid values are `OFF`, `ERROR`, and `INFO`. Choose `ERROR` to write only error-level entries to CloudWatch Logs, or choose `INFO` to include all `ERROR` events as well as extra informational events.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MetricsEnabled`

Specifies whether Amazon CloudWatch metrics are enabled for this method.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourcePath`

The resource path for this method. Forward slashes ( `/`) are encoded as `~1` and the initial slash must include a forward slash. For example, the path value `/resource/subresource` must be encoded as `/~1resource~1subresource`. To specify the root path, use only a slash ( `/`). To apply settings to multiple resources and methods, specify an asterisk ( `*`) for the `HttpMethod` and `/*` for the `ResourcePath`.
This parameter is required when you specify a `MethodSetting`.

_Required_: Conditional

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ThrottlingBurstLimit`

Specifies the throttling burst limit.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ThrottlingRateLimit`

Specifies the throttling rate limit.

_Required_: No

_Type_: Number

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [Stage](../../../apigateway/latest/api/api-stage.md) in the _Amazon API Gateway REST API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CanarySetting

Tag

All content copied from https://docs.aws.amazon.com/.
