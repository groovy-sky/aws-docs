# PutBucketNotification

###### Note

This operation is not supported for directory buckets.

No longer used, see the [PutBucketNotificationConfiguration](api-putbucketnotificationconfiguration.md) operation.

## Request Syntax

```nohighlight

PUT /?notification HTTP/1.1
Host: Bucket.s3.amazonaws.com
Content-MD5: ContentMD5
x-amz-sdk-checksum-algorithm: ChecksumAlgorithm
x-amz-expected-bucket-owner: ExpectedBucketOwner
<?xml version="1.0" encoding="UTF-8"?>
<NotificationConfiguration xmlns="http://s3.amazonaws.com/doc/2006-03-01/">
   <TopicConfiguration>
      <Event>string</Event>
      <Event>string</Event>
      ...
      <Id>string</Id>
      <Topic>string</Topic>
   </TopicConfiguration>
   <QueueConfiguration>
      <Event>string</Event>
      <Event>string</Event>
      ...
      <Id>string</Id>
      <Queue>string</Queue>
   </QueueConfiguration>
   <CloudFunctionConfiguration>
      <CloudFunction>string</CloudFunction>
      <Event>string</Event>
      <Event>string</Event>
      ...
      <Id>string</Id>
      <InvocationRole>string</InvocationRole>
   </CloudFunctionConfiguration>
</NotificationConfiguration>
```

## URI Request Parameters

The request uses the following URI parameters.

**[Bucket](#API_PutBucketNotification_RequestSyntax)**

The name of the bucket.

Required: Yes

**[Content-MD5](#API_PutBucketNotification_RequestSyntax)**

The MD5 hash of the `PutPublicAccessBlock` request body.

For requests made using the AWS Command Line Interface (CLI) or AWS SDKs, this field is calculated automatically.

**[x-amz-expected-bucket-owner](#API_PutBucketNotification_RequestSyntax)**

The account ID of the expected bucket owner. If the account ID that you provide does not match the actual owner of the bucket, the request fails with the HTTP status code `403 Forbidden` (access denied).

**[x-amz-sdk-checksum-algorithm](#API_PutBucketNotification_RequestSyntax)**

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

**[NotificationConfiguration](#API_PutBucketNotification_RequestSyntax)**

Root level tag for the NotificationConfiguration parameters.

Required: Yes

**[CloudFunctionConfiguration](#API_PutBucketNotification_RequestSyntax)**

Container for specifying the AWS Lambda notification configuration.

Type: [CloudFunctionConfiguration](api-cloudfunctionconfiguration.md) data type

Required: No

**[QueueConfiguration](#API_PutBucketNotification_RequestSyntax)**

This data type is deprecated. This data type specifies the configuration for publishing messages to
an Amazon Simple Queue Service (Amazon SQS) queue when Amazon S3 detects specified events.

Type: [QueueConfigurationDeprecated](api-queueconfigurationdeprecated.md) data type

Required: No

**[TopicConfiguration](#API_PutBucketNotification_RequestSyntax)**

This data type is deprecated. A container for specifying the configuration for publication of
messages to an Amazon Simple Notification Service (Amazon SNS) topic when Amazon S3 detects specified events.

Type: [TopicConfigurationDeprecated](api-topicconfigurationdeprecated.md) data type

Required: No

## Response Syntax

```

HTTP/1.1 200

```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/s3-2006-03-01/putbucketnotification.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3-2006-03-01/putbucketnotification.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3-2006-03-01/putbucketnotification.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3-2006-03-01/putbucketnotification.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3-2006-03-01/putbucketnotification.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3-2006-03-01/putbucketnotification.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3-2006-03-01/putbucketnotification.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3-2006-03-01/putbucketnotification.md)

- [AWS SDK for Python](../../../goto/boto3/s3-2006-03-01/putbucketnotification.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3-2006-03-01/putbucketnotification.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PutBucketMetricsConfiguration

PutBucketNotificationConfiguration

All content copied from https://docs.aws.amazon.com/.
