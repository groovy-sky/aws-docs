# CreateBucketMetadataConfiguration

Creates an S3 Metadata V2 metadata configuration for a general purpose bucket. For more information, see
[Accelerating\
data discovery with S3 Metadata](https://docs.aws.amazon.com/AmazonS3/latest/userguide/metadata-tables-overview.html) in the _Amazon S3 User Guide_.

Permissions

To use this operation, you must have the following permissions. For more information, see
[Setting up permissions for configuring metadata tables](https://docs.aws.amazon.com/AmazonS3/latest/userguide/metadata-tables-permissions.html) in the
_Amazon S3 User Guide_.

If you want to encrypt your metadata tables with server-side encryption with AWS Key Management Service
(AWS KMS) keys (SSE-KMS), you need additional permissions in your KMS key policy. For more
information, see [Setting up permissions for configuring metadata tables](https://docs.aws.amazon.com/AmazonS3/latest/userguide/metadata-tables-permissions.html) in the
_Amazon S3 User Guide_.

If you also want to integrate your table bucket with AWS analytics services so that you can
query your metadata table, you need additional permissions. For more information, see [Integrating\
Amazon S3 Tables with AWS analytics services](../userguide/s3-tables-integrating-aws.md) in the
_Amazon S3 User Guide_.

To query your metadata tables, you need additional permissions. For more information, see
[Permissions for querying metadata tables](https://docs.aws.amazon.com/AmazonS3/latest/userguide/metadata-tables-bucket-query-permissions.html) in the _Amazon S3 User Guide_.

- `s3:CreateBucketMetadataTableConfiguration`

###### Note

The IAM policy action name is the same for the V1 and V2 API operations.

- `s3tables:CreateTableBucket`

- `s3tables:CreateNamespace`

- `s3tables:GetTable`

- `s3tables:CreateTable`

- `s3tables:PutTablePolicy`

- `s3tables:PutTableEncryption`

- `kms:DescribeKey`

The following operations are related to `CreateBucketMetadataConfiguration`:

- [DeleteBucketMetadataConfiguration](https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketMetadataConfiguration.html)

- [GetBucketMetadataConfiguration](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketMetadataConfiguration.html)

- [UpdateBucketMetadataInventoryTableConfiguration](https://docs.aws.amazon.com/AmazonS3/latest/API/API_UpdateBucketMetadataInventoryTableConfiguration.html)

- [UpdateBucketMetadataJournalTableConfiguration](https://docs.aws.amazon.com/AmazonS3/latest/API/API_UpdateBucketMetadataJournalTableConfiguration.html)

###### Important

You must URL encode any signed header values that contain spaces. For example, if your header value is `my  file.txt`, containing two spaces after `my`, you must URL encode this value to `my%20%20file.txt`.

## Request Syntax

```nohighlight

POST /?metadataConfiguration HTTP/1.1
Host: Bucket.s3.amazonaws.com
Content-MD5: ContentMD5
x-amz-sdk-checksum-algorithm: ChecksumAlgorithm
x-amz-expected-bucket-owner: ExpectedBucketOwner
<?xml version="1.0" encoding="UTF-8"?>
<MetadataConfiguration xmlns="http://s3.amazonaws.com/doc/2006-03-01/">
   <JournalTableConfiguration>
      <EncryptionConfiguration>
         <KmsKeyArn>string</KmsKeyArn>
         <SseAlgorithm>string</SseAlgorithm>
      </EncryptionConfiguration>
      <RecordExpiration>
         <Days>integer</Days>
         <Expiration>string</Expiration>
      </RecordExpiration>
   </JournalTableConfiguration>
   <InventoryTableConfiguration>
      <ConfigurationState>string</ConfigurationState>
      <EncryptionConfiguration>
         <KmsKeyArn>string</KmsKeyArn>
         <SseAlgorithm>string</SseAlgorithm>
      </EncryptionConfiguration>
   </InventoryTableConfiguration>
</MetadataConfiguration>
```

## URI Request Parameters

The request uses the following URI parameters.

**[Bucket](#API_CreateBucketMetadataConfiguration_RequestSyntax)**

The general purpose bucket that you want to create the metadata configuration for.

Required: Yes

**[Content-MD5](#API_CreateBucketMetadataConfiguration_RequestSyntax)**

The `Content-MD5` header for the metadata configuration.

**[x-amz-expected-bucket-owner](#API_CreateBucketMetadataConfiguration_RequestSyntax)**

The expected owner of the general purpose bucket that corresponds to your metadata configuration.

**[x-amz-sdk-checksum-algorithm](#API_CreateBucketMetadataConfiguration_RequestSyntax)**

The checksum algorithm to use with your metadata configuration.

Valid Values: `CRC32 | CRC32C | SHA1 | SHA256 | CRC64NVME`

## Request Body

The request accepts the following data in XML format.

**[MetadataConfiguration](#API_CreateBucketMetadataConfiguration_RequestSyntax)**

Root level tag for the MetadataConfiguration parameters.

Required: Yes

**[InventoryTableConfiguration](#API_CreateBucketMetadataConfiguration_RequestSyntax)**

The inventory table configuration for a metadata configuration.

Type: [InventoryTableConfiguration](https://docs.aws.amazon.com/AmazonS3/latest/API/API_InventoryTableConfiguration.html) data type

Required: No

**[JournalTableConfiguration](#API_CreateBucketMetadataConfiguration_RequestSyntax)**

The journal table configuration for a metadata configuration.

Type: [JournalTableConfiguration](https://docs.aws.amazon.com/AmazonS3/latest/API/API_JournalTableConfiguration.html) data type

Required: Yes

## Response Syntax

```

HTTP/1.1 200

```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/s3-2006-03-01/CreateBucketMetadataConfiguration)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/s3-2006-03-01/CreateBucketMetadataConfiguration)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/s3-2006-03-01/CreateBucketMetadataConfiguration)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/s3-2006-03-01/CreateBucketMetadataConfiguration)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/s3-2006-03-01/CreateBucketMetadataConfiguration)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/s3-2006-03-01/CreateBucketMetadataConfiguration)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/s3-2006-03-01/CreateBucketMetadataConfiguration)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/s3-2006-03-01/CreateBucketMetadataConfiguration)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/s3-2006-03-01/CreateBucketMetadataConfiguration)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/s3-2006-03-01/CreateBucketMetadataConfiguration)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CreateBucket

CreateBucketMetadataTableConfiguration
