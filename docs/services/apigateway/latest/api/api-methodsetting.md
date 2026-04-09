# MethodSetting

Specifies the method setting properties.

## Contents

**cacheDataEncrypted**

Specifies whether the cached responses are encrypted.

Type: Boolean

Required: No

**cacheTtlInSeconds**

Specifies the time to live (TTL), in seconds, for cached responses. The higher the TTL, the longer the response will be cached.

Type: Integer

Required: No

**cachingEnabled**

Specifies whether responses should be cached and returned for requests. A cache cluster must be enabled on the stage for responses to be cached.

Type: Boolean

Required: No

**dataTraceEnabled**

Specifies whether data trace logging is enabled for this method, which affects the log entries pushed to Amazon CloudWatch Logs. This can be useful to troubleshoot APIs, but can result in logging sensitive data. We recommend that you don't enable this option for production APIs.

Type: Boolean

Required: No

**loggingLevel**

Specifies the logging level for this method, which affects the log entries pushed to Amazon CloudWatch Logs. Valid values are `OFF`, `ERROR`, and `INFO`. Choose `ERROR` to write only error-level entries to CloudWatch Logs, or choose `INFO` to include all `ERROR` events as well as extra informational events.

Type: String

Required: No

**metricsEnabled**

Specifies whether Amazon CloudWatch metrics are enabled for this method.

Type: Boolean

Required: No

**requireAuthorizationForCacheControl**

Specifies whether authorization is required for a cache invalidation request.

Type: Boolean

Required: No

**throttlingBurstLimit**

Specifies the throttling burst limit.

Type: Integer

Required: No

**throttlingRateLimit**

Specifies the throttling rate limit.

Type: Double

Required: No

**unauthorizedCacheControlHeaderStrategy**

Specifies how to handle unauthorized requests for cache invalidation.

Type: String

Valid Values: `FAIL_WITH_403 | SUCCEED_WITH_RESPONSE_HEADER | SUCCEED_WITHOUT_RESPONSE_HEADER`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apigateway-2015-07-09/methodsetting.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apigateway-2015-07-09/methodsetting.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apigateway-2015-07-09/methodsetting.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MethodResponse

MethodSnapshot

All content copied from https://docs.aws.amazon.com/.
