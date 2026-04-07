This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApplicationSignals::ServiceLevelObjective ExclusionWindow

The time window to be excluded from the SLO performance metrics.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Reason" : String,
  "RecurrenceRule" : RecurrenceRule,
  "StartTime" : String,
  "Window" : Window
}

```

### YAML

```yaml

  Reason: String
  RecurrenceRule:
    RecurrenceRule
  StartTime: String
  Window:
    Window

```

## Properties

`Reason`

The reason for the time exclusion windows. For example, maintenance.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RecurrenceRule`

The recurrence rule for the time exclusion window.

_Required_: No

_Type_: [RecurrenceRule](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-applicationsignals-servicelevelobjective-recurrencerule.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StartTime`

The start time of the time exclusion window.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Window`

The time exclusion window.

_Required_: Yes

_Type_: [Window](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-applicationsignals-servicelevelobjective-window.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Dimension

Goal
