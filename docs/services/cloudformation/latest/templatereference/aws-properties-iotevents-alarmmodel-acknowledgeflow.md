This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTEvents::AlarmModel AcknowledgeFlow

Specifies whether to get notified for alarm state changes.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Enabled" : Boolean
}

```

### YAML

```yaml

  Enabled: Boolean

```

## Properties

`Enabled`

The value must be `TRUE` or `FALSE`. If `TRUE`, you
receive a notification when the alarm state changes. You must choose to acknowledge the
notification before the alarm state can return to `NORMAL`. If `FALSE`,
you won't receive notifications. The alarm automatically changes to the `NORMAL`
state when the input property value returns to the specified range.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::IoTEvents::AlarmModel

AlarmAction

All content copied from https://docs.aws.amazon.com/.
