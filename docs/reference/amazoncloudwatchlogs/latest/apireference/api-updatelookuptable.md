# UpdateLookupTable

Updates an existing lookup table by replacing all of its CSV content. After the update
completes, queries that use this table will use the new data.

This is a full replacement operation. All existing content is replaced with the new CSV
data.

## Request Syntax

```nohighlight

{
   "description": "string",
   "kmsKeyId": "string",
   "lookupTableArn": "string",
   "tableBody": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[description](#API_UpdateLookupTable_RequestSyntax)**

An updated description of the lookup table.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1024.

Required: No

**[kmsKeyId](#API_UpdateLookupTable_RequestSyntax)**

The ARN of the AWS KMS key to use to encrypt the lookup table data. You can
use this parameter to add, update, or remove the KMS key. To remove the KMS key and use an
AWS-owned key instead, specify an empty string.

Type: String

Length Constraints: Maximum length of 256.

Required: No

**[lookupTableArn](#API_UpdateLookupTable_RequestSyntax)**

The ARN of the lookup table to update.

Type: String

Required: Yes

**[tableBody](#API_UpdateLookupTable_RequestSyntax)**

The new CSV content to replace the existing data. The first row must be a header row
with column names. The content must use UTF-8 encoding and not exceed 10 MB.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 10485760.

Required: Yes

## Response Syntax

```nohighlight

{
   "lastUpdatedTime": number,
   "lookupTableArn": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[lastUpdatedTime](#API_UpdateLookupTable_ResponseSyntax)**

The time when the lookup table was last updated, expressed as the number of
milliseconds after `Jan 1, 1970 00:00:00 UTC`.

Type: Long

Valid Range: Minimum value of 0.

**[lookupTableArn](#API_UpdateLookupTable_ResponseSyntax)**

The ARN of the lookup table that was updated.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**AccessDeniedException**

You don't have sufficient permissions to perform this action.

HTTP Status Code: 400

**InvalidParameterException**

A parameter is specified incorrectly.

HTTP Status Code: 400

**ResourceNotFoundException**

The specified resource does not exist.

HTTP Status Code: 400

**ServiceUnavailableException**

The service cannot complete the request.

HTTP Status Code: 500

**ValidationException**

One of the parameters for the request is not valid.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/logs-2014-03-28/UpdateLookupTable)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/logs-2014-03-28/UpdateLookupTable)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/logs-2014-03-28/UpdateLookupTable)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/logs-2014-03-28/UpdateLookupTable)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/logs-2014-03-28/UpdateLookupTable)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/logs-2014-03-28/UpdateLookupTable)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/logs-2014-03-28/UpdateLookupTable)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/logs-2014-03-28/UpdateLookupTable)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/logs-2014-03-28/UpdateLookupTable)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/logs-2014-03-28/UpdateLookupTable)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

UpdateLogAnomalyDetector

UpdateScheduledQuery
