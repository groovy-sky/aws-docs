# GetBucketMetadataConfiguration

Retrieves the S3 Metadata configuration for a general purpose bucket. For more information, see
[Accelerating\
data discovery with S3 Metadata](https://docs.aws.amazon.com/AmazonS3/latest/userguide/metadata-tables-overview.html) in the _Amazon S3 User Guide_.

###### Note

You can use the V2 `GetBucketMetadataConfiguration` API operation with V1 or V2
metadata configurations. However, if you try to use the V1
`GetBucketMetadataTableConfiguration` API operation with V2 configurations, you
will receive an HTTP `405 Method Not Allowed` error.

Permissions

To use this operation, you must have the `s3:GetBucketMetadataTableConfiguration`
permission. For more information, see [Setting up permissions for\
configuring metadata tables](../userguide/metadata-tables-permissions.md) in the _Amazon S3 User Guide_.

###### Note

The IAM policy action name is the same for the V1 and V2 API operations.

The following operations are related to `GetBucketMetadataConfiguration`:

- [CreateBucketMetadataConfiguration](api-createbucketmetadataconfiguration.md)

- [DeleteBucketMetadataConfiguration](api-deletebucketmetadataconfiguration.md)

- [UpdateBucketMetadataInventoryTableConfiguration](api-updatebucketmetadatainventorytableconfiguration.md)

- [UpdateBucketMetadataJournalTableConfiguration](api-updatebucketmetadatajournaltableconfiguration.md)

###### Important

You must URL encode any signed header values that contain spaces. For example, if your header value is `my  file.txt`, containing two spaces after `my`, you must URL encode this value to `my%20%20file.txt`.

## Request Syntax

```nohighlight

GET /?metadataConfiguration HTTP/1.1
Host: Bucket.s3.amazonaws.com
x-amz-expected-bucket-owner: ExpectedBucketOwner

```

## URI Request Parameters

The request uses the following URI parameters.

**[Bucket](#API_GetBucketMetadataConfiguration_RequestSyntax)**

The general purpose bucket that corresponds to the metadata configuration that you want to
retrieve.

Required: Yes

**[x-amz-expected-bucket-owner](#API_GetBucketMetadataConfiguration_RequestSyntax)**

The expected owner of the general purpose bucket that you want to retrieve the metadata table
configuration for.

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
<?xml version="1.0" encoding="UTF-8"?>
<GetBucketMetadataConfigurationResult>
   <MetadataConfigurationResult>
      <DestinationResult>
         <TableBucketArn>string</TableBucketArn>
         <TableBucketType>string</TableBucketType>
         <TableNamespace>string</TableNamespace>
      </DestinationResult>
      <InventoryTableConfigurationResult>
         <ConfigurationState>string</ConfigurationState>
         <Error>
            <ErrorCode>string</ErrorCode>
            <ErrorMessage>string</ErrorMessage>
         </Error>
         <TableArn>string</TableArn>
         <TableName>string</TableName>
         <TableStatus>string</TableStatus>
      </InventoryTableConfigurationResult>
      <JournalTableConfigurationResult>
         <Error>
            <ErrorCode>string</ErrorCode>
            <ErrorMessage>string</ErrorMessage>
         </Error>
         <RecordExpiration>
            <Days>integer</Days>
            <Expiration>string</Expiration>
         </RecordExpiration>
         <TableArn>string</TableArn>
         <TableName>string</TableName>
         <TableStatus>string</TableStatus>
      </JournalTableConfigurationResult>
   </MetadataConfigurationResult>
</GetBucketMetadataConfigurationResult>
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in XML format by the service.

**[GetBucketMetadataConfigurationResult](#API_GetBucketMetadataConfiguration_ResponseSyntax)**

Root level tag for the GetBucketMetadataConfigurationResult parameters.

Required: Yes

**[MetadataConfigurationResult](#API_GetBucketMetadataConfiguration_ResponseSyntax)**

The metadata configuration for a general purpose bucket.

Type: [MetadataConfigurationResult](https://docs.aws.amazon.com/AmazonS3/latest/API/API_MetadataConfigurationResult.html) data type

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/s3-2006-03-01/GetBucketMetadataConfiguration)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/s3-2006-03-01/GetBucketMetadataConfiguration)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/s3-2006-03-01/GetBucketMetadataConfiguration)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/s3-2006-03-01/GetBucketMetadataConfiguration)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/s3-2006-03-01/GetBucketMetadataConfiguration)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/s3-2006-03-01/GetBucketMetadataConfiguration)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/s3-2006-03-01/GetBucketMetadataConfiguration)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/s3-2006-03-01/GetBucketMetadataConfiguration)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/s3-2006-03-01/GetBucketMetadataConfiguration)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/s3-2006-03-01/GetBucketMetadataConfiguration)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

GetBucketLogging

GetBucketMetadataTableConfiguration
