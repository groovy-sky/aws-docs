# DeleteLookupTable

Deletes a lookup table permanently. This operation cannot be undone.

Queries that reference a deleted table will return an error. Before deleting a lookup
table, review any saved queries or dashboards that may reference it.

## Request Syntax

```nohighlight

{
   "lookupTableArn": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[lookupTableArn](#API_DeleteLookupTable_RequestSyntax)**

The ARN of the lookup table to delete.

Type: String

Required: Yes

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**AccessDeniedException**

You don't have sufficient permissions to perform this action.

HTTP Status Code: 400

**InvalidParameterException**

A parameter is specified incorrectly.

HTTP Status Code: 400

**ResourceNotFoundException**

The specified resource does not exist.

HTTP Status Code: 400

**ServiceUnavailableException**

The service cannot complete the request.

HTTP Status Code: 500

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/logs-2014-03-28/deletelookuptable.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/logs-2014-03-28/deletelookuptable.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/deletelookuptable.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/logs-2014-03-28/deletelookuptable.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/deletelookuptable.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/logs-2014-03-28/deletelookuptable.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/logs-2014-03-28/deletelookuptable.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/logs-2014-03-28/deletelookuptable.md)

- [AWS SDK for Python](../../../../services/goto/boto3/logs-2014-03-28/deletelookuptable.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/deletelookuptable.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteLogStream

DeleteMetricFilter

All content copied from https://docs.aws.amazon.com/.
