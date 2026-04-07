This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::PinpointEmail::ConfigurationSetEventDestination EventDestination

In Amazon Pinpoint, _events_ include message sends, deliveries, opens,
clicks, bounces, and complaints. _Event destinations_ are places that
you can send information about these events to. For example, you can send event data to
Amazon SNS to receive notifications when you receive bounces or complaints, or you can use
Amazon Kinesis Data Firehose to stream data to Amazon S3 for long-term storage.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CloudWatchDestination" : CloudWatchDestination,
  "Enabled" : Boolean,
  "KinesisFirehoseDestination" : KinesisFirehoseDestination,
  "MatchingEventTypes" : [ String, ... ],
  "PinpointDestination" : PinpointDestination,
  "SnsDestination" : SnsDestination
}

```

### YAML

```yaml

  CloudWatchDestination:
    CloudWatchDestination
  Enabled: Boolean
  KinesisFirehoseDestination:
    KinesisFirehoseDestination
  MatchingEventTypes:
    - String
  PinpointDestination:
    PinpointDestination
  SnsDestination:
    SnsDestination

```

## Properties

`CloudWatchDestination`

An object that defines an Amazon CloudWatch destination for email events. You can use Amazon CloudWatch to
monitor and gain insights on your email sending metrics.

_Required_: No

_Type_: [CloudWatchDestination](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-pinpointemail-configurationseteventdestination-cloudwatchdestination.html)

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

`KinesisFirehoseDestination`

An object that defines an Amazon Kinesis Data Firehose destination for email events. You can use Amazon Kinesis Data Firehose to
stream data to other services, such as Amazon S3 and Amazon Redshift.

_Required_: No

_Type_: [KinesisFirehoseDestination](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-pinpointemail-configurationseteventdestination-kinesisfirehosedestination.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MatchingEventTypes`

The types of events that Amazon Pinpoint sends to the specified event destinations.
Acceptable values: `SEND`, `REJECT`, `BOUNCE`,
`COMPLAINT`, `DELIVERY`, `OPEN`, `CLICK`, and
`RENDERING_FAILURE`.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PinpointDestination`

An object that defines a Amazon Pinpoint destination for email events. You can use Amazon Pinpoint events
to create attributes in Amazon Pinpoint projects. You can use these attributes to create segments
for your campaigns.

_Required_: No

_Type_: [PinpointDestination](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-pinpointemail-configurationseteventdestination-pinpointdestination.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SnsDestination`

An object that defines an Amazon SNS destination for email events. You can use Amazon SNS to
send notification when certain email events occur.

_Required_: No

_Type_: [SnsDestination](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-pinpointemail-configurationseteventdestination-snsdestination.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DimensionConfiguration

KinesisFirehoseDestination
