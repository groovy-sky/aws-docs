# PutObjectLockConfiguration

###### Note

This operation is not supported for directory buckets.

Places an Object Lock configuration on the specified bucket. The rule specified in the Object Lock
configuration will be applied by default to every new object placed in the specified bucket. For more
information, see [Locking\
Objects](../dev/object-lock.md).

###### Note

- The `DefaultRetention` settings require both a mode and a period.

- The `DefaultRetention` period can be either `Days` or `Years`
but you must select one. You cannot specify `Days` and `Years` at the same
time.

- You can enable Object Lock for new or existing buckets. For more information, see [Configuring\
Object Lock](../userguide/object-lock-configure.md).

###### Important

You must URL encode any signed header values that contain spaces. For example, if your header value is `my  file.txt`, containing two spaces after `my`, you must URL encode this value to `my%20%20file.txt`.

## Request Syntax

```nohighlight

PUT /?object-lock HTTP/1.1
Host: Bucket.s3.amazonaws.com
x-amz-request-payer: RequestPayer
x-amz-bucket-object-lock-token: Token
Content-MD5: ContentMD5
x-amz-sdk-checksum-algorithm: ChecksumAlgorithm
x-amz-expected-bucket-owner: ExpectedBucketOwner
<?xml version="1.0" encoding="UTF-8"?>
<ObjectLockConfiguration xmlns="http://s3.amazonaws.com/doc/2006-03-01/">
   <ObjectLockEnabled>string</ObjectLockEnabled>
   <Rule>
      <DefaultRetention>
         <Days>integer</Days>
         <Mode>string</Mode>
         <Years>integer</Years>
      </DefaultRetention>
   </Rule>
</ObjectLockConfiguration>
```

## URI Request Parameters

The request uses the following URI parameters.

**[Bucket](#API_PutObjectLockConfiguration_RequestSyntax)**

The bucket whose Object Lock configuration you want to create or replace.

Required: Yes

**[Content-MD5](#API_PutObjectLockConfiguration_RequestSyntax)**

The MD5 hash for the request body.

For requests made using the AWS Command Line Interface (CLI) or AWS SDKs, this field is calculated automatically.

**[x-amz-bucket-object-lock-token](#API_PutObjectLockConfiguration_RequestSyntax)**

A token to allow Object Lock to be enabled for an existing bucket.

**[x-amz-expected-bucket-owner](#API_PutObjectLockConfiguration_RequestSyntax)**

The account ID of the expected bucket owner. If the account ID that you provide does not match the actual owner of the bucket, the request fails with the HTTP status code `403 Forbidden` (access denied).

**[x-amz-request-payer](#API_PutObjectLockConfiguration_RequestSyntax)**

Confirms that the requester knows that they will be charged for the request. Bucket owners need not
specify this parameter in their requests. If either the source or destination S3 bucket has Requester
Pays enabled, the requester will pay for the corresponding charges. For information about
downloading objects from Requester Pays buckets, see [Downloading Objects in Requester Pays\
Buckets](../dev/objectsinrequesterpaysbuckets.md) in the _Amazon S3 User Guide_.

###### Note

This functionality is not supported for directory buckets.

Valid Values: `requester`

**[x-amz-sdk-checksum-algorithm](#API_PutObjectLockConfiguration_RequestSyntax)**

Indicates the algorithm used to create the checksum for the object when you use the SDK. This header will not provide any
additional functionality if you don't use the SDK. When you send this header, there must be a corresponding `x-amz-checksum` or
`x-amz-trailer` header sent. Otherwise, Amazon S3 fails the request with the HTTP status code `400 Bad Request`. For more
information, see [Checking object integrity](../userguide/checking-object-integrity.md) in
the _Amazon S3 User Guide_.

If you provide an individual checksum, Amazon S3 ignores any provided `ChecksumAlgorithm`
parameter.

Valid Values: `CRC32 | CRC32C | SHA1 | SHA256 | CRC64NVME`

## Request Body

The request accepts the following data in XML format.

**[ObjectLockConfiguration](#API_PutObjectLockConfiguration_RequestSyntax)**

Root level tag for the ObjectLockConfiguration parameters.

Required: Yes

**[ObjectLockEnabled](#API_PutObjectLockConfiguration_RequestSyntax)**

Indicates whether this bucket has an Object Lock configuration enabled. Enable
`ObjectLockEnabled` when you apply `ObjectLockConfiguration` to a bucket.

Type: String

Valid Values: `Enabled`

Required: No

**[Rule](#API_PutObjectLockConfiguration_RequestSyntax)**

Specifies the Object Lock rule for the specified object. Enable the this rule when you apply
`ObjectLockConfiguration` to a bucket. Bucket settings require both a mode and a period.
The period can be either `Days` or `Years` but you must select one. You cannot
specify `Days` and `Years` at the same time.

Type: [ObjectLockRule](api-objectlockrule.md) data type

Required: No

## Response Syntax

```nohighlight

HTTP/1.1 200
x-amz-request-charged: RequestCharged

```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The response returns the following HTTP headers.

**[x-amz-request-charged](#API_PutObjectLockConfiguration_ResponseSyntax)**

If present, indicates that the requester was successfully charged for the request. For more
information, see [Using Requester Pays buckets for storage transfers and usage](../userguide/requesterpaysbuckets.md) in the _Amazon Simple_
_Storage Service user guide_.

###### Note

This functionality is not supported for directory buckets.

Valid Values: `requester`

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/s3-2006-03-01/putobjectlockconfiguration.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3-2006-03-01/putobjectlockconfiguration.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3-2006-03-01/putobjectlockconfiguration.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3-2006-03-01/putobjectlockconfiguration.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3-2006-03-01/putobjectlockconfiguration.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3-2006-03-01/putobjectlockconfiguration.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3-2006-03-01/putobjectlockconfiguration.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3-2006-03-01/putobjectlockconfiguration.md)

- [AWS SDK for Python](../../../goto/boto3/s3-2006-03-01/putobjectlockconfiguration.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3-2006-03-01/putobjectlockconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PutObjectLegalHold

PutObjectRetention

All content copied from https://docs.aws.amazon.com/.
