# GetLogObject

Retrieves a large logging object (LLO) and streams it back. This API is used to fetch the
content of large portions of log events that have been ingested through the
PutOpenTelemetryLogs API. When log events contain fields that would cause the total event size
to exceed 1MB, CloudWatch Logs automatically processes up to 10 fields, starting with the
largest fields. Each field is truncated as needed to keep the total event size as close to 1MB
as possible. The excess portions are stored as Large Log Objects (LLOs) and these fields are
processed separately and LLO reference system fields (in the format
`@ptr.$[path.to.field]`) are added. The path in the reference field reflects the
original JSON structure where the large field was located. For example, this could be
`@ptr.$['input']['message']`, `@ptr.$['AAA']['BBB']['CCC']['DDD']`,
`@ptr.$['AAA']`, or any other path matching your log structure.

###### Note

The `GetLogObject` API routes requests using SDK host prefix injection. SDK versions released before April 1, 2026 route to
`streaming-logs.Region.amazonaws.com`, which does not support VPC endpoints. SDK versions released on or after April 1, 2026 route to
`stream-logs.Region.amazonaws.com`, which supports VPC endpoints. To set up a VPC endpoint for this API, see [Creating a VPC endpoint for CloudWatch Logs](../../../../services/amazoncloudwatch/latest/logs/cloudwatch-logs-and-interface-vpc.md#create-VPC-endpoint-for-CloudWatchLogs).

## Request Syntax

```nohighlight

{
   "logObjectPointer": "string",
   "unmask": boolean
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[logObjectPointer](#API_GetLogObject_RequestSyntax)**

A pointer to the specific log object to retrieve. This is a required parameter that
uniquely identifies the log object within CloudWatch Logs. The pointer is typically obtained
from a previous query or filter operation.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 512.

Required: Yes

**[unmask](#API_GetLogObject_RequestSyntax)**

A boolean flag that indicates whether to unmask sensitive log data. When set to true, any
masked or redacted data in the log object will be displayed in its original form. Default is
false.

Type: Boolean

Required: No

## Response Syntax

```nohighlight

{
   "fieldStream": {
      "fields": {
         "data": blob
      },
      "InternalStreamingException": {
      }
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[fieldStream](#API_GetLogObject_ResponseSyntax)**

A stream of structured log data returned by the GetLogObject operation. This stream
contains log events with their associated metadata and extracted fields.

Type: [GetLogObjectResponseStream](api-getlogobjectresponsestream.md) object

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/logs-2014-03-28/getlogobject.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/logs-2014-03-28/getlogobject.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/getlogobject.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/logs-2014-03-28/getlogobject.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/getlogobject.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/logs-2014-03-28/getlogobject.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/logs-2014-03-28/getlogobject.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/logs-2014-03-28/getlogobject.md)

- [AWS SDK for Python](../../../../services/goto/boto3/logs-2014-03-28/getlogobject.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/getlogobject.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

GetLogGroupFields

GetLogRecord
