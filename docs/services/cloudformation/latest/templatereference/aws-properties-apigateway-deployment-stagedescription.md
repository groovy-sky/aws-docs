This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApiGateway::Deployment StageDescription

`StageDescription` is a property of the [AWS::ApiGateway::Deployment](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-deployment.html) resource that configures a deployment stage.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AccessLogSetting" : AccessLogSetting,
  "CacheClusterEnabled" : Boolean,
  "CacheClusterSize" : String,
  "CacheDataEncrypted" : Boolean,
  "CacheTtlInSeconds" : Integer,
  "CachingEnabled" : Boolean,
  "CanarySetting" : CanarySetting,
  "ClientCertificateId" : String,
  "DataTraceEnabled" : Boolean,
  "Description" : String,
  "DocumentationVersion" : String,
  "LoggingLevel" : String,
  "MethodSettings" : [ MethodSetting, ... ],
  "MetricsEnabled" : Boolean,
  "Tags" : [ Tag, ... ],
  "ThrottlingBurstLimit" : Integer,
  "ThrottlingRateLimit" : Number,
  "TracingEnabled" : Boolean,
  "Variables" : {Key: Value, ...}
}

```

### YAML

```yaml

  AccessLogSetting:
    AccessLogSetting
  CacheClusterEnabled: Boolean
  CacheClusterSize: String
  CacheDataEncrypted: Boolean
  CacheTtlInSeconds: Integer
  CachingEnabled: Boolean
  CanarySetting:
    CanarySetting
  ClientCertificateId: String
  DataTraceEnabled: Boolean
  Description: String
  DocumentationVersion: String
  LoggingLevel: String
  MethodSettings:
    - MethodSetting
  MetricsEnabled: Boolean
  Tags:
    - Tag
  ThrottlingBurstLimit: Integer
  ThrottlingRateLimit: Number
  TracingEnabled: Boolean
  Variables:
    Key: Value

```

## Properties

`AccessLogSetting`

Specifies settings for logging access in this stage.

_Required_: No

_Type_: [AccessLogSetting](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-apigateway-deployment-accesslogsetting.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CacheClusterEnabled`

Specifies whether a cache cluster is enabled for the stage. To activate a method-level cache, set `CachingEnabled` to `true` for a method.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CacheClusterSize`

The size of the stage's cache cluster. For more information, see [cacheClusterSize](../../../apigateway/latest/api/api-createstage.md#apigw-CreateStage-request-cacheClusterSize) in the _API Gateway API Reference_.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CacheDataEncrypted`

Indicates whether the cached responses are encrypted.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CacheTtlInSeconds`

The time-to-live (TTL) period, in seconds, that specifies how long API Gateway caches responses.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CachingEnabled`

Indicates whether responses are cached and returned for requests. You must enable a cache cluster on the stage to cache responses. For more information, see [Enable API Gateway Caching in a Stage to Enhance API Performance](../../../apigateway/latest/developerguide/api-gateway-caching.md) in the _API Gateway Developer Guide_.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CanarySetting`

Specifies settings for the canary deployment in this stage.

_Required_: No

_Type_: [CanarySetting](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-apigateway-deployment-canarysetting.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClientCertificateId`

The identifier of the client certificate that API Gateway uses to call your integration endpoints in the stage.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataTraceEnabled`

Indicates whether data trace logging is enabled for methods in the stage. API Gateway pushes these logs to Amazon CloudWatch Logs.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description of the purpose of the stage.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DocumentationVersion`

The version identifier of the API documentation snapshot.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LoggingLevel`

The logging level for this method. For valid values, see the `loggingLevel` property of the [MethodSetting](https://docs.aws.amazon.com/apigateway/latest/api/API_MethodSetting.html) resource in the _Amazon API Gateway API Reference_.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MethodSettings`

Configures settings for all of the stage's methods.

_Required_: No

_Type_: Array of [MethodSetting](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-apigateway-deployment-methodsetting.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MetricsEnabled`

Indicates whether Amazon CloudWatch metrics are enabled for methods in the stage.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of arbitrary tags (key-value pairs) to associate with the stage.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-apigateway-deployment-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ThrottlingBurstLimit`

The target request burst rate limit. This allows more requests through for a period of time than the target rate limit. For more information, see [Manage API Request Throttling](../../../apigateway/latest/developerguide/api-gateway-request-throttling.md) in the _API Gateway Developer Guide_.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ThrottlingRateLimit`

The target request steady-state rate limit. For more information, see [Manage API Request Throttling](../../../apigateway/latest/developerguide/api-gateway-request-throttling.md) in the _API Gateway Developer Guide_.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TracingEnabled`

Specifies whether active tracing with X-ray is enabled for this stage.

For more information, see [Trace API Gateway API Execution with AWS X-Ray](../../../apigateway/latest/developerguide/apigateway-xray.md) in the _API Gateway Developer Guide_.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Variables`

A map that defines the stage variables. Variable names must consist of alphanumeric characters, and the values must match the following regular expression: `[A-Za-z0-9-._~:/?#&=,]+`.

_Required_: No

_Type_: Object of String

_Pattern_: `[a-zA-Z0-9]+`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [Stage](../../../apigateway/latest/api/api-stage.md) in the _Amazon API Gateway REST API Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

MethodSetting

Tag
