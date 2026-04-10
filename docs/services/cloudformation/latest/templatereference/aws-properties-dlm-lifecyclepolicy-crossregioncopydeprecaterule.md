This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DLM::LifecyclePolicy CrossRegionCopyDeprecateRule

**\[Custom AMI policies only\]** Specifies an AMI deprecation rule for cross-Region AMI copies created by an AMI policy.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Interval" : Integer,
  "IntervalUnit" : String
}

```

### YAML

```yaml

  Interval: Integer
  IntervalUnit: String

```

## Properties

`Interval`

The period after which to deprecate the cross-Region AMI copies. The period must be less than or
equal to the cross-Region AMI copy retention period, and it can't be greater than 10 years. This is
equivalent to 120 months, 520 weeks, or 3650 days.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IntervalUnit`

The unit of time in which to measure the **Interval**. For example,
to deprecate a cross-Region AMI copy after 3 months, specify `Interval=3` and
`IntervalUnit=MONTHS`.

_Required_: Yes

_Type_: String

_Allowed values_: `DAYS | WEEKS | MONTHS | YEARS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CrossRegionCopyAction

CrossRegionCopyRetainRule

All content copied from https://docs.aws.amazon.com/.
