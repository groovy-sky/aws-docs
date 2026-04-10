This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DLM::LifecyclePolicy CrossRegionCopyRetainRule

Specifies a retention rule for cross-Region snapshot copies created by snapshot or
event-based policies, or cross-Region AMI copies created by AMI policies. After the
retention period expires, the cross-Region copy is deleted.

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

The amount of time to retain a cross-Region snapshot or AMI copy. The maximum is 100 years.
This is equivalent to 1200 months, 5200 weeks, or 36500 days.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IntervalUnit`

The unit of time for time-based retention. For example, to retain a cross-Region copy for
3 months, specify `Interval=3` and `IntervalUnit=MONTHS`.

_Required_: Yes

_Type_: String

_Allowed values_: `DAYS | WEEKS | MONTHS | YEARS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CrossRegionCopyDeprecateRule

CrossRegionCopyRule

All content copied from https://docs.aws.amazon.com/.
