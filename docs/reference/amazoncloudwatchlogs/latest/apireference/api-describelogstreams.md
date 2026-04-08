# DescribeLogStreams

Lists the log streams for the specified log group. You can list all the log streams or
filter the results by prefix. You can also control how the results are ordered.

You can specify the log group to search by using either `logGroupIdentifier` or
`logGroupName`. You must include one of these two parameters, but you can't
include both.

This operation has a limit of 25 transactions per second, after which transactions are
throttled.

If you are using CloudWatch cross-account observability, you can use this operation
in a monitoring account and view data from the linked source accounts. For more information,
see [CloudWatch cross-account observability](../../../../services/amazoncloudwatch/latest/monitoring/cloudwatch-unified-cross-account.md).

## Request Syntax

```nohighlight

{
   "descending": boolean,
   "limit": number,
   "logGroupIdentifier": "string",
   "logGroupName": "string",
   "logStreamNamePrefix": "string",
   "nextToken": "string",
   "orderBy": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[descending](#API_DescribeLogStreams_RequestSyntax)**

If the value is true, results are returned in descending order. If the value is to
false, results are returned in ascending order. The default value is false.

Type: Boolean

Required: No

**[limit](#API_DescribeLogStreams_RequestSyntax)**

The maximum number of items returned. If you don't specify a value, the default is up
to 50 items.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 50.

Required: No

**[logGroupIdentifier](#API_DescribeLogStreams_RequestSyntax)**

Specify either the name or ARN of the log group to view. If the log group is in a source
account and you are using a monitoring account, you must use the log group ARN.

###### Note

You must include either `logGroupIdentifier` or `logGroupName`,
but not both.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Pattern: `[\w#+=/:,.@-]*`

Required: No

**[logGroupName](#API_DescribeLogStreams_RequestSyntax)**

The name of the log group.

###### Note

You must include either `logGroupIdentifier` or `logGroupName`,
but not both.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 512.

Pattern: `[\.\-_/#A-Za-z0-9]+`

Required: No

**[logStreamNamePrefix](#API_DescribeLogStreams_RequestSyntax)**

The prefix to match.

If `orderBy` is `LastEventTime`, you cannot specify this
parameter.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 512.

Pattern: `[^:*]*`

Required: No

**[nextToken](#API_DescribeLogStreams_RequestSyntax)**

The token for the next set of items to return. (You received this token from a previous
call.)

Type: String

Length Constraints: Minimum length of 1.

Required: No

**[orderBy](#API_DescribeLogStreams_RequestSyntax)**

If the value is `LogStreamName`, the results are ordered by log stream name.
If the value is `LastEventTime`, the results are ordered by the event time. The
default value is `LogStreamName`.

If you order the results by event time, you cannot specify the
`logStreamNamePrefix` parameter.

`lastEventTimestamp` represents the time of the most recent log event in the
log stream in CloudWatch Logs. This number is expressed as the number of milliseconds after
`Jan 1, 1970 00:00:00 UTC`. `lastEventTimestamp` updates on an
eventual consistency basis. It typically updates in less than an hour from ingestion, but in
rare situations might take longer.

Type: String

Valid Values: `LogStreamName | LastEventTime`

Required: No

## Response Syntax

```nohighlight

{
   "logStreams": [
      {
         "arn": "string",
         "creationTime": number,
         "firstEventTimestamp": number,
         "lastEventTimestamp": number,
         "lastIngestionTime": number,
         "logStreamName": "string",
         "storedBytes": number,
         "uploadSequenceToken": "string"
      }
   ],
   "nextToken": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[logStreams](#API_DescribeLogStreams_ResponseSyntax)**

The log streams.

Type: Array of [LogStream](api-logstream.md) objects

**[nextToken](#API_DescribeLogStreams_ResponseSyntax)**

The token for the next set of items to return. The token expires after 24
hours.

Type: String

Length Constraints: Minimum length of 1.

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

### To list the log streams for a log group

The following example lists the log streams associated with the specified log
group.

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
X-Amz-Target: Logs_20140328.DescribeLogStreams
{
  "logGroupName": "my-log-group"
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
  "logStreams": [
    {
      "storedBytes": 0,
      "arn": "arn:aws:logs:us-east-1:123456789012:log-group:my-log-group-1:log-stream:my-log-stream-1",
      "creationTime": 1393545600000,
      "firstEventTimestamp": 1393545600000,
      "lastEventTimestamp": 1393567800000,
      "lastIngestionTime": 1393589200000,
      "logStreamName": "my-log-stream-1"
    },
    {
      "storedBytes": 0,
      "arn": "arn:aws:logs:us-east-1:123456789012:log-group:my-log-group-2:log-stream:my-log-stream-2",
      "creationTime": 1396224000000,
      "firstEventTimestamp": 1396224000000,
      "lastEventTimestamp": 1396235500000,
      "lastIngestionTime": 1396225560000,
      "logStreamName": "my-log-stream-2"
    }
  ]
}
```

### Example

The following example lists the log streams associated with the specified log
group.

#### Sample Request

```

{
  "logGroupIdentifier": "arn:aws:logs:us-east-1:123456789012:log-group:my-log-group-1:dli"
}
```

#### Sample Response

```

{
  "logStreams": [
    {
      "storedBytes": 0,
      "arn": "arn:aws:logs:us-east-1:123456789012:log-group:my-log-group-1:log-stream:my-
log-stream-1",
      "creationTime": 1393545600000,
      "firstEventTimestamp": 1393545600000,
      "lastEventTimestamp": 1393567800000,
      "lastIngestionTime": 1393589200000,
      "logStreamName": "my-log-stream-1"
      },
      {
      "storedBytes": 0,
      "arn": "arn:aws:logs:us-east-1:123456789012:log-group:my-log-group-2:log-stream:my-
log-stream-2",
      "creationTime": 1396224000000,
      "firstEventTimestamp": 1396224000000,
      "lastEventTimestamp": 1396235500000,
      "lastIngestionTime": 1396225560000,
      "logStreamName": "my-log-stream-2"
      } ]
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/logs-2014-03-28/describelogstreams.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/logs-2014-03-28/describelogstreams.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/describelogstreams.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/logs-2014-03-28/describelogstreams.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/describelogstreams.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/logs-2014-03-28/describelogstreams.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/logs-2014-03-28/describelogstreams.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/logs-2014-03-28/describelogstreams.md)

- [AWS SDK for Python](../../../../services/goto/boto3/logs-2014-03-28/describelogstreams.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/describelogstreams.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeLogGroups

DescribeLookupTables
