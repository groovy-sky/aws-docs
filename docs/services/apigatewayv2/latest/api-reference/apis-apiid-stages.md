# Stages

Represents the collection of stages for an API.

## URI

`/v2/apis/apiId/stages`

## HTTP methods

### GET

**Operation ID:** `GetStages`

Gets the `Stages` for an API.

Path parametersNameTypeRequiredDescription`apiId`StringTrue

The API identifier.

Query parametersNameTypeRequiredDescription`nextToken`StringFalse

The next page of elements from this collection. Not valid for the last element of the collection.

`maxResults`StringFalse

The maximum number of elements to be returned for this resource.

ResponsesStatus codeResponse modelDescription`200``Stages`

Success

`400``BadRequestException`

One of the parameters in the request is invalid.

`404``NotFoundException`

The resource specified in the request was not found.

`429``LimitExceededException`

The client is sending more than the allowed number of requests per unit of time.

### POST

**Operation ID:** `CreateStage`

Creates a `Stage` for an API.

Path parametersNameTypeRequiredDescription`apiId`StringTrue

The API identifier.

ResponsesStatus codeResponse modelDescription`201``Stage`

The request has succeeded and has resulted in the creation of a resource.

`400``BadRequestException`

One of the parameters in the request is invalid.

`404``NotFoundException`

The resource specified in the request was not found.

`409``ConflictException`

The resource already exists.

`429``LimitExceededException`

The client is sending more than the allowed number of requests per unit of time.

## Schemas

### Request bodies

```json

{
  "stageName": "string",
  "description": "string",
  "deploymentId": "string",
  "clientCertificateId": "string",
  "defaultRouteSettings": {
    "detailedMetricsEnabled": boolean,
    "loggingLevel": enum,
    "dataTraceEnabled": boolean,
    "throttlingBurstLimit": integer,
    "throttlingRateLimit": number
  },
  "routeSettings": {
  },
  "stageVariables": {
  },
  "accessLogSettings": {
    "format": "string",
    "destinationArn": "string"
  },
  "tags": {
  },
  "autoDeploy": boolean
}
```

### Response bodies

```json

{
  "items": [
    {
      "stageName": "string",
      "description": "string",
      "deploymentId": "string",
      "clientCertificateId": "string",
      "defaultRouteSettings": {
        "detailedMetricsEnabled": boolean,
        "loggingLevel": enum,
        "dataTraceEnabled": boolean,
        "throttlingBurstLimit": integer,
        "throttlingRateLimit": number
      },
      "routeSettings": {
      },
      "stageVariables": {
      },
      "accessLogSettings": {
        "format": "string",
        "destinationArn": "string"
      },
      "autoDeploy": boolean,
      "lastDeploymentStatusMessage": "string",
      "createdDate": "string",
      "lastUpdatedDate": "string",
      "tags": {
      },
      "apiGatewayManaged": boolean
    }
  ],
  "nextToken": "string"
}
```

```json

{
  "stageName": "string",
  "description": "string",
  "deploymentId": "string",
  "clientCertificateId": "string",
  "defaultRouteSettings": {
    "detailedMetricsEnabled": boolean,
    "loggingLevel": enum,
    "dataTraceEnabled": boolean,
    "throttlingBurstLimit": integer,
    "throttlingRateLimit": number
  },
  "routeSettings": {
  },
  "stageVariables": {
  },
  "accessLogSettings": {
    "format": "string",
    "destinationArn": "string"
  },
  "autoDeploy": boolean,
  "lastDeploymentStatusMessage": "string",
  "createdDate": "string",
  "lastUpdatedDate": "string",
  "tags": {
  },
  "apiGatewayManaged": boolean
}
```

```json

{
  "message": "string"
}
```

```json

{
  "message": "string",
  "resourceType": "string"
}
```

```json

{
  "message": "string"
}
```

```json

{
  "message": "string",
  "limitType": "string"
}
```

## Properties

### AccessLogSettings

Settings for logging access in a stage.

PropertyTypeRequiredDescription`destinationArn`

string

False

The ARN of the CloudWatch Logs log group to receive access logs.

`format`

string

False

A single line format of the access logs of data, as specified by selected $context variables. The format must include at least $context.requestId.

### BadRequestException

The request is not valid, for example, the input is incomplete or incorrect. See the accompanying error message for details.

PropertyTypeRequiredDescription`message`

string

False

Describes the error encountered.

### ConflictException

The requested operation would cause a conflict with the current state of a service resource associated with the request. Resolve the conflict before retrying this request. See the accompanying error message for details.

PropertyTypeRequiredDescription`message`

string

False

Describes the error encountered.

### CreateStageInput

Represents the input parameters for a `CreateStage` request.

PropertyTypeRequiredDescription`accessLogSettings`

[AccessLogSettings](#apis-apiid-stages-model-accesslogsettings)

False

Settings for logging access in this stage.

`autoDeploy`

boolean

False

Specifies whether updates to an API automatically trigger a new deployment. The default value is `false`.

`clientCertificateId`

string

False

The identifier of a client certificate for a `Stage`. Supported only for WebSocket APIs.

`defaultRouteSettings`

[RouteSettings](#apis-apiid-stages-model-routesettings)

False

The default route settings for the stage.

`deploymentId`

string

False

The deployment identifier of the API stage.

`description`

string

False

The description for the API stage.

`routeSettings`

[RouteSettingsMap](#apis-apiid-stages-model-routesettingsmap)

False

Route settings for the stage, by `routeKey`.

`stageName`

string

True

The name of the stage.

`stageVariables`

[StageVariablesMap](#apis-apiid-stages-model-stagevariablesmap)

False

A map that defines the stage variables for a `Stage`. Variable names can have alphanumeric and underscore characters, and the values must match \[A-Za-z0-9-.\_~:/?#&=,\]+.

`tags`

[Tags](#apis-apiid-stages-model-tags)

False

The collection of tags. Each tag element is associated with a given resource.

### LimitExceededException

A limit has been exceeded. See the accompanying error message for details.

PropertyTypeRequiredDescription`limitType`

string

False

The limit type.

`message`

string

False

Describes the error encountered.

### LoggingLevel

The logging level.

- `ERROR`

- `INFO`

- `false`

### NotFoundException

The resource specified in the request was not found. See the `message` field for more information.

PropertyTypeRequiredDescription`message`

string

False

Describes the error encountered.

`resourceType`

string

False

The resource type.

### RouteSettings

Represents a collection of route settings.

PropertyTypeRequiredDescription`dataTraceEnabled`

boolean

False

Specifies whether ( `true`) or not ( `false`) data trace logging is enabled for this route. This property affects the log entries pushed to Amazon CloudWatch Logs. Supported only for WebSocket APIs.

`detailedMetricsEnabled`

boolean

False

Specifies whether detailed metrics are enabled.

`loggingLevel`

[LoggingLevel](#apis-apiid-stages-model-logginglevel)

False

Specifies the logging level for this route: `INFO`, `ERROR`, or `OFF`. This property affects the log entries pushed to Amazon CloudWatch Logs. Supported only for WebSocket APIs.

`throttlingBurstLimit`

integer

Format: int32

False

Specifies the throttling burst limit.

`throttlingRateLimit`

number

Format: double

False

Specifies the throttling rate limit.

### RouteSettingsMap

The route settings map.

PropertyTypeRequiredDescription

`*`

object

False

### Stage

Represents an API stage.

PropertyTypeRequiredDescription`accessLogSettings`

[AccessLogSettings](#apis-apiid-stages-model-accesslogsettings)

False

Settings for logging access in this stage.

`apiGatewayManaged`

boolean

False

Specifies whether a stage is managed by API Gateway. If you created an API using
quick create, the `$default` stage is managed by API Gateway. You can't
modify the `$default` stage.

`autoDeploy`

boolean

False

Specifies whether updates to an API automatically trigger a new deployment. The default value is `false`.

`clientCertificateId`

string

False

The identifier of a client certificate for a `Stage`. Supported only for WebSocket APIs.

`createdDate`

string

Format: date-time

False

The timestamp when the stage was created.

`defaultRouteSettings`

[RouteSettings](#apis-apiid-stages-model-routesettings)

False

Default route settings for the stage.

`deploymentId`

string

False

The identifier of the `Deployment` that the `Stage` is associated with. Can't be updated if `autoDeploy` is enabled.

`description`

string

False

The description of the stage.

`lastDeploymentStatusMessage`

string

False

Describes the status of the last deployment of a stage. Supported only for stages with `autoDeploy` enabled.

`lastUpdatedDate`

string

Format: date-time

False

The timestamp when the stage was last updated.

`routeSettings`

[RouteSettingsMap](#apis-apiid-stages-model-routesettingsmap)

False

Route settings for the stage, by `routeKey`.

`stageName`

string

True

The name of the stage.

`stageVariables`

[StageVariablesMap](#apis-apiid-stages-model-stagevariablesmap)

False

A map that defines the stage variables for a stage resource. Variable names can have alphanumeric and underscore characters, and the values must match \[A-Za-z0-9-.\_~:/?#&=,\]+.

`tags`

[Tags](#apis-apiid-stages-model-tags)

False

The collection of tags. Each tag element is associated with a given resource.

### StageVariablesMap

The stage variable map.

PropertyTypeRequiredDescription

`*`

string

False

### Stages

A collection of Stage resources that are associated with the ApiKey resource.

PropertyTypeRequiredDescription`items`

Array of type Stage

False

The elements from this collection.

`nextToken`

string

False

The next page of elements from this collection. Not valid for the last element of the collection.

### Tags

Represents a collection of tags associated with the resource.

PropertyTypeRequiredDescription

`*`

string

False

## See also

For more information about using this API in one of the language-specific AWS SDKs and references, see the following:

### GetStages

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/apigatewayv2-2018-11-29/GetStages)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/apigatewayv2-2018-11-29/GetStages)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/apigatewayv2-2018-11-29/GetStages)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/apigatewayv2-2018-11-29/GetStages)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/apigatewayv2-2018-11-29/GetStages)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/apigatewayv2-2018-11-29/GetStages)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/apigatewayv2-2018-11-29/GetStages)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/apigatewayv2-2018-11-29/GetStages)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/apigatewayv2-2018-11-29/GetStages)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/apigatewayv2-2018-11-29/GetStages)

### CreateStage

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/apigatewayv2-2018-11-29/CreateStage)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/apigatewayv2-2018-11-29/CreateStage)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/apigatewayv2-2018-11-29/CreateStage)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/apigatewayv2-2018-11-29/CreateStage)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/apigatewayv2-2018-11-29/CreateStage)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/apigatewayv2-2018-11-29/CreateStage)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/apigatewayv2-2018-11-29/CreateStage)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/apigatewayv2-2018-11-29/CreateStage)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/apigatewayv2-2018-11-29/CreateStage)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/apigatewayv2-2018-11-29/CreateStage)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Stage

Tags
