# UpdateArchiveRule

Updates the criteria and values for the specified archive rule.

## Request Syntax

```nohighlight

PUT /analyzer/analyzerName/archive-rule/ruleName HTTP/1.1
Content-type: application/json

{
   "clientToken": "string",
   "filter": {
      "string" : {
         "contains": [ "string" ],
         "eq": [ "string" ],
         "exists": boolean,
         "neq": [ "string" ]
      }
   }
}
```

## URI Request Parameters

The request uses the following URI parameters.

**[analyzerName](#API_UpdateArchiveRule_RequestSyntax)**

The name of the analyzer to update the archive rules for.

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `[A-Za-z][A-Za-z0-9_.-]*`

Required: Yes

**[ruleName](#API_UpdateArchiveRule_RequestSyntax)**

The name of the rule to update.

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `[A-Za-z][A-Za-z0-9_.-]*`

Required: Yes

## Request Body

The request accepts the following data in JSON format.

**[clientToken](#API_UpdateArchiveRule_RequestSyntax)**

A client token.

Type: String

Required: No

**[filter](#API_UpdateArchiveRule_RequestSyntax)**

A filter to match for the rules to update. Only rules that match the filter are
updated.

Type: String to [Criterion](api-criterion.md) object map

Required: Yes

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

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/accessanalyzer-2019-11-01/UpdateArchiveRule)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/accessanalyzer-2019-11-01/UpdateArchiveRule)

- [AWS SDK for C++](../../../goto/sdkforcpp/accessanalyzer-2019-11-01/updatearchiverule.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/accessanalyzer-2019-11-01/updatearchiverule.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/accessanalyzer-2019-11-01/updatearchiverule.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/accessanalyzer-2019-11-01/updatearchiverule.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/accessanalyzer-2019-11-01/updatearchiverule.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/accessanalyzer-2019-11-01/updatearchiverule.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/accessanalyzer-2019-11-01/UpdateArchiveRule)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/accessanalyzer-2019-11-01/updatearchiverule.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

UpdateAnalyzer

UpdateFindings
