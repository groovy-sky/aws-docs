# DescribeQueries

Returns a list of CloudWatch Logs Insights queries that are scheduled, running, or have
been run recently in this account. You can request all queries or limit it to queries of a
specific log group or queries with a certain status.

This operation includes both interactive queries started directly by users and automated
queries executed by scheduled query configurations. Scheduled query executions appear in the
results alongside manually initiated queries, providing visibility into all query activity in
your account.

## Request Syntax

```nohighlight

{
   "logGroupName": "string",
   "maxResults": number,
   "nextToken": "string",
   "queryLanguage": "string",
   "status": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[logGroupName](#API_DescribeQueries_RequestSyntax)**

Limits the returned queries to only those for the specified log group.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 512.

Pattern: `[\.\-_/#A-Za-z0-9]+`

Required: No

**[maxResults](#API_DescribeQueries_RequestSyntax)**

Limits the number of returned queries to the specified number.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 1000.

Required: No

**[nextToken](#API_DescribeQueries_RequestSyntax)**

The token for the next set of items to return. The token expires after 24
hours.

Type: String

Length Constraints: Minimum length of 1.

Required: No

**[queryLanguage](#API_DescribeQueries_RequestSyntax)**

Limits the returned queries to only the queries that use the specified query
language.

Type: String

Valid Values: `CWLI | SQL | PPL`

Required: No

**[status](#API_DescribeQueries_RequestSyntax)**

Limits the returned queries to only those that have the specified status. Valid values are
`Cancelled`, `Complete`, `Failed`, `Running`,
and `Scheduled`.

Type: String

Valid Values: `Scheduled | Running | Complete | Failed | Cancelled | Timeout | Unknown`

Required: No

## Response Syntax

```nohighlight

{
   "nextToken": "string",
   "queries": [
      {
         "bytesScanned": number,
         "createTime": number,
         "logGroupName": "string",
         "queryDuration": number,
         "queryId": "string",
         "queryLanguage": "string",
         "queryString": "string",
         "status": "string",
         "userIdentity": "string"
      }
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[nextToken](#API_DescribeQueries_ResponseSyntax)**

The token for the next set of items to return. The token expires after 24
hours.

Type: String

Length Constraints: Minimum length of 1.

**[queries](#API_DescribeQueries_ResponseSyntax)**

The list of queries that match the request.

Type: Array of [QueryInfo](api-queryinfo.md) objects

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

### List the CloudWatch Logs Insights queries for a specific log group

The following example lists the successfully completed queries of the log group
named `MyLogGroup`.

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
X-Amz-Target: Logs_20140328.DescribeQueries
{
   "logGroupName": "MyLogGroup",
   "status": "Completed"
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
    "nextToken": "string",
    "queries": [
        {
            "createTime": 1540923785,
            "logGroupName": "MyLogGroup",
            "queryId": "12ab3456-12ab-123a-789e-1234567890ab",
            "queryString": "filter @message like /Exception/ | stats count(*) as @exceptionCount by date_floor(@timestamp, 5m) | sort @exceptionCount desc",
            "status": "Completed",
            "queryDuration": 5200,
            "bytesScanned": 1048576.0,
            "userIdentity": "arn:aws:iam::123456789012:user/example-user"
        },
        {
            "createTime": 1540025601,
            "logGroupName": "MyLogGroup",
            "queryId": "98ab3456-12ab-123a-789e-1234567890ab",
            "queryString": "stats count(*) by eventSource, eventName, awsRegion",
            "status": "Running",
            "queryDuration": 1500,
            "bytesScanned": 524288.0,
            "userIdentity": "arn:aws:iam::123456789012:user/example-user"
        }
    ]
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/logs-2014-03-28/describequeries.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/logs-2014-03-28/describequeries.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/describequeries.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/logs-2014-03-28/describequeries.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/describequeries.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/logs-2014-03-28/describequeries.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/logs-2014-03-28/describequeries.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/logs-2014-03-28/describequeries.md)

- [AWS SDK for Python](../../../../services/goto/boto3/logs-2014-03-28/describequeries.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/describequeries.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DescribeMetricFilters

DescribeQueryDefinitions

All content copied from https://docs.aws.amazon.com/.
