# GetLogRecord

Retrieves all of the fields and values of a single log event. All fields are retrieved,
even if the original query that produced the `logRecordPointer` retrieved only a
subset of fields. Fields are returned as field name/field value pairs.

The full unparsed log event is returned within `@message`.

## Request Syntax

```nohighlight

{
   "logRecordPointer": "string",
   "unmask": boolean
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[logRecordPointer](#API_GetLogRecord_RequestSyntax)**

The pointer corresponding to the log event record you want to retrieve. You get this from
the response of a `GetQueryResults` operation. In that response, the value of the
`@ptr` field for a log event is the value to use as `logRecordPointer`
to retrieve that complete log event record.

Type: String

Required: Yes

**[unmask](#API_GetLogRecord_RequestSyntax)**

Specify `true` to display the log event fields with all sensitive data unmasked
and visible. The default is `false`.

To use this operation with this parameter, you must be signed into an account with the
`logs:Unmask` permission.

Type: Boolean

Required: No

## Response Syntax

```nohighlight

{
   "logRecord": {
      "string" : "string"
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[logRecord](#API_GetLogRecord_ResponseSyntax)**

The requested log event, as a JSON string.

Type: String to string map

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidParameterException**

A parameter is specified incorrectly.

HTTP Status Code: 400

**LimitExceededException**

You have reached the maximum number of resources that can be created.

HTTP Status Code: 400

**ResourceNotFoundException**

The specified resource does not exist.

HTTP Status Code: 400

**ServiceUnavailableException**

The service cannot complete the request.

HTTP Status Code: 500

## Examples

### To retrieve all fields for a specified log event

The following example retrieves the fields for a specified log event.

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
X-Amz-Target: Logs_20140328.GetLogRecord
{
   "logRecordPointer": "123456789"
}
```

#### Sample Response

```

HTTP/1.1 200 OK
x-amzn-RequestId: <RequestId>
Content-Type: application/x-amz-json-1.1
Content-Length: <PayloadSizeBytes>
Date: <Date>
{
   "logRecord": {
      "@timestamp" : "1536857812",
      "@message" : "123456789012 eni-1234567890abcde123 6 33 ACCEPT"
      "accountId" : "123456789012",
      "interfaceId" : "eni-1234567890abcde123",
      "protocol" : "6",
      "packets" : "33",
      "action" : "ACCEPT"
   }
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/logs-2014-03-28/getlogrecord.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/logs-2014-03-28/getlogrecord.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/getlogrecord.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/logs-2014-03-28/getlogrecord.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/getlogrecord.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/logs-2014-03-28/getlogrecord.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/logs-2014-03-28/getlogrecord.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/logs-2014-03-28/getlogrecord.md)

- [AWS SDK for Python](../../../../services/goto/boto3/logs-2014-03-28/getlogrecord.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/getlogrecord.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetLogObject

GetLookupTable

All content copied from https://docs.aws.amazon.com/.
