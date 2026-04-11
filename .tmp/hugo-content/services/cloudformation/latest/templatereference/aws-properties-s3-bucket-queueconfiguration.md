This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::Bucket QueueConfiguration

Specifies the configuration for publishing messages to an Amazon Simple Queue Service (Amazon SQS)
queue when Amazon S3 detects specified events.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Event" : String,
  "Filter" : NotificationFilter,
  "Queue" : String
}

```

### YAML

```yaml

  Event: String
  Filter:
    NotificationFilter
  Queue: String

```

## Properties

`Event`

The Amazon S3 bucket event about which you want to publish messages to Amazon SQS. For
more information, see [Supported Event Types](../../../s3/latest/dev/notificationhowto.md) in the
_Amazon S3 User Guide_.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Filter`

The filtering rules that determine which objects trigger notifications. For example, you
can create a filter so that Amazon S3 sends notifications only when image files with a
`.jpg` extension are added to the bucket. For more information, see [Configuring event notifications using object key name filtering](../../../s3/latest/user-guide/notification-how-to-filtering.md) in the
_Amazon S3 User Guide_.

_Required_: No

_Type_: [NotificationFilter](aws-properties-s3-bucket-notificationfilter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Queue`

The Amazon Resource Name (ARN) of the Amazon SQS queue to which Amazon S3 publishes a
message when it detects events of the specified type. FIFO queues are not allowed when
enabling an SQS queue as the event notification destination.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PublicAccessBlockConfiguration

RecordExpiration

All content copied from https://docs.aws.amazon.com/.
