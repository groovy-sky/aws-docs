# DescribeImportTasks

Lists and describes import tasks, with optional filtering by import status and source ARN.

## Request Syntax

```nohighlight

{
   "importId": "string",
   "importSourceArn": "string",
   "importStatus": "string",
   "limit": number,
   "nextToken": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[importId](#API_DescribeImportTasks_RequestSyntax)**

Optional filter to describe a specific import task by its ID.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Pattern: `[\-a-zA-Z0-9]+`

Required: No

**[importSourceArn](#API_DescribeImportTasks_RequestSyntax)**

Optional filter to list imports from a specific source

Type: String

Required: No

**[importStatus](#API_DescribeImportTasks_RequestSyntax)**

Optional filter to list imports by their status. Valid values are IN\_PROGRESS, CANCELLED, COMPLETED and FAILED.

Type: String

Valid Values: `IN_PROGRESS | CANCELLED | COMPLETED | FAILED`

Required: No

**[limit](#API_DescribeImportTasks_RequestSyntax)**

The maximum number of import tasks to return in the response. Default: 50

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 50.

Required: No

**[nextToken](#API_DescribeImportTasks_RequestSyntax)**

The pagination token for the next set of results.

Type: String

Length Constraints: Minimum length of 1.

Required: No

## Response Syntax

```nohighlight

{
   "imports": [
      {
         "creationTime": number,
         "errorMessage": "string",
         "importDestinationArn": "string",
         "importFilter": {
            "endEventTime": number,
            "startEventTime": number
         },
         "importId": "string",
         "importSourceArn": "string",
         "importStatistics": {
            "bytesImported": number
         },
         "importStatus": "string",
         "lastUpdatedTime": number
      }
   ],
   "nextToken": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[imports](#API_DescribeImportTasks_ResponseSyntax)**

The list of import tasks that match the request filters.

Type: Array of [Import](api-import.md) objects

**[nextToken](#API_DescribeImportTasks_ResponseSyntax)**

The token to use when requesting the next set of results. Not present if there are no additional results to retrieve.

Type: String

Length Constraints: Minimum length of 1.

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

**ResourceNotFoundException**

The specified resource does not exist.

HTTP Status Code: 400

**ThrottlingException**

The request was throttled because of quota limits.

HTTP Status Code: 400

## Examples

### To describe import tasks

The following example retrieves a list of import tasks with filters.

#### Sample Request

```

POST / HTTP/1.1
          Host: logs.<region>.<domain>
          X-Amz-Target: Logs_20140328.DescribeImportTasks
          {
          "importId": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
          "importStatus": "IN_PROGRESS",
          "importSourceArn": "arn:aws:cloudtrail:us-west-2:123456789012:eventdatastore/f1d45bff-d0e3-4868-b5d9-2eb678aa32fb",
          "limit": 50,
          "nextToken": "eyJOZXh0VG9rZW4iOiJudWxsIiwiYm90b190cnVuY2F0ZV9hbW91bnQiOjF9"
          }
```

#### Sample Response

```

HTTP/1.1 200 OK
          {
          "imports": [
          {
          "importId": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
          "importSourceArn": "arn:aws:cloudtrail:us-east-1:123456789012:eventdatastore/f1d45bff-d0e3-4868-b5d9-2eb678aa32fb",
          "importStatus": "IN_PROGRESS",
          "importDestinationArn": "arn:aws:logs:us-east-1:123456789012:log-group:aws/cloudtrail/f1d45bff-d0e3-4868-b5d9-2eb678aa32fb",
          "importStatistics": {
          "bytesImported": 1048576
          },
          "importFilter": {
          "startEventTime": 1640995200000,
          "endEventTime": 1641081600000
          },
          "creationTime": 1641168000000,
          "lastUpdatedTime": 1641171600000
          }
          ],
          "nextToken": "eyJOZXh0VG9rZW4iOiJudWxsIiwiYm90b190cnVuY2F0ZV9hbW91bnQiOjJ9"
          }
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/logs-2014-03-28/describeimporttasks.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/logs-2014-03-28/describeimporttasks.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/describeimporttasks.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/logs-2014-03-28/describeimporttasks.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/describeimporttasks.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/logs-2014-03-28/describeimporttasks.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/logs-2014-03-28/describeimporttasks.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/logs-2014-03-28/describeimporttasks.md)

- [AWS SDK for Python](../../../../services/goto/boto3/logs-2014-03-28/describeimporttasks.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/describeimporttasks.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeImportTaskBatches

DescribeIndexPolicies
