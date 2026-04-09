# DeleteImagePermissions

Deletes permissions for the specified private image. After you delete permissions for an image, AWS accounts to which you previously granted these permissions can no longer use the image.

## Request Syntax

```nohighlight

{
   "Name": "string",
   "SharedAccountId": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[Name](#API_DeleteImagePermissions_RequestSyntax)**

The name of the private image.

Type: String

Pattern: `^[a-zA-Z0-9][a-zA-Z0-9_.-]{0,100}$`

Required: Yes

**[SharedAccountId](#API_DeleteImagePermissions_RequestSyntax)**

The 12-digit identifier of the AWS account for which to delete image permissions.

Type: String

Pattern: `^\d+$`

Required: Yes

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**ResourceNotAvailableException**

The specified resource exists and is not in use, but isn't available.

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/appstream-2016-12-01/deleteimagepermissions.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/appstream-2016-12-01/deleteimagepermissions.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/appstream-2016-12-01/deleteimagepermissions.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/appstream-2016-12-01/deleteimagepermissions.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appstream-2016-12-01/deleteimagepermissions.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/appstream-2016-12-01/deleteimagepermissions.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/appstream-2016-12-01/deleteimagepermissions.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/appstream-2016-12-01/deleteimagepermissions.md)

- [AWS SDK for Python](../../../../services/goto/boto3/appstream-2016-12-01/deleteimagepermissions.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appstream-2016-12-01/deleteimagepermissions.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteImageBuilder

DeleteStack

All content copied from https://docs.aws.amazon.com/.
