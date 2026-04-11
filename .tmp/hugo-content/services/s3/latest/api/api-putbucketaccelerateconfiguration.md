# PutBucketAccelerateConfiguration

###### Note

This operation is not supported for directory buckets.

Sets the accelerate configuration of an existing bucket. Amazon S3 Transfer Acceleration is a
bucket-level feature that enables you to perform faster data transfers to Amazon S3.

To use this operation, you must have permission to perform the
`s3:PutAccelerateConfiguration` action. The bucket owner has this permission by default.
The bucket owner can grant this permission to others. For more information about permissions, see [Permissions Related to Bucket Subresource Operations](../userguide/using-with-s3-actions.md#using-with-s3-actions-related-to-bucket-subresources) and [Managing Access Permissions to Your Amazon S3\
Resources](../userguide/s3-access-control.md).

The Transfer Acceleration state of a bucket can be set to one of the following two values:

- Enabled – Enables accelerated data transfers to the bucket.

- Suspended – Disables accelerated data transfers to the bucket.

The [GetBucketAccelerateConfiguration](api-getbucketaccelerateconfiguration.md) action returns the transfer acceleration state of a
bucket.

After setting the Transfer Acceleration state of a bucket to Enabled, it might take up to thirty
minutes before the data transfer rates to the bucket increase.

The name of the bucket used for Transfer Acceleration must be DNS-compliant and must not contain
periods (".").

For more information about transfer acceleration, see [Transfer Acceleration](../dev/transfer-acceleration.md).

The following operations are related to `PutBucketAccelerateConfiguration`:

- [GetBucketAccelerateConfiguration](api-getbucketaccelerateconfiguration.md)

- [CreateBucket](api-createbucket.md)

###### Important

You must URL encode any signed header values that contain spaces. For example, if your header value is `my  file.txt`, containing two spaces after `my`, you must URL encode this value to `my%20%20file.txt`.

## Request Syntax

```nohighlight

PUT /?accelerate HTTP/1.1
Host: Bucket.s3.amazonaws.com
x-amz-expected-bucket-owner: ExpectedBucketOwner
x-amz-sdk-checksum-algorithm: ChecksumAlgorithm
<?xml version="1.0" encoding="UTF-8"?>
<AccelerateConfiguration xmlns="http://s3.amazonaws.com/doc/2006-03-01/">
   <Status>string</Status>
</AccelerateConfiguration>
```

## URI Request Parameters

The request uses the following URI parameters.

**[Bucket](#API_PutBucketAccelerateConfiguration_RequestSyntax)**

The name of the bucket for which the accelerate configuration is set.

Required: Yes

**[x-amz-expected-bucket-owner](#API_PutBucketAccelerateConfiguration_RequestSyntax)**

The account ID of the expected bucket owner. If the account ID that you provide does not match the actual owner of the bucket, the request fails with the HTTP status code `403 Forbidden` (access denied).

**[x-amz-sdk-checksum-algorithm](#API_PutBucketAccelerateConfiguration_RequestSyntax)**

Indicates the algorithm used to create the checksum for the request when you use the SDK. This header will not provide any
additional functionality if you don't use the SDK. When you send this header, there must be a corresponding `x-amz-checksum` or
`x-amz-trailer` header sent. Otherwise, Amazon S3 fails the request with the HTTP status code `400 Bad Request`. For more
information, see [Checking object integrity](../userguide/checking-object-integrity.md) in
the _Amazon S3 User Guide_.

If you provide an individual checksum, Amazon S3 ignores any provided `ChecksumAlgorithm`
parameter.

Valid Values: `CRC32 | CRC32C | SHA1 | SHA256 | CRC64NVME`

## Request Body

The request accepts the following data in XML format.

**[AccelerateConfiguration](#API_PutBucketAccelerateConfiguration_RequestSyntax)**

Root level tag for the AccelerateConfiguration parameters.

Required: Yes

**[Status](#API_PutBucketAccelerateConfiguration_RequestSyntax)**

Specifies the transfer acceleration status of the bucket.

Type: String

Valid Values: `Enabled | Suspended`

Required: No

## Response Syntax

```

HTTP/1.1 200

```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## Examples

### Sample Request: Add transfer acceleration configuration to set acceleration status

The following is an example of a `PUT /?accelerate` request that enables transfer
acceleration for the bucket named `examplebucket`.

```

PUT /?accelerate HTTP/1.1
Host: examplebucket.s3.<Region>.amazonaws.com
Date: Mon, 11 Apr 2016 12:00:00 GMT
Authorization: authorization string
Content-Type: text/plain
Content-Length: length

<AccelerateConfiguration xmlns="http://s3.amazonaws.com/doc/2006-03-01/">
  <Status>Enabled</Status>
</AccelerateConfiguration>

```

### Sample Response

This example illustrates one usage of PutBucketAccelerateConfiguration.

```

HTTP/1.1 200 OK
x-amz-id-2: YgIPIfBiKa2bj0KMg95r/0zo3emzU4dzsD4rcKCHQUAdQkf3ShJTOOpXUueF6QKo
x-amz-request-id: 236A8905248E5A01
Date: Mon, 11 Apr 2016 12:00:00 GMT
Content-Length: 0
Server: AmazonS3

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/s3-2006-03-01/putbucketaccelerateconfiguration.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3-2006-03-01/putbucketaccelerateconfiguration.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3-2006-03-01/putbucketaccelerateconfiguration.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3-2006-03-01/putbucketaccelerateconfiguration.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3-2006-03-01/putbucketaccelerateconfiguration.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3-2006-03-01/putbucketaccelerateconfiguration.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3-2006-03-01/putbucketaccelerateconfiguration.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3-2006-03-01/putbucketaccelerateconfiguration.md)

- [AWS SDK for Python](../../../goto/boto3/s3-2006-03-01/putbucketaccelerateconfiguration.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3-2006-03-01/putbucketaccelerateconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PutBucketAbac

PutBucketAcl

All content copied from https://docs.aws.amazon.com/.
