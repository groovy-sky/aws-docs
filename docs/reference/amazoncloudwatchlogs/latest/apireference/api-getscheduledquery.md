# GetScheduledQuery

Retrieves details about a specific scheduled query, including its configuration, execution
status, and metadata.

## Request Syntax

```nohighlight

{
   "identifier": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[identifier](#API_GetScheduledQuery_RequestSyntax)**

The ARN or name of the scheduled query to retrieve.

Type: String

Pattern: `[\w#+=/:,.@-]*`

Required: Yes

## Response Syntax

```nohighlight

{
   "creationTime": number,
   "description": "string",
   "destinationConfiguration": {
      "s3Configuration": {
         "destinationIdentifier": "string",
         "kmsKeyId": "string",
         "ownerAccountId": "string",
         "roleArn": "string"
      }
   },
   "executionRoleArn": "string",
   "lastExecutionStatus": "string",
   "lastTriggeredTime": number,
   "lastUpdatedTime": number,
   "logGroupIdentifiers": [ "string" ],
   "name": "string",
   "queryLanguage": "string",
   "queryString": "string",
   "scheduledQueryArn": "string",
   "scheduleEndTime": number,
   "scheduleExpression": "string",
   "scheduleStartTime": number,
   "startTimeOffset": number,
   "state": "string",
   "timezone": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[creationTime](#API_GetScheduledQuery_ResponseSyntax)**

The timestamp when the scheduled query was created.

Type: Long

Valid Range: Minimum value of 0.

**[description](#API_GetScheduledQuery_ResponseSyntax)**

The description of the scheduled query.

Type: String

Length Constraints: Maximum length of 1024.

**[destinationConfiguration](#API_GetScheduledQuery_ResponseSyntax)**

Configuration for where query results are delivered.

Type: [DestinationConfiguration](api-destinationconfiguration.md) object

**[executionRoleArn](#API_GetScheduledQuery_ResponseSyntax)**

The ARN of the IAM role used to execute the query and deliver results.

Type: String

Length Constraints: Minimum length of 1.

**[lastExecutionStatus](#API_GetScheduledQuery_ResponseSyntax)**

The status of the most recent execution of the scheduled query.

Type: String

Valid Values: `Running | InvalidQuery | Complete | Failed | Timeout`

**[lastTriggeredTime](#API_GetScheduledQuery_ResponseSyntax)**

The timestamp when the scheduled query was last executed.

Type: Long

Valid Range: Minimum value of 0.

**[lastUpdatedTime](#API_GetScheduledQuery_ResponseSyntax)**

The timestamp when the scheduled query was last updated.

Type: Long

Valid Range: Minimum value of 0.

**[logGroupIdentifiers](#API_GetScheduledQuery_ResponseSyntax)**

The log groups queried by the scheduled query.

Type: Array of strings

Array Members: Minimum number of 1 item. Maximum number of 50 items.

Length Constraints: Minimum length of 1. Maximum length of 2048.

Pattern: `[\w#+=/:,.@-]*`

**[name](#API_GetScheduledQuery_ResponseSyntax)**

The name of the scheduled query.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `^[a-zA-Z0-9_\-/.#]+$`

**[queryLanguage](#API_GetScheduledQuery_ResponseSyntax)**

The query language used by the scheduled query.

Type: String

Valid Values: `CWLI | SQL | PPL`

**[queryString](#API_GetScheduledQuery_ResponseSyntax)**

The query string executed by the scheduled query.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 10000.

**[scheduledQueryArn](#API_GetScheduledQuery_ResponseSyntax)**

The ARN of the scheduled query.

Type: String

**[scheduleEndTime](#API_GetScheduledQuery_ResponseSyntax)**

The end time for the scheduled query in Unix epoch format.

Type: Long

Valid Range: Minimum value of 0.

**[scheduleExpression](#API_GetScheduledQuery_ResponseSyntax)**

The cron expression that defines when the scheduled query runs.

Type: String

Length Constraints: Maximum length of 256.

**[scheduleStartTime](#API_GetScheduledQuery_ResponseSyntax)**

The start time for the scheduled query in Unix epoch format.

Type: Long

Valid Range: Minimum value of 0.

**[startTimeOffset](#API_GetScheduledQuery_ResponseSyntax)**

The time offset in seconds that defines the lookback period for the query.

Type: Long

**[state](#API_GetScheduledQuery_ResponseSyntax)**

The current state of the scheduled query.

Type: String

Valid Values: `ENABLED | DISABLED`

**[timezone](#API_GetScheduledQuery_ResponseSyntax)**

The timezone used for evaluating the schedule expression.

Type: String

Length Constraints: Minimum length of 1.

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/logs-2014-03-28/getscheduledquery.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/logs-2014-03-28/getscheduledquery.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/getscheduledquery.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/logs-2014-03-28/getscheduledquery.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/getscheduledquery.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/logs-2014-03-28/getscheduledquery.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/logs-2014-03-28/getscheduledquery.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/logs-2014-03-28/getscheduledquery.md)

- [AWS SDK for Python](../../../../services/goto/boto3/logs-2014-03-28/getscheduledquery.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/getscheduledquery.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

GetQueryResults

GetScheduledQueryHistory
