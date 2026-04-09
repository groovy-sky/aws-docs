# DescribeImages

Retrieves a list that describes one or more specified images, if the image names or image ARNs are provided. Otherwise, all images in the account are described.

## Request Syntax

```nohighlight

{
   "Arns": [ "string" ],
   "MaxResults": number,
   "Names": [ "string" ],
   "NextToken": "string",
   "Type": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[Arns](#API_DescribeImages_RequestSyntax)**

The ARNs of the public, private, and shared images to describe.

Type: Array of strings

Pattern: `^arn:aws(?:\-cn|\-iso\-b|\-iso|\-us\-gov)?:[A-Za-z0-9][A-Za-z0-9_/.-]{0,62}:[A-Za-z0-9_/.-]{0,63}:[A-Za-z0-9_/.-]{0,63}:[A-Za-z0-9][A-Za-z0-9:_/+=,@.\\-]{0,1023}$`

Required: No

**[MaxResults](#API_DescribeImages_RequestSyntax)**

The maximum size of each page of results.

Type: Integer

Valid Range: Minimum value of 0. Maximum value of 25.

Required: No

**[Names](#API_DescribeImages_RequestSyntax)**

The names of the public or private images to describe.

Type: Array of strings

Length Constraints: Minimum length of 1.

Required: No

**[NextToken](#API_DescribeImages_RequestSyntax)**

The pagination token to use to retrieve the next page of results for this operation. If this value is null, it retrieves the first page.

Type: String

Length Constraints: Minimum length of 1.

Required: No

**[Type](#API_DescribeImages_RequestSyntax)**

The type of image (public, private, or shared) to describe.

Type: String

Valid Values: `PUBLIC | PRIVATE | SHARED`

Required: No

## Response Syntax

```nohighlight

{
   "Images": [
      {
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
   ],
   "NextToken": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[Images](#API_DescribeImages_ResponseSyntax)**

Information about the images.

Type: Array of [Image](api-image.md) objects

**[NextToken](#API_DescribeImages_ResponseSyntax)**

The pagination token to use to retrieve the next page of results for this operation. If there are no more pages, this value is null.

Type: String

Length Constraints: Minimum length of 1.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidParameterCombinationException**

Indicates an incorrect combination of parameters, or a missing parameter.

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/appstream-2016-12-01/describeimages.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/appstream-2016-12-01/describeimages.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/appstream-2016-12-01/describeimages.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/appstream-2016-12-01/describeimages.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appstream-2016-12-01/describeimages.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/appstream-2016-12-01/describeimages.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/appstream-2016-12-01/describeimages.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/appstream-2016-12-01/describeimages.md)

- [AWS SDK for Python](../../../../services/goto/boto3/appstream-2016-12-01/describeimages.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appstream-2016-12-01/describeimages.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DescribeImagePermissions

DescribeSessions

All content copied from https://docs.aws.amazon.com/.
