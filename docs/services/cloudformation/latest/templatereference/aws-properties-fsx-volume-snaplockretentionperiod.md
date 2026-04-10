This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FSx::Volume SnaplockRetentionPeriod

The configuration to set the retention period of an FSx for ONTAP SnapLock volume. The retention
period includes default, maximum, and minimum settings. For more information, see
[Working with the retention period \
in SnapLock](../../../fsx/latest/ontapguide/snaplock-retention.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DefaultRetention" : RetentionPeriod,
  "MaximumRetention" : RetentionPeriod,
  "MinimumRetention" : RetentionPeriod
}

```

### YAML

```yaml

  DefaultRetention:
    RetentionPeriod
  MaximumRetention:
    RetentionPeriod
  MinimumRetention:
    RetentionPeriod

```

## Properties

`DefaultRetention`

The retention period assigned to a write once, read many (WORM) file by default if an explicit retention period is not set for an
FSx for ONTAP SnapLock volume. The default retention period must be greater than or equal to
the minimum retention period and less than or equal to the maximum retention period.

_Required_: Yes

_Type_: [RetentionPeriod](aws-properties-fsx-volume-retentionperiod.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaximumRetention`

The longest retention period that can be assigned to a WORM file on
an FSx for ONTAP SnapLock volume.

_Required_: Yes

_Type_: [RetentionPeriod](aws-properties-fsx-volume-retentionperiod.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinimumRetention`

The shortest retention period that can be assigned to a WORM file on an FSx for ONTAP SnapLock volume.

_Required_: Yes

_Type_: [RetentionPeriod](aws-properties-fsx-volume-retentionperiod.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SnaplockConfiguration

Tag

All content copied from https://docs.aws.amazon.com/.
