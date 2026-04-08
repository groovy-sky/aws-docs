# CreateAnalyzer

Creates an analyzer for your account.

## Request Syntax

```nohighlight

PUT /analyzer HTTP/1.1
Content-type: application/json

{
   "analyzerName": "string",
   "archiveRules": [
      {
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
   ],
   "clientToken": "string",
   "configuration": { ... },
   "tags": {
      "string" : "string"
   },
   "type": "string"
}
```

## URI Request Parameters

The request does not use any URI parameters.

## Request Body

The request accepts the following data in JSON format.

**[analyzerName](#API_CreateAnalyzer_RequestSyntax)**

The name of the analyzer to create.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `[A-Za-z][A-Za-z0-9_.-]*`

Required: Yes

**[archiveRules](#API_CreateAnalyzer_RequestSyntax)**

Specifies the archive rules to add for the analyzer. Archive rules automatically archive
findings that meet the criteria you define for the rule.

Type: Array of [InlineArchiveRule](api-inlinearchiverule.md) objects

Required: No

**[clientToken](#API_CreateAnalyzer_RequestSyntax)**

A client token.

Type: String

Required: No

**[configuration](#API_CreateAnalyzer_RequestSyntax)**

Specifies the configuration of the analyzer. If the analyzer is an unused access
analyzer, the specified scope of unused access is used for the configuration. If the
analyzer is an internal access analyzer, the specified internal access analysis rules are
used for the configuration.

Type: [AnalyzerConfiguration](api-analyzerconfiguration.md) object

**Note:** This object is a Union. Only one member of this object can be specified or returned.

Required: No

**[tags](#API_CreateAnalyzer_RequestSyntax)**

An array of key-value pairs to apply to the analyzer. You can use the set of Unicode
letters, digits, whitespace, `_`, `.`, `/`,
`=`, `+`, and `-`.

For the tag key, you can specify a value that is 1 to 128 characters in length and
cannot be prefixed with `aws:`.

For the tag value, you can specify a value that is 0 to 256 characters in length.

Type: String to string map

Required: No

**[type](#API_CreateAnalyzer_RequestSyntax)**

The type of analyzer to create. You can create only one analyzer per account per Region.
You can create up to 5 analyzers per organization per Region.

Type: String

Valid Values: `ACCOUNT | ORGANIZATION | ACCOUNT_UNUSED_ACCESS | ORGANIZATION_UNUSED_ACCESS | ACCOUNT_INTERNAL_ACCESS | ORGANIZATION_INTERNAL_ACCESS`

Required: Yes

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "arn": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[arn](#API_CreateAnalyzer_ResponseSyntax)**

The ARN of the analyzer that was created by the request.

Type: String

Pattern: `[^:]*:[^:]*:[^:]*:[^:]*:[^:]*:analyzer/.{1,255}`

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/accessanalyzer-2019-11-01/createanalyzer.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/accessanalyzer-2019-11-01/createanalyzer.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/accessanalyzer-2019-11-01/createanalyzer.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/accessanalyzer-2019-11-01/createanalyzer.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/accessanalyzer-2019-11-01/createanalyzer.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/accessanalyzer-2019-11-01/createanalyzer.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/accessanalyzer-2019-11-01/createanalyzer.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/accessanalyzer-2019-11-01/createanalyzer.md)

- [AWS SDK for Python](../../../../services/goto/boto3/accessanalyzer-2019-11-01/createanalyzer.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/accessanalyzer-2019-11-01/createanalyzer.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CreateAccessPreview

CreateArchiveRule
