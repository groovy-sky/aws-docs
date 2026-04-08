# ListLogGroupsForQuery

Returns a list of the log groups that were analyzed during a single CloudWatch Logs
Insights query. This can be useful for queries that use log group name prefixes or the
`filterIndex` command, because the log groups are dynamically selected in these
cases.

For more information about field indexes, see [Create field indexes\
to improve query performance and reduce costs](../../../../services/amazoncloudwatch/latest/logs/cloudwatchlogs-field-indexing.md).

## Request Syntax

```nohighlight

{
   "maxResults": number,
   "nextToken": "string",
   "queryId": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[maxResults](#API_ListLogGroupsForQuery_RequestSyntax)**

Limits the number of returned log groups to the specified number.

Type: Integer

Valid Range: Minimum value of 50. Maximum value of 500.

Required: No

**[nextToken](#API_ListLogGroupsForQuery_RequestSyntax)**

The token for the next set of items to return. The token expires after 24
hours.

Type: String

Length Constraints: Minimum length of 1.

Required: No

**[queryId](#API_ListLogGroupsForQuery_RequestSyntax)**

The ID of the query to use. This query ID is from the response to your [StartQuery](api-startquery.md) operation.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Required: Yes

## Response Syntax

```nohighlight

{
   "logGroupIdentifiers": [ "string" ],
   "nextToken": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[logGroupIdentifiers](#API_ListLogGroupsForQuery_ResponseSyntax)**

An array of the names and ARNs of the log groups that were processed in the query.

Type: Array of strings

Length Constraints: Minimum length of 1. Maximum length of 2048.

Pattern: `[\w#+=/:,.@-]*`

**[nextToken](#API_ListLogGroupsForQuery_ResponseSyntax)**

The token for the next set of items to return. The token expires after 24
hours.

Type: String

Length Constraints: Minimum length of 1.

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

## Examples

### To list the log groups that were analyzed during a specific query

The following examplereturns the log groups that were analyzed during the query
with the `71bacb5a-30f1-4ed6-9959-2797EXAMPLE` ID.

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
X-Amz-Target: Logs_20140328.ListLogGroupsForQuery
{
  "queryId": "71bacb5a-30f1-4ed6-9959-2797EXAMPLE"
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
  "logGroupIdentifiers": [
    "arn:aws:logs:us-east-1:112233445566:log-group:/aws/lambda/applogs",
    "arn:aws:logs:us-east-1:112233445566:log-group:/aws/lambda/servicelogs",
    "arn:aws:logs:us-east-1:112233445566:log-group:/aws/lambda/errorlogs"
  ]
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/logs-2014-03-28/listloggroupsforquery.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/logs-2014-03-28/listloggroupsforquery.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/listloggroupsforquery.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/logs-2014-03-28/listloggroupsforquery.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/listloggroupsforquery.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/logs-2014-03-28/listloggroupsforquery.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/logs-2014-03-28/listloggroupsforquery.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/logs-2014-03-28/listloggroupsforquery.md)

- [AWS SDK for Python](../../../../services/goto/boto3/logs-2014-03-28/listloggroupsforquery.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/listloggroupsforquery.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ListLogGroups

ListScheduledQueries
