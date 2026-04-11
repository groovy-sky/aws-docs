This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTEvents::AlarmModel AssetPropertyVariant

A structure that contains an asset property value. For more information, see [Variant](../../../../reference/iot-sitewise/latest/apireference/api-variant.md)
in the _AWS IoT SiteWise API Reference_.

You must use expressions for all parameters in `AssetPropertyVariant`. The
expressions accept literals, operators, functions, references, and substitution
templates.

###### Examples

- For literal values, the expressions must contain single quotes. For example, the value
for the `integerValue` parameter can be `'100'`.

- For references, you must specify either variables or parameters. For example, the
value for the `booleanValue` parameter can be
`$variable.offline`.

- For a substitution template, you must use `${}`, and the template must be
in single quotes. A substitution template can also contain a combination of literals,
operators, functions, references, and substitution templates.

In the following example, the value for the `doubleValue` parameter uses a
substitution template.

`'${$input.TemperatureInput.sensorData.temperature * 6 / 5 + 32}'`

For more information,
see [Expressions](../../../iotevents/latest/developerguide/iotevents-expressions.md)
in the _AWS IoT Events Developer Guide_.

You must specify one of the following value types, depending on the `dataType`
of the specified asset property. For more information, see [AssetProperty](../../../../reference/iot-sitewise/latest/apireference/api-assetproperty.md) in the
_AWS IoT SiteWise API Reference_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BooleanValue" : String,
  "DoubleValue" : String,
  "IntegerValue" : String,
  "StringValue" : String
}

```

### YAML

```yaml

  BooleanValue: String
  DoubleValue: String
  IntegerValue: String
  StringValue:
    String

```

## Properties

`BooleanValue`

The asset property value is a Boolean value that must be `'TRUE'` or
`'FALSE'`. You must use an expression, and the evaluated result should be a
Boolean value.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DoubleValue`

The asset property value is a double. You must use an expression, and the evaluated result
should be a double.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IntegerValue`

The asset property value is an integer. You must use an expression, and the evaluated
result should be an integer.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StringValue`

The asset property value is a string. You must use an expression, and the evaluated result
should be a string.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AssetPropertyValue

DynamoDB

All content copied from https://docs.aws.amazon.com/.
