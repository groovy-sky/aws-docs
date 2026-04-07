# ListScheduledQueries

Lists all scheduled queries in your account and region. You can filter results by state to
show only enabled or disabled queries.

## Request Syntax

```nohighlight

{
   "maxResults": number,
   "nextToken": "string",
   "state": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[maxResults](#API_ListScheduledQueries_RequestSyntax)**

The maximum number of scheduled queries to return. Valid range is 1 to 1000.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 1000.

Required: No

**[nextToken](#API_ListScheduledQueries_RequestSyntax)**

The token for the next set of items to return. The token expires after 24
hours.

Type: String

Length Constraints: Minimum length of 1.

Required: No

**[state](#API_ListScheduledQueries_RequestSyntax)**

Filter scheduled queries by state. Valid values are `ENABLED` and
`DISABLED`. If not specified, all scheduled queries are returned.

Type: String

Valid Values: `ENABLED | DISABLED`

Required: No

## Response Syntax

```nohighlight

{
   "nextToken": "string",
   "scheduledQueries": [
      {
         "creationTime": number,
         "destinationConfiguration": {
            "s3Configuration": {
               "destinationIdentifier": "string",
               "kmsKeyId": "string",
               "ownerAccountId": "string",
               "roleArn": "string"
            }
         },
         "lastExecutionStatus": "string",
         "lastTriggeredTime": number,
         "lastUpdatedTime": number,
         "name": "string",
         "scheduledQueryArn": "string",
         "scheduleExpression": "string",
         "state": "string",
         "timezone": "string"
      }
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[nextToken](#API_ListScheduledQueries_ResponseSyntax)**

The token for the next set of items to return. The token expires after 24
hours.

Type: String

Length Constraints: Minimum length of 1.

**[scheduledQueries](#API_ListScheduledQueries_ResponseSyntax)**

An array of scheduled query summary information.

Type: Array of [ScheduledQuerySummary](https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_ScheduledQuerySummary.html) objects

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

**ThrottlingException**

The request was throttled because of quota limits.

HTTP Status Code: 400

**ValidationException**

One of the parameters for the request is not valid.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/logs-2014-03-28/ListScheduledQueries)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/logs-2014-03-28/ListScheduledQueries)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/logs-2014-03-28/ListScheduledQueries)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/logs-2014-03-28/ListScheduledQueries)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/logs-2014-03-28/ListScheduledQueries)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/logs-2014-03-28/ListScheduledQueries)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/logs-2014-03-28/ListScheduledQueries)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/logs-2014-03-28/ListScheduledQueries)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/logs-2014-03-28/ListScheduledQueries)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/logs-2014-03-28/ListScheduledQueries)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ListLogGroupsForQuery

ListSourcesForS3TableIntegration
