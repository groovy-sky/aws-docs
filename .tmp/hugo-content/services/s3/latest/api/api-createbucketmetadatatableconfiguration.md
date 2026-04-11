# CreateBucketMetadataTableConfiguration

###### Important

We recommend that you create your S3 Metadata configurations by using the V2
[CreateBucketMetadataConfiguration](api-createbucketmetadataconfiguration.md) API operation. We no longer recommend using the V1
`CreateBucketMetadataTableConfiguration` API operation.

If you created your S3 Metadata configuration before July 15, 2025, we recommend that you delete
and re-create your configuration by using [CreateBucketMetadataConfiguration](api-createbucketmetadataconfiguration.md) so that you can expire journal table records and create
a live inventory table.

Creates a V1 S3 Metadata configuration for a general purpose bucket. For more information, see
[Accelerating\
data discovery with S3 Metadata](../userguide/metadata-tables-overview.md) in the _Amazon S3 User Guide_.

Permissions

To use this operation, you must have the following permissions. For more information, see
[Setting up permissions for configuring metadata tables](../userguide/metadata-tables-permissions.md) in the
_Amazon S3 User Guide_.

If you want to encrypt your metadata tables with server-side encryption with AWS Key Management Service
(AWS KMS) keys (SSE-KMS), you need additional permissions. For more
information, see [Setting up permissions for configuring metadata tables](../userguide/metadata-tables-permissions.md) in the
_Amazon S3 User Guide_.

If you also want to integrate your table bucket with AWS analytics services so that you can
query your metadata table, you need additional permissions. For more information, see [Integrating\
Amazon S3 Tables with AWS analytics services](../userguide/s3-tables-integrating-aws.md) in the
_Amazon S3 User Guide_.

- `s3:CreateBucketMetadataTableConfiguration`

- `s3tables:CreateNamespace`

- `s3tables:GetTable`

- `s3tables:CreateTable`

- `s3tables:PutTablePolicy`

The following operations are related to `CreateBucketMetadataTableConfiguration`:

- [DeleteBucketMetadataTableConfiguration](api-deletebucketmetadatatableconfiguration.md)

- [GetBucketMetadataTableConfiguration](api-getbucketmetadatatableconfiguration.md)

###### Important

You must URL encode any signed header values that contain spaces. For example, if your header value is `my  file.txt`, containing two spaces after `my`, you must URL encode this value to `my%20%20file.txt`.

## Request Syntax

```nohighlight

POST /?metadataTable HTTP/1.1
Host: Bucket.s3.amazonaws.com
Content-MD5: ContentMD5
x-amz-sdk-checksum-algorithm: ChecksumAlgorithm
x-amz-expected-bucket-owner: ExpectedBucketOwner
<?xml version="1.0" encoding="UTF-8"?>
<MetadataTableConfiguration xmlns="http://s3.amazonaws.com/doc/2006-03-01/">
   <S3TablesDestination>
      <TableBucketArn>string</TableBucketArn>
      <TableName>string</TableName>
   </S3TablesDestination>
</MetadataTableConfiguration>
```

## URI Request Parameters

The request uses the following URI parameters.

**[Bucket](#API_CreateBucketMetadataTableConfiguration_RequestSyntax)**

The general purpose bucket that you want to create the metadata table configuration for.

Required: Yes

**[Content-MD5](#API_CreateBucketMetadataTableConfiguration_RequestSyntax)**

The `Content-MD5` header for the metadata table configuration.

**[x-amz-expected-bucket-owner](#API_CreateBucketMetadataTableConfiguration_RequestSyntax)**

The expected owner of the general purpose bucket that corresponds to your metadata table configuration.

**[x-amz-sdk-checksum-algorithm](#API_CreateBucketMetadataTableConfiguration_RequestSyntax)**

The checksum algorithm to use with your metadata table configuration.

Valid Values: `CRC32 | CRC32C | SHA1 | SHA256 | CRC64NVME`

## Request Body

The request accepts the following data in XML format.

**[MetadataTableConfiguration](#API_CreateBucketMetadataTableConfiguration_RequestSyntax)**

Root level tag for the MetadataTableConfiguration parameters.

Required: Yes

**[S3TablesDestination](#API_CreateBucketMetadataTableConfiguration_RequestSyntax)**

The destination information for the metadata table configuration. The destination table bucket must
be in the same Region and AWS account as the general purpose bucket. The specified metadata table name
must be unique within the `aws_s3_metadata` namespace in the destination table bucket.

Type: [S3TablesDestination](api-s3tablesdestination.md) data type

Required: Yes

## Response Syntax

```

HTTP/1.1 200

```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/s3-2006-03-01/createbucketmetadatatableconfiguration.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3-2006-03-01/createbucketmetadatatableconfiguration.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3-2006-03-01/createbucketmetadatatableconfiguration.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3-2006-03-01/createbucketmetadatatableconfiguration.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3-2006-03-01/createbucketmetadatatableconfiguration.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3-2006-03-01/createbucketmetadatatableconfiguration.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3-2006-03-01/createbucketmetadatatableconfiguration.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3-2006-03-01/createbucketmetadatatableconfiguration.md)

- [AWS SDK for Python](../../../goto/boto3/s3-2006-03-01/createbucketmetadatatableconfiguration.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3-2006-03-01/createbucketmetadatatableconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CreateBucketMetadataConfiguration

CreateMultipartUpload

All content copied from https://docs.aws.amazon.com/.
