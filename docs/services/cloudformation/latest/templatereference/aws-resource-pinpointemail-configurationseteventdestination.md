This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::PinpointEmail::ConfigurationSetEventDestination

Create an event destination. In Amazon Pinpoint, _events_ include message
sends, deliveries, opens, clicks, bounces, and complaints. _Event_
_destinations_ are places that you can send information about these events
to. For example, you can send event data to Amazon SNS to receive notifications when you
receive bounces or complaints, or you can use Amazon Kinesis Data Firehose to stream data to Amazon S3 for long-term
storage.

A single configuration set can include more than one event destination.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::PinpointEmail::ConfigurationSetEventDestination",
  "Properties" : {
      "ConfigurationSetName" : String,
      "EventDestination" : EventDestination,
      "EventDestinationName" : String
    }
}

```

### YAML

```yaml

Type: AWS::PinpointEmail::ConfigurationSetEventDestination
Properties:
  ConfigurationSetName: String
  EventDestination:
    EventDestination
  EventDestinationName: String

```

## Properties

`ConfigurationSetName`

The name of the configuration set that contains the event destination that you want to
modify.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EventDestination`

An object that defines the event destination.

_Required_: No

_Type_: [EventDestination](aws-properties-pinpointemail-configurationseteventdestination-eventdestination.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EventDestinationName`

The name of the event destination that you want to modify.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource name. For example:

`{ "Ref": "myEventDestination" }`

For the Amazon Pinpoint event destination `myEventDestination`, Ref returns the
name of the configuration set event destination.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TrackingOptions

CloudWatchDestination

All content copied from https://docs.aws.amazon.com/.
