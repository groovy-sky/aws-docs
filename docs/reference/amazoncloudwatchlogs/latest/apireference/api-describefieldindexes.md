# DescribeFieldIndexes

Returns a list of custom and default field indexes which are discovered in log data. For
more information about field index policies, see [PutIndexPolicy](api-putindexpolicy.md).

## Request Syntax

```nohighlight

{
   "logGroupIdentifiers": [ "string" ],
   "nextToken": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[logGroupIdentifiers](#API_DescribeFieldIndexes_RequestSyntax)**

An array containing the names or ARNs of the log groups that you want to retrieve field
indexes for.

Type: Array of strings

Array Members: Minimum number of 1 item. Maximum number of 100 items.

Length Constraints: Minimum length of 1. Maximum length of 2048.

Pattern: `[\w#+=/:,.@-]*`

Required: Yes

**[nextToken](#API_DescribeFieldIndexes_RequestSyntax)**

The token for the next set of items to return. The token expires after 24
hours.

Type: String

Length Constraints: Minimum length of 1.

Required: No

## Response Syntax

```nohighlight

{
   "fieldIndexes": [
      {
         "fieldIndexName": "string",
         "firstEventTime": number,
         "lastEventTime": number,
         "lastScanTime": number,
         "logGroupIdentifier": "string",
         "type": "string"
      }
   ],
   "nextToken": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[fieldIndexes](#API_DescribeFieldIndexes_ResponseSyntax)**

An array containing the field index information.

Type: Array of [FieldIndex](api-fieldindex.md) objects

**[nextToken](#API_DescribeFieldIndexes_ResponseSyntax)**

The token for the next set of items to return. The token expires after 24
hours.

Type: String

Length Constraints: Minimum length of 1.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidParameterException**

A parameter is specified incorrectly.

HTTP Status Code: 400

**LimitExceededException**

You have reached the maximum number of resources that can be created.

HTTP Status Code: 400

**OperationAbortedException**

Multiple concurrent requests to update the same resource were in conflict.

HTTP Status Code: 400

**ResourceNotFoundException**

The specified resource does not exist.

HTTP Status Code: 400

**ServiceUnavailableException**

The service cannot complete the request.

HTTP Status Code: 500

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/logs-2014-03-28/describefieldindexes.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/logs-2014-03-28/describefieldindexes.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/describefieldindexes.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/logs-2014-03-28/describefieldindexes.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/describefieldindexes.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/logs-2014-03-28/describefieldindexes.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/logs-2014-03-28/describefieldindexes.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/logs-2014-03-28/describefieldindexes.md)

- [AWS SDK for Python](../../../../services/goto/boto3/logs-2014-03-28/describefieldindexes.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/describefieldindexes.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeExportTasks

DescribeImportTaskBatches
