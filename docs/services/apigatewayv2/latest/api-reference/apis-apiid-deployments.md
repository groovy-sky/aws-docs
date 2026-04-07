# Deployments

Represents the collection of deployments for an API.

## URI

`/v2/apis/apiId/deployments`

## HTTP methods

### GET

**Operation ID:** `GetDeployments`

Gets the `Deployments` for an API.

Path parametersNameTypeRequiredDescription`apiId`StringTrue

The API identifier.

Query parametersNameTypeRequiredDescription`nextToken`StringFalse

The next page of elements from this collection. Not valid for the last element of the collection.

`maxResults`StringFalse

The maximum number of elements to be returned for this resource.

ResponsesStatus codeResponse modelDescription`200``Deployments`

Success

`400``BadRequestException`

One of the parameters in the request is invalid.

`404``NotFoundException`

The resource specified in the request was not found.

`429``LimitExceededException`

The client is sending more than the allowed number of requests per unit of time.

### POST

**Operation ID:** `CreateDeployment`

Creates a `Deployment` for an API.

Path parametersNameTypeRequiredDescription`apiId`StringTrue

The API identifier.

ResponsesStatus codeResponse modelDescription`201``Deployment`

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
  "description": "string",
  "stageName": "string"
}
```

### Response bodies

```json

{
  "items": [
    {
      "deploymentId": "string",
      "description": "string",
      "createdDate": "string",
      "deploymentStatus": enum,
      "deploymentStatusMessage": "string",
      "autoDeployed": boolean
    }
  ],
  "nextToken": "string"
}
```

```json

{
  "deploymentId": "string",
  "description": "string",
  "createdDate": "string",
  "deploymentStatus": enum,
  "deploymentStatusMessage": "string",
  "autoDeployed": boolean
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

### CreateDeploymentInput

Represents the input parameters for a `CreateDeployment` request.

PropertyTypeRequiredDescription`description`

string

False

The description for the deployment resource.

`stageName`

string

False

The name of an existing stage to associate with the deployment.

### Deployment

An immutable representation of an API that can be called by users. A `Deployment` must be associated with a `Stage` for it to be callable over the internet.

PropertyTypeRequiredDescription`autoDeployed`

boolean

False

Specifies whether a deployment was automatically released.

`createdDate`

string

Format: date-time

False

The date and time when the `Deployment` resource was created.

`deploymentId`

string

False

The identifier for the deployment.

`deploymentStatus`

[DeploymentStatus](#apis-apiid-deployments-model-deploymentstatus)

False

The status of the deployment: `PENDING`, `FAILED`, or `SUCCEEDED`.

`deploymentStatusMessage`

string

False

May contain additional feedback on the status of an API deployment.

`description`

string

False

The description for the deployment.

### DeploymentStatus

Represents a deployment status.

- `PENDING`

- `FAILED`

- `DEPLOYED`

### Deployments

A collection resource that contains zero or more references to your existing deployments, and links that guide you on how to interact with your collection. The collection offers a paginated view of the contained deployments.

PropertyTypeRequiredDescription`items`

Array of type Deployment

False

The elements from this collection.

`nextToken`

string

False

The next page of elements from this collection. Not valid for the last element of the collection.

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

## See also

For more information about using this API in one of the language-specific AWS SDKs and references, see the following:

### GetDeployments

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/apigatewayv2-2018-11-29/GetDeployments)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/apigatewayv2-2018-11-29/GetDeployments)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/apigatewayv2-2018-11-29/GetDeployments)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/apigatewayv2-2018-11-29/GetDeployments)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/apigatewayv2-2018-11-29/GetDeployments)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/apigatewayv2-2018-11-29/GetDeployments)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/apigatewayv2-2018-11-29/GetDeployments)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/apigatewayv2-2018-11-29/GetDeployments)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/apigatewayv2-2018-11-29/GetDeployments)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/apigatewayv2-2018-11-29/GetDeployments)

### CreateDeployment

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/apigatewayv2-2018-11-29/CreateDeployment)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/apigatewayv2-2018-11-29/CreateDeployment)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/apigatewayv2-2018-11-29/CreateDeployment)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/apigatewayv2-2018-11-29/CreateDeployment)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/apigatewayv2-2018-11-29/CreateDeployment)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/apigatewayv2-2018-11-29/CreateDeployment)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/apigatewayv2-2018-11-29/CreateDeployment)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/apigatewayv2-2018-11-29/CreateDeployment)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/apigatewayv2-2018-11-29/CreateDeployment)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/apigatewayv2-2018-11-29/CreateDeployment)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Deployment

DomainName
