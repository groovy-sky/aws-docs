This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTEvents::AlarmModel InitializationConfiguration

Specifies the default alarm state.
The configuration applies to all alarms that were created based on this alarm model.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DisabledOnInitialization" : Boolean
}

```

### YAML

```yaml

  DisabledOnInitialization: Boolean

```

## Properties

`DisabledOnInitialization`

The value must be `TRUE` or `FALSE`. If `FALSE`, all
alarm instances created based on the alarm model are activated. The default value is
`TRUE`.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Firehose

IotEvents

All content copied from https://docs.aws.amazon.com/.
