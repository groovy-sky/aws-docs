---
title: "PutBucketInventoryConfiguration"
---

# PutBucketInventoryConfiguration

This implementation of the `PUT` action adds an S3 Inventory configuration (identified by
the inventory ID) to the bucket. You can have up to 1,000 inventory configurations per bucket.

Amazon S3 inventory generates inventories of the objects in the bucket on a daily or weekly basis, and
the results are published to a flat file. The bucket that is inventoried is called the
_source_ bucket, and the bucket where the inventory flat file is stored is called
the _destination_ bucket. The _destination_ bucket must be in the
same AWS Region as the _source_ bucket.

When you configure an inventory for a _source_ bucket, you specify the
_destination_ bucket where you want the inventory to be stored, and whether to
generate the inventory daily or weekly. You can also configure what object metadata to include and
whether to inventory all object versions or only current versions. For more information, see [Amazon S3 Inventory](../dev/storage-inventory.md) in the
Amazon S3 User Guide.

###### Important

You must create a bucket policy on the _destination_ bucket to grant
permissions to Amazon S3 to write objects to the bucket in the defined location. For an example policy, see
[Granting\
Permissions for Amazon S3 Inventory and Storage Class Analysis](../dev/example-bucket-policies.md#example-bucket-policies-use-case-9).

###### Note

**Directory buckets** \- For directory buckets, you must make requests for this API operation to the Regional endpoint. These endpoints support path-style requests in the format `https://s3express-control.region-code.amazonaws.com/bucket-name
         `. Virtual-hosted-style requests aren't supported.
For more information about endpoints in Availability Zones, see [Regional and Zonal endpoints for directory buckets in Availability Zones](../userguide/endpoint-directory-buckets-az.md) in the
_Amazon S3 User Guide_. For more information about endpoints in Local Zones, see [Concepts for directory buckets in Local Zones](../userguide/s3-lzs-for-directory-buckets.md) in the
_Amazon S3 User Guide_.

Permissions

To use this operation, you must have permission to perform the
`s3:PutInventoryConfiguration` action. The bucket owner has this permission by
default and can grant this permission to others.

The `s3:PutInventoryConfiguration` permission allows a user to create an [S3 Inventory](../userguide/storage-inventory.md)
report that includes all object metadata fields available and to specify the destination bucket to
store the inventory. A user with read access to objects in the destination bucket can also access
all object metadata fields that are available in the inventory report.

- **General purpose bucket permissions** \- The
`s3:PutInventoryConfiguration` permission is required in a policy. For more information
about general purpose buckets permissions, see [Using Bucket Policies and User\
Policies](../dev/using-iam-policies.md) in the _Amazon S3 User Guide_.

- **Directory bucket permissions** \- To grant access to
this API operation, you must have the `s3express:PutInventoryConfiguration` permission in
an IAM identity-based policy instead of a bucket policy.
For more information about directory bucket policies and permissions, see [AWS Identity and Access Management (IAM) for S3 Express One Zone](../userguide/s3-express-security-iam.md) in the _Amazon S3 User Guide_.

To restrict access to an inventory report, see [Restricting access to an Amazon S3 Inventory report](../userguide/example-bucket-policies.md#example-bucket-policies-s3-inventory) in the
_Amazon S3 User Guide_. For more information about the metadata fields available
in S3 Inventory, see [Amazon S3 Inventory\
lists](../userguide/storage-inventory.md#storage-inventory-contents) in the _Amazon S3 User Guide_. For more information about
permissions, see [Permissions related to bucket subresource operations](../userguide/using-with-s3-actions.md#using-with-s3-actions-related-to-bucket-subresources) and [Identity and access management in\
Amazon S3](../userguide/s3-access-control.md) in the _Amazon S3 User Guide_.

HTTP Host header syntax

**Directory buckets** \- The HTTP Host header syntax is `s3express-control.region-code.amazonaws.com`.

`PutBucketInventoryConfiguration` has the following special errors:

HTTP 400 Bad Request Error

_Code:_ InvalidArgument

_Cause:_ Invalid Argument

HTTP 400 Bad Request Error

_Code:_ TooManyConfigurations

_Cause:_ You are attempting to create a new configuration but have already
reached the 1,000-configuration limit.

HTTP 403 Forbidden Error

_Cause:_ You are not the owner of the specified bucket, or you do not have
the `s3:PutInventoryConfiguration` bucket permission to set the configuration on the
bucket.

The following operations are related to `PutBucketInventoryConfiguration`:

- [GetBucketInventoryConfiguration](api-getbucketinventoryconfiguration.md)

- [DeleteBucketInventoryConfiguration](api-deletebucketinventoryconfiguration.md)

- [ListBucketInventoryConfigurations](api-listbucketinventoryconfigurations.md)

###### Important

You must URL encode any signed header values that contain spaces. For example, if your header value is `my  file.txt`, containing two spaces after `my`, you must URL encode this value to `my%20%20file.txt`.

## Request Syntax

```nohighlight

PUT /?inventory&id=Id HTTP/1.1
Host: Bucket.s3.amazonaws.com
x-amz-expected-bucket-owner: ExpectedBucketOwner
<?xml version="1.0" encoding="UTF-8"?>
<InventoryConfiguration xmlns="http://s3.amazonaws.com/doc/2006-03-01/">
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

## URI Request Parameters

The request uses the following URI parameters.

**[Bucket](#API_PutBucketInventoryConfiguration_RequestSyntax)**

The name of the bucket where the inventory configuration will be stored.

**Directory buckets** \- When you use this operation with a directory bucket, you must use path-style requests in the format `https://s3express-control.region-code.amazonaws.com/bucket-name
                  `. Virtual-hosted-style requests aren't supported. Directory bucket names must be unique in the chosen Zone (Availability Zone or Local Zone). Bucket names must also follow the format `
                     bucket-base-name--zone-id--x-s3` (for example, `
                     DOC-EXAMPLE-BUCKET--usw2-az1--x-s3`). For information about bucket naming restrictions, see [Directory bucket naming rules](../userguide/directory-bucket-naming-rules.md) in the _Amazon S3 User Guide_

Required: Yes

**[id](#API_PutBucketInventoryConfiguration_RequestSyntax)**

The ID used to identify the inventory configuration.

Required: Yes

**[x-amz-expected-bucket-owner](#API_PutBucketInventoryConfiguration_RequestSyntax)**

The account ID of the expected bucket owner. If the account ID that you provide does not match the actual owner of the bucket, the request fails with the HTTP status code `403 Forbidden` (access denied).

###### Note

For directory buckets, this header is not supported in this API operation. If you specify this header, the request fails with the HTTP status code
`501 Not Implemented`.

## Request Body

The request accepts the following data in XML format.

**[InventoryConfiguration](#API_PutBucketInventoryConfiguration_RequestSyntax)**

Root level tag for the InventoryConfiguration parameters.

Required: Yes

**[Destination](#API_PutBucketInventoryConfiguration_RequestSyntax)**

Contains information about where to publish the inventory results.

Type: [InventoryDestination](api-inventorydestination.md) data type

Required: Yes

**[Filter](#API_PutBucketInventoryConfiguration_RequestSyntax)**

Specifies an inventory filter. The inventory only includes objects that meet the filter's
criteria.

Type: [InventoryFilter](api-inventoryfilter.md) data type

Required: No

**[Id](#API_PutBucketInventoryConfiguration_RequestSyntax)**

The ID used to identify the inventory configuration.

Type: String

Required: Yes

**[IncludedObjectVersions](#API_PutBucketInventoryConfiguration_RequestSyntax)**

Object versions to include in the inventory list. If set to `All`, the list includes all
the object versions, which adds the version-related fields `VersionId`,
`IsLatest`, and `DeleteMarker` to the list. If set to `Current`, the
list does not contain these version-related fields.

Type: String

Valid Values: `All | Current`

Required: Yes

**[IsEnabled](#API_PutBucketInventoryConfiguration_RequestSyntax)**

Specifies whether the inventory is enabled or disabled. If set to `True`, an inventory
list is generated. If set to `False`, no inventory list is generated.

Type: Boolean

Required: Yes

**[OptionalFields](#API_PutBucketInventoryConfiguration_RequestSyntax)**

Contains the optional fields that are included in the inventory results.

###### Note

The following optional fields are supported for directory buckets `Size | LastModifiedDate | StorageClass | ETag | IsMultipartUploaded |
      EncryptionStatus | BucketKeyStatus | ChecksumAlgorithm | LifecycleExpirationDate.` Throws MalformedXML error if unsupported optional field is provided.

Type: Array of strings

Valid Values: `Size | LastModifiedDate | StorageClass | ETag | IsMultipartUploaded | ReplicationStatus | EncryptionStatus | ObjectLockRetainUntilDate | ObjectLockMode | ObjectLockLegalHoldStatus | IntelligentTieringAccessTier | BucketKeyStatus | ChecksumAlgorithm | ObjectAccessControlList | ObjectOwner | LifecycleExpirationDate`

Required: No

**[Schedule](#API_PutBucketInventoryConfiguration_RequestSyntax)**

Specifies the schedule for generating inventory results.

Type: [InventorySchedule](api-inventoryschedule.md) data type

Required: Yes

## Response Syntax

```

HTTP/1.1 200

```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## Examples

### Example: Create an inventory configuration

The following `PUT` request and response for the bucket `examplebucket`
creates a new or replaces an existing inventory configuration with the ID `report1`. The
configuration is defined in the request body.

```xml

PUT /?inventory&id=report1 HTTP/1.1
Host: examplebucket.s3.<Region>.amazonaws.com
Date: Mon, 31 Oct 2016 12:00:00 GMT
Authorization: authorization string
Content-Length: length

<?xml version="1.0" encoding="UTF-8"?>
<InventoryConfiguration xmlns="http://s3.amazonaws.com/doc/2006-03-01/">
   <Id>report1</Id>
   <IsEnabled>true</IsEnabled>
   <Filter>
      <Prefix>filterPrefix</Prefix>
   </Filter>
   <Destination>
      <S3BucketDestination>
         <Format>CSV</Format>
         <AccountId>123456789012</AccountId>
         <Bucket>arn:aws:s3:::destination-bucket</Bucket>
         <Prefix>prefix1</Prefix>
         <Encryption>
            <SSE-KMS>
               <KeyId>arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab</KeyId>
            </SSE-KMS>
         </Encryption>
      </S3BucketDestination>
   </Destination>
   <Schedule>
      <Frequency>Daily</Frequency>
   </Schedule>
   <IncludedObjectVersions>All</IncludedObjectVersions>
   <OptionalFields>
      <Field>Size</Field>
      <Field>LastModifiedDate</Field>
      <Field>ETag</Field>
      <Field>StorageClass</Field>
      <Field>IsMultipartUploaded</Field>
      <Field>ReplicationStatus</Field>
      <Field>EncryptionStatus</Field>
      <Field>ObjectLockRetainUntilDate</Field>
      <Field>ObjectLockMode</Field>
      <Field>ObjectLockLegalHoldStatus</Field>
   </OptionalFields>
</InventoryConfiguration>

```

```HTTP

HTTP/1.1 200 OK
x-amz-id-2: YgIPIfBiKa2bj0KMg95r/0zo3emzU4dzsD4rcKCHQUAdQkf3ShJTOOpXUueF6QKo
x-amz-request-id: 236A8905248E5A01
Date: Mon, 31 Oct 2016 12:00:00 GMT
Content-Length: 0
Server: AmazonS3

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/s3-2006-03-01/PutBucketInventoryConfiguration)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/s3-2006-03-01/PutBucketInventoryConfiguration)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/s3-2006-03-01/PutBucketInventoryConfiguration)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/s3-2006-03-01/PutBucketInventoryConfiguration)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/s3-2006-03-01/PutBucketInventoryConfiguration)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/s3-2006-03-01/PutBucketInventoryConfiguration)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/s3-2006-03-01/PutBucketInventoryConfiguration)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/s3-2006-03-01/PutBucketInventoryConfiguration)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/s3-2006-03-01/PutBucketInventoryConfiguration)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/s3-2006-03-01/PutBucketInventoryConfiguration)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PutBucketIntelligentTieringConfiguration

PutBucketLifecycle

All content copied from https://docs.aws.amazon.com/.
