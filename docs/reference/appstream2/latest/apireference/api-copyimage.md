# CopyImage

Copies the image within the same region or to a new region within the same AWS account. Note that any tags you added to the image will not be copied.

## Request Syntax

```nohighlight

{
   "DestinationImageDescription": "string",
   "DestinationImageName": "string",
   "DestinationRegion": "string",
   "SourceImageName": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[DestinationImageDescription](#API_CopyImage_RequestSyntax)**

The description that the image will have when it is copied to the destination.

Type: String

Length Constraints: Maximum length of 256.

Required: No

**[DestinationImageName](#API_CopyImage_RequestSyntax)**

The name that the image will have when it is copied to the destination.

Type: String

Pattern: `^[a-zA-Z0-9][a-zA-Z0-9_.-]{0,100}$`

Required: Yes

**[DestinationRegion](#API_CopyImage_RequestSyntax)**

The destination region to which the image will be copied. This parameter is required, even if you are copying an image within the same region.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 32.

Required: Yes

**[SourceImageName](#API_CopyImage_RequestSyntax)**

The name of the image to copy.

Type: String

Pattern: `^[a-zA-Z0-9][a-zA-Z0-9_.-]{0,100}$`

Required: Yes

## Response Syntax

```nohighlight

{
   "DestinationImageName": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[DestinationImageName](#API_CopyImage_ResponseSyntax)**

The name of the destination image.

Type: String

Pattern: `^[a-zA-Z0-9][a-zA-Z0-9_.-]{0,100}$`

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**IncompatibleImageException**

The image can't be updated because it's not compatible for updates.

**Message**

The error message in the exception.

HTTP Status Code: 400

**InvalidAccountStatusException**

The resource cannot be created because your AWS account is suspended. For assistance, contact AWS Support.

**Message**

The error message in the exception.

HTTP Status Code: 400

**LimitExceededException**

The requested limit exceeds the permitted limit for an account.

**Message**

The error message in the exception.

HTTP Status Code: 400

**ResourceAlreadyExistsException**

The specified resource already exists.

**Message**

The error message in the exception.

HTTP Status Code: 400

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/appstream-2016-12-01/copyimage.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/appstream-2016-12-01/copyimage.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/appstream-2016-12-01/copyimage.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/appstream-2016-12-01/copyimage.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appstream-2016-12-01/copyimage.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/appstream-2016-12-01/copyimage.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/appstream-2016-12-01/copyimage.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/appstream-2016-12-01/copyimage.md)

- [AWS SDK for Python](../../../../services/goto/boto3/appstream-2016-12-01/copyimage.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appstream-2016-12-01/copyimage.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BatchDisassociateUserStack

CreateAppBlock

All content copied from https://docs.aws.amazon.com/.
