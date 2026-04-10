This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTEvents::AlarmModel AssetPropertyTimestamp

A structure that contains timestamp information. For more information, see [TimeInNanos](../../../../reference/iot-sitewise/latest/apireference/api-timeinnanos.md) in the _AWS IoT SiteWise API Reference_.

You must use expressions for all parameters in `AssetPropertyTimestamp`. The
expressions accept literals, operators, functions, references, and substitution
templates.

###### Examples

- For literal values, the expressions must contain single quotes. For example, the value
for the `timeInSeconds` parameter can be `'1586400675'`.

- For references, you must specify either variables or input values. For example, the
value for the `offsetInNanos` parameter can be
`$variable.time`.

- For a substitution template, you must use `${}`, and the template must be
in single quotes. A substitution template can also contain a combination of literals,
operators, functions, references, and substitution templates.

In the following example, the value for the `timeInSeconds` parameter uses
a substitution template.

`'${$input.TemperatureInput.sensorData.timestamp / 1000}'`

For more information,
see [Expressions](../../../iotevents/latest/developerguide/iotevents-expressions.md)
in the _AWS IoT Events Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "OffsetInNanos" : String,
  "TimeInSeconds" : String
}

```

### YAML

```yaml

  OffsetInNanos: String
  TimeInSeconds: String

```

## Properties

`OffsetInNanos`

The nanosecond offset converted from `timeInSeconds`. The valid range is
between 0-999999999.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeInSeconds`

The timestamp, in seconds, in the Unix epoch format. The valid range is between
1-31556889864403199.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AlarmRule

AssetPropertyValue

All content copied from https://docs.aws.amazon.com/.
