This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SMSVOICE::ConfigurationSet EventDestination

Contains information about an event destination.

Event destinations are associated with configuration sets, which enable you to publish
message sending events to CloudWatch, Firehose, or Amazon SNS.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CloudWatchLogsDestination" : CloudWatchLogsDestination,
  "Enabled" : Boolean,
  "EventDestinationName" : String,
  "KinesisFirehoseDestination" : KinesisFirehoseDestination,
  "MatchingEventTypes" : [ String, ... ],
  "SnsDestination" : SnsDestination
}

```

### YAML

```yaml

  CloudWatchLogsDestination:
    CloudWatchLogsDestination
  Enabled: Boolean
  EventDestinationName: String
  KinesisFirehoseDestination:
    KinesisFirehoseDestination
  MatchingEventTypes:
    - String
  SnsDestination:
    SnsDestination

```

## Properties

`CloudWatchLogsDestination`

An object that contains information about an event destination that sends logging
events to Amazon CloudWatch logs.

_Required_: No

_Type_: [CloudWatchLogsDestination](aws-properties-smsvoice-configurationset-cloudwatchlogsdestination.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Enabled`

When set to true events will be logged.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EventDestinationName`

The name of the EventDestination.

_Required_: Yes

_Type_: String

_Pattern_: `^[A-Za-z0-9_-]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KinesisFirehoseDestination`

An object that contains information about an event destination for logging to Amazon Data Firehose.

_Required_: No

_Type_: [KinesisFirehoseDestination](aws-properties-smsvoice-configurationset-kinesisfirehosedestination.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MatchingEventTypes`

An array of event types that determine which events to log.

###### Note

The `TEXT_SENT` event type is not supported.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SnsDestination`

An object that contains information about an event destination that sends logging
events to Amazon SNS.

_Required_: No

_Type_: [SnsDestination](aws-properties-smsvoice-configurationset-snsdestination.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CloudWatchLogsDestination

KinesisFirehoseDestination

All content copied from https://docs.aws.amazon.com/.
