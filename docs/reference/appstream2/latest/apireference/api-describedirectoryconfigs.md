# DescribeDirectoryConfigs

Retrieves a list that describes one or more specified Directory Config objects for WorkSpaces Applications, if the names for these objects are provided. Otherwise, all Directory Config objects in the account are described. These objects include the configuration information required to join fleets and image builders to Microsoft Active Directory domains.

Although the response syntax in this topic includes the account password, this password is not returned in the actual response.

## Request Syntax

```nohighlight

{
   "DirectoryNames": [ "string" ],
   "MaxResults": number,
   "NextToken": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[DirectoryNames](#API_DescribeDirectoryConfigs_RequestSyntax)**

The directory names.

Type: Array of strings

Required: No

**[MaxResults](#API_DescribeDirectoryConfigs_RequestSyntax)**

The maximum size of each page of results.

Type: Integer

Required: No

**[NextToken](#API_DescribeDirectoryConfigs_RequestSyntax)**

The pagination token to use to retrieve the next page of results for this operation. If this value is null, it retrieves the first page.

Type: String

Length Constraints: Minimum length of 1.

Required: No

## Response Syntax

```nohighlight

{
   "DirectoryConfigs": [
      {
         "CertificateBasedAuthProperties": {
            "CertificateAuthorityArn": "string",
            "Status": "string"
         },
         "CreatedTime": number,
         "DirectoryName": "string",
         "OrganizationalUnitDistinguishedNames": [ "string" ],
         "ServiceAccountCredentials": {
            "AccountName": "string",
            "AccountPassword": "string"
         }
      }
   ],
   "NextToken": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[DirectoryConfigs](#API_DescribeDirectoryConfigs_ResponseSyntax)**

Information about the directory configurations. Note that although the response syntax in this topic includes the account password, this password is not returned in the actual response.

Type: Array of [DirectoryConfig](api-directoryconfig.md) objects

**[NextToken](#API_DescribeDirectoryConfigs_ResponseSyntax)**

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/appstream-2016-12-01/describedirectoryconfigs.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/appstream-2016-12-01/describedirectoryconfigs.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/appstream-2016-12-01/describedirectoryconfigs.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/appstream-2016-12-01/describedirectoryconfigs.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appstream-2016-12-01/describedirectoryconfigs.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/appstream-2016-12-01/describedirectoryconfigs.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/appstream-2016-12-01/describedirectoryconfigs.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/appstream-2016-12-01/describedirectoryconfigs.md)

- [AWS SDK for Python](../../../../services/goto/boto3/appstream-2016-12-01/describedirectoryconfigs.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appstream-2016-12-01/describedirectoryconfigs.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DescribeAppLicenseUsage

DescribeEntitlements

All content copied from https://docs.aws.amazon.com/.
