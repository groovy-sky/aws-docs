This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FSx::Volume SnaplockConfiguration

Specifies the SnapLock configuration for an FSx for ONTAP SnapLock volume.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AuditLogVolume" : String,
  "AutocommitPeriod" : AutocommitPeriod,
  "PrivilegedDelete" : String,
  "RetentionPeriod" : SnaplockRetentionPeriod,
  "SnaplockType" : String,
  "VolumeAppendModeEnabled" : String
}

```

### YAML

```yaml

  AuditLogVolume: String
  AutocommitPeriod:
    AutocommitPeriod
  PrivilegedDelete: String
  RetentionPeriod:
    SnaplockRetentionPeriod
  SnaplockType: String
  VolumeAppendModeEnabled: String

```

## Properties

`AuditLogVolume`

Enables or disables the audit log volume for an FSx for ONTAP SnapLock volume. The default
value is `false`. If you set `AuditLogVolume` to `true`, the SnapLock volume is
created as an audit log volume. The minimum retention period for an audit log volume is six months.

For more information, see
[SnapLock audit log volumes](../../../fsx/latest/ontapguide/how-snaplock-works.md#snaplock-audit-log-volume).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AutocommitPeriod`

The configuration object for setting the autocommit period of files in an FSx for ONTAP SnapLock volume.

_Required_: No

_Type_: [AutocommitPeriod](aws-properties-fsx-volume-autocommitperiod.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrivilegedDelete`

Enables, disables, or permanently disables privileged delete on an FSx for ONTAP SnapLock
Enterprise volume. Enabling privileged delete allows SnapLock administrators to delete write once, read
many (WORM) files even
if they have active retention periods. `PERMANENTLY_DISABLED` is a terminal state.
If privileged delete is permanently disabled on a SnapLock volume, you can't re-enable it. The default
value is `DISABLED`.

For more information, see [Privileged delete](../../../fsx/latest/ontapguide/snaplock-enterprise.md#privileged-delete).

_Required_: No

_Type_: String

_Allowed values_: `DISABLED | ENABLED | PERMANENTLY_DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RetentionPeriod`

Specifies the retention period of an FSx for ONTAP SnapLock volume.

_Required_: No

_Type_: [SnaplockRetentionPeriod](aws-properties-fsx-volume-snaplockretentionperiod.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SnaplockType`

Specifies the retention mode of an FSx for ONTAP SnapLock volume. After it is set, it can't be changed.
You can choose one of the following retention modes:

- `COMPLIANCE`: Files transitioned to write once, read many (WORM) on a Compliance volume can't be deleted
until their retention periods expire. This retention mode is used to address government or industry-specific mandates or to protect
against ransomware attacks. For more information,
see [SnapLock Compliance](../../../fsx/latest/ontapguide/snaplock-compliance.md).

- `ENTERPRISE`: Files transitioned to WORM on an Enterprise volume can be deleted by authorized users
before their retention periods expire using privileged delete. This retention mode is used to advance an organization's data integrity
and internal compliance or to test retention settings before using SnapLock Compliance. For more information, see
[SnapLock Enterprise](../../../fsx/latest/ontapguide/snaplock-enterprise.md).

_Required_: Yes

_Type_: String

_Allowed values_: `COMPLIANCE | ENTERPRISE`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VolumeAppendModeEnabled`

Enables or disables volume-append mode
on an FSx for ONTAP SnapLock volume. Volume-append mode allows you to
create WORM-appendable files and write data to them incrementally.
The default value is `false`.

For more information, see [Volume-append mode](../../../fsx/latest/ontapguide/worm-state.md#worm-state-append).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RetentionPeriod

SnaplockRetentionPeriod

All content copied from https://docs.aws.amazon.com/.
