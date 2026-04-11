# PutBucketNotificationConfiguration

###### Note

This operation is not supported for directory buckets.

Enables notifications of specified events for a bucket. For more information about event
notifications, see [Configuring Event Notifications](../dev/notificationhowto.md).

Using this API, you can replace an existing notification configuration. The configuration is an XML
file that defines the event types that you want Amazon S3 to publish and the destination where you want Amazon S3
to publish an event notification when it detects an event of the specified type.

By default, your bucket has no event notifications configured. That is, the notification
configuration will be an empty `NotificationConfiguration`.

`<NotificationConfiguration>`

`</NotificationConfiguration>`

This action replaces the existing notification configuration with the configuration you include in
the request body.

After Amazon S3 receives this request, it first verifies that any Amazon Simple Notification Service
(Amazon SNS) or Amazon Simple Queue Service (Amazon SQS) destination exists, and that the bucket owner
has permission to publish to it by sending a test notification. In the case of AWS Lambda destinations,
Amazon S3 verifies that the Lambda function permissions grant Amazon S3 permission to invoke the function from the
Amazon S3 bucket. For more information, see [Configuring Notifications for Amazon S3\
Events](../dev/notificationhowto.md).

You can disable notifications by adding the empty NotificationConfiguration element.

For more information about the number of event notification configurations that you can create per
bucket, see [Amazon S3 service\
quotas](../../../../general/latest/gr/s3.md#limits_s3) in _AWS General Reference_.

By default, only the bucket owner can configure notifications on a bucket. However, bucket owners
can use a bucket policy to grant permission to other users to set this configuration with the required
`s3:PutBucketNotification` permission.

###### Note

The PUT notification is an atomic operation. For example, suppose your notification configuration
includes SNS topic, SQS queue, and Lambda function configurations. When you send a PUT request with
this configuration, Amazon S3 sends test messages to your SNS topic. If the message fails, the entire PUT
action will fail, and Amazon S3 will not add the configuration to your bucket.

If the configuration in the request body includes only one `TopicConfiguration`
specifying only the `s3:ReducedRedundancyLostObject` event type, the response will also
include the `x-amz-sns-test-message-id` header containing the message ID of the test
notification sent to the topic.

The following action is related to `PutBucketNotificationConfiguration`:

- [GetBucketNotificationConfiguration](api-getbucketnotificationconfiguration.md)

###### Important

You must URL encode any signed header values that contain spaces. For example, if your header value is `my  file.txt`, containing two spaces after `my`, you must URL encode this value to `my%20%20file.txt`.

## Request Syntax

```nohighlight

PUT /?notification HTTP/1.1
Host: Bucket.s3.amazonaws.com
x-amz-expected-bucket-owner: ExpectedBucketOwner
x-amz-skip-destination-validation: SkipDestinationValidation
<?xml version="1.0" encoding="UTF-8"?>
<NotificationConfiguration xmlns="http://s3.amazonaws.com/doc/2006-03-01/">
   <TopicConfiguration>
      <Event>string</Event>
      ...
      <Filter>
         <S3Key>
            <FilterRule>
               <Name>string</Name>
               <Value>string</Value>
            </FilterRule>
            ...
         </S3Key>
      </Filter>
      <Id>string</Id>
      <Topic>string</Topic>
   </TopicConfiguration>
   ...
   <QueueConfiguration>
      <Event>string</Event>
      ...
      <Filter>
         <S3Key>
            <FilterRule>
               <Name>string</Name>
               <Value>string</Value>
            </FilterRule>
            ...
         </S3Key>
      </Filter>
      <Id>string</Id>
      <Queue>string</Queue>
   </QueueConfiguration>
   ...
   <CloudFunctionConfiguration>
      <Event>string</Event>
      ...
      <Filter>
         <S3Key>
            <FilterRule>
               <Name>string</Name>
               <Value>string</Value>
            </FilterRule>
            ...
         </S3Key>
      </Filter>
      <Id>string</Id>
      <CloudFunction>string</CloudFunction>
   </CloudFunctionConfiguration>
   ...
   <EventBridgeConfiguration>
   </EventBridgeConfiguration>
</NotificationConfiguration>
```

## URI Request Parameters

The request uses the following URI parameters.

**[Bucket](#API_PutBucketNotificationConfiguration_RequestSyntax)**

The name of the bucket.

Required: Yes

**[x-amz-expected-bucket-owner](#API_PutBucketNotificationConfiguration_RequestSyntax)**

The account ID of the expected bucket owner. If the account ID that you provide does not match the actual owner of the bucket, the request fails with the HTTP status code `403 Forbidden` (access denied).

**[x-amz-skip-destination-validation](#API_PutBucketNotificationConfiguration_RequestSyntax)**

Skips validation of Amazon SQS, Amazon SNS, and AWS Lambda destinations.
True or false value.

## Request Body

The request accepts the following data in XML format.

**[NotificationConfiguration](#API_PutBucketNotificationConfiguration_RequestSyntax)**

Root level tag for the NotificationConfiguration parameters.

Required: Yes

**[CloudFunctionConfiguration](#API_PutBucketNotificationConfiguration_RequestSyntax)**

Describes the AWS Lambda functions to invoke and the events for which to invoke them.

Type: Array of [LambdaFunctionConfiguration](api-lambdafunctionconfiguration.md) data types

Required: No

**[EventBridgeConfiguration](#API_PutBucketNotificationConfiguration_RequestSyntax)**

Enables delivery of events to Amazon EventBridge.

Type: [EventBridgeConfiguration](api-eventbridgeconfiguration.md) data type

Required: No

**[QueueConfiguration](#API_PutBucketNotificationConfiguration_RequestSyntax)**

The Amazon Simple Queue Service queues to publish messages to and the events for which to publish
messages.

Type: Array of [QueueConfiguration](api-queueconfiguration.md) data types

Required: No

**[TopicConfiguration](#API_PutBucketNotificationConfiguration_RequestSyntax)**

The topic to which notifications are sent and the events for which notifications are
generated.

Type: Array of [TopicConfiguration](api-topicconfiguration.md) data types

Required: No

## Response Syntax

```

HTTP/1.1 200

```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## Examples

### Example 1: Configure notification to invoke a cloud function in Lambda

The following notification configuration includes CloudFunctionConfiguration, which identifies
the event type for which Amazon S3 can invoke a cloud function and the name of the cloud function to
invoke.

```

<NotificationConfiguration>
  <CloudFunctionConfiguration>
    <Id>ObjectCreatedEvents</Id>
    <CloudFunction>arn:aws:lambda:us-west-2:35667example:function:CreateThumbnail</CloudFunction>
    <Event>s3:ObjectCreated:*</Event>
  </CloudFunctionConfiguration>
</NotificationConfiguration>

```

### Example

The following PUT uploads the notification configuration. The action replaces the existing
notification configuration.

```

PUT http://s3.<Region>.amazonaws.com/examplebucket?notification= HTTP/1.1
User-Agent: s3curl 2.0
Host: s3.amazonaws.com
Pragma: no-cache
Accept: */*
Proxy-Connection: Keep-Alive
Authorization: authorization string
Date: Mon, 13 Oct 2014 23:14:52 +0000
Content-Length: length

[request body]

```

### Sample Response

This example illustrates one usage of PutBucketNotificationConfiguration.

```

HTTP/1.1 200 OK
x-amz-id-2: 8+FlwagBSoT2qpMaGlfCUkRkFR5W3OeS7UhhoBb17j+kqvpS2cSFlgJ5coLd53d2
x-amz-request-id: E5BA4600A3937335
Date: Fri, 31 Oct 2014 01:49:50 GMT
Content-Length: 0
Server: AmazonS3

```

### Example 2: Configure a notification with multiple destinations

The following notification configuration includes the topic and queue configurations:

- A topic configuration identifying an SNS topic for Amazon S3 to publish events of the
`s3:ReducedRedundancyLostObject` type.

- A queue configuration identifying an SQS queue for Amazon S3 to publish events of the
`s3:ObjectCreated:*` type.

```

<NotificationConfiguration>
  <TopicConfiguration>
    <Topic>arn:aws:sns:us-east-1:356671443308:s3notificationtopic2</Topic>
    <Event>s3:ReducedRedundancyLostObject</Event>
  </TopicConfiguration>
  <QueueConfiguration>
    <Queue>arn:aws:sqs:us-east-1:356671443308:s3notificationqueue</Queue>
    <Event>s3:ObjectCreated:*</Event>
  </QueueConfiguration>
</NotificationConfiguration>

```

### Example

The following PUT request against the notification subresource of the `examplebucket`
bucket sends the preceding notification configuration in the request body. The action replaces the
existing notification configuration on the bucket.

```

PUT http://s3.<Region>.amazonaws.com/examplebucket?notification= HTTP/1.1
User-Agent: s3curl 2.0
Host: s3.amazonaws.com
Pragma: no-cache
Accept: */*
Proxy-Connection: Keep-Alive
Authorization: authorization string
Date: Mon, 13 Oct 2014 22:58:43 +0000
Content-Length: 391
Expect: 100-continue

```

### Example 3: Configure a notification with object key name filtering

The following notification configuration contains a queue configuration identifying an Amazon
SQS queue for Amazon S3 to publish events to of the `s3:ObjectCreated:Put` type. The events
will be published whenever an object that has a prefix of `images/` and a
`.jpg` suffix is PUT to a bucket. For more examples of notification configurations that
use filtering, see [Configuring Event Notifications](../dev/notificationhowto.md).

```

<NotificationConfiguration>
  <QueueConfiguration>
      <Id>1</Id>
      <Filter>
          <S3Key>
              <FilterRule>
                  <Name>prefix</Name>
                  <Value>images/</Value>
              </FilterRule>
              <FilterRule>
                  <Name>suffix</Name>
                  <Value>.jpg</Value>
              </FilterRule>
          </S3Key>
     </Filter>
     <Queue>arn:aws:sqs:us-west-2:444455556666:s3notificationqueue</Queue>
     <Event>s3:ObjectCreated:Put</Event>
  </QueueConfiguration>
</NotificationConfiguration>

```

### Example

The following PUT request against the notification subresource of the `examplebucket`
bucket sends the preceding notification configuration in the request body. The action replaces the
existing notification configuration on the bucket.

```

PUT http://s3.<Region>.amazonaws.com/examplebucket?notification= HTTP/1.1
User-Agent: s3curl 2.0
Host: s3.amazonaws.com
Pragma: no-cache
Accept: */*
Proxy-Connection: Keep-Alive
Authorization: authorization string
Date: Mon, 13 Oct 2014 22:58:43 +0000
Content-Length: length
Expect: 100-continue

```

### Sample Response

This example illustrates one usage of PutBucketNotificationConfiguration.

```

HTTP/1.1 200 OK
x-amz-id-2: SlvJLkfunoAGILZK3KqHSSUq4kwbudkrROmESoHOpDacULy+cxRoR1Svrfoyvg2A
x-amz-request-id: BB1BA8E12D6A80B7
Date: Mon, 13 Oct 2014 22:58:44 GMT
Content-Length: 0
Server: AmazonS3

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/s3-2006-03-01/putbucketnotificationconfiguration.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3-2006-03-01/putbucketnotificationconfiguration.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3-2006-03-01/putbucketnotificationconfiguration.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3-2006-03-01/putbucketnotificationconfiguration.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3-2006-03-01/putbucketnotificationconfiguration.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3-2006-03-01/putbucketnotificationconfiguration.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3-2006-03-01/putbucketnotificationconfiguration.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3-2006-03-01/putbucketnotificationconfiguration.md)

- [AWS SDK for Python](../../../goto/boto3/s3-2006-03-01/putbucketnotificationconfiguration.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3-2006-03-01/putbucketnotificationconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PutBucketNotification

PutBucketOwnershipControls

All content copied from https://docs.aws.amazon.com/.
