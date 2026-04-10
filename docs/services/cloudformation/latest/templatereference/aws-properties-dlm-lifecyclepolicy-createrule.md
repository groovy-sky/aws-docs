This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DLM::LifecyclePolicy CreateRule

**\[Custom snapshot and AMI policies only\]** Specifies when the policy should create snapshots or AMIs.

###### Note

- You must specify either **CronExpression**, or
**Interval**, **IntervalUnit**,
and **Times**.

- If you need to specify an [ArchiveRule](../../../../reference/dlm/latest/apireference/api-archiverule.md)
for the schedule, then you must specify a creation frequency of at least
28 days.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CronExpression" : String,
  "Interval" : Integer,
  "IntervalUnit" : String,
  "Location" : String,
  "Scripts" : [ Script, ... ],
  "Times" : [ String, ... ]
}

```

### YAML

```yaml

  CronExpression: String
  Interval: Integer
  IntervalUnit: String
  Location: String
  Scripts:
    - Script
  Times:
    - String

```

## Properties

`CronExpression`

The schedule, as a Cron expression. The schedule interval must be between 1 hour and 1
year. For more information, see the [Cron and rate expressions](../../../eventbridge/latest/userguide/eb-scheduled-rule-pattern.md)
in the _Amazon EventBridge User Guide_.

_Required_: No

_Type_: String

_Pattern_: `cron\([^\n]{11,100}\)`

_Minimum_: `17`

_Maximum_: `106`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Interval`

The interval between snapshots. The supported values are 1, 2, 3, 4, 6, 8, 12, and 24.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IntervalUnit`

The interval unit.

_Required_: No

_Type_: String

_Allowed values_: `HOURS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Location`

**\[Custom snapshot policies only\]** Specifies the destination for snapshots created by the policy. The
allowed destinations depend on the location of the targeted resources.

- If the policy targets resources in a Region, then you must create snapshots
in the same Region as the source resource.

- If the policy targets resources in a Local Zone, you can create snapshots in
the same Local Zone or in its parent Region.

- If the policy targets resources on an Outpost, then you can create snapshots
on the same Outpost or in its parent Region.

Specify one of the following values:

- To create snapshots in the same Region as the source resource, specify
`CLOUD`.

- To create snapshots in the same Local Zone as the source resource, specify
`LOCAL_ZONE`.

- To create snapshots on the same Outpost as the source resource, specify
`OUTPOST_LOCAL`.

Default: `CLOUD`

_Required_: No

_Type_: String

_Allowed values_: `CLOUD | OUTPOST_LOCAL | LOCAL_ZONE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Scripts`

**\[Custom snapshot policies that target instances only\]** Specifies pre and/or post scripts for a snapshot lifecycle policy
that targets instances. This is useful for creating application-consistent snapshots, or for
performing specific administrative tasks before or after Amazon Data Lifecycle Manager initiates snapshot creation.

For more information, see [Automating \
application-consistent snapshots with pre and post scripts](../../../ec2/latest/userguide/automate-app-consistent-backups.md).

_Required_: No

_Type_: Array of [Script](aws-properties-dlm-lifecyclepolicy-script.md)

_Minimum_: `0`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Times`

The time, in UTC, to start the operation. The supported format is hh:mm.

The operation occurs within a one-hour window following the specified time. If you do
not specify a time, Amazon Data Lifecycle Manager selects a time within the next 24 hours.

_Required_: No

_Type_: Array of String

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ArchiveRule

CrossRegionCopyAction

All content copied from https://docs.aws.amazon.com/.
