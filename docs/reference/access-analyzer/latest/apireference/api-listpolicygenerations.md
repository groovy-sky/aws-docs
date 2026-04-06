# ListPolicyGenerations

Lists all of the policy generations requested in the last seven days.

## Request Syntax

```nohighlight

GET /policy/generation?maxResults=maxResults&nextToken=nextToken&principalArn=principalArn HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[maxResults](#API_ListPolicyGenerations_RequestSyntax)**

The maximum number of results to return in the response.

Valid Range: Minimum value of 1.

**[nextToken](#API_ListPolicyGenerations_RequestSyntax)**

A token used for pagination of results returned.

**[principalArn](#API_ListPolicyGenerations_RequestSyntax)**

The ARN of the IAM entity (user or role) for which you are generating a policy. Use
this with `ListGeneratedPolicies` to filter the results to only include results
for a specific principal.

Pattern: `arn:[^:]*:iam::[^:]*:(role|user)/.{1,576}`

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "nextToken": "string",
   "policyGenerations": [
      {
         "completedOn": "string",
         "jobId": "string",
         "principalArn": "string",
         "startedOn": "string",
         "status": "string"
      }
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[nextToken](#API_ListPolicyGenerations_ResponseSyntax)**

A token used for pagination of results returned.

Type: String

**[policyGenerations](#API_ListPolicyGenerations_ResponseSyntax)**

A `PolicyGeneration` object that contains details about the generated
policy.

Type: Array of [PolicyGeneration](api-policygeneration.md) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**AccessDeniedException**

You do not have sufficient access to perform this action.

HTTP Status Code: 403

**InternalServerException**

Internal server error.

**retryAfterSeconds**

The seconds to wait to retry.

HTTP Status Code: 500

**ThrottlingException**

Throttling limit exceeded error.

**retryAfterSeconds**

The seconds to wait to retry.

HTTP Status Code: 429

**ValidationException**

Validation exception error.

**fieldList**

A list of fields that didn't validate.

**reason**

The reason for the exception.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/accessanalyzer-2019-11-01/ListPolicyGenerations)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/accessanalyzer-2019-11-01/ListPolicyGenerations)

- [AWS SDK for C++](../../../goto/sdkforcpp/accessanalyzer-2019-11-01/listpolicygenerations.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/accessanalyzer-2019-11-01/listpolicygenerations.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/accessanalyzer-2019-11-01/listpolicygenerations.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/accessanalyzer-2019-11-01/listpolicygenerations.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/accessanalyzer-2019-11-01/listpolicygenerations.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/accessanalyzer-2019-11-01/listpolicygenerations.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/accessanalyzer-2019-11-01/ListPolicyGenerations)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/accessanalyzer-2019-11-01/listpolicygenerations.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ListFindingsV2

ListTagsForResource
