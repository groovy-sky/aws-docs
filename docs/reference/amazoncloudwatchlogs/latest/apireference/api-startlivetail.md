# StartLiveTail

Starts a Live Tail streaming session for one or more log groups. A Live Tail session
returns a stream of log events that have been recently ingested in the log groups. For more
information, see [Use Live Tail to view logs\
in near real time](../../../../services/amazoncloudwatch/latest/logs/cloudwatchlogs-livetail.md).

The response to this operation is a response stream, over which the server sends live log
events and the client receives them.

The following objects are sent over the stream:

- A single [LiveTailSessionStart](api-livetailsessionstart.md) object is sent at the start of the session.

- Every second, a [LiveTailSessionUpdate](api-livetailsessionupdate.md) object is sent. Each of these objects contains an array
of the actual log events.

If no new log events were ingested in the past second, the
`LiveTailSessionUpdate` object will contain an empty array.

The array of log events contained in a `LiveTailSessionUpdate` can include
as many as 500 log events. If the number of log events matching the request exceeds 500
per second, the log events are sampled down to 500 log events to be included in each
`LiveTailSessionUpdate` object.

If your client consumes the log events slower than the server produces them, CloudWatch Logs buffers up to 10 `LiveTailSessionUpdate` events or 5000 log
events, after which it starts dropping the oldest events.

- A [SessionStreamingException](api-startlivetailresponsestream-cwl-type-startlivetailresponsestream-sessionstreamingexception.md) object is returned if an unknown error occurs on the
server side.

- A [SessionTimeoutException](api-startlivetailresponsestream-cwl-type-startlivetailresponsestream-sessiontimeoutexception.md) object is returned when the session times out, after it
has been kept open for three hours.

###### Note

The `StartLiveTail` API routes requests using SDK host prefix injection. SDK versions released before April 1, 2026 route to
`streaming-logs.Region.amazonaws.com`, which does not support VPC endpoints. SDK versions released on or after April 1, 2026 route to
`stream-logs.Region.amazonaws.com`, which supports VPC endpoints. To set up a VPC endpoint for this API, see [Creating a VPC endpoint for CloudWatch Logs](../../../../services/amazoncloudwatch/latest/logs/cloudwatch-logs-and-interface-vpc-create-vpc-endpoint-for-cloudwatchlogs.md).

###### Important

You can end a session before it times out by closing the session stream or by closing
the client that is receiving the stream. The session also ends if the established connection
between the client and the server breaks.

For examples of using an SDK to start a Live Tail session, see [Start\
a Live Tail session using an AWS SDK](../../../../services/amazoncloudwatch/latest/logs/example-cloudwatch-logs-startlivetail-section.md).

## Request Syntax

```nohighlight

{
   "logEventFilterPattern": "string",
   "logGroupIdentifiers": [ "string" ],
   "logStreamNamePrefixes": [ "string" ],
   "logStreamNames": [ "string" ]
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[logEventFilterPattern](#API_StartLiveTail_RequestSyntax)**

An optional pattern to use to filter the results to include only log events that match the
pattern. For example, a filter pattern of `error 404` causes only log events that
include both `error` and `404` to be included in the Live Tail
stream.

Regular expression filter patterns are supported.

For more information about filter pattern syntax, see [Filter and Pattern\
Syntax](../../../../services/amazoncloudwatch/latest/logs/filterandpatternsyntax.md).

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1024.

Required: No

**[logGroupIdentifiers](#API_StartLiveTail_RequestSyntax)**

An array where each item in the array is a log group to include in the Live Tail
session.

Specify each log group by its ARN.

If you specify an ARN, the ARN can't end with an asterisk (\*).

###### Note

You can include up to 10 log groups.

Type: Array of strings

Array Members: Minimum number of 1 item. Maximum number of 10 items.

Length Constraints: Minimum length of 1. Maximum length of 2048.

Pattern: `[\w#+=/:,.@-]*`

Required: Yes

**[logStreamNamePrefixes](#API_StartLiveTail_RequestSyntax)**

If you specify this parameter, then only log events in the log streams that have names
that start with the prefixes that you specify here are included in the Live Tail
session.

If you specify this field, you can't also specify the `logStreamNames`
field.

###### Note

You can specify this parameter only if you specify only one log group in
`logGroupIdentifiers`.

Type: Array of strings

Array Members: Minimum number of 1 item. Maximum number of 100 items.

Length Constraints: Minimum length of 1. Maximum length of 512.

Pattern: `[^:*]*`

Required: No

**[logStreamNames](#API_StartLiveTail_RequestSyntax)**

If you specify this parameter, then only log events in the log streams that you specify
here are included in the Live Tail session.

If you specify this field, you can't also specify the `logStreamNamePrefixes`
field.

###### Note

You can specify this parameter only if you specify only one log group in
`logGroupIdentifiers`.

Type: Array of strings

Array Members: Minimum number of 1 item. Maximum number of 100 items.

Length Constraints: Minimum length of 1. Maximum length of 512.

Pattern: `[^:*]*`

Required: No

## Response Syntax

```nohighlight

{
   "responseStream": {
      "sessionStart": {
         "logEventFilterPattern": "string",
         "logGroupIdentifiers": [ "string" ],
         "logStreamNamePrefixes": [ "string" ],
         "logStreamNames": [ "string" ],
         "requestId": "string",
         "sessionId": "string"
      },
      "SessionStreamingException": {
      },
      "SessionTimeoutException": {
      },
      "sessionUpdate": {
         "sessionMetadata": {
            "sampled": boolean
         },
         "sessionResults": [
            {
               "ingestionTime": number,
               "logGroupIdentifier": "string",
               "logStreamName": "string",
               "message": "string",
               "timestamp": number
            }
         ]
      }
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[responseStream](#API_StartLiveTail_ResponseSyntax)**

An object that includes the stream returned by your request. It can include both log
events and exceptions.

Type: [StartLiveTailResponseStream](api-startlivetailresponsestream.md) object

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

**LimitExceededException**

You have reached the maximum number of resources that can be created.

HTTP Status Code: 400

**ResourceNotFoundException**

The specified resource does not exist.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/logs-2014-03-28/startlivetail.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/logs-2014-03-28/startlivetail.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/startlivetail.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/logs-2014-03-28/startlivetail.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/startlivetail.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/logs-2014-03-28/startlivetail.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/logs-2014-03-28/startlivetail.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/logs-2014-03-28/startlivetail.md)

- [AWS SDK for Python](../../../../services/goto/boto3/logs-2014-03-28/startlivetail.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/startlivetail.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

PutTransformer

StartQuery
