This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DLM::LifecyclePolicy RetainRule

**\[Custom snapshot and AMI policies only\]** Specifies a retention rule for snapshots created by snapshot policies, or for AMIs
created by AMI policies.

###### Note

For snapshot policies that have an [ArchiveRule](../../../../reference/dlm/latest/apireference/api-archiverule.md), this retention rule
applies to standard tier retention. When the retention threshold is met, snapshots
are moved from the standard to the archive tier.

For snapshot policies that do not have an **ArchiveRule**, snapshots
are permanently deleted when this retention threshold is met.

You can retain snapshots based on either a count or a time interval.

- **Count-based retention**

You must specify **Count**.
If you specify an [ArchiveRule](../../../../reference/dlm/latest/apireference/api-archiverule.md) for the schedule, then you can specify a retention count of
`0` to archive snapshots immediately after creation. If you specify a [FastRestoreRule](../../../../reference/dlm/latest/apireference/api-fastrestorerule.md),
[ShareRule](../../../../reference/dlm/latest/apireference/api-sharerule.md), or a
[CrossRegionCopyRule](../../../../reference/dlm/latest/apireference/api-crossregioncopyrule.md), then you must specify a retention count
of `1` or more.

- **Age-based retention**

You must specify **Interval**
and **IntervalUnit**. If you specify an [ArchiveRule](../../../../reference/dlm/latest/apireference/api-archiverule.md) for the
schedule, then you can specify a retention interval of `0` days to archive snapshots immediately
after creation. If you specify a [FastRestoreRule](../../../../reference/dlm/latest/apireference/api-fastrestorerule.md), [ShareRule](../../../../reference/dlm/latest/apireference/api-sharerule.md), or a
[CrossRegionCopyRule](../../../../reference/dlm/latest/apireference/api-crossregioncopyrule.md),
then you must specify a retention interval of `1` day or
more.

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

The number of snapshots to retain for each volume, up to a maximum of 1000. For example if you want to
retain a maximum of three snapshots, specify `3`. When the fourth snapshot is created, the
oldest retained snapshot is deleted, or it is moved to the archive tier if you have specified an
[ArchiveRule](../../../../reference/dlm/latest/apireference/api-archiverule.md).

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Interval`

The amount of time to retain each snapshot. The maximum is 100 years. This is
equivalent to 1200 months, 5200 weeks, or 36500 days.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IntervalUnit`

The unit of time for time-based retention. For example, to retain snapshots for 3 months, specify
`Interval=3` and `IntervalUnit=MONTHS`. Once the snapshot has been retained for
3 months, it is deleted, or it is moved to the archive tier if you have specified an
[ArchiveRule](../../../../reference/dlm/latest/apireference/api-archiverule.md).

_Required_: No

_Type_: String

_Allowed values_: `DAYS | WEEKS | MONTHS | YEARS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PolicyDetails

RetentionArchiveTier

All content copied from https://docs.aws.amazon.com/.
