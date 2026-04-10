This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTEvents::AlarmModel AssetPropertyValue

A structure that contains value information. For more information, see [AssetPropertyValue](../../../../reference/iot-sitewise/latest/apireference/api-assetpropertyvalue.md) in the _AWS IoT SiteWise API Reference_.

You must use expressions for all parameters in `AssetPropertyValue`. The
expressions accept literals, operators, functions, references, and substitution
templates.

###### Examples

- For literal values, the expressions must contain single quotes. For example, the value
for the `quality` parameter can be `'GOOD'`.

- For references, you must specify either variables or input values. For example, the
value for the `quality` parameter can be
`$input.TemperatureInput.sensorData.quality`.

For more information,
see [Expressions](../../../iotevents/latest/developerguide/iotevents-expressions.md)
in the _AWS IoT Events Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Quality" : String,
  "Timestamp" : AssetPropertyTimestamp,
  "Value" : AssetPropertyVariant
}

```

### YAML

```yaml

  Quality: String
  Timestamp:
    AssetPropertyTimestamp
  Value:
    AssetPropertyVariant

```

## Properties

`Quality`

The quality of the asset property value. The value must be `'GOOD'`,
`'BAD'`, or `'UNCERTAIN'`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Timestamp`

The timestamp associated with the asset property value. The default is the current event
time.

_Required_: No

_Type_: [AssetPropertyTimestamp](aws-properties-iotevents-alarmmodel-assetpropertytimestamp.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The value to send to an asset property.

_Required_: Yes

_Type_: [AssetPropertyVariant](aws-properties-iotevents-alarmmodel-assetpropertyvariant.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AssetPropertyTimestamp

AssetPropertyVariant

All content copied from https://docs.aws.amazon.com/.
