# DeleteBucketMetadataConfiguration

Deletes an S3 Metadata configuration from a general purpose bucket. For more information, see
[Accelerating\
data discovery with S3 Metadata](../userguide/metadata-tables-overview.md) in the _Amazon S3 User Guide_.

###### Note

You can use the V2 `DeleteBucketMetadataConfiguration` API operation with V1 or V2
metadata configurations. However, if you try to use the V1
`DeleteBucketMetadataTableConfiguration` API operation with V2 configurations, you
will receive an HTTP `405 Method Not Allowed` error.

Permissions

To use this operation, you must have the
`s3:DeleteBucketMetadataTableConfiguration` permission. For more information, see
[Setting up permissions for configuring metadata tables](../userguide/metadata-tables-permissions.md) in the
_Amazon S3 User Guide_.

###### Note

The IAM policy action name is the same for the V1 and V2 API operations.

The following operations are related to `DeleteBucketMetadataConfiguration`:

- [CreateBucketMetadataConfiguration](api-createbucketmetadataconfiguration.md)

- [GetBucketMetadataConfiguration](api-getbucketmetadataconfiguration.md)

- [UpdateBucketMetadataInventoryTableConfiguration](api-updatebucketmetadatainventorytableconfiguration.md)

- [UpdateBucketMetadataJournalTableConfiguration](api-updatebucketmetadatajournaltableconfiguration.md)

###### Important

You must URL encode any signed header values that contain spaces. For example, if your header value is `my  file.txt`, containing two spaces after `my`, you must URL encode this value to `my%20%20file.txt`.

## Request Syntax

```nohighlight

DELETE /?metadataConfiguration HTTP/1.1
Host: Bucket.s3.amazonaws.com
x-amz-expected-bucket-owner: ExpectedBucketOwner

```

## URI Request Parameters

The request uses the following URI parameters.

**[Bucket](#API_DeleteBucketMetadataConfiguration_RequestSyntax)**

The general purpose bucket that you want to remove the metadata configuration from.

Required: Yes

**[x-amz-expected-bucket-owner](#API_DeleteBucketMetadataConfiguration_RequestSyntax)**

The expected bucket owner of the general purpose bucket that you want to remove the metadata table
configuration from.

## Request Body

The request does not have a request body.

## Response Syntax

```

HTTP/1.1 204

```

## Response Elements

If the action is successful, the service sends back an HTTP 204 response with an empty HTTP body.

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/s3-2006-03-01/deletebucketmetadataconfiguration.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3-2006-03-01/deletebucketmetadataconfiguration.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3-2006-03-01/deletebucketmetadataconfiguration.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3-2006-03-01/deletebucketmetadataconfiguration.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3-2006-03-01/deletebucketmetadataconfiguration.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3-2006-03-01/deletebucketmetadataconfiguration.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3-2006-03-01/deletebucketmetadataconfiguration.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3-2006-03-01/deletebucketmetadataconfiguration.md)

- [AWS SDK for Python](../../../goto/boto3/s3-2006-03-01/deletebucketmetadataconfiguration.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3-2006-03-01/deletebucketmetadataconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteBucketLifecycle

DeleteBucketMetadataTableConfiguration

All content copied from https://docs.aws.amazon.com/.
