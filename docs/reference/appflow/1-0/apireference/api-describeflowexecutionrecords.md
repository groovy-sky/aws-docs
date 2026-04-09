# DescribeFlowExecutionRecords

Fetches the execution history of the flow.

## Request Syntax

```nohighlight

POST /describe-flow-execution-records HTTP/1.1
Content-type: application/json

{
   "flowName": "string",
   "maxResults": number,
   "nextToken": "string"
}
```

## URI Request Parameters

The request does not use any URI parameters.

## Request Body

The request accepts the following data in JSON format.

**[flowName](#API_DescribeFlowExecutionRecords_RequestSyntax)**

The specified name of the flow. Spaces are not allowed. Use underscores (\_) or hyphens
(-) only.

Type: String

Length Constraints: Maximum length of 256.

Pattern: `[a-zA-Z0-9][\w!@#.-]+`

Required: Yes

**[maxResults](#API_DescribeFlowExecutionRecords_RequestSyntax)**

Specifies the maximum number of items that should be returned in the result set. The
default for `maxResults` is 20 (for all paginated API operations).

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 100.

Required: No

**[nextToken](#API_DescribeFlowExecutionRecords_RequestSyntax)**

The pagination token for the next page of data.

Type: String

Length Constraints: Maximum length of 2048.

Pattern: `\S+`

Required: No

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "flowExecutions": [
      {
         "dataPullEndTime": number,
         "dataPullStartTime": number,
         "executionId": "string",
         "executionResult": {
            "bytesProcessed": number,
            "bytesWritten": number,
            "errorInfo": {
               "executionMessage": "string",
               "putFailuresCount": number
            },
            "maxPageSize": number,
            "numParallelProcesses": number,
            "recordsProcessed": number
         },
         "executionStatus": "string",
         "lastUpdatedAt": number,
         "metadataCatalogDetails": [
            {
               "catalogType": "string",
               "partitionRegistrationOutput": {
                  "message": "string",
                  "result": "string",
                  "status": "string"
               },
               "tableName": "string",
               "tableRegistrationOutput": {
                  "message": "string",
                  "result": "string",
                  "status": "string"
               }
            }
         ],
         "startedAt": number
      }
   ],
   "nextToken": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[flowExecutions](#API_DescribeFlowExecutionRecords_ResponseSyntax)**

Returns a list of all instances when this flow was run.

Type: Array of [ExecutionRecord](api-executionrecord.md) objects

**[nextToken](#API_DescribeFlowExecutionRecords_ResponseSyntax)**

The pagination token for the next page of data.

Type: String

Length Constraints: Maximum length of 2048.

Pattern: `\S+`

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InternalServerException**

An internal service error occurred during the processing of your request. Try again
later.

HTTP Status Code: 500

**ResourceNotFoundException**

The resource specified in the request (such as the source or destination connector
profile) is not found.

HTTP Status Code: 404

**ValidationException**

The request has invalid or missing parameters.

HTTP Status Code: 400

## Examples

### DescribeFlowExecutionRecords examples

This example shows sample requests and responses for the
`DescribeFlowExecutionRecords` API.

#### Sample Request

```json

{
  "flowName": "name",
  "maxResults": 1
}
```

#### Sample Response

```json

{
  "flowExecutionList": [
    {
      "executionId": "Execution_ID",
      "executionMessage": "Request failed with Trendmicro Status Code 404: , RequestId: RequestId_value",
      "executionMetadata":
      {
        "bytesProcessed": 234,
        "bytesWritten": 0,
        "numFailures": 878,
        "numFilteredRecords": 0,
        "numPutFailures": 978978,
        "reason": null,
        "recordsProcessed": 23342,
        "terminateFlow": false
      },
      "executionResult":
      {
        "bytesProcessed": 234,
        "bytesWritten": 0,
        "errorInfo":
        {
          "executionMessage": "Request failed with Trendmicro Status Code 404: , RequestId: RequestId_value",
          "putFailuresCount": 978978
        },
        "recordsProcessed": 23342
      },
      "executionStartTime": "2022-02-22T15:31:41.467000-08:00",
      "executionStatus": "Error",
      "lastUpdatedAt": "2022-02-22T15:31:41.467000-08:00",
      "startedAt": "2022-02-22T15:31:41.467000-08:00"
    }
  ],
  "nextToken": "next_token_value"
}
```

#### Sample Request

```json

{
  "flowName": "test-new-create",
  "maxResults": 1,
  "nextToken": "next_token_value"
}
```

#### Sample Response

```json

{
  "flowExecutionList": [
    {
      "executionId": "execution_id_value",
      "executionMessage": "Request failed with Trendmicro Status Code 404: , RequestId: RequestId_value",
      "executionMetadata": {
        "bytesProcessed": 234,
        "bytesWritten": 0,
        "numFailures": 0,
        "numFilteredRecords": 0,
        "numPutFailures": 0,
        "reason": null,
        "recordsProcessed": 23342,
        "terminateFlow": false
      },
      "executionResult": {
        "bytesProcessed": 234,
        "bytesWritten": 0,
        "errorInfo": null,
        "recordsProcessed": 23342
      },
      "executionStartTime": "execution_start_time_value",
      "executionStatus": "Successful",
      "lastUpdatedAt": "lastupdated_at_value",
      "lastUpdatedTime": "lastupdated_time_value",
      "startedAt": "started_at_value"
    }
  ],
  "nextToken": "next_token_value"
}
```

#### Sample Request

```json

{
  "flowName": "name",
  "maxResults": 1,
  "nextToken": "next_token_value"
}
```

#### Sample Response

```json

{
  "flowExecutionList": [
    {
      "executionId": "execution_id_value-delete",
      "executionMessage": "Request failed with Trendmicro Status Code 404: , RequestId: RequestId_value",
      "executionMetadata": {
        "bytesProcessed": 0,
        "bytesWritten": 0,
        "numFailures": 0,
        "numFilteredRecords": 0,
        "numPutFailures": 0,
        "reason": null,
        "recordsProcessed": 0,
        "terminateFlow": false
      },
      "executionResult": {
        "bytesProcessed": 0,
        "bytesWritten": 0,
        "errorInfo": {
          "executionMessage": "Request failed with Trendmicro Status Code 404: , RequestId: RequestId_value",
          "putFailuresCount": 0
        },
        "recordsProcessed": 0
      },
      "executionStartTime": "execution_start_time_value",
      "executionStatus": "Error",
      "lastUpdatedAt": "lastupdated_at_value",
      "lastUpdatedTime": "lastupdated_time_value",
      "startedAt": "started_at_value"
    }
  ],
  "nextToken": "next_token_value"
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/appflow-2020-08-23/describeflowexecutionrecords.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/appflow-2020-08-23/describeflowexecutionrecords.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/appflow-2020-08-23/describeflowexecutionrecords.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/appflow-2020-08-23/describeflowexecutionrecords.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appflow-2020-08-23/describeflowexecutionrecords.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/appflow-2020-08-23/describeflowexecutionrecords.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/appflow-2020-08-23/describeflowexecutionrecords.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/appflow-2020-08-23/describeflowexecutionrecords.md)

- [AWS SDK for Python](../../../../services/goto/boto3/appflow-2020-08-23/describeflowexecutionrecords.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appflow-2020-08-23/describeflowexecutionrecords.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DescribeFlow

ListConnectorEntities

All content copied from https://docs.aws.amazon.com/.
