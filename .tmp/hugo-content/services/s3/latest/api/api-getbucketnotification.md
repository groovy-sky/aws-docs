# GetBucketNotification

###### Note

This operation is not supported for directory buckets.

No longer used, see [GetBucketNotificationConfiguration](api-getbucketnotificationconfiguration.md).

## Request Syntax

```nohighlight

GET /?notification HTTP/1.1
Host: Bucket.s3.amazonaws.com
x-amz-expected-bucket-owner: ExpectedBucketOwner

```

## URI Request Parameters

The request uses the following URI parameters.

**[Bucket](#API_GetBucketNotification_RequestSyntax)**

The name of the bucket for which to get the notification configuration.

When you use this API operation with an access point, provide the alias of the access point in place of the bucket name.

When you use this API operation with an Object Lambda access point, provide the alias of the Object Lambda access point in place of the bucket name.
If the Object Lambda access point alias in a request is not valid, the error code `InvalidAccessPointAliasError` is returned.
For more information about `InvalidAccessPointAliasError`, see [List of\
Error Codes](errorresponses.md#ErrorCodeList).

Required: Yes

**[x-amz-expected-bucket-owner](#API_GetBucketNotification_RequestSyntax)**

The account ID of the expected bucket owner. If the account ID that you provide does not match the actual owner of the bucket, the request fails with the HTTP status code `403 Forbidden` (access denied).

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
<?xml version="1.0" encoding="UTF-8"?>
<NotificationConfiguration>
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

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in XML format by the service.

**[NotificationConfiguration](#API_GetBucketNotification_ResponseSyntax)**

Root level tag for the NotificationConfiguration parameters.

Required: Yes

**[CloudFunctionConfiguration](#API_GetBucketNotification_ResponseSyntax)**

Container for specifying the AWS Lambda notification configuration.

Type: [CloudFunctionConfiguration](api-cloudfunctionconfiguration.md) data type

**[QueueConfiguration](#API_GetBucketNotification_ResponseSyntax)**

This data type is deprecated. This data type specifies the configuration for publishing messages to
an Amazon Simple Queue Service (Amazon SQS) queue when Amazon S3 detects specified events.

Type: [QueueConfigurationDeprecated](api-queueconfigurationdeprecated.md) data type

**[TopicConfiguration](#API_GetBucketNotification_ResponseSyntax)**

This data type is deprecated. A container for specifying the configuration for publication of
messages to an Amazon Simple Notification Service (Amazon SNS) topic when Amazon S3 detects specified events.

Type: [TopicConfigurationDeprecated](api-topicconfigurationdeprecated.md) data type

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/s3-2006-03-01/getbucketnotification.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3-2006-03-01/getbucketnotification.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3-2006-03-01/getbucketnotification.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3-2006-03-01/getbucketnotification.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3-2006-03-01/getbucketnotification.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3-2006-03-01/getbucketnotification.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3-2006-03-01/getbucketnotification.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3-2006-03-01/getbucketnotification.md)

- [AWS SDK for Python](../../../goto/boto3/s3-2006-03-01/getbucketnotification.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3-2006-03-01/getbucketnotification.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetBucketMetricsConfiguration

GetBucketNotificationConfiguration

All content copied from https://docs.aws.amazon.com/.
