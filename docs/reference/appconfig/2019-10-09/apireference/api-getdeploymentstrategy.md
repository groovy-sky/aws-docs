# GetDeploymentStrategy

Retrieves information about a deployment strategy. A deployment strategy defines
important criteria for rolling out your configuration to the designated targets. A
deployment strategy includes the overall duration required, a percentage of targets to
receive the deployment during each interval, an algorithm that defines how percentage
grows, and bake time.

## Request Syntax

```nohighlight

GET /deploymentstrategies/DeploymentStrategyId HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[DeploymentStrategyId](#API_GetDeploymentStrategy_RequestSyntax)**

The ID of the deployment strategy to get.

Pattern: `(^[a-z0-9]{4,7}$|^AppConfig\.[A-Za-z0-9]{9,40}$)`

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "DeploymentDurationInMinutes": number,
   "Description": "string",
   "FinalBakeTimeInMinutes": number,
   "GrowthFactor": number,
   "GrowthType": "string",
   "Id": "string",
   "Name": "string",
   "ReplicateTo": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[DeploymentDurationInMinutes](#API_GetDeploymentStrategy_ResponseSyntax)**

Total amount of time the deployment lasted.

Type: Integer

Valid Range: Minimum value of 0. Maximum value of 1440.

**[Description](#API_GetDeploymentStrategy_ResponseSyntax)**

The description of the deployment strategy.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1024.

**[FinalBakeTimeInMinutes](#API_GetDeploymentStrategy_ResponseSyntax)**

The amount of time that AWS AppConfig monitored for alarms before considering the
deployment to be complete and no longer eligible for automatic rollback.

Type: Integer

Valid Range: Minimum value of 0. Maximum value of 1440.

**[GrowthFactor](#API_GetDeploymentStrategy_ResponseSyntax)**

The percentage of targets that received a deployed configuration during each
interval.

Type: Float

Valid Range: Minimum value of 1.0. Maximum value of 100.0.

**[GrowthType](#API_GetDeploymentStrategy_ResponseSyntax)**

The algorithm used to define how percentage grew over time.

Type: String

Valid Values: `LINEAR | EXPONENTIAL`

**[Id](#API_GetDeploymentStrategy_ResponseSyntax)**

The deployment strategy ID.

Type: String

Pattern: `[a-z0-9]{4,7}`

**[Name](#API_GetDeploymentStrategy_ResponseSyntax)**

The name of the deployment strategy.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

**[ReplicateTo](#API_GetDeploymentStrategy_ResponseSyntax)**

Save the deployment strategy to a Systems Manager (SSM) document.

Type: String

Valid Values: `NONE | SSM_DOCUMENT`

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**BadRequestException**

The input fails to satisfy the constraints specified by an AWS service.

**Details**

Detailed information about the input that failed to satisfy the constraints specified by
a call.

HTTP Status Code: 400

**InternalServerException**

There was an internal failure in the AWS AppConfig service.

HTTP Status Code: 500

**ResourceNotFoundException**

The requested resource could not be found.

HTTP Status Code: 404

## Examples

### Example

This example illustrates one usage of GetDeploymentStrategy.

#### Sample Request

```

GET /deploymentstrategies/1225qzk HTTP/1.1
Host: appconfig.us-east-1.amazonaws.com
Accept-Encoding: identity
User-Agent: aws-cli/2.2.4 Python/3.8.8 Linux/5.4.134-73.228.amzn2int.x86_64 exe/x86_64.amzn.2 prompt/off command/appconfig.get-deployment-strategy
X-Amz-Date: 20210917T224000Z
Authorization: AWS4-HMAC-SHA256 Credential=AKIAIOSFODNN7EXAMPLE/20210917/us-east-1/appconfig/aws4_request, SignedHeaders=host;x-amz-date, Signature=39c3b3042cd2aEXAMPLE
```

#### Sample Response

```

{
    "Id": "1225qzk",
    "Name": "Example-Deployment",
    "DeploymentDurationInMinutes": 15,
    "GrowthType": "LINEAR",
    "GrowthFactor": 25.0,
    "FinalBakeTimeInMinutes": 0,
    "ReplicateTo": "SSM_DOCUMENT"
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/appconfig-2019-10-09/getdeploymentstrategy.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/appconfig-2019-10-09/getdeploymentstrategy.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/appconfig-2019-10-09/getdeploymentstrategy.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/appconfig-2019-10-09/getdeploymentstrategy.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appconfig-2019-10-09/getdeploymentstrategy.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/appconfig-2019-10-09/getdeploymentstrategy.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/appconfig-2019-10-09/getdeploymentstrategy.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/appconfig-2019-10-09/getdeploymentstrategy.md)

- [AWS SDK for Python](../../../../services/goto/boto3/appconfig-2019-10-09/getdeploymentstrategy.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appconfig-2019-10-09/getdeploymentstrategy.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

GetDeployment

GetEnvironment
