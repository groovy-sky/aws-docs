# ListDeploymentStrategies

Lists deployment strategies.

## Request Syntax

```nohighlight

GET /deploymentstrategies?max_results=MaxResults&next_token=NextToken HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[MaxResults](#API_ListDeploymentStrategies_RequestSyntax)**

The maximum number of items to return for this call. The call also returns a token that
you can specify in a subsequent call to get the next set of results.

Valid Range: Minimum value of 1. Maximum value of 50.

**[NextToken](#API_ListDeploymentStrategies_RequestSyntax)**

A token to start the list. Use this token to get the next set of results.

Length Constraints: Minimum length of 1. Maximum length of 2048.

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "Items": [
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
   ],
   "NextToken": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[Items](#API_ListDeploymentStrategies_ResponseSyntax)**

The elements from this collection.

Type: Array of [DeploymentStrategy](https://docs.aws.amazon.com/appconfig/2019-10-09/APIReference/API_DeploymentStrategy.html) objects

**[NextToken](#API_ListDeploymentStrategies_ResponseSyntax)**

The token for the next set of items to return. Use this token to get the next set of
results.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

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

## Examples

### Example

This example illustrates one usage of ListDeploymentStrategies.

#### Sample Request

```

GET /deploymentstrategies HTTP/1.1
Host: appconfig.us-east-1.amazonaws.com
Accept-Encoding: identity
User-Agent: aws-cli/2.2.4 Python/3.8.8 Linux/5.4.134-73.228.amzn2int.x86_64 exe/x86_64.amzn.2 prompt/off command/appconfig.list-deployment-strategies
X-Amz-Date: 20210920T174939Z
Authorization: AWS4-HMAC-SHA256 Credential=AKIAIOSFODNN7EXAMPLE/20210920/us-east-1/appconfig/aws4_request, SignedHeaders=host;x-amz-date, Signature=39c3b3042cd2aEXAMPLE
```

#### Sample Response

```

{
    "Items": [
        {
            "Id": "1225qzk",
            "Name": "Example-Deployment",
            "DeploymentDurationInMinutes": 15,
            "GrowthType": "LINEAR",
            "GrowthFactor": 25.0,
            "FinalBakeTimeInMinutes": 0,
            "ReplicateTo": "SSM_DOCUMENT"
        },
        {
            "Id": "AppConfig.AllAtOnce",
            "Name": "AppConfig.AllAtOnce",
            "Description": "Quick",
            "DeploymentDurationInMinutes": 0,
            "GrowthType": "LINEAR",
            "GrowthFactor": 100.0,
            "FinalBakeTimeInMinutes": 10,
            "ReplicateTo": "NONE"
        },
        {
            "Id": "AppConfig.Linear50PercentEvery30Seconds",
            "Name": "AppConfig.Linear50PercentEvery30Seconds",
            "Description": "Test/Demo",
            "DeploymentDurationInMinutes": 1,
            "GrowthType": "LINEAR",
            "GrowthFactor": 50.0,
            "FinalBakeTimeInMinutes": 1,
            "ReplicateTo": "NONE"
        },
        {
            "Id": "AppConfig.Canary10Percent20Minutes",
            "Name": "AppConfig.Canary10Percent20Minutes",
            "Description": "AWS Recommended",
            "DeploymentDurationInMinutes": 20,
            "GrowthType": "EXPONENTIAL",
            "GrowthFactor": 10.0,
            "FinalBakeTimeInMinutes": 10,
            "ReplicateTo": "NONE"
        }
    ]
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/appconfig-2019-10-09/ListDeploymentStrategies)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/appconfig-2019-10-09/ListDeploymentStrategies)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/appconfig-2019-10-09/ListDeploymentStrategies)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/appconfig-2019-10-09/ListDeploymentStrategies)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/appconfig-2019-10-09/ListDeploymentStrategies)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/appconfig-2019-10-09/ListDeploymentStrategies)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/appconfig-2019-10-09/ListDeploymentStrategies)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/appconfig-2019-10-09/ListDeploymentStrategies)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/appconfig-2019-10-09/ListDeploymentStrategies)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/appconfig-2019-10-09/ListDeploymentStrategies)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ListDeployments

ListEnvironments
