# GetLogFields

Discovers available fields for a specific data source and type. The response includes any
field modifications introduced through pipelines, such as new fields or changed field types.

## Request Syntax

```nohighlight

{
   "dataSourceName": "string",
   "dataSourceType": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[dataSourceName](#API_GetLogFields_RequestSyntax)**

The name of the data source to retrieve log fields for.

Type: String

Required: Yes

**[dataSourceType](#API_GetLogFields_RequestSyntax)**

The type of the data source to retrieve log fields for.

Type: String

Required: Yes

## Response Syntax

```nohighlight

{
   "logFields": [
      {
         "logFieldName": "string",
         "logFieldType": {
            "element": "LogFieldType",
            "fields": [
               "LogFieldsListItem"
            ],
            "type": "string"
         }
      }
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[logFields](#API_GetLogFields_ResponseSyntax)**

The list of log fields for the specified data source, including field names and their data
types.

Type: Array of [LogFieldsListItem](api-logfieldslistitem.md) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidParameterException**

A parameter is specified incorrectly.

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/logs-2014-03-28/getlogfields.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/logs-2014-03-28/getlogfields.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/getlogfields.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/logs-2014-03-28/getlogfields.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/getlogfields.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/logs-2014-03-28/getlogfields.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/logs-2014-03-28/getlogfields.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/logs-2014-03-28/getlogfields.md)

- [AWS SDK for Python](../../../../services/goto/boto3/logs-2014-03-28/getlogfields.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/getlogfields.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetLogEvents

GetLogGroupFields

All content copied from https://docs.aws.amazon.com/.
