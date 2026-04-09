# ListStorageLensConfigurations

###### Note

This operation is not supported by directory buckets.

Gets a list of Amazon S3 Storage Lens configurations. For more information about S3 Storage Lens, see
[Assessing your\
storage activity and usage with Amazon S3 Storage Lens](../dev/storage-lens.md) in the
_Amazon S3 User Guide_.

###### Note

To use this action, you must have permission to perform the
`s3:ListStorageLensConfigurations` action. For more information, see
[Setting permissions to\
use Amazon S3 Storage Lens](../dev/storage-lens-iam-permissions.md) in the _Amazon S3 User Guide_.

## Request Syntax

```nohighlight

GET /v20180820/storagelens?nextToken=NextToken HTTP/1.1
Host: s3-control.amazonaws.com
x-amz-account-id: AccountId

```

## URI Request Parameters

The request uses the following URI parameters.

**[nextToken](#API_control_ListStorageLensConfigurations_RequestSyntax)**

A pagination token to request the next page of results.

**[x-amz-account-id](#API_control_ListStorageLensConfigurations_RequestSyntax)**

The account ID of the requester.

Length Constraints: Maximum length of 64.

Pattern: `^\d{12}$`

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
<?xml version="1.0" encoding="UTF-8"?>
<ListStorageLensConfigurationsResult>
   <NextToken>string</NextToken>
   <StorageLensConfiguration>
      <HomeRegion>string</HomeRegion>
      <Id>string</Id>
      <IsEnabled>boolean</IsEnabled>
      <StorageLensArn>string</StorageLensArn>
   </StorageLensConfiguration>
   ...
</ListStorageLensConfigurationsResult>
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in XML format by the service.

**[ListStorageLensConfigurationsResult](#API_control_ListStorageLensConfigurations_ResponseSyntax)**

Root level tag for the ListStorageLensConfigurationsResult parameters.

Required: Yes

**[NextToken](#API_control_ListStorageLensConfigurations_ResponseSyntax)**

If the request produced more than the maximum number of S3 Storage Lens configuration results,
you can pass this value into a subsequent request to retrieve the next page of
results.

Type: String

**[StorageLensConfiguration](#API_control_ListStorageLensConfigurations_ResponseSyntax)**

A list of S3 Storage Lens configurations.

Type: Array of [ListStorageLensConfigurationEntry](api-control-liststoragelensconfigurationentry.md) data types

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/s3control-2018-08-20/liststoragelensconfigurations.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3control-2018-08-20/liststoragelensconfigurations.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/liststoragelensconfigurations.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3control-2018-08-20/liststoragelensconfigurations.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/liststoragelensconfigurations.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3control-2018-08-20/liststoragelensconfigurations.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3control-2018-08-20/liststoragelensconfigurations.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3control-2018-08-20/liststoragelensconfigurations.md)

- [AWS SDK for Python](../../../goto/boto3/s3control-2018-08-20/liststoragelensconfigurations.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/liststoragelensconfigurations.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ListRegionalBuckets

ListStorageLensGroups

All content copied from https://docs.aws.amazon.com/.
