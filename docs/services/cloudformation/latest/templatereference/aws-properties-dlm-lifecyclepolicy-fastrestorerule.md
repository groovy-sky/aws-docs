This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DLM::LifecyclePolicy FastRestoreRule

**\[Custom snapshot policies only\]** Specifies a rule for enabling fast snapshot restore for snapshots created by
snapshot policies. You can enable fast snapshot restore based on either a count or a
time interval.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AvailabilityZoneIds" : [ String, ... ],
  "AvailabilityZones" : [ String, ... ],
  "Count" : Integer,
  "Interval" : Integer,
  "IntervalUnit" : String
}

```

### YAML

```yaml

  AvailabilityZoneIds:
    - String
  AvailabilityZones:
    - String
  Count: Integer
  Interval: Integer
  IntervalUnit: String

```

## Properties

`AvailabilityZoneIds`

The Availability Zone Ids in which to enable fast snapshot restore.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AvailabilityZones`

The Availability Zones in which to enable fast snapshot restore.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Count`

The number of snapshots to be enabled with fast snapshot restore.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Interval`

The amount of time to enable fast snapshot restore. The maximum is 100 years. This is
equivalent to 1200 months, 5200 weeks, or 36500 days.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IntervalUnit`

The unit of time for enabling fast snapshot restore.

_Required_: No

_Type_: String

_Allowed values_: `HOURS | DAYS | WEEKS | MONTHS | YEARS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Exclusions

Parameters

All content copied from https://docs.aws.amazon.com/.
