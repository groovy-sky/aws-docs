This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTEvents::DetectorModel ResetTimer

Information required to reset the timer. The timer is reset to the previously evaluated
result of the duration. The duration expression isn't reevaluated when you reset the
timer.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "TimerName" : String
}

```

### YAML

```yaml

  TimerName: String

```

## Properties

`TimerName`

The name of the timer to reset.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Payload

SetTimer

All content copied from https://docs.aws.amazon.com/.
