# ListStorageLensGroups

Lists all the Storage Lens groups in the specified home Region.

To use this operation, you must have the permission to perform the
`s3:ListStorageLensGroups` action. For more information about the required Storage Lens
Groups permissions, see [Setting account permissions to use S3 Storage Lens groups](../userguide/storage-lens-iam-permissions.md#storage_lens_groups_permissions).

For information about Storage Lens groups errors, see [List of Amazon S3 Storage\
Lens error codes](errorresponses.md#S3LensErrorCodeList).

## Request Syntax

```nohighlight

GET /v20180820/storagelensgroup?nextToken=NextToken HTTP/1.1
Host: s3-control.amazonaws.com
x-amz-account-id: AccountId

```

## URI Request Parameters

The request uses the following URI parameters.

**[nextToken](#API_control_ListStorageLensGroups_RequestSyntax)**

The token for the next set of results, or `null` if there are no more results.

**[x-amz-account-id](#API_control_ListStorageLensGroups_RequestSyntax)**

The AWS account ID that owns the Storage Lens groups.

Length Constraints: Maximum length of 64.

Pattern: `^\d{12}$`

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
<?xml version="1.0" encoding="UTF-8"?>
<ListStorageLensGroupsResult>
   <NextToken>string</NextToken>
   <StorageLensGroup>
      <HomeRegion>string</HomeRegion>
      <Name>string</Name>
      <StorageLensGroupArn>string</StorageLensGroupArn>
   </StorageLensGroup>
   ...
</ListStorageLensGroupsResult>
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in XML format by the service.

**[ListStorageLensGroupsResult](#API_control_ListStorageLensGroups_ResponseSyntax)**

Root level tag for the ListStorageLensGroupsResult parameters.

Required: Yes

**[NextToken](#API_control_ListStorageLensGroups_ResponseSyntax)**

If `NextToken` is returned, there are more Storage Lens groups results available. The value of `NextToken` is a
unique pagination token for each page. Make the call again using the returned token to
retrieve the next page. Keep all other arguments unchanged. Each pagination token expires
after 24 hours.

Type: String

**[StorageLensGroup](#API_control_ListStorageLensGroups_ResponseSyntax)**

The list of Storage Lens groups that exist in the specified home Region.

Type: Array of [ListStorageLensGroupEntry](api-control-liststoragelensgroupentry.md) data types

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/s3control-2018-08-20/liststoragelensgroups.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3control-2018-08-20/liststoragelensgroups.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/liststoragelensgroups.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3control-2018-08-20/liststoragelensgroups.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/liststoragelensgroups.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3control-2018-08-20/liststoragelensgroups.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3control-2018-08-20/liststoragelensgroups.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3control-2018-08-20/liststoragelensgroups.md)

- [AWS SDK for Python](../../../goto/boto3/s3control-2018-08-20/liststoragelensgroups.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/liststoragelensgroups.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ListStorageLensConfigurations

ListTagsForResource

All content copied from https://docs.aws.amazon.com/.
