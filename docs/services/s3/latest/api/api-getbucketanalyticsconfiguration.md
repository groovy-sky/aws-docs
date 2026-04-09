# GetBucketAnalyticsConfiguration

###### Note

This operation is not supported for directory buckets.

This implementation of the GET action returns an analytics configuration (identified by the
analytics configuration ID) from the bucket.

To use this operation, you must have permissions to perform the
`s3:GetAnalyticsConfiguration` action. The bucket owner has this permission by default. The
bucket owner can grant this permission to others. For more information about permissions, see [Permissions Related to Bucket Subresource Operations](../userguide/using-with-s3-actions.md#using-with-s3-actions-related-to-bucket-subresources) and [Managing Access Permissions to Your Amazon S3\
Resources](../userguide/s3-access-control.md) in the _Amazon S3 User Guide_.

For information about Amazon S3 analytics feature, see [Amazon S3 Analytics – Storage Class Analysis](../dev/analytics-storage-class.md)
in the _Amazon S3 User Guide_.

The following operations are related to `GetBucketAnalyticsConfiguration`:

- [DeleteBucketAnalyticsConfiguration](api-deletebucketanalyticsconfiguration.md)

- [ListBucketAnalyticsConfigurations](api-listbucketanalyticsconfigurations.md)

- [PutBucketAnalyticsConfiguration](api-putbucketanalyticsconfiguration.md)

###### Important

You must URL encode any signed header values that contain spaces. For example, if your header value is `my  file.txt`, containing two spaces after `my`, you must URL encode this value to `my%20%20file.txt`.

## Request Syntax

```nohighlight

GET /?analytics&id=Id HTTP/1.1
Host: Bucket.s3.amazonaws.com
x-amz-expected-bucket-owner: ExpectedBucketOwner

```

## URI Request Parameters

The request uses the following URI parameters.

**[Bucket](#API_GetBucketAnalyticsConfiguration_RequestSyntax)**

The name of the bucket from which an analytics configuration is retrieved.

Required: Yes

**[id](#API_GetBucketAnalyticsConfiguration_RequestSyntax)**

The ID that identifies the analytics configuration.

Required: Yes

**[x-amz-expected-bucket-owner](#API_GetBucketAnalyticsConfiguration_RequestSyntax)**

The account ID of the expected bucket owner. If the account ID that you provide does not match the actual owner of the bucket, the request fails with the HTTP status code `403 Forbidden` (access denied).

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
<?xml version="1.0" encoding="UTF-8"?>
<AnalyticsConfiguration>
   <Id>string</Id>
   <Filter>
      <And>
         <Prefix>string</Prefix>
         <Tag>
            <Key>string</Key>
            <Value>string</Value>
         </Tag>
         ...
      </And>
      <Prefix>string</Prefix>
      <Tag>
         <Key>string</Key>
         <Value>string</Value>
      </Tag>
   </Filter>
   <StorageClassAnalysis>
      <DataExport>
         <Destination>
            <S3BucketDestination>
               <Bucket>string</Bucket>
               <BucketAccountId>string</BucketAccountId>
               <Format>string</Format>
               <Prefix>string</Prefix>
            </S3BucketDestination>
         </Destination>
         <OutputSchemaVersion>string</OutputSchemaVersion>
      </DataExport>
   </StorageClassAnalysis>
</AnalyticsConfiguration>
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in XML format by the service.

**[AnalyticsConfiguration](#API_GetBucketAnalyticsConfiguration_ResponseSyntax)**

Root level tag for the AnalyticsConfiguration parameters.

Required: Yes

**[Filter](#API_GetBucketAnalyticsConfiguration_ResponseSyntax)**

The filter used to describe a set of objects for analyses. A filter must have exactly one prefix,
one tag, or one conjunction (AnalyticsAndOperator). If no filter is provided, all objects will be
considered in any analysis.

Type: [AnalyticsFilter](api-analyticsfilter.md) data type

**[Id](#API_GetBucketAnalyticsConfiguration_ResponseSyntax)**

The ID that identifies the analytics configuration.

Type: String

**[StorageClassAnalysis](#API_GetBucketAnalyticsConfiguration_ResponseSyntax)**

Contains data related to access patterns to be collected and made available to analyze the
tradeoffs between different storage classes.

Type: [StorageClassAnalysis](api-storageclassanalysis.md) data type

## Examples

### Configure an Analytics Report

The following GET request for the bucket `amzn-s3-demo-bucket` returns the inventory
configuration with the ID `list1`:

```HTTP

GET /?analytics&id=list1 HTTP/1.1
Host: amzn-s3-demo-bucket.s3.<Region>.amazonaws.com
Date: Mon, 31 Oct 2016 12:00:00 GMT
Authorization: authorization string

```

### Example

The following is a sample response to the preceding GET request.

```xml

HTTP/1.1 200 OK
x-amz-id-2: YgIPIfBiKa2bj0KMgUAdQkf3ShJTOOpXUueF6QKo
x-amz-request-id: 236A8905248E5A02
Date: Mon, 31 Oct 2016 12:00:00 GMT
Server: AmazonS3
Content-Length: length

<?xml version="1.0" encoding="UTF-8"?>
<AnalyticsConfiguration xmlns="http://s3.amazonaws.com/doc/2006-03-01/">
  <Id>list1</Id>
  <Filter>
    <And>
      <Prefix>images/</Prefix>
      <Tag>
        <Key>dog</Key>
        <Value>corgi</Value>
      </Tag>
    </And>
  </Filter>
  <StorageClassAnalysis>
    <DataExport>
      <OutputSchemaVersion>V_1</OutputSchemaVersion>
      <Destination>
        <S3BucketDestination>
          <Format>CSV</Format>
          <BucketAccountId>123456789012</BucketAccountId>
          <Bucket>arn:aws:s3:::destination-bucket</Bucket>
          <Prefix>destination-prefix</Prefix>
        </S3BucketDestination>
      </Destination>
    </DataExport>
  </StorageClassAnalysis>
</AnalyticsConfiguration>

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/s3-2006-03-01/getbucketanalyticsconfiguration.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3-2006-03-01/getbucketanalyticsconfiguration.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3-2006-03-01/getbucketanalyticsconfiguration.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3-2006-03-01/getbucketanalyticsconfiguration.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3-2006-03-01/getbucketanalyticsconfiguration.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3-2006-03-01/getbucketanalyticsconfiguration.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3-2006-03-01/getbucketanalyticsconfiguration.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3-2006-03-01/getbucketanalyticsconfiguration.md)

- [AWS SDK for Python](../../../goto/boto3/s3-2006-03-01/getbucketanalyticsconfiguration.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3-2006-03-01/getbucketanalyticsconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetBucketAcl

GetBucketCors

All content copied from https://docs.aws.amazon.com/.
