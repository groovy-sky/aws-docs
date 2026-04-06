# ListArchiveRules

Retrieves a list of archive rules created for the specified analyzer.

## Request Syntax

```nohighlight

GET /analyzer/analyzerName/archive-rule?maxResults=maxResults&nextToken=nextToken HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[analyzerName](#API_ListArchiveRules_RequestSyntax)**

The name of the analyzer to retrieve rules from.

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `[A-Za-z][A-Za-z0-9_.-]*`

Required: Yes

**[maxResults](#API_ListArchiveRules_RequestSyntax)**

The maximum number of results to return in the request.

**[nextToken](#API_ListArchiveRules_RequestSyntax)**

A token used for pagination of results returned.

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "archiveRules": [
      {
         "createdAt": "string",
         "filter": {
            "string" : {
               "contains": [ "string" ],
               "eq": [ "string" ],
               "exists": boolean,
               "neq": [ "string" ]
            }
         },
         "ruleName": "string",
         "updatedAt": "string"
      }
   ],
   "nextToken": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[archiveRules](#API_ListArchiveRules_ResponseSyntax)**

A list of archive rules created for the specified analyzer.

Type: Array of [ArchiveRuleSummary](api-archiverulesummary.md) objects

**[nextToken](#API_ListArchiveRules_ResponseSyntax)**

A token used for pagination of results returned.

Type: String

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

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/accessanalyzer-2019-11-01/ListArchiveRules)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/accessanalyzer-2019-11-01/ListArchiveRules)

- [AWS SDK for C++](../../../goto/sdkforcpp/accessanalyzer-2019-11-01/listarchiverules.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/accessanalyzer-2019-11-01/listarchiverules.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/accessanalyzer-2019-11-01/listarchiverules.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/accessanalyzer-2019-11-01/listarchiverules.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/accessanalyzer-2019-11-01/listarchiverules.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/accessanalyzer-2019-11-01/listarchiverules.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/accessanalyzer-2019-11-01/ListArchiveRules)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/accessanalyzer-2019-11-01/listarchiverules.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ListAnalyzers

ListFindings
