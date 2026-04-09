# DeleteStack

Deletes the specified stack. After the stack is deleted, the application streaming environment provided by the stack is no longer available to users. Also, any reservations made for application streaming sessions for the stack are released.

## Request Syntax

```nohighlight

{
   "Name": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[Name](#API_DeleteStack_RequestSyntax)**

The name of the stack.

Type: String

Length Constraints: Minimum length of 1.

Required: Yes

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**ConcurrentModificationException**

An API error occurred. Wait a few minutes and try again.

**Message**

The error message in the exception.

HTTP Status Code: 400

**OperationNotPermittedException**

The attempted operation is not permitted.

**Message**

The error message in the exception.

HTTP Status Code: 400

**ResourceInUseException**

The specified resource is in use.

**Message**

The error message in the exception.

HTTP Status Code: 400

**ResourceNotFoundException**

The specified resource was not found.

**Message**

The error message in the exception.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/appstream-2016-12-01/deletestack.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/appstream-2016-12-01/deletestack.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/appstream-2016-12-01/deletestack.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/appstream-2016-12-01/deletestack.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appstream-2016-12-01/deletestack.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/appstream-2016-12-01/deletestack.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/appstream-2016-12-01/deletestack.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/appstream-2016-12-01/deletestack.md)

- [AWS SDK for Python](../../../../services/goto/boto3/appstream-2016-12-01/deletestack.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appstream-2016-12-01/deletestack.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteImagePermissions

DeleteThemeForStack

All content copied from https://docs.aws.amazon.com/.
