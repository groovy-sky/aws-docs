# GetFindingRecommendation

Retrieves information about a finding recommendation for the specified analyzer.

## Request Syntax

```nohighlight

GET /recommendation/id?analyzerArn=analyzerArn&maxResults=maxResults&nextToken=nextToken HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[analyzerArn](#API_GetFindingRecommendation_RequestSyntax)**

The [ARN of\
the analyzer](../../../../services/iam/latest/userguide/access-analyzer-getting-started.md#permission-resources) used to generate the finding recommendation.

Pattern: `[^:]*:[^:]*:[^:]*:[^:]*:[^:]*:analyzer/.{1,255}`

Required: Yes

**[id](#API_GetFindingRecommendation_RequestSyntax)**

The unique ID for the finding recommendation.

Length Constraints: Minimum length of 1. Maximum length of 2048.

Required: Yes

**[maxResults](#API_GetFindingRecommendation_RequestSyntax)**

The maximum number of results to return in the response.

Valid Range: Minimum value of 1. Maximum value of 1000.

**[nextToken](#API_GetFindingRecommendation_RequestSyntax)**

A token used for pagination of results returned.

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "completedAt": "string",
   "error": {
      "code": "string",
      "message": "string"
   },
   "nextToken": "string",
   "recommendationType": "string",
   "recommendedSteps": [
      { ... }
   ],
   "resourceArn": "string",
   "startedAt": "string",
   "status": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[completedAt](#API_GetFindingRecommendation_ResponseSyntax)**

The time at which the retrieval of the finding recommendation was completed.

Type: Timestamp

**[error](#API_GetFindingRecommendation_ResponseSyntax)**

Detailed information about the reason that the retrieval of a recommendation for the
finding failed.

Type: [RecommendationError](api-recommendationerror.md) object

**[nextToken](#API_GetFindingRecommendation_ResponseSyntax)**

A token used for pagination of results returned.

Type: String

**[recommendationType](#API_GetFindingRecommendation_ResponseSyntax)**

The type of recommendation for the finding.

Type: String

Valid Values: `UnusedPermissionRecommendation`

**[recommendedSteps](#API_GetFindingRecommendation_ResponseSyntax)**

A group of recommended steps for the finding.

Type: Array of [RecommendedStep](api-recommendedstep.md) objects

**[resourceArn](#API_GetFindingRecommendation_ResponseSyntax)**

The ARN of the resource of the finding.

Type: String

Pattern: `arn:[^:]*:[^:]*:[^:]*:[^:]*:.*`

**[startedAt](#API_GetFindingRecommendation_ResponseSyntax)**

The time at which the retrieval of the finding recommendation was started.

Type: Timestamp

**[status](#API_GetFindingRecommendation_ResponseSyntax)**

The status of the retrieval of the finding recommendation.

Type: String

Valid Values: `SUCCEEDED | FAILED | IN_PROGRESS`

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

**ResourceNotFoundException**

The specified resource could not be found.

**resourceId**

The ID of the resource.

**resourceType**

The type of the resource.

HTTP Status Code: 404

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/accessanalyzer-2019-11-01/getfindingrecommendation.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/accessanalyzer-2019-11-01/getfindingrecommendation.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/accessanalyzer-2019-11-01/getfindingrecommendation.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/accessanalyzer-2019-11-01/getfindingrecommendation.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/accessanalyzer-2019-11-01/getfindingrecommendation.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/accessanalyzer-2019-11-01/getfindingrecommendation.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/accessanalyzer-2019-11-01/getfindingrecommendation.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/accessanalyzer-2019-11-01/getfindingrecommendation.md)

- [AWS SDK for Python](../../../../services/goto/boto3/accessanalyzer-2019-11-01/getfindingrecommendation.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/accessanalyzer-2019-11-01/getfindingrecommendation.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

GetFinding

GetFindingsStatistics
