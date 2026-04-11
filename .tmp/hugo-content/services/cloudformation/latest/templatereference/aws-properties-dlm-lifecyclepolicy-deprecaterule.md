This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DLM::LifecyclePolicy DeprecateRule

**\[Custom AMI policies only\]** Specifies an AMI deprecation rule for AMIs created by an AMI lifecycle policy.

For age-based schedules, you must specify **Interval** and
**IntervalUnit**. For count-based schedules, you must specify
**Count**.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Count" : Integer,
  "Interval" : Integer,
  "IntervalUnit" : String
}

```

### YAML

```yaml

  Count: Integer
  Interval: Integer
  IntervalUnit: String

```

## Properties

`Count`

If the schedule has a count-based retention rule, this parameter specifies the number of oldest
AMIs to deprecate. The count must be less than or equal to the schedule's retention count, and it
can't be greater than 1000.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Interval`

If the schedule has an age-based retention rule, this parameter specifies the period after which
to deprecate AMIs created by the schedule. The period must be less than or equal to the schedule's
retention period, and it can't be greater than 10 years. This is equivalent to 120 months, 520
weeks, or 3650 days.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IntervalUnit`

The unit of time in which to measure the **Interval**.

_Required_: No

_Type_: String

_Allowed values_: `DAYS | WEEKS | MONTHS | YEARS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CrossRegionCopyTarget

EncryptionConfiguration

All content copied from https://docs.aws.amazon.com/.
