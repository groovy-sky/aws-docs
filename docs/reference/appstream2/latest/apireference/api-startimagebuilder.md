# StartImageBuilder

Starts the specified image builder.

## Request Syntax

```nohighlight

{
   "AppstreamAgentVersion": "string",
   "Name": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[AppstreamAgentVersion](#API_StartImageBuilder_RequestSyntax)**

The version of the WorkSpaces Applications agent to use for this image builder. To use the latest version of the WorkSpaces Applications agent, specify \[LATEST\].

Type: String

Length Constraints: Minimum length of 1. Maximum length of 100.

Required: No

**[Name](#API_StartImageBuilder_RequestSyntax)**

The name of the image builder.

Type: String

Length Constraints: Minimum length of 1.

Required: Yes

## Response Syntax

```nohighlight

{
   "ImageBuilder": {
      "AccessEndpoints": [
         {
            "EndpointType": "string",
            "VpceId": "string"
         }
      ],
      "AppstreamAgentVersion": "string",
      "Arn": "string",
      "CreatedTime": number,
      "Description": "string",
      "DisableIMDSV1": boolean,
      "DisplayName": "string",
      "DomainJoinInfo": {
         "DirectoryName": "string",
         "OrganizationalUnitDistinguishedName": "string"
      },
      "EnableDefaultInternetAccess": boolean,
      "IamRoleArn": "string",
      "ImageArn": "string",
      "ImageBuilderErrors": [
         {
            "ErrorCode": "string",
            "ErrorMessage": "string",
            "ErrorTimestamp": number
         }
      ],
      "InstanceType": "string",
      "LatestAppstreamAgentVersion": "string",
      "Name": "string",
      "NetworkAccessConfiguration": {
         "EniId": "string",
         "EniIpv6Addresses": [ "string" ],
         "EniPrivateIpAddress": "string"
      },
      "Platform": "string",
      "RootVolumeConfig": {
         "VolumeSizeInGb": number
      },
      "State": "string",
      "StateChangeReason": {
         "Code": "string",
         "Message": "string"
      },
      "VpcConfig": {
         "SecurityGroupIds": [ "string" ],
         "SubnetIds": [ "string" ]
      }
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[ImageBuilder](#API_StartImageBuilder_ResponseSyntax)**

Information about the image builder.

Type: [ImageBuilder](api-imagebuilder.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**ConcurrentModificationException**

An API error occurred. Wait a few minutes and try again.

**Message**

The error message in the exception.

HTTP Status Code: 400

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/appstream-2016-12-01/startimagebuilder.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/appstream-2016-12-01/startimagebuilder.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/appstream-2016-12-01/startimagebuilder.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/appstream-2016-12-01/startimagebuilder.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appstream-2016-12-01/startimagebuilder.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/appstream-2016-12-01/startimagebuilder.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/appstream-2016-12-01/startimagebuilder.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/appstream-2016-12-01/startimagebuilder.md)

- [AWS SDK for Python](../../../../services/goto/boto3/appstream-2016-12-01/startimagebuilder.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appstream-2016-12-01/startimagebuilder.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StartFleet

StartSoftwareDeploymentToImageBuilder

All content copied from https://docs.aws.amazon.com/.
