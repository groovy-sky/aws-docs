# ListDeployments

Lists the deployments for an environment in descending deployment number order.

## Request Syntax

```nohighlight

GET /applications/ApplicationId/environments/EnvironmentId/deployments?max_results=MaxResults&next_token=NextToken HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[ApplicationId](#API_ListDeployments_RequestSyntax)**

The application ID.

Pattern: `[a-z0-9]{4,7}`

Required: Yes

**[EnvironmentId](#API_ListDeployments_RequestSyntax)**

The environment ID.

Pattern: `[a-z0-9]{4,7}`

Required: Yes

**[MaxResults](#API_ListDeployments_RequestSyntax)**

The maximum number of items that may be returned for this call. If there are items that
have not yet been returned, the response will include a non-null `NextToken`
that you can provide in a subsequent call to get the next set of results.

Valid Range: Minimum value of 1. Maximum value of 50.

**[NextToken](#API_ListDeployments_RequestSyntax)**

The token returned by a prior call to this operation indicating the next set of results
to be returned. If not specified, the operation will return the first set of
results.

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
         "CompletedAt": "string",
         "ConfigurationName": "string",
         "ConfigurationVersion": "string",
         "DeploymentDurationInMinutes": number,
         "DeploymentNumber": number,
         "FinalBakeTimeInMinutes": number,
         "GrowthFactor": number,
         "GrowthType": "string",
         "PercentageComplete": number,
         "StartedAt": "string",
         "State": "string",
         "VersionLabel": "string"
      }
   ],
   "NextToken": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[Items](#API_ListDeployments_ResponseSyntax)**

The elements from this collection.

Type: Array of [DeploymentSummary](api-deploymentsummary.md) objects

**[NextToken](#API_ListDeployments_ResponseSyntax)**

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

**ResourceNotFoundException**

The requested resource could not be found.

HTTP Status Code: 404

## Examples

### Example

This example illustrates one usage of ListDeployments.

#### Sample Request

```

GET /applications/abc1234/environments/54j1r29/deployments HTTP/1.1
Host: appconfig.us-east-1.amazonaws.com
Accept-Encoding: identity
User-Agent: aws-cli/2.2.4 Python/3.8.8 Linux/5.4.134-73.228.amzn2int.x86_64 exe/x86_64.amzn.2 prompt/off command/appconfig.list-deployments
X-Amz-Date: 20210920T182141Z
Authorization: AWS4-HMAC-SHA256 Credential=AWS_ACCESS_KEY_ID_REDACTED/20210920/us-east-1/appconfig/aws4_request, SignedHeaders=host;x-amz-date, Signature=39c3b3042cd2aEXAMPLE
```

#### Sample Response

```

{
    "Items": [
        {
            "DeploymentNumber": 1,
            "ConfigurationName": "Example-Configuration-Profile",
            "ConfigurationVersion": "1",
            "DeploymentDurationInMinutes": 15,
            "GrowthType": "LINEAR",
            "GrowthFactor": 25.0,
            "FinalBakeTimeInMinutes": 0,
            "State": "COMPLETE",
            "PercentageComplete": 100.0,
            "StartedAt": "2021-09-17T21:43:54.205000+00:00",
            "CompletedAt": "2021-09-17T21:59:03.888000+00:00"
        }
    ]
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/appconfig-2019-10-09/listdeployments.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/appconfig-2019-10-09/listdeployments.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/appconfig-2019-10-09/listdeployments.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/appconfig-2019-10-09/listdeployments.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appconfig-2019-10-09/listdeployments.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/appconfig-2019-10-09/listdeployments.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/appconfig-2019-10-09/listdeployments.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/appconfig-2019-10-09/listdeployments.md)

- [AWS SDK for Python](../../../../services/goto/boto3/appconfig-2019-10-09/listdeployments.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appconfig-2019-10-09/listdeployments.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ListConfigurationProfiles

ListDeploymentStrategies

All content copied from https://docs.aws.amazon.com/.
