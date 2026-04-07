# PutLogGroupDeletionProtection

Enables or disables deletion protection for the specified log group. When enabled on a
log group, deletion protection blocks all deletion operations until it is explicitly
disabled.

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

## Request Syntax

```nohighlight

{
   "deletionProtectionEnabled": boolean,
   "logGroupIdentifier": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[deletionProtectionEnabled](#API_PutLogGroupDeletionProtection_RequestSyntax)**

Whether to enable deletion protection.

Type: Boolean

Required: Yes

Type: Boolean

Required: Yes

**[logGroupIdentifier](#API_PutLogGroupDeletionProtection_RequestSyntax)**

The name or ARN of the log group.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 512.

Pattern: `[\.\-_/#A-Za-z0-9]+`

Required: Yes

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Pattern: `[\w#+=/:,.@-]*`

Required: Yes

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**AccessDeniedException**

You don't have sufficient permissions to perform this action.

HTTP Status Code: 400

**InvalidOperationException**

The operation is not valid on the specified resource.

HTTP Status Code: 400

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

## Examples

### Sample Request

This example illustrates one usage of PutLogGroupDeletionProtection.

#### Sample Request

```

POST / HTTP/1.1
Host: logs.<region>.<domain>
X-Amz-Date: <DATE>
Authorization: AWS4-HMAC-SHA256 Credential=<Credential>, SignedHeaders=content-type;date;host;user-agent;x-amz-date;x-amz-target;x-amzn-requestid, Signature=<Signature>
User-Agent: <UserAgentString>
Accept: application/json
Content-Type: application/x-amz-json-1.1
Content-Length: <PayloadSizeBytes>
Connection: Keep-Alive
X-Amz-Target: Logs_20140328.PutLogGroupDeletionProtection
{
  "logGroupIdentifier": "my-log-group",
  "deletionProtectionEnabled": true
}
```

#### Sample Response

```

HTTP/1.1 200 OK
x-amzn-RequestId: <RequestId>
Content-Type: application/x-amz-json-1.1
Content-Length: 0
Date: <Date>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/logs-2014-03-28/PutLogGroupDeletionProtection)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/logs-2014-03-28/PutLogGroupDeletionProtection)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/logs-2014-03-28/PutLogGroupDeletionProtection)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/logs-2014-03-28/PutLogGroupDeletionProtection)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/logs-2014-03-28/PutLogGroupDeletionProtection)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/logs-2014-03-28/PutLogGroupDeletionProtection)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/logs-2014-03-28/PutLogGroupDeletionProtection)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/logs-2014-03-28/PutLogGroupDeletionProtection)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/logs-2014-03-28/PutLogGroupDeletionProtection)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/logs-2014-03-28/PutLogGroupDeletionProtection)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

PutLogEvents

PutMetricFilter
