This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::Flow Maintenance

The maintenance setting of a flow.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaintenanceDay" : String,
  "MaintenanceStartHour" : String
}

```

### YAML

```yaml

  MaintenanceDay: String
  MaintenanceStartHour: String

```

## Properties

`MaintenanceDay`

A day of a week when the maintenance will happen. Use Monday/Tuesday/Wednesday/Thursday/Friday/Saturday/Sunday.

_Required_: Yes

_Type_: String

_Allowed values_: `Monday | Tuesday | Wednesday | Thursday | Friday | Saturday | Sunday`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaintenanceStartHour`

UTC time when the maintenance will happen. Use 24-hour HH:MM format. Minutes must be 00. Example: 13:00. The default value is 02:00.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Interface

MediaStream
