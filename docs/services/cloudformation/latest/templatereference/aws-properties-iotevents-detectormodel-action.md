This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTEvents::DetectorModel Action

An action to be performed when the `condition` is TRUE.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ClearTimer" : ClearTimer,
  "DynamoDB" : DynamoDB,
  "DynamoDBv2" : DynamoDBv2,
  "Firehose" : Firehose,
  "IotEvents" : IotEvents,
  "IotSiteWise" : IotSiteWise,
  "IotTopicPublish" : IotTopicPublish,
  "Lambda" : Lambda,
  "ResetTimer" : ResetTimer,
  "SetTimer" : SetTimer,
  "SetVariable" : SetVariable,
  "Sns" : Sns,
  "Sqs" : Sqs
}

```

### YAML

```yaml

  ClearTimer:
    ClearTimer
  DynamoDB:
    DynamoDB
  DynamoDBv2:
    DynamoDBv2
  Firehose:
    Firehose
  IotEvents:
    IotEvents
  IotSiteWise:
    IotSiteWise
  IotTopicPublish:
    IotTopicPublish
  Lambda:
    Lambda
  ResetTimer:
    ResetTimer
  SetTimer:
    SetTimer
  SetVariable:
    SetVariable
  Sns:
    Sns
  Sqs:
    Sqs

```

## Properties

`ClearTimer`

Information needed to clear the timer.

_Required_: No

_Type_: [ClearTimer](aws-properties-iotevents-detectormodel-cleartimer.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DynamoDB`

Writes to the DynamoDB table that you created. The default action payload contains all
attribute-value pairs that have the information about the detector model instance and the
event that triggered the action. You can customize the [payload](../../../../reference/iotevents/latest/apireference/api-payload.md). One column of the
DynamoDB table receives all attribute-value pairs in the payload that you specify. For more
information, see [Actions](../../../iotevents/latest/developerguide/iotevents-event-actions.md) in
_AWS IoT Events Developer Guide_.

_Required_: No

_Type_: [DynamoDB](aws-properties-iotevents-detectormodel-dynamodb.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DynamoDBv2`

Writes to the DynamoDB table that you created. The default action payload contains all
attribute-value pairs that have the information about the detector model instance and the
event that triggered the action. You can customize the [payload](../../../../reference/iotevents/latest/apireference/api-payload.md). A separate column of
the DynamoDB table receives one attribute-value pair in the payload that you specify. For more
information, see [Actions](../../../iotevents/latest/developerguide/iotevents-event-actions.md) in
_AWS IoT Events Developer Guide_.

_Required_: No

_Type_: [DynamoDBv2](aws-properties-iotevents-detectormodel-dynamodbv2.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Firehose`

Sends information about the detector model instance and the event that triggered the
action to an Amazon Kinesis Data Firehose delivery stream.

_Required_: No

_Type_: [Firehose](aws-properties-iotevents-detectormodel-firehose.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IotEvents`

Sends AWS IoT Events input, which passes information about the detector model instance and the
event that triggered the action.

_Required_: No

_Type_: [IotEvents](aws-properties-iotevents-detectormodel-iotevents.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IotSiteWise`

Sends information about the detector model instance and the event that triggered the
action to an asset property in AWS IoT SiteWise .

_Required_: No

_Type_: [IotSiteWise](aws-properties-iotevents-detectormodel-iotsitewise.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IotTopicPublish`

Publishes an MQTT message with the given topic to the AWS IoT message broker.

_Required_: No

_Type_: [IotTopicPublish](aws-properties-iotevents-detectormodel-iottopicpublish.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Lambda`

Calls a Lambda function, passing in information about the detector model instance and the
event that triggered the action.

_Required_: No

_Type_: [Lambda](aws-properties-iotevents-detectormodel-lambda.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResetTimer`

Information needed to reset the timer.

_Required_: No

_Type_: [ResetTimer](aws-properties-iotevents-detectormodel-resettimer.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SetTimer`

Information needed to set the timer.

_Required_: No

_Type_: [SetTimer](aws-properties-iotevents-detectormodel-settimer.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SetVariable`

Sets a variable to a specified value.

_Required_: No

_Type_: [SetVariable](aws-properties-iotevents-detectormodel-setvariable.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Sns`

Sends an Amazon SNS message.

_Required_: No

_Type_: [Sns](aws-properties-iotevents-detectormodel-sns.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Sqs`

Sends an Amazon SNS message.

_Required_: No

_Type_: [Sqs](aws-properties-iotevents-detectormodel-sqs.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::IoTEvents::DetectorModel

AssetPropertyTimestamp

All content copied from https://docs.aws.amazon.com/.
