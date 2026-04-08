# FilterLogEvents

Lists log events from the specified log group. You can list all the log events or
filter the results using one or more of the following:

- A filter pattern

- A time range

- The log stream name, or a log stream name prefix that matches multiple log
streams

You must have the `logs:FilterLogEvents` permission to perform this
operation.

You can specify the log group to search by using either `logGroupIdentifier` or
`logGroupName`. You must include one of these two parameters, but you can't
include both.

`FilterLogEvents` is a paginated operation. Each page returned can contain up
to 1 MB of log events or up to 10,000 log events. A returned page might only be partially
full, or even empty. For example, if the result of a query would return 15,000 log events, the
first page isn't guaranteed to have 10,000 log events even if they all fit into 1 MB.

Partially full or empty pages don't necessarily mean that pagination is finished. If the
results include a `nextToken`, there might be more log events available. You can
return these additional log events by providing the nextToken in a subsequent
`FilterLogEvents` operation. If the results don't include a
`nextToken`, then pagination is finished.

Specifying the `limit` parameter only guarantees that a single page doesn't
return more log events than the specified limit, but it might return fewer events than the
limit. This is the expected API behavior.

The returned log events are sorted by event timestamp, the timestamp when the event was
ingested by CloudWatch Logs, and the ID of the `PutLogEvents` request.

If you are using CloudWatch cross-account observability, you can use this operation
in a monitoring account and view data from the linked source accounts. For more information,
see [CloudWatch cross-account observability](../../../../services/amazoncloudwatch/latest/monitoring/cloudwatch-unified-cross-account.md).

###### Note

If you are using [log\
transformation](../../../../services/amazoncloudwatch/latest/logs/cloudwatch-logs-transformation.md), the `FilterLogEvents` operation returns only the
original versions of log events, before they were transformed. To view the transformed
versions, you must use a [CloudWatch Logs\
query.](../../../../services/amazoncloudwatch/latest/logs/analyzinglogdata.md)

## Request Syntax

```nohighlight

{
   "endTime": number,
   "filterPattern": "string",
   "interleaved": boolean,
   "limit": number,
   "logGroupIdentifier": "string",
   "logGroupName": "string",
   "logStreamNamePrefix": "string",
   "logStreamNames": [ "string" ],
   "nextToken": "string",
   "startTime": number,
   "unmask": boolean
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[endTime](#API_FilterLogEvents_RequestSyntax)**

The end of the time range, expressed as the number of milliseconds after `Jan 1,
        1970 00:00:00 UTC`. Events with a timestamp later than this time are not
returned.

Type: Long

Valid Range: Minimum value of 0.

Required: No

**[filterPattern](#API_FilterLogEvents_RequestSyntax)**

The filter pattern to use. For more information, see [Filter and Pattern\
Syntax](../../../../services/amazoncloudwatch/latest/logs/filterandpatternsyntax.md).

If not provided, all the events are matched.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1024.

Required: No

**[interleaved](#API_FilterLogEvents_RequestSyntax)**

_This parameter has been deprecated._

If the value is true, the operation attempts to provide responses that contain events
from multiple log streams within the log group, interleaved in a single response. If the value
is false, all the matched log events in the first log stream are searched first, then those in
the next log stream, and so on.

**Important** As of June 17, 2019, this parameter is
ignored and the value is assumed to be true. The response from this operation always
interleaves events from multiple log streams within a log group.

Type: Boolean

Required: No

**[limit](#API_FilterLogEvents_RequestSyntax)**

The maximum number of events to return. The default is 10,000 events.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 10000.

Required: No

**[logGroupIdentifier](#API_FilterLogEvents_RequestSyntax)**

Specify either the name or ARN of the log group to view log events from. If the log group
is in a source account and you are using a monitoring account, you must use the log group
ARN.

###### Note

You must include either `logGroupIdentifier` or `logGroupName`,
but not both.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Pattern: `[\w#+=/:,.@-]*`

Required: No

**[logGroupName](#API_FilterLogEvents_RequestSyntax)**

The name of the log group to search.

###### Note

You must include either `logGroupIdentifier` or `logGroupName`,
but not both.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 512.

Pattern: `[\.\-_/#A-Za-z0-9]+`

Required: No

**[logStreamNamePrefix](#API_FilterLogEvents_RequestSyntax)**

Filters the results to include only events from log streams that have names starting with
this prefix.

If you specify a value for both `logStreamNamePrefix` and
`logStreamNames`, the action returns an `InvalidParameterException`
error.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 512.

Pattern: `[^:*]*`

Required: No

**[logStreamNames](#API_FilterLogEvents_RequestSyntax)**

Filters the results to only logs from the log streams in this list.

If you specify a value for both `logStreamNames` and
`logStreamNamePrefix`, the action returns an
`InvalidParameterException` error.

Type: Array of strings

Array Members: Minimum number of 1 item. Maximum number of 100 items.

Length Constraints: Minimum length of 1. Maximum length of 512.

Pattern: `[^:*]*`

Required: No

**[nextToken](#API_FilterLogEvents_RequestSyntax)**

The token for the next set of events to return. (You received this token from a
previous call.)

Type: String

Length Constraints: Minimum length of 1.

Required: No

**[startTime](#API_FilterLogEvents_RequestSyntax)**

The start of the time range, expressed as the number of milliseconds after `Jan 1,
        1970 00:00:00 UTC`. Events with a timestamp before this time are not
returned.

Type: Long

Valid Range: Minimum value of 0.

Required: No

**[unmask](#API_FilterLogEvents_RequestSyntax)**

Specify `true` to display the log event fields with all sensitive data unmasked
and visible. The default is `false`.

To use this operation with this parameter, you must be signed into an account with the
`logs:Unmask` permission.

Type: Boolean

Required: No

## Response Syntax

```nohighlight

{
   "events": [
      {
         "eventId": "string",
         "ingestionTime": number,
         "logStreamName": "string",
         "message": "string",
         "timestamp": number
      }
   ],
   "nextToken": "string",
   "searchedLogStreams": [
      {
         "logStreamName": "string",
         "searchedCompletely": boolean
      }
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[events](#API_FilterLogEvents_ResponseSyntax)**

The matched events.

Type: Array of [FilteredLogEvent](api-filteredlogevent.md) objects

**[nextToken](#API_FilterLogEvents_ResponseSyntax)**

The token to use when requesting the next set of items. The token expires after 24
hours.

If the results don't include a `nextToken`, then pagination is finished.

Type: String

Length Constraints: Minimum length of 1.

**[searchedLogStreams](#API_FilterLogEvents_ResponseSyntax)**

**Important** As of May 15, 2020, this parameter is no longer
supported. This parameter returns an empty list.

Indicates which log streams have been searched and whether each has been searched
completely.

Type: Array of [SearchedLogStream](api-searchedlogstream.md) objects

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

### To list the events in a log group that contain a pattern

The following example lists the events for the specified log group that contain
`ERROR`.

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
X-Amz-Target: Logs_20140328.FilterLogEvents
{
  "logGroupName": "my-log-group",
  "filterPattern": "ERROR"
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
  "events": [
    {
      "ingestionTime": 1396035394997,
      "timestamp": 1396035378988,
      "message": "ERROR Event 1",
      "logStreamName": "my-log-stream-1",
      "eventId": "31132629274945519779805322857203735586714454643391594505"
    },
    {
      "ingestionTime": 1396035394997,
      "timestamp": 1396035378988,
      "message": "ERROR Event 2",
      "logStreamName": "my-log-stream-2",
      "eventId": "31132629274945519779805322857203735586814454643391594505"
    },
    {
      "ingestionTime": 1396035394997,
      "timestamp": 1396035378989,
      "message": "ERROR Event 3",
      "logStreamName": "my-log-stream-3"
      "eventId": "31132629274945519779805322857203735586824454643391594505"
    }
  ],
  "searchedLogStreams": [
    {
      "searchedCompletely": true,
      "logStreamName": "my-log-stream-1"
    },
    {
      "searchedCompletely": true,
      "logStreamName": "my-log-stream-2"
    },
    {
      "searchedCompletely": false,
      "logStreamName": "my-log-stream-3"
    },
  ],
  "nextToken": "ZNUEPl7FcQuXbIH4Swk9D9eFu2XBg-ijZIZlvzz4ea9zZRjw-MMtQtvcoMdmq4T29K7Q6Y1e_KvyfpcT_f_tUw"
}
```

### Example

The following example lists the events for the specified log group that contain
`ERROR`.

#### Sample Request

```

{
  "logGroupIdentifier": "arn:aws:logs:us-east-1:123456789012:log-group:monitoring-logGroup-1234:*",
  "filterPattern": "ERROR"
}
```

#### Sample Response

```

{
  "events": [
  {
      "ingestionTime": 1396035394997,
      "timestamp": 1396035378988,
      "message": "ERROR Event 1",
      "logStreamName": "my-log-stream-1",
      "eventId": "31132629274945519779805322857203735586714454643391594505"
  },
  {
     "ingestionTime": 1396035394997,
      "timestamp": 1396035378988,
      "message": "ERROR Event 2",
      "logStreamName": "my-log-stream-2",
      "eventId": "31132629274945519779805322857203735586814454643391594505"
  },
  {
      "ingestionTime": 1396035394997,
      "timestamp": 1396035378989,
      "message": "ERROR Event 3",
      "logStreamName": "my-log-stream-3"
      "eventId": "31132629274945519779805322857203735586824454643391594505"
   } ],
   "searchedLogStreams": [
   {
      "searchedCompletely": true,
      "logStreamName": "my-log-stream-1"
    },
    {
      "searchedCompletely": true,
      "logStreamName": "my-log-stream-2"
    },
    {
      "searchedCompletely": false,
      "logStreamName": "my-log-stream-3"
    }
    ],
    "nextToken": "ZNUEPl7FcQuXbIH4Swk9D9eFu2XBg-ijZIZlvzz4ea9zZRjw-
MMtQtvcoMdmq4T29K7Q6Y1e_KvyfpcT_f_tUw"
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/logs-2014-03-28/filterlogevents.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/logs-2014-03-28/filterlogevents.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/filterlogevents.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/logs-2014-03-28/filterlogevents.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/filterlogevents.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/logs-2014-03-28/filterlogevents.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/logs-2014-03-28/filterlogevents.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/logs-2014-03-28/filterlogevents.md)

- [AWS SDK for Python](../../../../services/goto/boto3/logs-2014-03-28/filterlogevents.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/filterlogevents.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DisassociateSourceFromS3TableIntegration

GetDataProtectionPolicy
