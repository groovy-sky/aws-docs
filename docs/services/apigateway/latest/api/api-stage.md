# Stage

Represents a unique identifier for a version of a deployed RestApi that is callable by users.

## Contents

**accessLogSettings**

Settings for logging access in this stage.

Type: [AccessLogSettings](api-accesslogsettings.md) object

Required: No

**cacheClusterEnabled**

Specifies whether a cache cluster is enabled for the stage. To activate a method-level cache, set `CachingEnabled` to `true` for a method.

Type: Boolean

Required: No

**cacheClusterSize**

The stage's cache capacity in GB. For more information about choosing a cache size, see [Enabling API caching to enhance responsiveness](../developerguide/api-gateway-caching.md).

Type: String

Valid Values: `0.5 | 1.6 | 6.1 | 13.5 | 28.4 | 58.2 | 118 | 237`

Required: No

**cacheClusterStatus**

The status of the cache cluster for the stage, if enabled.

Type: String

Valid Values: `CREATE_IN_PROGRESS | AVAILABLE | DELETE_IN_PROGRESS | NOT_AVAILABLE | FLUSH_IN_PROGRESS`

Required: No

**canarySettings**

Settings for the canary deployment in this stage.

Type: [CanarySettings](api-canarysettings.md) object

Required: No

**clientCertificateId**

The identifier of a client certificate for an API stage.

Type: String

Required: No

**createdDate**

The timestamp when the stage was created.

Type: Timestamp

Required: No

**deploymentId**

The identifier of the Deployment that the stage points to.

Type: String

Required: No

**description**

The stage's description.

Type: String

Required: No

**documentationVersion**

The version of the associated API documentation.

Type: String

Required: No

**lastUpdatedDate**

The timestamp when the stage last updated.

Type: Timestamp

Required: No

**methodSettings**

A map that defines the method settings for a Stage resource. Keys (designated as `/{method_setting_key` below) are method paths defined as `{resource_path}/{http_method}` for an individual method override, or `/\*/\*` for overriding all methods in the stage.

Type: String to [MethodSetting](api-methodsetting.md) object map

Required: No

**stageName**

The name of the stage is the first path segment in the Uniform Resource Identifier (URI) of a call to API Gateway. Stage names can only contain alphanumeric characters, hyphens, and underscores. Maximum length is 128 characters.

Type: String

Required: No

**tags**

The collection of tags. Each tag element is associated with a given resource.

Type: String to string map

Required: No

**tracingEnabled**

Specifies whether active tracing with X-ray is enabled for the Stage.

Type: Boolean

Required: No

**variables**

A map that defines the stage variables for a Stage resource. Variable names can
have alphanumeric and underscore characters, and the values must match `[A-Za-z0-9-._~:/?#&=,]+`.

Type: String to string map

Required: No

**webAclArn**

The ARN of the WebAcl associated with the Stage.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apigateway-2015-07-09/stage.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apigateway-2015-07-09/stage.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apigateway-2015-07-09/stage.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SdkType

StageKey

All content copied from https://docs.aws.amazon.com/.
