# GetDataProtectionPolicy

Returns information about a log group data protection policy.

## Request Syntax

```nohighlight

{
   "logGroupIdentifier": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[logGroupIdentifier](#API_GetDataProtectionPolicy_RequestSyntax)**

The name or ARN of the log group that contains the data protection policy that you want to
see.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Pattern: `[\w#+=/:,.@-]*`

Required: Yes

## Response Syntax

```nohighlight

{
   "lastUpdatedTime": number,
   "logGroupIdentifier": "string",
   "policyDocument": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[lastUpdatedTime](#API_GetDataProtectionPolicy_ResponseSyntax)**

The date and time that this policy was most recently updated.

Type: Long

Valid Range: Minimum value of 0.

**[logGroupIdentifier](#API_GetDataProtectionPolicy_ResponseSyntax)**

The log group name or ARN that you specified in your request.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Pattern: `[\w#+=/:,.@-]*`

**[policyDocument](#API_GetDataProtectionPolicy_ResponseSyntax)**

The data protection policy document for this log group.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidParameterException**

A parameter is specified incorrectly.

HTTP Status Code: 400

**OperationAbortedException**

Multiple concurrent requests to update the same resource were in conflict.

HTTP Status Code: 400

**ResourceNotFoundException**

The specified resource does not exist.

HTTP Status Code: 400

**ServiceUnavailableException**

The service cannot complete the request.

HTTP Status Code: 500

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/logs-2014-03-28/getdataprotectionpolicy.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/logs-2014-03-28/getdataprotectionpolicy.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/getdataprotectionpolicy.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/logs-2014-03-28/getdataprotectionpolicy.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/getdataprotectionpolicy.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/logs-2014-03-28/getdataprotectionpolicy.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/logs-2014-03-28/getdataprotectionpolicy.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/logs-2014-03-28/getdataprotectionpolicy.md)

- [AWS SDK for Python](../../../../services/goto/boto3/logs-2014-03-28/getdataprotectionpolicy.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/getdataprotectionpolicy.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FilterLogEvents

GetDelivery

All content copied from https://docs.aws.amazon.com/.
