# UpdateAnalyzer

Modifies the configuration of an existing analyzer.

###### Note

This action is not supported for external access analyzers.

## Request Syntax

```nohighlight

PUT /analyzer/analyzerName HTTP/1.1
Content-type: application/json

{
   "configuration": { ... }
}
```

## URI Request Parameters

The request uses the following URI parameters.

**[analyzerName](#API_UpdateAnalyzer_RequestSyntax)**

The name of the analyzer to modify.

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `[A-Za-z][A-Za-z0-9_.-]*`

Required: Yes

## Request Body

The request accepts the following data in JSON format.

**[configuration](#API_UpdateAnalyzer_RequestSyntax)**

Contains information about the configuration of an analyzer for an AWS organization or
account.

Type: [AnalyzerConfiguration](api-analyzerconfiguration.md) object

**Note:** This object is a Union. Only one member of this object can be specified or returned.

Required: No

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "configuration": { ... }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[configuration](#API_UpdateAnalyzer_ResponseSyntax)**

Contains information about the configuration of an analyzer for an AWS organization or
account.

Type: [AnalyzerConfiguration](api-analyzerconfiguration.md) object

**Note:** This object is a Union. Only one member of this object can be specified or returned.

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/accessanalyzer-2019-11-01/updateanalyzer.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/accessanalyzer-2019-11-01/updateanalyzer.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/accessanalyzer-2019-11-01/updateanalyzer.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/accessanalyzer-2019-11-01/updateanalyzer.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/accessanalyzer-2019-11-01/updateanalyzer.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/accessanalyzer-2019-11-01/updateanalyzer.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/accessanalyzer-2019-11-01/updateanalyzer.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/accessanalyzer-2019-11-01/updateanalyzer.md)

- [AWS SDK for Python](../../../../services/goto/boto3/accessanalyzer-2019-11-01/updateanalyzer.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/accessanalyzer-2019-11-01/updateanalyzer.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

UntagResource

UpdateArchiveRule
