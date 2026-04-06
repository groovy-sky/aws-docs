# Using Amazon SQS, Amazon SNS, and Lambda

Enabling notifications is a bucket-level operation. You store notification
configuration information in the _notification_ subresource that's
associated with a bucket. After you create or change the bucket notification
configuration, it usually takes about five minutes for the changes to take effect. When
the notification is first enabled, an `s3:TestEvent` occurs. You can use any
of the following methods to manage notification configuration:

- **Using the Amazon S3 console** — You can use
the console UI to set a notification configuration on a bucket without having to
write any code. For more information, see [Enabling and configuring event notifications using the Amazon S3 console](https://docs.aws.amazon.com/AmazonS3/latest/userguide/enable-event-notifications.html).

- **Programmatically using the AWS SDKs**
— Internally, both the console and the SDKs call the Amazon S3 REST API to
manage _notification_ subresources that are associated with
the bucket. For examples of notification configurations that use AWS
SDK, see [Walkthrough: Configuring a bucket for notifications (SNS topic or SQS queue)](https://docs.aws.amazon.com/AmazonS3/latest/userguide/ways-to-add-notification-config-to-bucket.html).

###### Note

You can also make the Amazon S3 REST API calls directly from your code.
However, this can be cumbersome because to do so you must write code to
authenticate your requests.

Regardless of the method that you use, Amazon S3 stores the notification configuration as
XML in the _notification_ subresource that's associated with a
bucket. For information about bucket subresources, see [General purpose buckets configuration options](usingbucket.md#bucket-config-options-intro).

###### Note

If you have multiple failed event notifications due to deleted destinations you may receive the **Unable to validate the following destination configurations** when trying to delete them. You can resolve this in the S3 console by deleting all the failed notifications at the same time.

###### Topics

- [Granting permissions to publish event notification messages to a destination](https://docs.aws.amazon.com/AmazonS3/latest/userguide/grant-destinations-permissions-to-s3.html)

- [Enabling and configuring event notifications using the Amazon S3 console](https://docs.aws.amazon.com/AmazonS3/latest/userguide/enable-event-notifications.html)

- [Configuring event notifications programmatically](#event-notification-configuration)

- [Walkthrough: Configuring a bucket for notifications (SNS topic or SQS queue)](https://docs.aws.amazon.com/AmazonS3/latest/userguide/ways-to-add-notification-config-to-bucket.html)

- [Configuring event notifications using object key name filtering](https://docs.aws.amazon.com/AmazonS3/latest/userguide/notification-how-to-filtering.html)

- [Event message structure](https://docs.aws.amazon.com/AmazonS3/latest/userguide/notification-content-structure.html)

## Configuring event notifications programmatically

By default, notifications aren't enabled for any type of event. Therefore, the
_notification_ subresource initially stores an empty
configuration.

```

<NotificationConfiguration xmlns="http://s3.amazonaws.com/doc/2006-03-01/">
</NotificationConfiguration>
```

To enable notifications for events of specific types, you replace the XML with the
appropriate configuration that identifies the event types you want Amazon S3 to publish
and the destination where you want the events published. For each destination, you
add a corresponding XML configuration.

###### To publish event messages to an SQS queue

To set an SQS queue as the notification destination for one or more event
types, add the `QueueConfiguration`.

```nohighlight

<NotificationConfiguration>
  <QueueConfiguration>
    <Id>optional-id-string</Id>
    <Queue>sqs-queue-arn</Queue>
    <Event>event-type</Event>
    <Event>event-type</Event>
     ...
  </QueueConfiguration>
   ...
</NotificationConfiguration>
```

###### To publish event messages to an SNS topic

To set an SNS topic as the notification destination for specific event types,
add the `TopicConfiguration`.

```nohighlight

<NotificationConfiguration>
  <TopicConfiguration>
     <Id>optional-id-string</Id>
     <Topic>sns-topic-arn</Topic>
     <Event>event-type</Event>
     <Event>event-type</Event>
      ...
  </TopicConfiguration>
   ...
</NotificationConfiguration>
```

###### To invoke the AWS Lambda function and provide an event message as an argument

To set a Lambda function as the notification destination for specific event
types, add the `CloudFunctionConfiguration`.

```nohighlight

<NotificationConfiguration>
  <CloudFunctionConfiguration>
     <Id>optional-id-string</Id>
     <CloudFunction>cloud-function-arn</CloudFunction>
     <Event>event-type</Event>
     <Event>event-type</Event>
      ...
  </CloudFunctionConfiguration>
   ...
</NotificationConfiguration>
```

###### To remove all notifications configured on a bucket

To remove all notifications configured on a bucket, save an empty
`<NotificationConfiguration/>` element in the
_notification_ subresource.

When Amazon S3 detects an event of the specific type, it publishes a message with the
event information. For more information, see [Event message structure](https://docs.aws.amazon.com/AmazonS3/latest/userguide/notification-content-structure.html).

For more information about configuring event notifications, see the following topics:

- [Walkthrough: Configuring a bucket for notifications (SNS topic or SQS queue)](https://docs.aws.amazon.com/AmazonS3/latest/userguide/ways-to-add-notification-config-to-bucket.html).

- [Configuring event notifications using object key name filtering](https://docs.aws.amazon.com/AmazonS3/latest/userguide/notification-how-to-filtering.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Notification types and destinations

Granting permissions
