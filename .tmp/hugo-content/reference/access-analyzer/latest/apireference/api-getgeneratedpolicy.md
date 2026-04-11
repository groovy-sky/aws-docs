# GetGeneratedPolicy

Retrieves the policy that was generated using `StartPolicyGeneration`.

## Request Syntax

```nohighlight

GET /policy/generation/jobId?includeResourcePlaceholders=includeResourcePlaceholders&includeServiceLevelTemplate=includeServiceLevelTemplate HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[includeResourcePlaceholders](#API_GetGeneratedPolicy_RequestSyntax)**

The level of detail that you want to generate. You can specify whether to generate
policies with placeholders for resource ARNs for actions that support resource level
granularity in policies.

For example, in the resource section of a policy, you can receive a placeholder such as
`"Resource":"arn:aws:s3:::${BucketName}"` instead of `"*"`.

**[includeServiceLevelTemplate](#API_GetGeneratedPolicy_RequestSyntax)**

The level of detail that you want to generate. You can specify whether to generate
service-level policies.

IAM Access Analyzer uses `iam:servicelastaccessed` to identify services that have
been used recently to create this service-level template.

**[jobId](#API_GetGeneratedPolicy_RequestSyntax)**

The `JobId` that is returned by the `StartPolicyGeneration`
operation. The `JobId` can be used with `GetGeneratedPolicy` to
retrieve the generated policies or used with `CancelPolicyGeneration` to cancel
the policy generation request.

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "generatedPolicyResult": {
      "generatedPolicies": [
         {
            "policy": "string"
         }
      ],
      "properties": {
         "cloudTrailProperties": {
            "endTime": "string",
            "startTime": "string",
            "trailProperties": [
               {
                  "allRegions": boolean,
                  "cloudTrailArn": "string",
                  "regions": [ "string" ]
               }
            ]
         },
         "isComplete": boolean,
         "principalArn": "string"
      }
   },
   "jobDetails": {
      "completedOn": "string",
      "jobError": {
         "code": "string",
         "message": "string"
      },
      "jobId": "string",
      "startedOn": "string",
      "status": "string"
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[generatedPolicyResult](#API_GetGeneratedPolicy_ResponseSyntax)**

A `GeneratedPolicyResult` object that contains the generated policies and
associated details.

Type: [GeneratedPolicyResult](api-generatedpolicyresult.md) object

**[jobDetails](#API_GetGeneratedPolicy_ResponseSyntax)**

A `GeneratedPolicyDetails` object that contains details about the generated
policy.

Type: [JobDetails](api-jobdetails.md) object

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/accessanalyzer-2019-11-01/getgeneratedpolicy.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/accessanalyzer-2019-11-01/getgeneratedpolicy.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/accessanalyzer-2019-11-01/getgeneratedpolicy.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/accessanalyzer-2019-11-01/getgeneratedpolicy.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/accessanalyzer-2019-11-01/getgeneratedpolicy.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/accessanalyzer-2019-11-01/getgeneratedpolicy.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/accessanalyzer-2019-11-01/getgeneratedpolicy.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/accessanalyzer-2019-11-01/getgeneratedpolicy.md)

- [AWS SDK for Python](../../../../services/goto/boto3/accessanalyzer-2019-11-01/getgeneratedpolicy.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/accessanalyzer-2019-11-01/getgeneratedpolicy.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

GetFindingV2

ListAccessPreviewFindings
