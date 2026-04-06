# GetBucketInventoryConfiguration

###### Note

This operation is not supported for directory buckets.

Returns an S3 Inventory configuration (identified by the inventory configuration ID) from the
bucket.

To use this operation, you must have permissions to perform the
`s3:GetInventoryConfiguration` action. The bucket owner has this permission by default and
can grant this permission to others. For more information about permissions, see [Permissions Related to Bucket Subresource Operations](https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources) and [Managing Access Permissions to Your Amazon S3\
Resources](https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-access-control.html).

For information about the Amazon S3 inventory feature, see [Amazon S3 Inventory](https://docs.aws.amazon.com/AmazonS3/latest/dev/storage-inventory.html).

The following operations are related to `GetBucketInventoryConfiguration`:

- [DeleteBucketInventoryConfiguration](api-deletebucketinventoryconfiguration.md)

- [ListBucketInventoryConfigurations](api-listbucketinventoryconfigurations.md)

- [PutBucketInventoryConfiguration](api-putbucketinventoryconfiguration.md)

###### Important

You must URL encode any signed header values that contain spaces. For example, if your header value is `my  file.txt`, containing two spaces after `my`, you must URL encode this value to `my%20%20file.txt`.

## Request Syntax

```nohighlight

GET /?inventory&id=Id HTTP/1.1
Host: Bucket.s3.amazonaws.com
x-amz-expected-bucket-owner: ExpectedBucketOwner

```

## URI Request Parameters

The request uses the following URI parameters.

**[Bucket](#API_GetBucketInventoryConfiguration_RequestSyntax)**

The name of the bucket containing the inventory configuration to retrieve.

Required: Yes

**[id](#API_GetBucketInventoryConfiguration_RequestSyntax)**

The ID used to identify the inventory configuration.

Required: Yes

**[x-amz-expected-bucket-owner](#API_GetBucketInventoryConfiguration_RequestSyntax)**

The account ID of the expected bucket owner. If the account ID that you provide does not match the actual owner of the bucket, the request fails with the HTTP status code `403 Forbidden` (access denied).

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
<?xml version="1.0" encoding="UTF-8"?>
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
   <IsEnabled>boolean</IsEnabled>
   <Filter>
      <Prefix>string</Prefix>
   </Filter>
   <Id>string</Id>
   <IncludedObjectVersions>string</IncludedObjectVersions>
   <OptionalFields>
      <Field>string</Field>
   </OptionalFields>
   <Schedule>
      <Frequency>string</Frequency>
   </Schedule>
</InventoryConfiguration>
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in XML format by the service.

**[InventoryConfiguration](#API_GetBucketInventoryConfiguration_ResponseSyntax)**

Root level tag for the InventoryConfiguration parameters.

Required: Yes

**[Destination](#API_GetBucketInventoryConfiguration_ResponseSyntax)**

Contains information about where to publish the inventory results.

Type: [InventoryDestination](https://docs.aws.amazon.com/AmazonS3/latest/API/API_InventoryDestination.html) data type

**[Filter](#API_GetBucketInventoryConfiguration_ResponseSyntax)**

Specifies an inventory filter. The inventory only includes objects that meet the filter's
criteria.

Type: [InventoryFilter](https://docs.aws.amazon.com/AmazonS3/latest/API/API_InventoryFilter.html) data type

**[Id](#API_GetBucketInventoryConfiguration_ResponseSyntax)**

The ID used to identify the inventory configuration.

Type: String

**[IncludedObjectVersions](#API_GetBucketInventoryConfiguration_ResponseSyntax)**

Object versions to include in the inventory list. If set to `All`, the list includes all
the object versions, which adds the version-related fields `VersionId`,
`IsLatest`, and `DeleteMarker` to the list. If set to `Current`, the
list does not contain these version-related fields.

Type: String

Valid Values: `All | Current`

**[IsEnabled](#API_GetBucketInventoryConfiguration_ResponseSyntax)**

Specifies whether the inventory is enabled or disabled. If set to `True`, an inventory
list is generated. If set to `False`, no inventory list is generated.

Type: Boolean

**[OptionalFields](#API_GetBucketInventoryConfiguration_ResponseSyntax)**

Contains the optional fields that are included in the inventory results.

Type: Array of strings

Valid Values: `Size | LastModifiedDate | StorageClass | ETag | IsMultipartUploaded | ReplicationStatus | EncryptionStatus | ObjectLockRetainUntilDate | ObjectLockMode | ObjectLockLegalHoldStatus | IntelligentTieringAccessTier | BucketKeyStatus | ChecksumAlgorithm | ObjectAccessControlList | ObjectOwner | LifecycleExpirationDate`

**[Schedule](#API_GetBucketInventoryConfiguration_ResponseSyntax)**

Specifies the schedule for generating inventory results.

Type: [InventorySchedule](https://docs.aws.amazon.com/AmazonS3/latest/API/API_InventorySchedule.html) data type

## Examples

### Sample Request: Configure an inventory report

The following GET request for the bucket `examplebucket` returns the inventory
configuration with the ID `list1`.

```

            GET /?inventory&id=list1 HTTP/1.1
            Host: examplebucket.s3.<Region>.amazonaws.com
            Date: Mon, 31 Oct 2016 12:00:00 GMT
            Authorization: authorization string

```

### Sample Response

This example illustrates one usage of GetBucketInventoryConfiguration.

```

         HTTP/1.1 200 OK
         x-amz-id-2: YgIPIfBiKa2bj0KMgUAdQkf3ShJTOOpXUueF6QKo
         x-amz-request-id: 236A8905248E5A02
         Date: Mon, 31 Oct 2016 12:00:00 GMT
         Server: AmazonS3
         Content-Length: length

         <?xml version="1.0" encoding="UTF-8"?>
         <InventoryConfiguration xmlns="http://s3.amazonaws.com/doc/2006-03-01/">
            <Id>report1</Id>
           <IsEnabled>true</IsEnabled>
           <Destination>
              <S3BucketDestination>
                 <Format>CSV</Format>
                  <AccountId>123456789012</AccountId>
                  <Bucket>arn:aws:s3:::destination-bucket</Bucket>
                 <Prefix>prefix1</Prefix>
                 <SSE-S3/>
               </S3BucketDestination>
            </Destination>
            <Schedule>
               <Frequency>Daily</Frequency>
           </Schedule>
           <Filter>
             <Prefix>myprefix/</Prefix>
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

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/s3-2006-03-01/GetBucketInventoryConfiguration)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/s3-2006-03-01/GetBucketInventoryConfiguration)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/s3-2006-03-01/GetBucketInventoryConfiguration)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/s3-2006-03-01/GetBucketInventoryConfiguration)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/s3-2006-03-01/GetBucketInventoryConfiguration)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/s3-2006-03-01/GetBucketInventoryConfiguration)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/s3-2006-03-01/GetBucketInventoryConfiguration)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/s3-2006-03-01/GetBucketInventoryConfiguration)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/s3-2006-03-01/GetBucketInventoryConfiguration)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/s3-2006-03-01/GetBucketInventoryConfiguration)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

GetBucketIntelligentTieringConfiguration

GetBucketLifecycle
