This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTEvents::DetectorModel SetTimer

Information needed to set the timer.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DurationExpression" : String,
  "Seconds" : Integer,
  "TimerName" : String
}

```

### YAML

```yaml

  DurationExpression: String
  Seconds: Integer
  TimerName: String

```

## Properties

`DurationExpression`

The duration of the timer, in seconds. You can use a string expression that includes
numbers, variables ( `$variable.<variable-name>`), and input values
( `$input.<input-name>.<path-to-datum>`) as the duration. The range of
the duration is 1-31622400 seconds. To ensure accuracy, the minimum duration is 60 seconds.
The evaluated result of the duration is rounded down to the nearest whole number.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Seconds`

The number of seconds until the timer expires. The minimum value is 60 seconds to ensure
accuracy. The maximum value is 31622400 seconds.

_Required_: No

_Type_: Integer

_Minimum_: `60`

_Maximum_: `31622400`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimerName`

The name of the timer.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ResetTimer

SetVariable

All content copied from https://docs.aws.amazon.com/.
