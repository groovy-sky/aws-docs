# DeleteWorkGroup

Deletes the workgroup with the specified name. The primary workgroup cannot be
deleted.

## Request Syntax

```nohighlight

{
   "RecursiveDeleteOption": boolean,
   "WorkGroup": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[RecursiveDeleteOption](#API_DeleteWorkGroup_RequestSyntax)**

The option to delete the workgroup and its contents even if the workgroup contains any
named queries, query executions, or notebooks.

Type: Boolean

Required: No

**[WorkGroup](#API_DeleteWorkGroup_RequestSyntax)**

The unique name of the workgroup to delete.

Type: String

Pattern: `[a-zA-Z0-9._-]{1,128}`

Required: Yes

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/athena-2017-05-18/deleteworkgroup.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/athena-2017-05-18/deleteworkgroup.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/athena-2017-05-18/deleteworkgroup.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/athena-2017-05-18/deleteworkgroup.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/athena-2017-05-18/deleteworkgroup.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/athena-2017-05-18/deleteworkgroup.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/athena-2017-05-18/deleteworkgroup.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/athena-2017-05-18/deleteworkgroup.md)

- [AWS SDK for Python](../../../../services/goto/boto3/athena-2017-05-18/deleteworkgroup.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/athena-2017-05-18/deleteworkgroup.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DeletePreparedStatement

ExportNotebook
