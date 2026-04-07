# GetScheduledQueryHistory

Retrieves the execution history of a scheduled query within a specified time range,
including query results and destination processing status.

## Request Syntax

```nohighlight

{
   "endTime": number,
   "executionStatuses": [ "string" ],
   "identifier": "string",
   "maxResults": number,
   "nextToken": "string",
   "startTime": number
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[endTime](#API_GetScheduledQueryHistory_RequestSyntax)**

The end time for the history query in Unix epoch format.

Type: Long

Valid Range: Minimum value of 0.

Required: Yes

**[executionStatuses](#API_GetScheduledQueryHistory_RequestSyntax)**

An array of execution statuses to filter the history results. Only executions with the
specified statuses are returned.

Type: Array of strings

Valid Values: `Running | InvalidQuery | Complete | Failed | Timeout`

Required: No

**[identifier](#API_GetScheduledQueryHistory_RequestSyntax)**

The ARN or name of the scheduled query to retrieve history for.

Type: String

Pattern: `[\w#+=/:,.@-]*`

Required: Yes

**[maxResults](#API_GetScheduledQueryHistory_RequestSyntax)**

The maximum number of history records to return. Valid range is 1 to 1000.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 1000.

Required: No

**[nextToken](#API_GetScheduledQueryHistory_RequestSyntax)**

The token for the next set of items to return. The token expires after 24
hours.

Type: String

Length Constraints: Minimum length of 1.

Required: No

**[startTime](#API_GetScheduledQueryHistory_RequestSyntax)**

The start time for the history query in Unix epoch format.

Type: Long

Valid Range: Minimum value of 0.

Required: Yes

## Response Syntax

```nohighlight

{
   "name": "string",
   "nextToken": "string",
   "scheduledQueryArn": "string",
   "triggerHistory": [
      {
         "destinations": [
            {
               "destinationIdentifier": "string",
               "destinationType": "string",
               "errorMessage": "string",
               "processedIdentifier": "string",
               "status": "string"
            }
         ],
         "errorMessage": "string",
         "executionStatus": "string",
         "queryId": "string",
         "triggeredTimestamp": number
      }
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[name](#API_GetScheduledQueryHistory_ResponseSyntax)**

The name of the scheduled query.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `^[a-zA-Z0-9_\-/.#]+$`

**[nextToken](#API_GetScheduledQueryHistory_ResponseSyntax)**

The token for the next set of items to return. The token expires after 24
hours.

Type: String

Length Constraints: Minimum length of 1.

**[scheduledQueryArn](#API_GetScheduledQueryHistory_ResponseSyntax)**

The ARN of the scheduled query.

Type: String

**[triggerHistory](#API_GetScheduledQueryHistory_ResponseSyntax)**

An array of execution history records for the scheduled query.

Type: Array of [TriggerHistoryRecord](https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_TriggerHistoryRecord.html) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**AccessDeniedException**

You don't have sufficient permissions to perform this action.

HTTP Status Code: 400

**InternalServerException**

An internal server error occurred while processing the request. This exception is returned
when the service encounters an unexpected condition that prevents it from fulfilling the
request.

HTTP Status Code: 500

**ResourceNotFoundException**

The specified resource does not exist.

HTTP Status Code: 400

**ThrottlingException**

The request was throttled because of quota limits.

HTTP Status Code: 400

**ValidationException**

One of the parameters for the request is not valid.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/logs-2014-03-28/GetScheduledQueryHistory)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/logs-2014-03-28/GetScheduledQueryHistory)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/logs-2014-03-28/GetScheduledQueryHistory)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/logs-2014-03-28/GetScheduledQueryHistory)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/logs-2014-03-28/GetScheduledQueryHistory)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/logs-2014-03-28/GetScheduledQueryHistory)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/logs-2014-03-28/GetScheduledQueryHistory)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/logs-2014-03-28/GetScheduledQueryHistory)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/logs-2014-03-28/GetScheduledQueryHistory)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/logs-2014-03-28/GetScheduledQueryHistory)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

GetScheduledQuery

GetTransformer
