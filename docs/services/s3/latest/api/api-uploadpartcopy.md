# UploadPartCopy

Uploads a part by copying data from an existing object as data source. To specify the data source,
you add the request header `x-amz-copy-source` in your request. To specify a byte range, you
add the request header `x-amz-copy-source-range` in your request.

For information about maximum and minimum part sizes and other multipart upload specifications, see
[Multipart upload\
limits](../userguide/qfacts.md) in the _Amazon S3 User Guide_.

###### Note

Instead of copying data from an existing object as part data, you might use the [UploadPart](api-uploadpart.md) action to
upload new data as a part of an object in your request.

You must initiate a multipart upload before you can upload any part. In response to your initiate
request, Amazon S3 returns the upload ID, a unique identifier that you must include in your upload part
request.

For conceptual information about multipart uploads, see [Uploading Objects Using Multipart Upload](https://docs.aws.amazon.com/AmazonS3/latest/dev/uploadobjusingmpu.html) in
the _Amazon S3 User Guide_. For information about copying objects using a single atomic
action vs. a multipart upload, see [Operations on Objects](https://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectOperations.html) in the
_Amazon S3 User Guide_.

###### Note

**Directory buckets** \- For directory buckets, you must make requests for this API operation to the Zonal endpoint. These endpoints support virtual-hosted-style requests in the format `https://amzn-s3-demo-bucket.s3express-zone-id.region-code.amazonaws.com/key-name
         `. Path-style requests are not supported. For more information about endpoints in Availability Zones, see [Regional and Zonal endpoints for directory buckets in Availability Zones](../userguide/endpoint-directory-buckets-az.md) in the
_Amazon S3 User Guide_. For more information about endpoints in Local Zones, see [Concepts for directory buckets in Local Zones](../userguide/s3-lzs-for-directory-buckets.md) in the
_Amazon S3 User Guide_.

Authentication and authorization

All `UploadPartCopy` requests must be authenticated and signed by using IAM
credentials (access key ID and secret access key for the IAM identities). All headers with the
`x-amz-` prefix, including `x-amz-copy-source`, must be signed. For more
information, see [REST Authentication](https://docs.aws.amazon.com/AmazonS3/latest/dev/RESTAuthentication.html).

**Directory buckets** \- You must use IAM credentials to
authenticate and authorize your access to the `UploadPartCopy` API operation, instead
of using the temporary security credentials through the `CreateSession` API
operation.

AWS CLI or SDKs handles authentication and authorization on your behalf.

Permissions

You must have `READ` access to the source object and `WRITE` access to
the destination bucket.

- **General purpose bucket permissions** \- You must have the
permissions in a policy based on the bucket types of your source bucket and destination bucket
in an `UploadPartCopy` operation.

- If the source object is in a general purpose bucket, you must have the **`s3:GetObject`** permission to read the source object that is
being copied.

- If the destination bucket is a general purpose bucket, you must have the **`s3:PutObject`** permission to write the object copy to
the destination bucket.

- To perform a multipart upload with encryption using an AWS Key Management Service key, the requester
must have permission to the `kms:Decrypt` and `kms:GenerateDataKey`
actions on the key. The requester must also have permissions for the
`kms:GenerateDataKey` action for the `CreateMultipartUpload` API.
Then, the requester needs permissions for the `kms:Decrypt` action on the
`UploadPart` and `UploadPartCopy` APIs. These permissions are
required because Amazon S3 must decrypt and read data from the encrypted file parts before it
completes the multipart upload. For more information about KMS permissions, see [Protecting\
data using server-side encryption with AWS KMS](../userguide/usingkmsencryption.md) in the
_Amazon S3 User Guide_. For information about the permissions required to
use the multipart upload API, see [Multipart upload and\
permissions](https://docs.aws.amazon.com/AmazonS3/latest/dev/mpuAndPermissions.html) and [Multipart upload API\
and permissions](../userguide/mpuoverview.md#mpuAndPermissions) in the _Amazon S3 User Guide_.

- **Directory bucket permissions** \- You must have
permissions in a bucket policy or an IAM identity-based policy based on the source and destination bucket types
in an `UploadPartCopy` operation.

- If the source object that you want to copy is in a directory bucket, you must have
the **`s3express:CreateSession`** permission in
the `Action` element of a policy to read the object. If no session mode is specified,
the session will be created with the maximum allowable privilege, attempting
`ReadWrite` first, then `ReadOnly` if `ReadWrite` is not permitted.
If you want to explicitly restrict the access to be read-only, you can set the `s3express:SessionMode`
condition key to `ReadOnly` on the copy source bucket.

- If the copy destination is a directory bucket, you must have the **`s3express:CreateSession`** permission in the
`Action` element of a policy to write the object to the destination. The
`s3express:SessionMode` condition key cannot be set to `ReadOnly`
on the copy destination.

If the object is encrypted with SSE-KMS, you must also have the
`kms:GenerateDataKey` and `kms:Decrypt` permissions in IAM
identity-based policies and AWS KMS key policies for the AWS KMS key.

For example policies, see [Example\
bucket policies for S3 Express One Zone](../userguide/s3-express-security-iam-example-bucket-policies.md) and [AWS\
Identity and Access Management (IAM) identity-based policies for S3 Express One Zone](../userguide/s3-express-security-iam-identity-policies.md) in the
_Amazon S3 User Guide_.

Encryption

- **General purpose buckets** -
For information about using server-side encryption with
customer-provided encryption keys with the `UploadPartCopy` operation, see [CopyObject](api-copyobject.md) and
[UploadPart](api-uploadpart.md).

###### Note

If you have server-side encryption with customer-provided keys (SSE-C) blocked for your general purpose bucket, you will get an HTTP 403 Access Denied error when you specify the SSE-C request headers while writing new data to your bucket. For more information, see [Blocking or unblocking SSE-C for a general purpose bucket](../userguide/blocking-unblocking-s3-c-encryption-gpb.md).

- **Directory buckets** -
For directory buckets, there are only two supported options for server-side encryption: server-side encryption with Amazon S3 managed keys (SSE-S3) ( `AES256`) and server-side encryption with AWS KMS keys (SSE-KMS) ( `aws:kms`). For more
information, see [Protecting data with server-side encryption](https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-serv-side-encryption.html) in the _Amazon S3 User Guide_.

###### Note

For directory buckets, when you perform a `CreateMultipartUpload` operation
and an `UploadPartCopy` operation, the request headers you provide in the
`CreateMultipartUpload` request must match the default encryption configuration
of the destination bucket.

S3 Bucket Keys aren't supported, when you copy SSE-KMS encrypted objects from general purpose buckets
to directory buckets, from directory buckets to general purpose buckets, or between directory buckets, through [UploadPartCopy](api-uploadpartcopy.md). In this case, Amazon S3 makes a call to AWS KMS every time a copy request is made for a KMS-encrypted object.

Special errors

- Error Code: `NoSuchUpload`

- Description: The specified multipart upload does not exist. The upload ID might be
invalid, or the multipart upload might have been aborted or completed.

- HTTP Status Code: 404 Not Found

- Error Code: `InvalidRequest`

- Description: The specified copy source is not supported as a byte-range copy
source.

- HTTP Status Code: 400 Bad Request

HTTP Host header syntax

**Directory buckets** \- The HTTP Host header syntax is `
                  Bucket-name.s3express-zone-id.region-code.amazonaws.com`.

The following operations are related to `UploadPartCopy`:

- [CreateMultipartUpload](api-createmultipartupload.md)

- [UploadPart](api-uploadpart.md)

- [CompleteMultipartUpload](api-completemultipartupload.md)

- [AbortMultipartUpload](api-abortmultipartupload.md)

- [ListParts](api-listparts.md)

- [ListMultipartUploads](api-listmultipartuploads.md)

###### Important

You must URL encode any signed header values that contain spaces. For example, if your header value is `my  file.txt`, containing two spaces after `my`, you must URL encode this value to `my%20%20file.txt`.

## Request Syntax

```nohighlight

PUT /Key+?partNumber=PartNumber&uploadId=UploadId HTTP/1.1
Host: Bucket.s3.amazonaws.com
x-amz-copy-source: CopySource
x-amz-copy-source-if-match: CopySourceIfMatch
x-amz-copy-source-if-modified-since: CopySourceIfModifiedSince
x-amz-copy-source-if-none-match: CopySourceIfNoneMatch
x-amz-copy-source-if-unmodified-since: CopySourceIfUnmodifiedSince
x-amz-copy-source-range: CopySourceRange
x-amz-server-side-encryption-customer-algorithm: SSECustomerAlgorithm
x-amz-server-side-encryption-customer-key: SSECustomerKey
x-amz-server-side-encryption-customer-key-MD5: SSECustomerKeyMD5
x-amz-copy-source-server-side-encryption-customer-algorithm: CopySourceSSECustomerAlgorithm
x-amz-copy-source-server-side-encryption-customer-key: CopySourceSSECustomerKey
x-amz-copy-source-server-side-encryption-customer-key-MD5: CopySourceSSECustomerKeyMD5
x-amz-request-payer: RequestPayer
x-amz-expected-bucket-owner: ExpectedBucketOwner
x-amz-source-expected-bucket-owner: ExpectedSourceBucketOwner

```

## URI Request Parameters

The request uses the following URI parameters.

**[Bucket](#API_UploadPartCopy_RequestSyntax)**

The bucket name.

**Directory buckets** \- When you use this operation with a directory bucket, you must use virtual-hosted-style requests in the format `
                     Bucket-name.s3express-zone-id.region-code.amazonaws.com`. Path-style requests are not supported. Directory bucket names must be unique in the chosen Zone (Availability Zone or Local Zone). Bucket names must follow the format `
                     bucket-base-name--zone-id--x-s3` (for example, `
                     amzn-s3-demo-bucket--usw2-az1--x-s3`). For information about bucket naming
restrictions, see [Directory bucket naming\
rules](../userguide/directory-bucket-naming-rules.md) in the _Amazon S3 User Guide_.

###### Note

Copying objects across different AWS Regions isn't supported when the source or destination
bucket is in AWS Local Zones. The source and destination buckets must have the same parent AWS Region.
Otherwise, you get an HTTP `400 Bad Request` error with the error code
`InvalidRequest`.

**Access points** \- When you use this action with an access point for general purpose buckets, you must provide the alias of the access point in place of the bucket name or specify the access point ARN. When you use this action with an access point for directory buckets, you must provide the access point name in place of the bucket name. When using the access point ARN, you must direct requests to the access point hostname. The access point hostname takes the form _AccessPointName_- _AccountId_.s3-accesspoint. _Region_.amazonaws.com. When using this action with an access point through the AWS SDKs, you provide the access point ARN in place of the bucket name. For more information about access point ARNs, see [Using access points](../userguide/using-access-points.md) in the _Amazon S3 User Guide_.

###### Note

Object Lambda access points are not supported by directory buckets.

**S3 on Outposts** \- When you use this action with S3 on Outposts, you must direct requests to the S3 on Outposts hostname. The S3 on Outposts hostname takes the
form `
                     AccessPointName-AccountId.outpostID.s3-outposts.Region.amazonaws.com`. When you use this action with S3 on Outposts, the destination bucket must be the Outposts access point ARN or the access point alias. For more information about S3 on Outposts, see [What is S3 on Outposts?](https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3onOutposts.html) in the _Amazon S3 User Guide_.

Required: Yes

**[Key](#API_UploadPartCopy_RequestSyntax)**

Object key for which the multipart upload was initiated.

Length Constraints: Minimum length of 1.

Required: Yes

**[partNumber](#API_UploadPartCopy_RequestSyntax)**

Part number of part being copied. This is a positive integer between 1 and 10,000.

Required: Yes

**[uploadId](#API_UploadPartCopy_RequestSyntax)**

Upload ID identifying the multipart upload whose part is being copied.

Required: Yes

**[x-amz-copy-source](#API_UploadPartCopy_RequestSyntax)**

Specifies the source object for the copy operation. You specify the value in one of two formats,
depending on whether you want to access the source object through an [access point](../userguide/access-points.md):

- For objects not accessed through an access point, specify the name of the source bucket and key of the
source object, separated by a slash (/). For example, to copy the object
`reports/january.pdf` from the bucket `awsexamplebucket`, use
`awsexamplebucket/reports/january.pdf`. The value must be URL-encoded.

- For objects accessed through access points, specify the Amazon Resource Name (ARN) of the object as accessed through the access point, in the format `arn:aws:s3:<Region>:<account-id>:accesspoint/<access-point-name>/object/<key>`. For example, to copy the object `reports/january.pdf` through access point `my-access-point` owned by account `123456789012` in Region `us-west-2`, use the URL encoding of `arn:aws:s3:us-west-2:123456789012:accesspoint/my-access-point/object/reports/january.pdf`. The value must be URL encoded.

###### Note

- Amazon S3 supports copy operations using Access points only when the source and destination buckets are in the same AWS Region.

- Access points are not supported by directory buckets.

Alternatively, for objects accessed through Amazon S3 on Outposts, specify the ARN of the object as accessed in the format `arn:aws:s3-outposts:<Region>:<account-id>:outpost/<outpost-id>/object/<key>`. For example, to copy the object `reports/january.pdf` through outpost `my-outpost` owned by account `123456789012` in Region `us-west-2`, use the URL encoding of `arn:aws:s3-outposts:us-west-2:123456789012:outpost/my-outpost/object/reports/january.pdf`. The value must be URL-encoded.

If your bucket has versioning enabled, you could have multiple versions of the same object. By
default, `x-amz-copy-source` identifies the current version of the source object to copy. To
copy a specific version of the source object to copy, append `?versionId=<version-id>` to
the `x-amz-copy-source` request header (for example, `x-amz-copy-source:
        /awsexamplebucket/reports/january.pdf?versionId=QUpfdndhfd8438MNFDN93jdnJFkdmqnh893`).

If the current version is a delete marker and you don't specify a versionId in the
`x-amz-copy-source` request header, Amazon S3 returns a `404 Not Found` error,
because the object does not exist. If you specify versionId in the `x-amz-copy-source` and
the versionId is a delete marker, Amazon S3 returns an HTTP `400 Bad Request` error, because you
are not allowed to specify a delete marker as a version for the `x-amz-copy-source`.

###### Note

**Directory buckets** \- S3 Versioning isn't enabled and supported for directory buckets.

Pattern: `\/?.+\/.+`

Required: Yes

**[x-amz-copy-source-if-match](#API_UploadPartCopy_RequestSyntax)**

Copies the object if its entity tag (ETag) matches the specified tag.

If both of the `x-amz-copy-source-if-match` and
`x-amz-copy-source-if-unmodified-since` headers are present in the request as
follows:

`x-amz-copy-source-if-match` condition evaluates to `true`, and;

`x-amz-copy-source-if-unmodified-since` condition evaluates to `false`;

Amazon S3 returns `200 OK` and copies the data.

**[x-amz-copy-source-if-modified-since](#API_UploadPartCopy_RequestSyntax)**

Copies the object if it has been modified since the specified time.

If both of the `x-amz-copy-source-if-none-match` and
`x-amz-copy-source-if-modified-since` headers are present in the request as follows:

`x-amz-copy-source-if-none-match` condition evaluates to `false`, and;

`x-amz-copy-source-if-modified-since` condition evaluates to `true`;

Amazon S3 returns `412 Precondition Failed` response code.

**[x-amz-copy-source-if-none-match](#API_UploadPartCopy_RequestSyntax)**

Copies the object if its entity tag (ETag) is different than the specified ETag.

If both of the `x-amz-copy-source-if-none-match` and
`x-amz-copy-source-if-modified-since` headers are present in the request as follows:

`x-amz-copy-source-if-none-match` condition evaluates to `false`, and;

`x-amz-copy-source-if-modified-since` condition evaluates to `true`;

Amazon S3 returns `412 Precondition Failed` response code.

**[x-amz-copy-source-if-unmodified-since](#API_UploadPartCopy_RequestSyntax)**

Copies the object if it hasn't been modified since the specified time.

If both of the `x-amz-copy-source-if-match` and
`x-amz-copy-source-if-unmodified-since` headers are present in the request as
follows:

`x-amz-copy-source-if-match` condition evaluates to `true`, and;

`x-amz-copy-source-if-unmodified-since` condition evaluates to `false`;

Amazon S3 returns `200 OK` and copies the data.

**[x-amz-copy-source-range](#API_UploadPartCopy_RequestSyntax)**

The range of bytes to copy from the source object. The range value must use the form
bytes=first-last, where the first and last are the zero-based byte offsets to copy. For example,
bytes=0-9 indicates that you want to copy the first 10 bytes of the source. You can copy a range only if
the source object is greater than 5 MB.

**[x-amz-copy-source-server-side-encryption-customer-algorithm](#API_UploadPartCopy_RequestSyntax)**

Specifies the algorithm to use when decrypting the source object (for example,
`AES256`).

###### Note

This functionality is not supported when the source object is in a directory bucket.

**[x-amz-copy-source-server-side-encryption-customer-key](#API_UploadPartCopy_RequestSyntax)**

Specifies the customer-provided encryption key for Amazon S3 to use to decrypt the source object. The
encryption key provided in this header must be one that was used when the source object was
created.

###### Note

This functionality is not supported when the source object is in a directory bucket.

**[x-amz-copy-source-server-side-encryption-customer-key-MD5](#API_UploadPartCopy_RequestSyntax)**

Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321. Amazon S3 uses this header
for a message integrity check to ensure that the encryption key was transmitted without error.

###### Note

This functionality is not supported when the source object is in a directory bucket.

**[x-amz-expected-bucket-owner](#API_UploadPartCopy_RequestSyntax)**

The account ID of the expected destination bucket owner. If the account ID that you provide does not match the actual owner of the destination bucket, the request fails with the HTTP status code `403 Forbidden` (access denied).

**[x-amz-request-payer](#API_UploadPartCopy_RequestSyntax)**

Confirms that the requester knows that they will be charged for the request. Bucket owners need not
specify this parameter in their requests. If either the source or destination S3 bucket has Requester
Pays enabled, the requester will pay for the corresponding charges. For information about
downloading objects from Requester Pays buckets, see [Downloading Objects in Requester Pays\
Buckets](https://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html) in the _Amazon S3 User Guide_.

###### Note

This functionality is not supported for directory buckets.

Valid Values: `requester`

**[x-amz-server-side-encryption-customer-algorithm](#API_UploadPartCopy_RequestSyntax)**

Specifies the algorithm to use when encrypting the object (for example, AES256).

###### Note

This functionality is not supported when the destination bucket is a directory bucket.

**[x-amz-server-side-encryption-customer-key](#API_UploadPartCopy_RequestSyntax)**

Specifies the customer-provided encryption key for Amazon S3 to use in encrypting data. This value is
used to store the object and then it is discarded; Amazon S3 does not store the encryption key. The key must
be appropriate for use with the algorithm specified in the
`x-amz-server-side-encryption-customer-algorithm` header. This must be the same encryption
key specified in the initiate multipart upload request.

###### Note

This functionality is not supported when the destination bucket is a directory bucket.

**[x-amz-server-side-encryption-customer-key-MD5](#API_UploadPartCopy_RequestSyntax)**

Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321. Amazon S3 uses this header
for a message integrity check to ensure that the encryption key was transmitted without error.

###### Note

This functionality is not supported when the destination bucket is a directory bucket.

**[x-amz-source-expected-bucket-owner](#API_UploadPartCopy_RequestSyntax)**

The account ID of the expected source bucket owner. If the account ID that you provide does not match the actual owner of the source bucket, the request fails with the HTTP status code `403 Forbidden` (access denied).

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
x-amz-copy-source-version-id: CopySourceVersionId
x-amz-server-side-encryption: ServerSideEncryption
x-amz-server-side-encryption-customer-algorithm: SSECustomerAlgorithm
x-amz-server-side-encryption-customer-key-MD5: SSECustomerKeyMD5
x-amz-server-side-encryption-aws-kms-key-id: SSEKMSKeyId
x-amz-server-side-encryption-bucket-key-enabled: BucketKeyEnabled
x-amz-request-charged: RequestCharged
<?xml version="1.0" encoding="UTF-8"?>
<CopyPartResult>
   <ETag>string</ETag>
   <LastModified>timestamp</LastModified>
   <ChecksumCRC32>string</ChecksumCRC32>
   <ChecksumCRC32C>string</ChecksumCRC32C>
   <ChecksumCRC64NVME>string</ChecksumCRC64NVME>
   <ChecksumSHA1>string</ChecksumSHA1>
   <ChecksumSHA256>string</ChecksumSHA256>
</CopyPartResult>
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The response returns the following HTTP headers.

**[x-amz-copy-source-version-id](#API_UploadPartCopy_ResponseSyntax)**

The version of the source object that was copied, if you have enabled versioning on the source
bucket.

###### Note

This functionality is not supported when the source object is in a directory bucket.

**[x-amz-request-charged](#API_UploadPartCopy_ResponseSyntax)**

If present, indicates that the requester was successfully charged for the request. For more
information, see [Using Requester Pays buckets for storage transfers and usage](../userguide/requesterpaysbuckets.md) in the _Amazon Simple_
_Storage Service user guide_.

###### Note

This functionality is not supported for directory buckets.

Valid Values: `requester`

**[x-amz-server-side-encryption](#API_UploadPartCopy_ResponseSyntax)**

The server-side encryption algorithm used when you store this object in Amazon S3 or Amazon FSx.

###### Note

When accessing data stored in Amazon FSx file systems using S3 access points, the only valid server side
encryption option is `aws:fsx`.

Valid Values: `AES256 | aws:fsx | aws:kms | aws:kms:dsse`

**[x-amz-server-side-encryption-aws-kms-key-id](#API_UploadPartCopy_ResponseSyntax)**

If present, indicates the ID of the KMS key that was used for object encryption.

**[x-amz-server-side-encryption-bucket-key-enabled](#API_UploadPartCopy_ResponseSyntax)**

Indicates whether the multipart upload uses an S3 Bucket Key for server-side encryption with
AWS Key Management Service (AWS KMS) keys (SSE-KMS).

**[x-amz-server-side-encryption-customer-algorithm](#API_UploadPartCopy_ResponseSyntax)**

If server-side encryption with a customer-provided encryption key was requested, the response will
include this header to confirm the encryption algorithm that's used.

###### Note

This functionality is not supported for directory buckets.

**[x-amz-server-side-encryption-customer-key-MD5](#API_UploadPartCopy_ResponseSyntax)**

If server-side encryption with a customer-provided encryption key was requested, the response will
include this header to provide the round-trip message integrity verification of the customer-provided
encryption key.

###### Note

This functionality is not supported for directory buckets.

The following data is returned in XML format by the service.

**[CopyPartResult](#API_UploadPartCopy_ResponseSyntax)**

Root level tag for the CopyPartResult parameters.

Required: Yes

**[ChecksumCRC32](#API_UploadPartCopy_ResponseSyntax)**

This header can be used as a data integrity check to verify that the data received is the same data
that was originally sent. This header specifies the Base64 encoded, 32-bit `CRC32` checksum
of the part. For more information, see [Checking object integrity](../userguide/checking-object-integrity.md) in
the _Amazon S3 User Guide_.

Type: String

**[ChecksumCRC32C](#API_UploadPartCopy_ResponseSyntax)**

This header can be used as a data integrity check to verify that the data received is the same data
that was originally sent. This header specifies the Base64 encoded, 32-bit `CRC32C` checksum
of the part. For more information, see [Checking object integrity](../userguide/checking-object-integrity.md) in
the _Amazon S3 User Guide_.

Type: String

**[ChecksumCRC64NVME](#API_UploadPartCopy_ResponseSyntax)**

The Base64 encoded, 64-bit `CRC64NVME` checksum of the part. This checksum is present if
the multipart upload request was created with the `CRC64NVME` checksum algorithm to the
uploaded object). For more information, see [Checking object integrity](../userguide/checking-object-integrity.md) in
the _Amazon S3 User Guide_.

Type: String

**[ChecksumSHA1](#API_UploadPartCopy_ResponseSyntax)**

This header can be used as a data integrity check to verify that the data received is the same data
that was originally sent. This header specifies the Base64 encoded, 160-bit `SHA1` checksum
of the part. For more information, see [Checking object integrity](../userguide/checking-object-integrity.md) in
the _Amazon S3 User Guide_.

Type: String

**[ChecksumSHA256](#API_UploadPartCopy_ResponseSyntax)**

This header can be used as a data integrity check to verify that the data received is the same data
that was originally sent. This header specifies the Base64 encoded, 256-bit `SHA256` checksum
of the part. For more information, see [Checking object integrity](../userguide/checking-object-integrity.md) in
the _Amazon S3 User Guide_.

Type: String

**[ETag](#API_UploadPartCopy_ResponseSyntax)**

Entity tag of the object.

Type: String

**[LastModified](#API_UploadPartCopy_ResponseSyntax)**

Date and time at which the object was uploaded.

Type: Timestamp

## Examples

### Sample Request for general purpose buckets

The following PUT request uploads a part (part number 2) in a multipart upload. The request
specifies a byte range from an existing object as the source of this upload. The request includes
the upload ID that you get in response to your Initiate Multipart Upload request.

```

PUT /newobject?partNumber=2&uploadId=VCVsb2FkIElEIGZvciBlbZZpbmcncyBteS1tb3ZpZS5tMnRzIHVwbG9hZR HTTP/1.1
Host: target-bucket.s3.<Region>.amazonaws.com
Date:  Mon, 11 Apr 2011 20:34:56 GMT
x-amz-copy-source: /source-bucket/sourceobject
x-amz-copy-source-range:bytes=500-6291456
Authorization: authorization string

```

### Sample Response for general purpose buckets

The response includes the ETag value. You need to retain this value to use when you send the
Complete Multipart Upload request.

```

HTTP/1.1 200 OK
x-amz-id-2: Vvag1LuByRx9e6j5Onimru9pO4ZVKnJ2Qz7/C1NPcfTWAtRPfTaOFg==
x-amz-request-id: 656c76696e6727732072657175657374
Date:  Mon, 11 Apr 2011 20:34:56 GMT
Server: AmazonS3

<CopyPartResult>
   <LastModified>2011-04-11T20:34:56.000Z</LastModified>
   <ETag>"9b2cf535f27731c974343645a3985328"</ETag>
</CopyPartResult>

```

### Sample Request for general purpose buckets

The following PUT request uploads a part (part number 2) in a multipart upload. The request does
not specify the optional byte range header, but requests the entire source object copy as part 2.
The request includes the upload ID that you got in response to your Initiate Multipart Upload
request.

```

PUT /newobject?partNumber=2&uploadId=VCVsb2FkIElEIGZvciBlbZZpbmcncyBteS1tb3ZpZS5tMnRzIHVwbG9hZR HTTP/1.1
Host: target-bucket.s3.<Region>.amazonaws.com
Date:  Mon, 11 Apr 2011 20:34:56 GMT
x-amz-copy-source: /source-bucket/sourceobject?versionId=3/L4kqtJlcpXroDTDmJ+rmSpXd3dIbrHY+MTRCxf3vjVBH40Nr8X8gdRQBpUMLUo
Authorization: authorization string

```

### Sample Response for general purpose buckets

The response includes the ETag value. You need to retain this value to use when you send the
Complete Multipart Upload request.

```

HTTP/1.1 200 OK
x-amz-id-2: Vvag1LuByRx9e6j5Onimru9pO4ZVKnJ2Qz7/C1NPcfTWAtRPfTaOFg==
x-amz-request-id: 656c76696e6727732072657175657374
x-amz-copy-source-version-id: 3/L4kqtJlcpXroDTDmJ+rmSpXd3dIbrHY+MTRCxf3vjVBH40Nr8X8gdRQBpUMLUo
Date:  Mon, 11 Apr 2011 20:34:56 GMT
Server: AmazonS3

<CopyPartResult>
   <LastModified>2011-04-11T20:34:56.000Z</LastModified>
   <ETag>"9b2cf535f27731c974343645a3985328"</ETag>
</CopyPartResult>

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/s3-2006-03-01/UploadPartCopy)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/s3-2006-03-01/UploadPartCopy)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/s3-2006-03-01/UploadPartCopy)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/s3-2006-03-01/UploadPartCopy)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/s3-2006-03-01/UploadPartCopy)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/s3-2006-03-01/UploadPartCopy)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/s3-2006-03-01/UploadPartCopy)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/s3-2006-03-01/UploadPartCopy)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/s3-2006-03-01/UploadPartCopy)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/s3-2006-03-01/UploadPartCopy)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

UploadPart

WriteGetObjectResponse
