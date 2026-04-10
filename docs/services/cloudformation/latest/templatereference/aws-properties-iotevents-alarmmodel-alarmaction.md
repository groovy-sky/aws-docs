This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTEvents::AlarmModel AlarmAction

Specifies one of the following actions to receive notifications when the alarm state
changes.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DynamoDB" : DynamoDB,
  "DynamoDBv2" : DynamoDBv2,
  "Firehose" : Firehose,
  "IotEvents" : IotEvents,
  "IotSiteWise" : IotSiteWise,
  "IotTopicPublish" : IotTopicPublish,
  "Lambda" : Lambda,
  "Sns" : Sns,
  "Sqs" : Sqs
}

```

### YAML

```yaml

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
  Sns:
    Sns
  Sqs:
    Sqs

```

## Properties

`DynamoDB`

Defines an action to write to the Amazon DynamoDB table that you created. The standard action
payload contains all the information about the detector model instance and the event that
triggered the action. You can customize the [payload](../../../../reference/iotevents/latest/apireference/api-payload.md). One column of the
DynamoDB table receives all attribute-value pairs in the payload that you specify.

You must use expressions for all parameters in `DynamoDBAction`. The expressions
accept literals, operators, functions, references, and substitution templates.

###### Examples

- For literal values, the expressions must contain single quotes. For example, the value
for the `hashKeyType` parameter can be `'STRING'`.

- For references, you must specify either variables or input values. For example, the
value for the `hashKeyField` parameter can be
`$input.GreenhouseInput.name`.

- For a substitution template, you must use `${}`, and the template must be
in single quotes. A substitution template can also contain a combination of literals,
operators, functions, references, and substitution templates.

In the following example, the value for the `hashKeyValue` parameter uses a
substitution template.

`'${$input.GreenhouseInput.temperature * 6 / 5 + 32} in Fahrenheit'`

- For a string concatenation, you must use `+`. A string concatenation can
also contain a combination of literals, operators, functions, references, and substitution
templates.

In the following example, the value for the `tableName` parameter uses a
string concatenation.

`'GreenhouseTemperatureTable ' + $input.GreenhouseInput.date`

For more information,
see [Expressions](../../../iotevents/latest/developerguide/iotevents-expressions.md)
in the _AWS IoT Events Developer Guide_.

If the defined payload type is a string, `DynamoDBAction` writes non-JSON data to
the DynamoDB table as binary data. The DynamoDB console displays the data as Base64-encoded text.
The value for the `payloadField` parameter is
`<payload-field>_raw`.

_Required_: No

_Type_: [DynamoDB](aws-properties-iotevents-alarmmodel-dynamodb.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DynamoDBv2`

Defines an action to write to the Amazon DynamoDB table that you created. The default action
payload contains all the information about the detector model instance and the event that
triggered the action. You can customize the [payload](../../../../reference/iotevents/latest/apireference/api-payload.md). A separate column of
the DynamoDB table receives one attribute-value pair in the payload that you specify.

You must use expressions for all parameters in `DynamoDBv2Action`. The expressions
accept literals, operators, functions, references, and substitution templates.

###### Examples

- For literal values, the expressions must contain single quotes. For example, the value
for the `tableName` parameter can be
`'GreenhouseTemperatureTable'`.

- For references, you must specify either variables or input values. For example, the
value for the `tableName` parameter can be
`$variable.ddbtableName`.

- For a substitution template, you must use `${}`, and the template must be
in single quotes. A substitution template can also contain a combination of literals,
operators, functions, references, and substitution templates.

In the following example, the value for the `contentExpression` parameter
in `Payload` uses a substitution template.

`'{\"sensorID\": \"${$input.GreenhouseInput.sensor_id}\", \"temperature\":
              \"${$input.GreenhouseInput.temperature * 9 / 5 + 32}\"}'`

- For a string concatenation, you must use `+`. A string concatenation can
also contain a combination of literals, operators, functions, references, and substitution
templates.

In the following example, the value for the `tableName` parameter uses a
string concatenation.

`'GreenhouseTemperatureTable ' + $input.GreenhouseInput.date`

For more information,
see [Expressions](../../../iotevents/latest/developerguide/iotevents-expressions.md)
in the _AWS IoT Events Developer Guide_.

The value for the `type` parameter in `Payload` must be
`JSON`.

_Required_: No

_Type_: [DynamoDBv2](aws-properties-iotevents-alarmmodel-dynamodbv2.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Firehose`

Sends information about the detector model instance and the event that triggered the
action to an Amazon Kinesis Data Firehose delivery stream.

_Required_: No

_Type_: [Firehose](aws-properties-iotevents-alarmmodel-firehose.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IotEvents`

Sends an AWS IoT Events input, passing in information about the detector model instance and the
event that triggered the action.

_Required_: No

_Type_: [IotEvents](aws-properties-iotevents-alarmmodel-iotevents.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IotSiteWise`

Sends information about the detector model instance and the event that triggered the
action to a specified asset property in AWS IoT SiteWise.

You must use expressions for all parameters in `IotSiteWiseAction`. The
expressions accept literals, operators, functions, references, and substitutions
templates.

###### Examples

- For literal values, the expressions must contain single quotes. For example, the value
for the `propertyAlias` parameter can be
`'/company/windfarm/3/turbine/7/temperature'`.

- For references, you must specify either variables or input values. For example, the
value for the `assetId` parameter can be
`$input.TurbineInput.assetId1`.

- For a substitution template, you must use `${}`, and the template must be
in single quotes. A substitution template can also contain a combination of literals,
operators, functions, references, and substitution templates.

In the following example, the value for the `propertyAlias` parameter uses
a substitution template.

`'company/windfarm/${$input.TemperatureInput.sensorData.windfarmID}/turbine/
              ${$input.TemperatureInput.sensorData.turbineID}/temperature'`

You must specify either `propertyAlias` or both `assetId` and
`propertyId` to identify the target asset property in AWS IoT SiteWise.

For more information,
see [Expressions](../../../iotevents/latest/developerguide/iotevents-expressions.md)
in the _AWS IoT Events Developer Guide_.

_Required_: No

_Type_: [IotSiteWise](aws-properties-iotevents-alarmmodel-iotsitewise.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IotTopicPublish`

Information required to publish the MQTT message through the AWS IoT message broker.

_Required_: No

_Type_: [IotTopicPublish](aws-properties-iotevents-alarmmodel-iottopicpublish.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Lambda`

Calls a Lambda function, passing in information about the detector model instance and the
event that triggered the action.

_Required_: No

_Type_: [Lambda](aws-properties-iotevents-alarmmodel-lambda.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Sns`

Information required to publish the Amazon SNS message.

_Required_: No

_Type_: [Sns](aws-properties-iotevents-alarmmodel-sns.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Sqs`

Sends information about the detector model instance and the event that triggered the
action to an Amazon SQS queue.

_Required_: No

_Type_: [Sqs](aws-properties-iotevents-alarmmodel-sqs.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AcknowledgeFlow

AlarmCapabilities

All content copied from https://docs.aws.amazon.com/.
