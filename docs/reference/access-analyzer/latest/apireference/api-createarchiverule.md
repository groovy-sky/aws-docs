# CreateArchiveRule

Creates an archive rule for the specified analyzer. Archive rules automatically archive
new findings that meet the criteria you define when you create the rule.

To learn about filter keys that you can use to create an archive rule, see [IAM Access Analyzer filter keys](../../../../services/iam/latest/userguide/access-analyzer-reference-filter-keys.md) in the **IAM User Guide**.

## Request Syntax

```nohighlight

PUT /analyzer/analyzerName/archive-rule HTTP/1.1
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
   },
   "ruleName": "string"
}
```

## URI Request Parameters

The request uses the following URI parameters.

**[analyzerName](#API_CreateArchiveRule_RequestSyntax)**

The name of the created analyzer.

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `[A-Za-z][A-Za-z0-9_.-]*`

Required: Yes

## Request Body

The request accepts the following data in JSON format.

**[clientToken](#API_CreateArchiveRule_RequestSyntax)**

A client token.

Type: String

Required: No

**[filter](#API_CreateArchiveRule_RequestSyntax)**

The criteria for the rule.

Type: String to [Criterion](api-criterion.md) object map

Required: Yes

**[ruleName](#API_CreateArchiveRule_RequestSyntax)**

The name of the rule to create.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `[A-Za-z][A-Za-z0-9_.-]*`

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

**ConflictException**

A conflict exception error.

**resourceId**

The ID of the resource.

**resourceType**

The resource type.

HTTP Status Code: 409

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

**ServiceQuotaExceededException**

Service quote met error.

**resourceId**

The resource ID.

**resourceType**

The resource type.

HTTP Status Code: 402

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

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/accessanalyzer-2019-11-01/CreateArchiveRule)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/accessanalyzer-2019-11-01/CreateArchiveRule)

- [AWS SDK for C++](../../../goto/sdkforcpp/accessanalyzer-2019-11-01/createarchiverule.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/accessanalyzer-2019-11-01/createarchiverule.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/accessanalyzer-2019-11-01/createarchiverule.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/accessanalyzer-2019-11-01/createarchiverule.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/accessanalyzer-2019-11-01/createarchiverule.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/accessanalyzer-2019-11-01/createarchiverule.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/accessanalyzer-2019-11-01/CreateArchiveRule)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/accessanalyzer-2019-11-01/createarchiverule.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CreateAnalyzer

DeleteAnalyzer
