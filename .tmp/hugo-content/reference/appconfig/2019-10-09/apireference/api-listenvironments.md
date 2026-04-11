# ListEnvironments

Lists the environments for an application.

## Request Syntax

```nohighlight

GET /applications/ApplicationId/environments?max_results=MaxResults&next_token=NextToken HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[ApplicationId](#API_ListEnvironments_RequestSyntax)**

The application ID.

Pattern: `[a-z0-9]{4,7}`

Required: Yes

**[MaxResults](#API_ListEnvironments_RequestSyntax)**

The maximum number of items to return for this call. The call also returns a token that
you can specify in a subsequent call to get the next set of results.

Valid Range: Minimum value of 1. Maximum value of 50.

**[NextToken](#API_ListEnvironments_RequestSyntax)**

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
         "ApplicationId": "string",
         "Description": "string",
         "Id": "string",
         "Monitors": [
            {
               "AlarmArn": "string",
               "AlarmRoleArn": "string"
            }
         ],
         "Name": "string",
         "State": "string"
      }
   ],
   "NextToken": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[Items](#API_ListEnvironments_ResponseSyntax)**

The elements from this collection.

Type: Array of [Environment](api-environment.md) objects

**[NextToken](#API_ListEnvironments_ResponseSyntax)**

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

This example illustrates one usage of ListEnvironments.

#### Sample Request

```

GET /applications/abc1234/environments HTTP/1.1
Host: appconfig.us-east-1.amazonaws.com
Accept-Encoding: identity
User-Agent: aws-cli/2.2.4 Python/3.8.8 Linux/5.4.134-73.228.amzn2int.x86_64 exe/x86_64.amzn.2 prompt/off command/appconfig.list-environments
X-Amz-Date: 20210920T182621Z
Authorization: AWS4-HMAC-SHA256 Credential=AWS_ACCESS_KEY_ID_REDACTED/20210920/us-east-1/appconfig/aws4_request, SignedHeaders=host;x-amz-date, Signature=39c3b3042cd2aEXAMPLE
```

#### Sample Response

```

{
    "Items": [
        {
            "ApplicationId": "abc1234",
            "Id": "54j1r29",
            "Name": "Example-Environment",
            "State": "ReadyForDeployment"
        }
    ]
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/appconfig-2019-10-09/listenvironments.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/appconfig-2019-10-09/listenvironments.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/appconfig-2019-10-09/listenvironments.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/appconfig-2019-10-09/listenvironments.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appconfig-2019-10-09/listenvironments.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/appconfig-2019-10-09/listenvironments.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/appconfig-2019-10-09/listenvironments.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/appconfig-2019-10-09/listenvironments.md)

- [AWS SDK for Python](../../../../services/goto/boto3/appconfig-2019-10-09/listenvironments.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appconfig-2019-10-09/listenvironments.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ListDeploymentStrategies

ListExtensionAssociations

All content copied from https://docs.aws.amazon.com/.
