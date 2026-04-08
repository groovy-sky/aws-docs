# PutBearerTokenAuthentication

Enables or disables bearer token authentication for the specified log group. When enabled on a
log group, bearer token authentication is enabled on operations until it is explicitly
disabled.

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

## Request Syntax

```nohighlight

{
   "bearerTokenAuthenticationEnabled": boolean,
   "logGroupIdentifier": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[bearerTokenAuthenticationEnabled](#API_PutBearerTokenAuthentication_RequestSyntax)**

Whether to enable bearer token authentication.

Type: Boolean

Required: Yes

Type: Boolean

Required: Yes

**[logGroupIdentifier](#API_PutBearerTokenAuthentication_RequestSyntax)**

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

This example illustrates one usage of PutBearerTokenAuthentication.

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
          X-Amz-Target: Logs_20140328.PutBearerTokenAuthentication
          {
          "logGroupIdentifier": "my-log-group",
          "bearerTokenAuthenticationEnabled": true
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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/logs-2014-03-28/putbearertokenauthentication.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/logs-2014-03-28/putbearertokenauthentication.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/putbearertokenauthentication.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/logs-2014-03-28/putbearertokenauthentication.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/putbearertokenauthentication.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/logs-2014-03-28/putbearertokenauthentication.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/logs-2014-03-28/putbearertokenauthentication.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/logs-2014-03-28/putbearertokenauthentication.md)

- [AWS SDK for Python](../../../../services/goto/boto3/logs-2014-03-28/putbearertokenauthentication.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/putbearertokenauthentication.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

PutAccountPolicy

PutDataProtectionPolicy
