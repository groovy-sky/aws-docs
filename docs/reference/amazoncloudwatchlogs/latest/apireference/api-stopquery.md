# StopQuery

Stops a CloudWatch Logs Insights query that is in progress. If the query has already
ended, the operation returns an error indicating that the specified query is not
running.

This operation can be used to cancel both interactive queries and individual scheduled
query executions. When used with scheduled queries, `StopQuery` cancels only the
specific execution identified by the query ID, not the scheduled query configuration
itself.

## Request Syntax

```nohighlight

{
   "queryId": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[queryId](#API_StopQuery_RequestSyntax)**

The ID number of the query to stop. To find this ID number, use
`DescribeQueries`.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Required: Yes

## Response Syntax

```nohighlight

{
   "success": boolean
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[success](#API_StopQuery_ResponseSyntax)**

This is true if the query was stopped by the `StopQuery` operation.

Type: Boolean

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidParameterException**

A parameter is specified incorrectly.

HTTP Status Code: 400

**ResourceNotFoundException**

The specified resource does not exist.

HTTP Status Code: 400

**ServiceUnavailableException**

The service cannot complete the request.

HTTP Status Code: 500

## Examples

### Stop a query that is currently running

The following example stops the specified query, if it is currently
running.

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
X-Amz-Target: Logs_20140328.StopQuery
{
   "queryId": "12ab3456-12ab-123a-789e-1234567890ab"
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
   "success": True
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/logs-2014-03-28/stopquery.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/logs-2014-03-28/stopquery.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/stopquery.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/logs-2014-03-28/stopquery.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/stopquery.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/logs-2014-03-28/stopquery.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/logs-2014-03-28/stopquery.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/logs-2014-03-28/stopquery.md)

- [AWS SDK for Python](../../../../services/goto/boto3/logs-2014-03-28/stopquery.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/stopquery.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StartQuery

TagLogGroup

All content copied from https://docs.aws.amazon.com/.
