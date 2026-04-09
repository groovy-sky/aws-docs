# DescribeImageBuilders

Retrieves a list that describes one or more specified image builders, if the image builder names are provided. Otherwise, all image builders in the account are described.

## Request Syntax

```nohighlight

{
   "MaxResults": number,
   "Names": [ "string" ],
   "NextToken": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[MaxResults](#API_DescribeImageBuilders_RequestSyntax)**

The maximum size of each page of results.

Type: Integer

Required: No

**[Names](#API_DescribeImageBuilders_RequestSyntax)**

The names of the image builders to describe.

Type: Array of strings

Length Constraints: Minimum length of 1.

Required: No

**[NextToken](#API_DescribeImageBuilders_RequestSyntax)**

The pagination token to use to retrieve the next page of results for this operation. If this value is null, it retrieves the first page.

Type: String

Length Constraints: Minimum length of 1.

Required: No

## Response Syntax

```nohighlight

{
   "ImageBuilders": [
      {
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
   ],
   "NextToken": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[ImageBuilders](#API_DescribeImageBuilders_ResponseSyntax)**

Information about the image builders.

Type: Array of [ImageBuilder](api-imagebuilder.md) objects

**[NextToken](#API_DescribeImageBuilders_ResponseSyntax)**

The pagination token to use to retrieve the next page of results for this operation. If there are no more pages, this value is null.

Type: String

Length Constraints: Minimum length of 1.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**ResourceNotFoundException**

The specified resource was not found.

**Message**

The error message in the exception.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/appstream-2016-12-01/describeimagebuilders.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/appstream-2016-12-01/describeimagebuilders.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/appstream-2016-12-01/describeimagebuilders.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/appstream-2016-12-01/describeimagebuilders.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appstream-2016-12-01/describeimagebuilders.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/appstream-2016-12-01/describeimagebuilders.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/appstream-2016-12-01/describeimagebuilders.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/appstream-2016-12-01/describeimagebuilders.md)

- [AWS SDK for Python](../../../../services/goto/boto3/appstream-2016-12-01/describeimagebuilders.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appstream-2016-12-01/describeimagebuilders.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DescribeFleets

DescribeImagePermissions

All content copied from https://docs.aws.amazon.com/.
