# CreateLogStream

Creates a log stream for the specified log group. A log stream is a sequence of log
events that originate from a single source, such as an application instance or a resource that
is being monitored.

There is no limit on the number of log streams that you can create for a log group.
There is a limit of 50 TPS on `CreateLogStream` operations, after which
transactions are throttled.

You must use the following guidelines when naming a log stream:

- Log stream names must be unique within the log group.

- Log stream names can be between 1 and 512 characters long.

- Don't use ':' (colon) or '\*' (asterisk) characters.

## Request Syntax

```nohighlight

{
   "logGroupName": "string",
   "logStreamName": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[logGroupName](#API_CreateLogStream_RequestSyntax)**

The name of the log group.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 512.

Pattern: `[\.\-_/#A-Za-z0-9]+`

Required: Yes

**[logStreamName](#API_CreateLogStream_RequestSyntax)**

The name of the log stream.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 512.

Pattern: `[^:*]*`

Required: Yes

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidParameterException**

A parameter is specified incorrectly.

HTTP Status Code: 400

**ResourceAlreadyExistsException**

The specified resource already exists.

HTTP Status Code: 400

**ResourceNotFoundException**

The specified resource does not exist.

HTTP Status Code: 400

**ServiceUnavailableException**

The service cannot complete the request.

HTTP Status Code: 500

## Examples

### To create a log stream

The following example creates a log stream for the specified log group.

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
X-Amz-Target: Logs_20140328.CreateLogStream
{
  "logGroupName": "my-log-group",
  "logStreamName": "my-log-stream"
}
```

#### Sample Response

```

HTTP/1.1 200 OK
x-amzn-RequestId: <RequestId>
Content-Type: application/x-amz-json-1.1
Content-Length: <PayloadSizeBytes>
Date: <Date>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/logs-2014-03-28/createlogstream.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/logs-2014-03-28/createlogstream.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/createlogstream.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/logs-2014-03-28/createlogstream.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/createlogstream.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/logs-2014-03-28/createlogstream.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/logs-2014-03-28/createlogstream.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/logs-2014-03-28/createlogstream.md)

- [AWS SDK for Python](../../../../services/goto/boto3/logs-2014-03-28/createlogstream.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/createlogstream.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CreateLogGroup

CreateLookupTable
