# DeleteImage

Deletes the specified image. You cannot delete an image when it is in use.
After you delete an image, you cannot provision new capacity using the image.

## Request Syntax

```nohighlight

{
   "Name": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[Name](#API_DeleteImage_RequestSyntax)**

The name of the image.

Type: String

Pattern: `^[a-zA-Z0-9][a-zA-Z0-9_.-]{0,100}$`

Required: Yes

## Response Syntax

```nohighlight

{
   "Image": {
      "Applications": [
         {
            "AppBlockArn": "string",
            "Arn": "string",
            "CreatedTime": number,
            "Description": "string",
            "DisplayName": "string",
            "Enabled": boolean,
            "IconS3Location": {
               "S3Bucket": "string",
               "S3Key": "string"
            },
            "IconURL": "string",
            "InstanceFamilies": [ "string" ],
            "LaunchParameters": "string",
            "LaunchPath": "string",
            "Metadata": {
               "string" : "string"
            },
            "Name": "string",
            "Platforms": [ "string" ],
            "WorkingDirectory": "string"
         }
      ],
      "AppstreamAgentVersion": "string",
      "Arn": "string",
      "BaseImageArn": "string",
      "CreatedTime": number,
      "Description": "string",
      "DisplayName": "string",
      "DynamicAppProvidersEnabled": "string",
      "ImageBuilderName": "string",
      "ImageBuilderSupported": boolean,
      "ImageErrors": [
         {
            "ErrorCode": "string",
            "ErrorMessage": "string",
            "ErrorTimestamp": number
         }
      ],
      "ImagePermissions": {
         "allowFleet": boolean,
         "allowImageBuilder": boolean
      },
      "ImageSharedWithOthers": "string",
      "ImageType": "string",
      "LatestAppstreamAgentVersion": "string",
      "ManagedSoftwareIncluded": boolean,
      "Name": "string",
      "Platform": "string",
      "PublicBaseImageReleasedDate": number,
      "State": "string",
      "StateChangeReason": {
         "Code": "string",
         "Message": "string"
      },
      "SupportedInstanceFamilies": [ "string" ],
      "Visibility": "string"
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[Image](#API_DeleteImage_ResponseSyntax)**

Information about the image.

Type: [Image](api-image.md) object

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/appstream-2016-12-01/deleteimage.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/appstream-2016-12-01/deleteimage.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/appstream-2016-12-01/deleteimage.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/appstream-2016-12-01/deleteimage.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appstream-2016-12-01/deleteimage.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/appstream-2016-12-01/deleteimage.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/appstream-2016-12-01/deleteimage.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/appstream-2016-12-01/deleteimage.md)

- [AWS SDK for Python](../../../../services/goto/boto3/appstream-2016-12-01/deleteimage.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appstream-2016-12-01/deleteimage.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteFleet

DeleteImageBuilder

All content copied from https://docs.aws.amazon.com/.
