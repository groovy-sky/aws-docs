# GetQueryRuntimeStatistics

Returns query execution runtime statistics related to a single execution of a query if
you have access to the workgroup in which the query ran. Statistics from the
`Timeline` section of the response object are available as soon as [QueryExecutionStatus:State](api-queryexecutionstatus-athena-type-queryexecutionstatus-state.md) is in a SUCCEEDED or FAILED state. The
remaining non-timeline statistics in the response (like stage-level input and output row
count and data size) are updated asynchronously and may not be available immediately
after a query completes or, in some cases, may not be returned. The non-timeline
statistics are also not included when a query has row-level filters defined in Lake Formation.

## Request Syntax

```nohighlight

{
   "QueryExecutionId": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[QueryExecutionId](#API_GetQueryRuntimeStatistics_RequestSyntax)**

The unique ID of the query execution.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Pattern: `\S+`

Required: Yes

## Response Syntax

```nohighlight

{
   "QueryRuntimeStatistics": {
      "OutputStage": {
         "ExecutionTime": number,
         "InputBytes": number,
         "InputRows": number,
         "OutputBytes": number,
         "OutputRows": number,
         "QueryStagePlan": {
            "Children": [
               "QueryStagePlanNode"
            ],
            "Identifier": "string",
            "Name": "string",
            "RemoteSources": [ "string" ]
         },
         "StageId": number,
         "State": "string",
         "SubStages": [
            "QueryStage"
         ]
      },
      "Rows": {
         "InputBytes": number,
         "InputRows": number,
         "OutputBytes": number,
         "OutputRows": number
      },
      "Timeline": {
         "EngineExecutionTimeInMillis": number,
         "QueryPlanningTimeInMillis": number,
         "QueryQueueTimeInMillis": number,
         "ServicePreProcessingTimeInMillis": number,
         "ServiceProcessingTimeInMillis": number,
         "TotalExecutionTimeInMillis": number
      }
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[QueryRuntimeStatistics](#API_GetQueryRuntimeStatistics_ResponseSyntax)**

Runtime statistics about the query execution.

Type: [QueryRuntimeStatistics](api-queryruntimestatistics.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InternalServerException**

Indicates a platform issue, which may be due to a transient condition or
outage.

HTTP Status Code: 500

**InvalidRequestException**

Indicates that something is wrong with the input to the request. For example, a
required parameter may be missing or out of range.

**AthenaErrorCode**

The error code returned when the query execution failed to process, or when the
processing request for the named query failed.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/athena-2017-05-18/getqueryruntimestatistics.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/athena-2017-05-18/getqueryruntimestatistics.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/athena-2017-05-18/getqueryruntimestatistics.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/athena-2017-05-18/getqueryruntimestatistics.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/athena-2017-05-18/getqueryruntimestatistics.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/athena-2017-05-18/getqueryruntimestatistics.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/athena-2017-05-18/getqueryruntimestatistics.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/athena-2017-05-18/getqueryruntimestatistics.md)

- [AWS SDK for Python](../../../../services/goto/boto3/athena-2017-05-18/getqueryruntimestatistics.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/athena-2017-05-18/getqueryruntimestatistics.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

GetQueryResults

GetResourceDashboard
