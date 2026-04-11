# ListBucketInventoryConfigurations

###### Note

This operation is not supported for directory buckets.

Returns a list of S3 Inventory configurations for the bucket. You can have up to 1,000 inventory
configurations per bucket.

This action supports list pagination and does not return more than 100 configurations at a time.
Always check the `IsTruncated` element in the response. If there are no more configurations
to list, `IsTruncated` is set to false. If there are more configurations to list,
`IsTruncated` is set to true, and there is a value in `NextContinuationToken`.
You use the `NextContinuationToken` value to continue the pagination of the list by passing
the value in continuation-token in the request to `GET` the next page.

To use this operation, you must have permissions to perform the
`s3:GetInventoryConfiguration` action. The bucket owner has this permission by default. The
bucket owner can grant this permission to others. For more information about permissions, see [Permissions Related to Bucket Subresource Operations](../userguide/using-with-s3-actions.md#using-with-s3-actions-related-to-bucket-subresources) and [Managing Access Permissions to Your Amazon S3\
Resources](../userguide/s3-access-control.md).

For information about the Amazon S3 inventory feature, see [Amazon S3 Inventory](../dev/storage-inventory.md)

The following operations are related to `ListBucketInventoryConfigurations`:

- [GetBucketInventoryConfiguration](api-getbucketinventoryconfiguration.md)

- [DeleteBucketInventoryConfiguration](api-deletebucketinventoryconfiguration.md)

- [PutBucketInventoryConfiguration](api-putbucketinventoryconfiguration.md)

###### Important

You must URL encode any signed header values that contain spaces. For example, if your header value is `my  file.txt`, containing two spaces after `my`, you must URL encode this value to `my%20%20file.txt`.

## Request Syntax

```nohighlight

GET /?inventory&continuation-token=ContinuationToken HTTP/1.1
Host: Bucket.s3.amazonaws.com
x-amz-expected-bucket-owner: ExpectedBucketOwner

```

## URI Request Parameters

The request uses the following URI parameters.

**[Bucket](#API_ListBucketInventoryConfigurations_RequestSyntax)**

The name of the bucket containing the inventory configurations to retrieve.

Required: Yes

**[continuation-token](#API_ListBucketInventoryConfigurations_RequestSyntax)**

The marker used to continue an inventory configuration listing that has been truncated. Use the
`NextContinuationToken` from a previously truncated list response to continue the listing.
The continuation token is an opaque value that Amazon S3 understands.

**[x-amz-expected-bucket-owner](#API_ListBucketInventoryConfigurations_RequestSyntax)**

The account ID of the expected bucket owner. If the account ID that you provide does not match the actual owner of the bucket, the request fails with the HTTP status code `403 Forbidden` (access denied).

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
<?xml version="1.0" encoding="UTF-8"?>
<ListInventoryConfigurationsResult>
   <ContinuationToken>string</ContinuationToken>
   <InventoryConfiguration>
      <Destination>
         <S3BucketDestination>
            <AccountId>string</AccountId>
            <Bucket>string</Bucket>
            <Encryption>
               <SSE-KMS>
                  <KeyId>string</KeyId>
               </SSE-KMS>
               <SSE-S3>
               </SSE-S3>
            </Encryption>
            <Format>string</Format>
            <Prefix>string</Prefix>
         </S3BucketDestination>
      </Destination>
      <Filter>
         <Prefix>string</Prefix>
      </Filter>
      <Id>string</Id>
      <IncludedObjectVersions>string</IncludedObjectVersions>
      <IsEnabled>boolean</IsEnabled>
      <OptionalFields>
         <Field>string</Field>
      </OptionalFields>
      <Schedule>
         <Frequency>string</Frequency>
      </Schedule>
   </InventoryConfiguration>
   ...
   <IsTruncated>boolean</IsTruncated>
   <NextContinuationToken>string</NextContinuationToken>
</ListInventoryConfigurationsResult>
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in XML format by the service.

**[ListInventoryConfigurationsResult](#API_ListBucketInventoryConfigurations_ResponseSyntax)**

Root level tag for the ListInventoryConfigurationsResult parameters.

Required: Yes

**[ContinuationToken](#API_ListBucketInventoryConfigurations_ResponseSyntax)**

If sent in the request, the marker that is used as a starting point for this inventory configuration
list response.

Type: String

**[InventoryConfiguration](#API_ListBucketInventoryConfigurations_ResponseSyntax)**

The list of inventory configurations for a bucket.

Type: Array of [InventoryConfiguration](api-inventoryconfiguration.md) data types

**[IsTruncated](#API_ListBucketInventoryConfigurations_ResponseSyntax)**

Tells whether the returned list of inventory configurations is complete. A value of true indicates
that the list is not complete and the NextContinuationToken is provided for a subsequent request.

Type: Boolean

**[NextContinuationToken](#API_ListBucketInventoryConfigurations_ResponseSyntax)**

The marker used to continue this inventory configuration listing. Use the
`NextContinuationToken` from this response to continue the listing in a subsequent request.
The continuation token is an opaque value that Amazon S3 understands.

Type: String

## Examples

### Sample Request

The following request returns the inventory configurations in `example-bucket`.

```

GET /?inventory HTTP/1.1
Host: example-bucket.s3.<Region>.amazonaws.com
x-amz-date: 20160430T233541Z
Authorization: authorization string
Content-Type: text/plain

```

### Sample Response

Delete the metric configuration with a specified ID, which disables the CloudWatch metrics with
the `ExampleMetrics` value for the `FilterId` dimension.

```

HTTP/1.1 200 OK
x-amz-id-2: gyB+3jRPnrkN98ZajxHXr3u7EFM67bNgSAxexeEHndCX/7GRnfTXxReKUQF28IfP
x-amz-request-id: 3B3C7C725673C630
Date: Sat, 30 Apr 2016 23:29:37 GMT
Content-Type: application/xml
Content-Length: length
Connection: close
Server: AmazonS3

<?xml version="1.0" encoding="UTF-8"?>
<ListInventoryConfigurationsResult xmlns="http://s3.amazonaws.com/doc/2006-03-01/">
   <InventoryConfiguration>
      <Id>report1</Id>
      <IsEnabled>true</IsEnabled>
      <Destination>
         <S3BucketDestination>
            <Format>CSV</Format>
            <AccountId>123456789012</AccountId>
            <Bucket>arn:aws:s3:::destination-bucket</Bucket>
            <Prefix>prefix1</Prefix>
         </S3BucketDestination>
      </Destination>
      <Schedule>
         <Frequency>Daily</Frequency>
      </Schedule>
      <Filter>
         <Prefix>prefix/One</Prefix>
      </Filter>
      <IncludedObjectVersions>All</IncludedObjectVersions>
      <OptionalFields>
         <Field>Size</Field>
         <Field>LastModifiedDate</Field>
         <Field>ETag</Field>
         <Field>StorageClass</Field>
         <Field>IsMultipartUploaded</Field>
         <Field>ReplicationStatus</Field>
      </OptionalFields>
   </InventoryConfiguration>
      <InventoryConfiguration>
      <Id>report2</Id>
      <IsEnabled>true</IsEnabled>
      <Destination>
         <S3BucketDestination>
            <Format>CSV</Format>
            <AccountId>123456789012</AccountId>
            <Bucket>arn:aws:s3:::bucket2</Bucket>
            <Prefix>prefix2</Prefix>
         </S3BucketDestination>
      </Destination>
      <Schedule>
         <Frequency>Daily</Frequency>
      </Schedule>
      <Filter>
         <Prefix>prefix/Two</Prefix>
      </Filter>
      <IncludedObjectVersions>All</IncludedObjectVersions>
      <OptionalFields>
         <Field>Size</Field>
         <Field>LastModifiedDate</Field>
         <Field>ETag</Field>
         <Field>StorageClass</Field>
         <Field>IsMultipartUploaded</Field>
         <Field>ReplicationStatus</Field>
         <Field>ObjectLockRetainUntilDate</Field>
         <Field>ObjectLockMode</Field>
         <Field>ObjectLockLegalHoldStatus</Field>
      </OptionalFields>
   </InventoryConfiguration>
   <InventoryConfiguration>
      <Id>report3</Id>
      <IsEnabled>true</IsEnabled>
      <Destination>
         <S3BucketDestination>
            <Format>CSV</Format>
            <AccountId>123456789012</AccountId>
            <Bucket>arn:aws:s3:::bucket3</Bucket>
            <Prefix>prefix3</Prefix>
         </S3BucketDestination>
      </Destination>
      <Schedule>
         <Frequency>Daily</Frequency>
      </Schedule>
      <Filter>
         <Prefix>prefix/Three</Prefix>
      </Filter>
      <IncludedObjectVersions>All</IncludedObjectVersions>
      <OptionalFields>
         <Field>Size</Field>
         <Field>LastModifiedDate</Field>
         <Field>ETag</Field>
         <Field>StorageClass</Field>
         <Field>IsMultipartUploaded</Field>
         <Field>ReplicationStatus</Field>
      </OptionalFields>
   </InventoryConfiguration>
    ...
    <IsTruncated>false</IsTruncated>
    <!-- If ContinuationToken was provided in the request. -->
    <ContinuationToken>...</ContinuationToken>
    <!-- if IsTruncated == true -->
    <IsTruncated>true</IsTruncated>
   <NextContinuationToken>...</NextContinuationToken>
</ListInventoryConfigurationsResult>

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/s3-2006-03-01/listbucketinventoryconfigurations.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3-2006-03-01/listbucketinventoryconfigurations.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3-2006-03-01/listbucketinventoryconfigurations.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3-2006-03-01/listbucketinventoryconfigurations.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3-2006-03-01/listbucketinventoryconfigurations.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3-2006-03-01/listbucketinventoryconfigurations.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3-2006-03-01/listbucketinventoryconfigurations.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3-2006-03-01/listbucketinventoryconfigurations.md)

- [AWS SDK for Python](../../../goto/boto3/s3-2006-03-01/listbucketinventoryconfigurations.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3-2006-03-01/listbucketinventoryconfigurations.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ListBucketIntelligentTieringConfigurations

ListBucketMetricsConfigurations

All content copied from https://docs.aws.amazon.com/.
