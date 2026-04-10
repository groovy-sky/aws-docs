This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DLM::LifecyclePolicy RetentionArchiveTier

**\[Custom snapshot policies only\]** Describes the retention rule for archived snapshots. Once the archive
retention threshold is met, the snapshots are permanently deleted from the archive tier.

###### Note

The archive retention rule must retain snapshots in the archive tier for a minimum
of 90 days.

For **count-based schedules**, you must specify **Count**. For **age-based**
**schedules**, you must specify **Interval** and
**IntervalUnit**.

For more information about using snapshot archiving, see [Considerations for \
snapshot lifecycle policies](../../../ec2/latest/userguide/snapshot-ami-policy.md#dlm-archive).

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

The maximum number of snapshots to retain in the archive storage tier for each volume.
The count must ensure that each snapshot remains in the archive tier for at least
90 days. For example, if the schedule creates snapshots every 30 days, you must specify a
count of 3 or more to ensure that each snapshot is archived for at least 90 days.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Interval`

Specifies the period of time to retain snapshots in the archive tier. After this period
expires, the snapshot is permanently deleted.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IntervalUnit`

The unit of time in which to measure the **Interval**. For
example, to retain a snapshots in the archive tier for 6 months, specify `Interval=6`
and `IntervalUnit=MONTHS`.

_Required_: No

_Type_: String

_Allowed values_: `DAYS | WEEKS | MONTHS | YEARS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RetainRule

Schedule

All content copied from https://docs.aws.amazon.com/.
