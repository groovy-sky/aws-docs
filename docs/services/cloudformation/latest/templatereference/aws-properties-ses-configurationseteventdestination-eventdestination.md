This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::ConfigurationSetEventDestination EventDestination

In the Amazon SES API v2, _events_ include message sends, deliveries, opens,
clicks, bounces, complaints and delivery delays. _Event destinations_
are places that you can send information about these events to. For example, you can
send event data to Amazon SNS to receive notifications when you receive bounces or
complaints, or you can use Amazon Kinesis Data Firehose to stream data to Amazon S3 for long-term storage.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CloudWatchDestination" : CloudWatchDestination,
  "Enabled" : Boolean,
  "EventBridgeDestination" : EventBridgeDestination,
  "KinesisFirehoseDestination" : KinesisFirehoseDestination,
  "MatchingEventTypes" : [ String, ... ],
  "Name" : String,
  "SnsDestination" : SnsDestination
}

```

### YAML

```yaml

  CloudWatchDestination:
    CloudWatchDestination
  Enabled: Boolean
  EventBridgeDestination:
    EventBridgeDestination
  KinesisFirehoseDestination:
    KinesisFirehoseDestination
  MatchingEventTypes:
    - String
  Name: String
  SnsDestination:
    SnsDestination

```

## Properties

`CloudWatchDestination`

An object that defines an Amazon CloudWatch destination for email events. You can use Amazon CloudWatch to
monitor and gain insights on your email sending metrics.

_Required_: No

_Type_: [CloudWatchDestination](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ses-configurationseteventdestination-cloudwatchdestination.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Enabled`

If `true`, the event destination is enabled. When the event destination is
enabled, the specified event types are sent to the destinations in this
`EventDestinationDefinition`.

If `false`, the event destination is disabled. When the event destination
is disabled, events aren't sent to the specified destinations.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EventBridgeDestination`

An object that defines an Amazon EventBridge destination for email events. You can use Amazon EventBridge to
send notifications when certain email events occur.

_Required_: No

_Type_: [EventBridgeDestination](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ses-configurationseteventdestination-eventbridgedestination.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KinesisFirehoseDestination`

An object that contains the delivery stream ARN and the IAM role ARN associated with
an Amazon Kinesis Firehose event destination.

_Required_: No

_Type_: [KinesisFirehoseDestination](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ses-configurationseteventdestination-kinesisfirehosedestination.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MatchingEventTypes`

The types of events that Amazon SES sends to the specified event destinations.

- `SEND` \- The send request was successful and SES will
attempt to deliver the message to the recipient’s mail server. (If account-level
or global suppression is being used, SES will still count it as a send,
but delivery is suppressed.)

- `REJECT` \- SES accepted the email, but determined that it
contained a virus and didn’t attempt to deliver it to the recipient’s mail
server.

- `BOUNCE` \- ( _Hard bounce_) The recipient's
mail server permanently rejected the email. ( _Soft bounces_
are only included when SES fails to deliver the email after retrying for
a period of time.)

- `COMPLAINT` \- The email was successfully delivered to the
recipient’s mail server, but the recipient marked it as spam.

- `DELIVERY` \- SES successfully delivered the email to the
recipient's mail server.

- `OPEN` \- The recipient received the message and opened it in their
email client.

- `CLICK` \- The recipient clicked one or more links in the
email.

- `RENDERING_FAILURE` \- The email wasn't sent because of a
template rendering issue. This event type can occur when template data is
missing, or when there is a mismatch between template parameters and data. (This
event type only occurs when you send email using the [`SendEmail`](https://docs.aws.amazon.com/ses/latest/APIReference-V2/API_SendEmail.html) or [`SendBulkEmail`](https://docs.aws.amazon.com/ses/latest/APIReference-V2/API_SendBulkEmail.html) API operations.)

- `DELIVERY_DELAY` \- The email couldn't be delivered to the
recipient’s mail server because a temporary issue occurred. Delivery delays can
occur, for example, when the recipient's inbox is full, or when the
receiving email server experiences a transient issue.

- `SUBSCRIPTION` \- The email was successfully delivered, but the
recipient updated their subscription preferences by clicking on an
_unsubscribe_ link as part of your [subscription\
management](https://docs.aws.amazon.com/ses/latest/dg/sending-email-subscription-management.html).

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the event destination. The name must meet the following
requirements:

- Contain only ASCII letters (a-z, A-Z), numbers (0-9), underscores (\_), or
dashes (-).

- Contain 64 characters or fewer.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]{0,64}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SnsDestination`

An object that contains the topic ARN associated with an Amazon Simple Notification
Service (Amazon SNS) event destination.

_Required_: No

_Type_: [SnsDestination](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ses-configurationseteventdestination-snsdestination.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

EventBridgeDestination

KinesisFirehoseDestination
