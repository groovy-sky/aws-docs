# CreateDeployment

Creates a Deployment resource, which makes a specified RestApi callable over the internet.

## Request Syntax

```nohighlight

POST /restapis/restapi_id/deployments HTTP/1.1
Content-type: application/json

{
   "cacheClusterEnabled": boolean,
   "cacheClusterSize": "string",
   "canarySettings": {
      "percentTraffic": number,
      "stageVariableOverrides": {
         "string" : "string"
      },
      "useStageCache": boolean
   },
   "description": "string",
   "stageDescription": "string",
   "stageName": "string",
   "tracingEnabled": boolean,
   "variables": {
      "string" : "string"
   }
}
```

## URI Request Parameters

The request uses the following URI parameters.

**[restapi\_id](#API_CreateDeployment_RequestSyntax)**

The string identifier of the associated RestApi.

Required: Yes

## Request Body

The request accepts the following data in JSON format.

**[cacheClusterEnabled](#API_CreateDeployment_RequestSyntax)**

Enables a cache cluster for the Stage resource specified in the input.

Type: Boolean

Required: No

**[cacheClusterSize](#API_CreateDeployment_RequestSyntax)**

The stage's cache capacity in GB. For more information about choosing a cache size, see [Enabling API caching to enhance responsiveness](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-caching.html).

Type: String

Valid Values: `0.5 | 1.6 | 6.1 | 13.5 | 28.4 | 58.2 | 118 | 237`

Required: No

**[canarySettings](#API_CreateDeployment_RequestSyntax)**

The input configuration for the canary deployment when the deployment is a canary release deployment.

Type: [DeploymentCanarySettings](https://docs.aws.amazon.com/apigateway/latest/api/API_DeploymentCanarySettings.html) object

Required: No

**[description](#API_CreateDeployment_RequestSyntax)**

The description for the Deployment resource to create.

Type: String

Required: No

**[stageDescription](#API_CreateDeployment_RequestSyntax)**

The description of the Stage resource for the Deployment resource to create.

Type: String

Required: No

**[stageName](#API_CreateDeployment_RequestSyntax)**

The name of the Stage resource for the Deployment resource to create.

Type: String

Required: No

**[tracingEnabled](#API_CreateDeployment_RequestSyntax)**

Specifies whether active tracing with X-ray is enabled for the Stage.

Type: Boolean

Required: No

**[variables](#API_CreateDeployment_RequestSyntax)**

A map that defines the stage variables for the Stage resource that is associated
with the new deployment. Variable names can have alphanumeric and underscore characters, and the values
must match `[A-Za-z0-9-._~:/?#&=,]+`.

Type: String to string map

Required: No

## Response Syntax

```nohighlight

HTTP/1.1 201
Content-type: application/json

{
   "apiSummary": {
      "string" : {
         "string" : {
            "apiKeyRequired": boolean,
            "authorizationType": "string"
         }
      }
   },
   "createdDate": number,
   "description": "string",
   "id": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 201 response.

The following data is returned in JSON format by the service.

**[apiSummary](#API_CreateDeployment_ResponseSyntax)**

A summary of the RestApi at the date and time that the deployment resource was created.

Type: String to string to [MethodSnapshot](https://docs.aws.amazon.com/apigateway/latest/api/API_MethodSnapshot.html) object map map

**[createdDate](#API_CreateDeployment_ResponseSyntax)**

The date and time that the deployment resource was created.

Type: Timestamp

**[description](#API_CreateDeployment_ResponseSyntax)**

The description for the deployment resource.

Type: String

**[id](#API_CreateDeployment_ResponseSyntax)**

The identifier for the deployment resource.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common Error Types](https://docs.aws.amazon.com/apigateway/latest/api/CommonErrors.html).

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

**ServiceUnavailableException**

The requested service is not available. For details see the accompanying error message. Retry after the specified time period.

HTTP Status Code: 503

**TooManyRequestsException**

The request has reached its throttling limit. Retry after the specified time period.

HTTP Status Code: 429

**UnauthorizedException**

The request is denied because the caller has insufficient permissions.

HTTP Status Code: 401

## Examples

### Create a deployment of an API

This example illustrates one usage of CreateDeployment.

#### Sample Request

```

POST /restapis/fugvjdxtri/deployments HTTP/1.1
Content-Type: application/json
Host: apigateway.us-east-1.amazonaws.com
X-Amz-Date: 20160603T175605Z
Authorization: AWS4-HMAC-SHA256 Credential={access_key_ID}/20160603/us-east-1/apigateway/aws4_request, SignedHeaders=content-length;content-type;host;x-amz-date, Signature={sig4_hash}

{
  "stageName" : "stage1",
  "stageDescription" : "First stage",
  "description" : "First deployment",
  "cacheClusterEnabled" : "false",
  "variables" : {
    "sv1" : "opVar"
  }
}
```

#### Sample Response

```

{
  "_links": {
    "curies": {
      "href": "https://docs.aws.amazon.com/apigateway/latest/developerguide/restapi-deployment-{rel}.html",
      "name": "deployment",
      "templated": true
    },
    "self": {
      "href": "/restapis/fugvjdxtri/deployments/dzacq7"
    },
    "deployment:delete": {
      "href": "/restapis/fugvjdxtri/deployments/dzacq7"
    },
    "deployment:stages": {
      "href": "/restapis/fugvjdxtri/stages?deployment_id=dzacq7"
    },
    "deployment:update": {
      "href": "/restapis/fugvjdxtri/deployments/dzacq7"
    }
  },
  "createdDate": "2016-06-03T17:56:06Z",
  "description": "First deployment",
  "id": "dzacq7"
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/apigateway-2015-07-09/CreateDeployment)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/apigateway-2015-07-09/CreateDeployment)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/apigateway-2015-07-09/CreateDeployment)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/apigateway-2015-07-09/CreateDeployment)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/apigateway-2015-07-09/CreateDeployment)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/apigateway-2015-07-09/CreateDeployment)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/apigateway-2015-07-09/CreateDeployment)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/apigateway-2015-07-09/CreateDeployment)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/apigateway-2015-07-09/CreateDeployment)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/apigateway-2015-07-09/CreateDeployment)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CreateBasePathMapping

CreateDocumentationPart
