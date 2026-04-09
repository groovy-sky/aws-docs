# GetQueryResults

Returns the results from the specified query.

Only the fields requested in the query are returned, along with a `@ptr` field,
which is the identifier for the log record. You can use the value of `@ptr` in a
[GetLogRecord](api-getlogrecord.md)
operation to get the full log record.

`GetQueryResults` does not start running a query. To run a query, use [StartQuery](api-startquery.md). For more information about how long results of previous queries are
available, see [CloudWatch Logs\
quotas](../../../../services/amazoncloudwatch/latest/logs/cloudwatch-limits-cwl.md).

If the value of the `Status` field in the output is `Running`, this
operation returns only partial results. If you see a value of `Scheduled` or
`Running` for the status, you can retry the operation later to see the final
results.

This operation is used both for retrieving results from interactive queries and from
automated scheduled query executions. Scheduled queries use `GetQueryResults`
internally to retrieve query results for processing and delivery to configured
destinations.

If you are using CloudWatch cross-account observability, you can use this operation
in a monitoring account to start queries in linked source accounts. For more information, see
[CloudWatch cross-account observability](../../../../services/amazoncloudwatch/latest/monitoring/cloudwatch-unified-cross-account.md).

## Request Syntax

```nohighlight

{
   "queryId": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[queryId](#API_GetQueryResults_RequestSyntax)**

The ID number of the query.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Required: Yes

## Response Syntax

```nohighlight

{
   "encryptionKey": "string",
   "queryLanguage": "string",
   "results": [
      [
         {
            "field": "string",
            "value": "string"
         }
      ]
   ],
   "statistics": {
      "bytesScanned": number,
      "estimatedBytesSkipped": number,
      "estimatedRecordsSkipped": number,
      "logGroupsScanned": number,
      "recordsMatched": number,
      "recordsScanned": number
   },
   "status": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[encryptionKey](#API_GetQueryResults_ResponseSyntax)**

If you associated an AWS KMS key with the CloudWatch Logs Insights
query results in this account, this field displays the ARN of the key that's used to encrypt
the query results when [StartQuery](api-startquery.md) stores
them.

Type: String

Length Constraints: Maximum length of 256.

**[queryLanguage](#API_GetQueryResults_ResponseSyntax)**

The query language used for this query. For more information about the query languages
that CloudWatch Logs supports, see [Supported query\
languages](../../../../services/amazoncloudwatch/latest/logs/cwl-analyzelogdata-languages.md).

Type: String

Valid Values: `CWLI | SQL | PPL`

**[results](#API_GetQueryResults_ResponseSyntax)**

The log events that matched the query criteria during the most recent time it ran.

The `results` value is an array of arrays. Each log event is one object in the
top-level array. Each of these log event objects is an array of
`field`/ `value` pairs.

Type: Array of arrays of [ResultField](api-resultfield.md) objects

**[statistics](#API_GetQueryResults_ResponseSyntax)**

Includes the number of log events scanned by the query, the number of log events that
matched the query criteria, and the total number of bytes in the scanned log events. These
values reflect the full raw results of the query.

Type: [QueryStatistics](api-querystatistics.md) object

**[status](#API_GetQueryResults_ResponseSyntax)**

The status of the most recent running of the query. Possible values are
`Cancelled`, `Complete`, `Failed`, `Running`,
`Scheduled`, `Timeout`, and `Unknown`.

Queries time out after 60 minutes of runtime. To avoid having your queries time out,
reduce the time range being searched or partition your query into a number of queries.

Type: String

Valid Values: `Scheduled | Running | Complete | Failed | Cancelled | Timeout | Unknown`

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

### Get results from a recent query

The following returns the results from a specified query.

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
X-Amz-Target: Logs_20140328.GetQueryResults
{
   "queryId": "12ab3456-12ab-123a-789e-1234567890ab"
}
```

#### Sample Response

```nohighlight

HTTP/1.1 200 OK
x-amzn-RequestId: <RequestId>
Content-Type: application/x-amz-json-1.1
Content-Length: <PayloadSizeBytes>
Date: <Date>
{
    "results": [
        [
            {
                "field": "LogEvent1-field1-name",
                "value": "LogEvent1-field1-value"
            },
            {
                "field": "LogEvent1-field2-name",
                "value": "LogEvent1-field2-value"
            },
            ...
            {
                "field": "LogEvent1-fieldX-name",
                "value": "LogEvent1-fieldX-value"
            }
        ],
        [
            {
                "field": "LogEvent2-field1-name",
                "value": "LogEvent2-field1-value"
            },
            {
                "field": "LogEvent2-field2-name",
                "value": "LogEvent2-field2-value"
            },
            ...
            {
                "field": "LogEvent2-fieldX-name",
                "value": "LogEvent2-fieldX-value"
            }
        ],
        [
            {
                "field": "LogEventZ-field1-name",
                "value": "LogEventZ-field1-value"
            },
            {
                "field": "LogEventZ-field2-name",
                "value": "LogEventZ-field2-value"
            },
            ...
            {
                "field": "LogEventZ-fieldX-name",
                "value": "LogEventZ-fieldX-value"
            }
        ]
    ],
    "statistics": {
        "bytesScanned": 81349723,
        "recordsMatched": 360851,
        "recordsScanned": 610956
    },
    "status": "Complete"
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/logs-2014-03-28/getqueryresults.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/logs-2014-03-28/getqueryresults.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/getqueryresults.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/logs-2014-03-28/getqueryresults.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/getqueryresults.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/logs-2014-03-28/getqueryresults.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/logs-2014-03-28/getqueryresults.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/logs-2014-03-28/getqueryresults.md)

- [AWS SDK for Python](../../../../services/goto/boto3/logs-2014-03-28/getqueryresults.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/getqueryresults.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetLookupTable

GetScheduledQuery

All content copied from https://docs.aws.amazon.com/.
