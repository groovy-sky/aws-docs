This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::HoursOfOperation HoursOfOperationOverride

Information about the hours of operations override.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EffectiveFrom" : String,
  "EffectiveTill" : String,
  "HoursOfOperationOverrideId" : String,
  "OverrideConfig" : [ HoursOfOperationOverrideConfig, ... ],
  "OverrideDescription" : String,
  "OverrideName" : String,
  "OverrideType" : String,
  "RecurrenceConfig" : RecurrenceConfig
}

```

### YAML

```yaml

  EffectiveFrom: String
  EffectiveTill: String
  HoursOfOperationOverrideId: String
  OverrideConfig:
    - HoursOfOperationOverrideConfig
  OverrideDescription: String
  OverrideName: String
  OverrideType: String
  RecurrenceConfig:
    RecurrenceConfig

```

## Properties

`EffectiveFrom`

The date from which the hours of operation override would be effective.

_Required_: Yes

_Type_: String

_Pattern_: `^\d{4}-\d{2}-\d{2}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EffectiveTill`

The date until the hours of operation override is effective.

_Required_: Yes

_Type_: String

_Pattern_: `^\d{4}-\d{2}-\d{2}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HoursOfOperationOverrideId`

The identifier for the hours of operation override.

_Required_: No

_Type_: String

_Pattern_: `^[-a-zA-Z0-9]*$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OverrideConfig`

Property description not available.

_Required_: Yes

_Type_: Array of [HoursOfOperationOverrideConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connect-hoursofoperation-hoursofoperationoverrideconfig.html)

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OverrideDescription`

Property description not available.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `250`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OverrideName`

Property description not available.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `127`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OverrideType`

Whether the override will be defined as a _standard_ or as a _recurring event_.

_Required_: No

_Type_: String

_Allowed values_: `STANDARD | OPEN | CLOSED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RecurrenceConfig`

Configuration for a recurring event.

_Required_: No

_Type_: [RecurrenceConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connect-hoursofoperation-recurrenceconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

HoursOfOperationConfig

HoursOfOperationOverrideConfig
