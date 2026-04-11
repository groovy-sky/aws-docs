This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSMContacts::Rotation WeeklySetting

Information about rotations that recur weekly.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DayOfWeek" : String,
  "HandOffTime" : String
}

```

### YAML

```yaml

  DayOfWeek: String
  HandOffTime: String

```

## Properties

`DayOfWeek`

The day of the week when weekly recurring on-call shift rotations begins.

_Required_: Yes

_Type_: String

_Allowed values_: `MON | TUE | WED | THU | FRI | SAT | SUN`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HandOffTime`

The time of day when a weekly recurring on-call shift rotation begins.

_Required_: Yes

_Type_: String

_Pattern_: `^([0-9]|0[0-9]|1[0-9]|2[0-3]):[0-5][0-9]$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Next

All content copied from https://docs.aws.amazon.com/.
