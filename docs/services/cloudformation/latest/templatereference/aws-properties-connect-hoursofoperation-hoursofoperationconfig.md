This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::HoursOfOperation HoursOfOperationConfig

Contains information about the hours of operation.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Day" : String,
  "EndTime" : HoursOfOperationTimeSlice,
  "StartTime" : HoursOfOperationTimeSlice
}

```

### YAML

```yaml

  Day: String
  EndTime:
    HoursOfOperationTimeSlice
  StartTime:
    HoursOfOperationTimeSlice

```

## Properties

`Day`

The day that the hours of operation applies to.

_Required_: Yes

_Type_: String

_Allowed values_: `SUNDAY | MONDAY | TUESDAY | WEDNESDAY | THURSDAY | FRIDAY | SATURDAY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EndTime`

The end time that your contact center closes.

_Required_: Yes

_Type_: [HoursOfOperationTimeSlice](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connect-hoursofoperation-hoursofoperationtimeslice.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StartTime`

The start time that your contact center opens.

_Required_: Yes

_Type_: [HoursOfOperationTimeSlice](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connect-hoursofoperation-hoursofoperationtimeslice.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Connect::HoursOfOperation

HoursOfOperationOverride
