This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudTrail::Dashboard Frequency

Specifies the frequency for a dashboard refresh schedule.

For a custom dashboard, you can schedule a refresh for every 1, 6, 12, or 24 hours, or every day.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Unit" : String,
  "Value" : Integer
}

```

### YAML

```yaml

  Unit: String
  Value: Integer

```

## Properties

`Unit`

The unit to use for the refresh.

For custom dashboards, the unit can be `HOURS` or `DAYS`.

For the Highlights dashboard, the `Unit` must be `HOURS`.

_Required_: Yes

_Type_: String

_Allowed values_: `HOURS | DAYS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The value for the refresh schedule.

For custom dashboards, the following values are valid when the unit is `HOURS`: `1`, `6`, `12`, `24`

For custom dashboards, the only valid value when the unit is `DAYS` is `1`.

For the Highlights dashboard, the `Value` must be `6`.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::CloudTrail::Dashboard

RefreshSchedule

All content copied from https://docs.aws.amazon.com/.
