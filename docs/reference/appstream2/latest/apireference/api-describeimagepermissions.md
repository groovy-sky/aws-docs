# DescribeImagePermissions

Retrieves a list that describes the permissions for shared AWS account IDs on a private image that you own.

## Request Syntax

```nohighlight

{
   "MaxResults": number,
   "Name": "string",
   "NextToken": "string",
   "SharedAwsAccountIds": [ "string" ]
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[MaxResults](#API_DescribeImagePermissions_RequestSyntax)**

The maximum size of each page of results.

Type: Integer

Valid Range: Minimum value of 0. Maximum value of 500.

Required: No

**[Name](#API_DescribeImagePermissions_RequestSyntax)**

The name of the private image for which to describe permissions. The image must be one that you own.

Type: String

Pattern: `^[a-zA-Z0-9][a-zA-Z0-9_.-]{0,100}$`

Required: Yes

**[NextToken](#API_DescribeImagePermissions_RequestSyntax)**

The pagination token to use to retrieve the next page of results for this operation. If this value is null, it retrieves the first page.

Type: String

Length Constraints: Minimum length of 1.

Required: No

**[SharedAwsAccountIds](#API_DescribeImagePermissions_RequestSyntax)**

The 12-digit identifier of one or more AWS accounts with which the image is shared.

Type: Array of strings

Array Members: Minimum number of 1 item. Maximum number of 5 items.

Pattern: `^\d+$`

Required: No

## Response Syntax

```nohighlight

{
   "Name": "string",
   "NextToken": "string",
   "SharedImagePermissionsList": [
      {
         "imagePermissions": {
            "allowFleet": boolean,
            "allowImageBuilder": boolean
         },
         "sharedAccountId": "string"
      }
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[Name](#API_DescribeImagePermissions_ResponseSyntax)**

The name of the private image.

Type: String

Pattern: `^[a-zA-Z0-9][a-zA-Z0-9_.-]{0,100}$`

**[NextToken](#API_DescribeImagePermissions_ResponseSyntax)**

The pagination token to use to retrieve the next page of results for this operation. If there are no more pages, this value is null.

Type: String

Length Constraints: Minimum length of 1.

**[SharedImagePermissionsList](#API_DescribeImagePermissions_ResponseSyntax)**

The permissions for a private image that you own.

Type: Array of [SharedImagePermissions](api-sharedimagepermissions.md) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**ResourceNotFoundException**

The specified resource was not found.

**Message**

The error message in the exception.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/appstream-2016-12-01/describeimagepermissions.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/appstream-2016-12-01/describeimagepermissions.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/appstream-2016-12-01/describeimagepermissions.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/appstream-2016-12-01/describeimagepermissions.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appstream-2016-12-01/describeimagepermissions.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/appstream-2016-12-01/describeimagepermissions.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/appstream-2016-12-01/describeimagepermissions.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/appstream-2016-12-01/describeimagepermissions.md)

- [AWS SDK for Python](../../../../services/goto/boto3/appstream-2016-12-01/describeimagepermissions.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appstream-2016-12-01/describeimagepermissions.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DescribeImageBuilders

DescribeImages

All content copied from https://docs.aws.amazon.com/.
