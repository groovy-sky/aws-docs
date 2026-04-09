# UpdateStage

Changes information about a Stage resource.

## Request Syntax

```nohighlight

PATCH /restapis/restapi_id/stages/stage_name HTTP/1.1
Content-type: application/json

{
   "patchOperations": [
      {
         "from": "string",
         "op": "string",
         "path": "string",
         "value": "string"
      }
   ]
}
```

## URI Request Parameters

The request uses the following URI parameters.

**[restapi\_id](#API_UpdateStage_RequestSyntax)**

The string identifier of the associated RestApi.

Required: Yes

**[stage\_name](#API_UpdateStage_RequestSyntax)**

The name of the Stage resource to change information about.

Required: Yes

## Request Body

The request accepts the following data in JSON format.

**[patchOperations](#API_UpdateStage_RequestSyntax)**

For more information about supported patch operations, see [Patch Operations](patch-operations.md).

Type: Array of [PatchOperation](api-patchoperation.md) objects

Required: No

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "accessLogSettings": {
      "destinationArn": "string",
      "format": "string"
   },
   "cacheClusterEnabled": boolean,
   "cacheClusterSize": "string",
   "cacheClusterStatus": "string",
   "canarySettings": {
      "deploymentId": "string",
      "percentTraffic": number,
      "stageVariableOverrides": {
         "string" : "string"
      },
      "useStageCache": boolean
   },
   "clientCertificateId": "string",
   "createdDate": number,
   "deploymentId": "string",
   "description": "string",
   "documentationVersion": "string",
   "lastUpdatedDate": number,
   "methodSettings": {
      "string" : {
         "cacheDataEncrypted": boolean,
         "cacheTtlInSeconds": number,
         "cachingEnabled": boolean,
         "dataTraceEnabled": boolean,
         "loggingLevel": "string",
         "metricsEnabled": boolean,
         "requireAuthorizationForCacheControl": boolean,
         "throttlingBurstLimit": number,
         "throttlingRateLimit": number,
         "unauthorizedCacheControlHeaderStrategy": "string"
      }
   },
   "stageName": "string",
   "tags": {
      "string" : "string"
   },
   "tracingEnabled": boolean,
   "variables": {
      "string" : "string"
   },
   "webAclArn": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[accessLogSettings](#API_UpdateStage_ResponseSyntax)**

Settings for logging access in this stage.

Type: [AccessLogSettings](api-accesslogsettings.md) object

**[cacheClusterEnabled](#API_UpdateStage_ResponseSyntax)**

Specifies whether a cache cluster is enabled for the stage. To activate a method-level cache, set `CachingEnabled` to `true` for a method.

Type: Boolean

**[cacheClusterSize](#API_UpdateStage_ResponseSyntax)**

The stage's cache capacity in GB. For more information about choosing a cache size, see [Enabling API caching to enhance responsiveness](../developerguide/api-gateway-caching.md).

Type: String

Valid Values: `0.5 | 1.6 | 6.1 | 13.5 | 28.4 | 58.2 | 118 | 237`

**[cacheClusterStatus](#API_UpdateStage_ResponseSyntax)**

The status of the cache cluster for the stage, if enabled.

Type: String

Valid Values: `CREATE_IN_PROGRESS | AVAILABLE | DELETE_IN_PROGRESS | NOT_AVAILABLE | FLUSH_IN_PROGRESS`

**[canarySettings](#API_UpdateStage_ResponseSyntax)**

Settings for the canary deployment in this stage.

Type: [CanarySettings](api-canarysettings.md) object

**[clientCertificateId](#API_UpdateStage_ResponseSyntax)**

The identifier of a client certificate for an API stage.

Type: String

**[createdDate](#API_UpdateStage_ResponseSyntax)**

The timestamp when the stage was created.

Type: Timestamp

**[deploymentId](#API_UpdateStage_ResponseSyntax)**

The identifier of the Deployment that the stage points to.

Type: String

**[description](#API_UpdateStage_ResponseSyntax)**

The stage's description.

Type: String

**[documentationVersion](#API_UpdateStage_ResponseSyntax)**

The version of the associated API documentation.

Type: String

**[lastUpdatedDate](#API_UpdateStage_ResponseSyntax)**

The timestamp when the stage last updated.

Type: Timestamp

**[methodSettings](#API_UpdateStage_ResponseSyntax)**

A map that defines the method settings for a Stage resource. Keys (designated as `/{method_setting_key` below) are method paths defined as `{resource_path}/{http_method}` for an individual method override, or `/\*/\*` for overriding all methods in the stage.

Type: String to [MethodSetting](api-methodsetting.md) object map

**[stageName](#API_UpdateStage_ResponseSyntax)**

The name of the stage is the first path segment in the Uniform Resource Identifier (URI) of a call to API Gateway. Stage names can only contain alphanumeric characters, hyphens, and underscores. Maximum length is 128 characters.

Type: String

**[tags](#API_UpdateStage_ResponseSyntax)**

The collection of tags. Each tag element is associated with a given resource.

Type: String to string map

**[tracingEnabled](#API_UpdateStage_ResponseSyntax)**

Specifies whether active tracing with X-ray is enabled for the Stage.

Type: Boolean

**[variables](#API_UpdateStage_ResponseSyntax)**

A map that defines the stage variables for a Stage resource. Variable names can
have alphanumeric and underscore characters, and the values must match `[A-Za-z0-9-._~:/?#&=,]+`.

Type: String to string map

**[webAclArn](#API_UpdateStage_ResponseSyntax)**

The ARN of the WebAcl associated with the Stage.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**BadRequestException**

The submitted request is not valid, for example, the input is incomplete or incorrect. See the accompanying error message for details.

HTTP Status Code: 400

**ConflictException**

The request configuration has conflicts. For details, see the accompanying error message.

HTTP Status Code: 409

**LimitExceededException**

The request exceeded the rate limit. Retry after the specified time period.

HTTP Status Code: 429

**NotFoundException**

The requested resource is not found. Make sure that the request URI is correct.

HTTP Status Code: 404

**TooManyRequestsException**

The request has reached its throttling limit. Retry after the specified time period.

HTTP Status Code: 429

**UnauthorizedException**

The request is denied because the caller has insufficient permissions.

HTTP Status Code: 401

## Examples

### Update a deployment stage

This example illustrates one usage of UpdateStage.

#### Sample Request

```

PATCH /restapis/fugvjdxtri/stages/stage1 HTTP/1.1
Content-Type: application/json
Host: apigateway.us-east-1.amazonaws.com
X-Amz-Date: 20160603T200400Z
Authorization: AWS4-HMAC-SHA256 Credential={access_key_ID}/20160603/us-east-1/apigateway/aws4_request, SignedHeaders=content-length;content-type;host;x-amz-date, Signature={sig4_hash}

{
  "patchOperations" : [
    {
        "op" : "replace",
        "path" : "/*/*/metrics/enabled",
        "value" : "true"
    },
    {
        "op" : "replace",
        "path" : "/cacheClusterEnabled",
        "value" : "true"
    },
    {
        "op" : "replace",
        "path" : "/cacheClusterSize",
        "value" : "0.5"
    },
    {
        "op" : "replace",
        "path" : "/variables/sv2",
        "value" : "svVar"
    }
  ]
}

```

#### Sample Response

```

{
  "_links": {
    "curies": {
      "href": "https://docs.aws.amazon.com/apigateway/latest/developerguide/restapi-stage-{rel}.html",
      "name": "stage",
      "templated": true
    },
    "self": {
      "href": "/restapis/fugvjdxtri/stages/stage1"
    },
    "stage:delete": {
      "href": "/restapis/fugvjdxtri/stages/stage1"
    },
    "stage:flush-authorizer-cache": {
      "href": "/restapis/fugvjdxtri/stages/stage1/cache/authorizers"
    },
    "stage:flush-cache": {
      "href": "/restapis/fugvjdxtri/stages/stage1/cache/data"
    },
    "stage:update": {
      "href": "/restapis/fugvjdxtri/stages/stage1"
    }
  },
  "cacheClusterEnabled": true,
  "cacheClusterSize": "0.5",
  "cacheClusterStatus": "CREATE_IN_PROGRESS",
  "createdDate": "2016-06-03T17:56:06Z",
  "deploymentId": "dzacq7",
  "description": "First stage",
  "lastUpdatedDate": "2016-06-03T20:04:00Z",
  "methodSettings": {
    "/": {
      "dataTraceEnabled": false,
      "throttlingRateLimit": 500,
      "cacheTtlInSeconds": 0,
      "cachingEnabled": true,
      "requireAuthorizationForCacheControl": true,
      "metricsEnabled": true,
      "loggingLevel": "OFF",
      "unauthorizedCacheControlHeaderStrategy": "SUCCEED_WITH_RESPONSE_HEADER",
      "throttlingBurstLimit": 1000,
      "cacheDataEncrypted": false
    }
  },
  "stageName": "stage1",
  "variables": {
    "sv2": "svVar",
    "sv1": "opVar"
  }
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/apigateway-2015-07-09/updatestage.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apigateway-2015-07-09/updatestage.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apigateway-2015-07-09/updatestage.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apigateway-2015-07-09/updatestage.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apigateway-2015-07-09/updatestage.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apigateway-2015-07-09/updatestage.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apigateway-2015-07-09/updatestage.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apigateway-2015-07-09/updatestage.md)

- [AWS SDK for Python](../../../goto/boto3/apigateway-2015-07-09/updatestage.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apigateway-2015-07-09/updatestage.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UpdateRestApi

UpdateUsage

All content copied from https://docs.aws.amazon.com/.
