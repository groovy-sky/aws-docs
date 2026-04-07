This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::InspectorV2::CisScanConfiguration Schedule

The schedule the CIS scan configuration runs on. Each CIS scan configuration has
exactly one type of schedule.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Daily" : DailySchedule,
  "Monthly" : MonthlySchedule,
  "OneTime" : Json,
  "Weekly" : WeeklySchedule
}

```

### YAML

```yaml

  Daily:
    DailySchedule
  Monthly:
    MonthlySchedule
  OneTime: Json
  Weekly:
    WeeklySchedule

```

## Properties

`Daily`

A daily schedule.

_Required_: No

_Type_: [DailySchedule](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-inspectorv2-cisscanconfiguration-dailyschedule.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Monthly`

A monthly schedule.

_Required_: No

_Type_: [MonthlySchedule](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-inspectorv2-cisscanconfiguration-monthlyschedule.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OneTime`

A one time schedule.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Weekly`

A weekly schedule.

_Required_: No

_Type_: [WeeklySchedule](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-inspectorv2-cisscanconfiguration-weeklyschedule.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

MonthlySchedule

Time
