This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTEvents::DetectorModel DynamoDBv2

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

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Payload" : Payload,
  "TableName" : String
}

```

### YAML

```yaml

  Payload:
    Payload
  TableName: String

```

## Properties

`Payload`

Information needed to configure the payload.

By default, AWS IoT Events generates a standard payload in JSON for any action. This action payload
contains all attribute-value pairs that have the information about the detector model instance
and the event triggered the action. To configure the action payload, you can use
`contentExpression`.

_Required_: No

_Type_: [Payload](aws-properties-iotevents-detectormodel-payload.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TableName`

The name of the DynamoDB table.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DynamoDB

Event

All content copied from https://docs.aws.amazon.com/.
