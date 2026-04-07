# BatchGetQueryExecution

Returns the details of a single query execution or a list of up to 50 query
executions, which you provide as an array of query execution ID strings. Requires you to
have access to the workgroup in which the queries ran. To get a list of query execution
IDs, use [ListQueryExecutions:WorkGroup](https://docs.aws.amazon.com/athena/latest/APIReference/API_ListQueryExecutions.html#athena-ListQueryExecutions-request-WorkGroup). Query executions differ
from named (saved) queries. Use [BatchGetNamedQueryInput](https://docs.aws.amazon.com/athena/latest/APIReference/API_BatchGetNamedQueryInput.html) to get details
about named queries.

## Request Syntax

```nohighlight

{
   "QueryExecutionIds": [ "string" ]
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[QueryExecutionIds](#API_BatchGetQueryExecution_RequestSyntax)**

An array of query execution IDs.

Type: Array of strings

Array Members: Minimum number of 1 item. Maximum number of 50 items.

Length Constraints: Minimum length of 1. Maximum length of 128.

Pattern: `\S+`

Required: Yes

## Response Syntax

```nohighlight

{
   "QueryExecutions": [
      {
         "EngineVersion": {
            "EffectiveEngineVersion": "string",
            "SelectedEngineVersion": "string"
         },
         "ExecutionParameters": [ "string" ],
         "ManagedQueryResultsConfiguration": {
            "Enabled": boolean,
            "EncryptionConfiguration": {
               "KmsKey": "string"
            }
         },
         "Query": "string",
         "QueryExecutionContext": {
            "Catalog": "string",
            "Database": "string"
         },
         "QueryExecutionId": "string",
         "QueryResultsS3AccessGrantsConfiguration": {
            "AuthenticationType": "string",
            "CreateUserLevelPrefix": boolean,
            "EnableS3AccessGrants": boolean
         },
         "ResultConfiguration": {
            "AclConfiguration": {
               "S3AclOption": "string"
            },
            "EncryptionConfiguration": {
               "EncryptionOption": "string",
               "KmsKey": "string"
            },
            "ExpectedBucketOwner": "string",
            "OutputLocation": "string"
         },
         "ResultReuseConfiguration": {
            "ResultReuseByAgeConfiguration": {
               "Enabled": boolean,
               "MaxAgeInMinutes": number
            }
         },
         "StatementType": "string",
         "Statistics": {
            "DataManifestLocation": "string",
            "DataScannedInBytes": number,
            "DpuCount": number,
            "EngineExecutionTimeInMillis": number,
            "QueryPlanningTimeInMillis": number,
            "QueryQueueTimeInMillis": number,
            "ResultReuseInformation": {
               "ReusedPreviousResult": boolean
            },
            "ServicePreProcessingTimeInMillis": number,
            "ServiceProcessingTimeInMillis": number,
            "TotalExecutionTimeInMillis": number
         },
         "Status": {
            "AthenaError": {
               "ErrorCategory": number,
               "ErrorMessage": "string",
               "ErrorType": number,
               "Retryable": boolean
            },
            "CompletionDateTime": number,
            "State": "string",
            "StateChangeReason": "string",
            "SubmissionDateTime": number
         },
         "SubstatementType": "string",
         "WorkGroup": "string"
      }
   ],
   "UnprocessedQueryExecutionIds": [
      {
         "ErrorCode": "string",
         "ErrorMessage": "string",
         "QueryExecutionId": "string"
      }
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[QueryExecutions](#API_BatchGetQueryExecution_ResponseSyntax)**

Information about a query execution.

Type: Array of [QueryExecution](api-queryexecution.md) objects

**[UnprocessedQueryExecutionIds](#API_BatchGetQueryExecution_ResponseSyntax)**

Information about the query executions that failed to run.

Type: Array of [UnprocessedQueryExecutionId](https://docs.aws.amazon.com/athena/latest/APIReference/API_UnprocessedQueryExecutionId.html) objects

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

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/athena-2017-05-18/BatchGetQueryExecution)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/athena-2017-05-18/BatchGetQueryExecution)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/athena-2017-05-18/BatchGetQueryExecution)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/athena-2017-05-18/BatchGetQueryExecution)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/athena-2017-05-18/BatchGetQueryExecution)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/athena-2017-05-18/BatchGetQueryExecution)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/athena-2017-05-18/BatchGetQueryExecution)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/athena-2017-05-18/BatchGetQueryExecution)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/athena-2017-05-18/BatchGetQueryExecution)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/athena-2017-05-18/BatchGetQueryExecution)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

BatchGetPreparedStatement

CancelCapacityReservation
