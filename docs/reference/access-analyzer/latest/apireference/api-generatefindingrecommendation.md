# GenerateFindingRecommendation

Creates a recommendation for an unused permissions finding.

## Request Syntax

```nohighlight

POST /recommendation/id?analyzerArn=analyzerArn HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[analyzerArn](#API_GenerateFindingRecommendation_RequestSyntax)**

The [ARN of\
the analyzer](../../../../services/iam/latest/userguide/access-analyzer-getting-started.md#permission-resources) used to generate the finding recommendation.

Pattern: `[^:]*:[^:]*:[^:]*:[^:]*:[^:]*:analyzer/.{1,255}`

Required: Yes

**[id](#API_GenerateFindingRecommendation_RequestSyntax)**

The unique ID for the finding recommendation.

Length Constraints: Minimum length of 1. Maximum length of 2048.

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```

HTTP/1.1 200

```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/accessanalyzer-2019-11-01/generatefindingrecommendation.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/accessanalyzer-2019-11-01/generatefindingrecommendation.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/accessanalyzer-2019-11-01/generatefindingrecommendation.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/accessanalyzer-2019-11-01/generatefindingrecommendation.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/accessanalyzer-2019-11-01/generatefindingrecommendation.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/accessanalyzer-2019-11-01/generatefindingrecommendation.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/accessanalyzer-2019-11-01/generatefindingrecommendation.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/accessanalyzer-2019-11-01/generatefindingrecommendation.md)

- [AWS SDK for Python](../../../../services/goto/boto3/accessanalyzer-2019-11-01/generatefindingrecommendation.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/accessanalyzer-2019-11-01/generatefindingrecommendation.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DeleteArchiveRule

GetAccessPreview
