# UpdateScheduledQuery

Updates an existing scheduled query with new configuration. This operation uses PUT
semantics, allowing modification of query parameters, schedule, and destinations.

## Request Syntax

```nohighlight

{
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
   "identifier": "string",
   "logGroupIdentifiers": [ "string" ],
   "queryLanguage": "string",
   "queryString": "string",
   "scheduleEndTime": number,
   "scheduleExpression": "string",
   "scheduleStartTime": number,
   "startTimeOffset": number,
   "state": "string",
   "timezone": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[description](#API_UpdateScheduledQuery_RequestSyntax)**

An updated description for the scheduled query.

Type: String

Length Constraints: Maximum length of 1024.

Required: No

**[destinationConfiguration](#API_UpdateScheduledQuery_RequestSyntax)**

The updated configuration for where to deliver query results.

Type: [DestinationConfiguration](https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_DestinationConfiguration.html) object

Required: No

**[executionRoleArn](#API_UpdateScheduledQuery_RequestSyntax)**

The updated ARN of the IAM role that grants permissions to execute the query and deliver
results.

Type: String

Length Constraints: Minimum length of 1.

Required: Yes

**[identifier](#API_UpdateScheduledQuery_RequestSyntax)**

The ARN or name of the scheduled query to update.

Type: String

Pattern: `[\w#+=/:,.@-]*`

Required: Yes

**[logGroupIdentifiers](#API_UpdateScheduledQuery_RequestSyntax)**

The updated array of log group names or ARNs to query.

Type: Array of strings

Array Members: Minimum number of 1 item. Maximum number of 50 items.

Length Constraints: Minimum length of 1. Maximum length of 2048.

Pattern: `[\w#+=/:,.@-]*`

Required: No

**[queryLanguage](#API_UpdateScheduledQuery_RequestSyntax)**

The updated query language for the scheduled query.

Type: String

Valid Values: `CWLI | SQL | PPL`

Required: Yes

**[queryString](#API_UpdateScheduledQuery_RequestSyntax)**

The updated query string to execute.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 10000.

Required: Yes

**[scheduleEndTime](#API_UpdateScheduledQuery_RequestSyntax)**

The updated end time for the scheduled query in Unix epoch format.

Type: Long

Valid Range: Minimum value of 0.

Required: No

**[scheduleExpression](#API_UpdateScheduledQuery_RequestSyntax)**

The updated cron expression that defines when the scheduled query runs.

Type: String

Length Constraints: Maximum length of 256.

Required: Yes

**[scheduleStartTime](#API_UpdateScheduledQuery_RequestSyntax)**

The updated start time for the scheduled query in Unix epoch format.

Type: Long

Valid Range: Minimum value of 0.

Required: No

**[startTimeOffset](#API_UpdateScheduledQuery_RequestSyntax)**

The updated time offset in seconds that defines the lookback period for the query.

Type: Long

Required: No

**[state](#API_UpdateScheduledQuery_RequestSyntax)**

The updated state of the scheduled query.

Type: String

Valid Values: `ENABLED | DISABLED`

Required: No

**[timezone](#API_UpdateScheduledQuery_RequestSyntax)**

The updated timezone for evaluating the schedule expression.

Type: String

Length Constraints: Minimum length of 1.

Required: No

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

**[creationTime](#API_UpdateScheduledQuery_ResponseSyntax)**

The timestamp when the scheduled query was originally created.

Type: Long

Valid Range: Minimum value of 0.

**[description](#API_UpdateScheduledQuery_ResponseSyntax)**

The description of the updated scheduled query.

Type: String

Length Constraints: Maximum length of 1024.

**[destinationConfiguration](#API_UpdateScheduledQuery_ResponseSyntax)**

The destination configuration of the updated scheduled query.

Type: [DestinationConfiguration](https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_DestinationConfiguration.html) object

**[executionRoleArn](#API_UpdateScheduledQuery_ResponseSyntax)**

The execution role ARN of the updated scheduled query.

Type: String

Length Constraints: Minimum length of 1.

**[lastExecutionStatus](#API_UpdateScheduledQuery_ResponseSyntax)**

The status of the most recent execution of the updated scheduled query.

Type: String

Valid Values: `Running | InvalidQuery | Complete | Failed | Timeout`

**[lastTriggeredTime](#API_UpdateScheduledQuery_ResponseSyntax)**

The timestamp when the updated scheduled query was last executed.

Type: Long

Valid Range: Minimum value of 0.

**[lastUpdatedTime](#API_UpdateScheduledQuery_ResponseSyntax)**

The timestamp when the scheduled query was last updated.

Type: Long

Valid Range: Minimum value of 0.

**[logGroupIdentifiers](#API_UpdateScheduledQuery_ResponseSyntax)**

The log groups queried by the updated scheduled query.

Type: Array of strings

Array Members: Minimum number of 1 item. Maximum number of 50 items.

Length Constraints: Minimum length of 1. Maximum length of 2048.

Pattern: `[\w#+=/:,.@-]*`

**[name](#API_UpdateScheduledQuery_ResponseSyntax)**

The name of the updated scheduled query.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `^[a-zA-Z0-9_\-/.#]+$`

**[queryLanguage](#API_UpdateScheduledQuery_ResponseSyntax)**

The query language of the updated scheduled query.

Type: String

Valid Values: `CWLI | SQL | PPL`

**[queryString](#API_UpdateScheduledQuery_ResponseSyntax)**

The query string of the updated scheduled query.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 10000.

**[scheduledQueryArn](#API_UpdateScheduledQuery_ResponseSyntax)**

The ARN of the updated scheduled query.

Type: String

**[scheduleEndTime](#API_UpdateScheduledQuery_ResponseSyntax)**

The end time of the updated scheduled query.

Type: Long

Valid Range: Minimum value of 0.

**[scheduleExpression](#API_UpdateScheduledQuery_ResponseSyntax)**

The cron expression of the updated scheduled query.

Type: String

Length Constraints: Maximum length of 256.

**[scheduleStartTime](#API_UpdateScheduledQuery_ResponseSyntax)**

The start time of the updated scheduled query.

Type: Long

Valid Range: Minimum value of 0.

**[startTimeOffset](#API_UpdateScheduledQuery_ResponseSyntax)**

The time offset of the updated scheduled query.

Type: Long

**[state](#API_UpdateScheduledQuery_ResponseSyntax)**

The state of the updated scheduled query.

Type: String

Valid Values: `ENABLED | DISABLED`

**[timezone](#API_UpdateScheduledQuery_ResponseSyntax)**

The timezone of the updated scheduled query.

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

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/logs-2014-03-28/UpdateScheduledQuery)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/logs-2014-03-28/UpdateScheduledQuery)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/logs-2014-03-28/UpdateScheduledQuery)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/logs-2014-03-28/UpdateScheduledQuery)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/logs-2014-03-28/UpdateScheduledQuery)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/logs-2014-03-28/UpdateScheduledQuery)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/logs-2014-03-28/UpdateScheduledQuery)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/logs-2014-03-28/UpdateScheduledQuery)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/logs-2014-03-28/UpdateScheduledQuery)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/logs-2014-03-28/UpdateScheduledQuery)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

UpdateLookupTable

Data Types
