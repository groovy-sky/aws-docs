This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FSx::Volume OntapConfiguration

Specifies the configuration of the ONTAP volume that you are creating.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AggregateConfiguration" : AggregateConfiguration,
  "CopyTagsToBackups" : String,
  "JunctionPath" : String,
  "OntapVolumeType" : String,
  "SecurityStyle" : String,
  "SizeInBytes" : String,
  "SizeInMegabytes" : String,
  "SnaplockConfiguration" : SnaplockConfiguration,
  "SnapshotPolicy" : String,
  "StorageEfficiencyEnabled" : String,
  "StorageVirtualMachineId" : String,
  "TieringPolicy" : TieringPolicy,
  "VolumeStyle" : String
}

```

### YAML

```yaml

  AggregateConfiguration:
    AggregateConfiguration
  CopyTagsToBackups: String
  JunctionPath: String
  OntapVolumeType: String
  SecurityStyle: String
  SizeInBytes: String
  SizeInMegabytes: String
  SnaplockConfiguration:
    SnaplockConfiguration
  SnapshotPolicy: String
  StorageEfficiencyEnabled: String
  StorageVirtualMachineId: String
  TieringPolicy:
    TieringPolicy
  VolumeStyle: String

```

## Properties

`AggregateConfiguration`

Used to specify the configuration options for an FSx for ONTAP volume's storage aggregate or aggregates.

_Required_: No

_Type_: [AggregateConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-fsx-volume-aggregateconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CopyTagsToBackups`

A boolean flag indicating whether tags for the volume should be copied to backups. This value defaults to
false. If it's set to true, all tags for the volume are copied to all automatic and user-initiated backups
where the user doesn't specify tags. If this value is true, and you specify one or more tags, only the
specified tags are copied to backups. If you specify one or more tags when creating a user-initiated
backup, no tags are copied from the volume, regardless of this value.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`JunctionPath`

Specifies the location in the SVM's namespace where the volume is mounted.
This parameter is required. The `JunctionPath` must have a leading
forward slash, such as `/vol3`.

_Required_: No

_Type_: String

_Pattern_: `^[^\u0000\u0085\u2028\u2029\r\n]{1,255}$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OntapVolumeType`

Specifies the type of volume you are creating. Valid values are the following:

- `RW` specifies a read/write volume. `RW` is the default.

- `DP` specifies a data-protection volume. A `DP` volume
is read-only and can be used as the destination of a NetApp SnapMirror relationship.

For more information, see [Volume types](https://docs.aws.amazon.com/fsx/latest/ONTAPGuide/managing-volumes.html#volume-types)
in the Amazon FSx for NetApp ONTAP User Guide.

_Required_: No

_Type_: String

_Allowed values_: `RW | DP`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SecurityStyle`

Specifies the security style for the volume. If a volume's security style is not specified,
it is automatically set to the root volume's security style. The security style determines the type of permissions
that FSx for ONTAP uses to control data access. Specify one of the following values:

- `UNIX` if the file system is managed by a UNIX
administrator, the majority of users are NFS clients, and an application
accessing the data uses a UNIX user as the service account.

- `NTFS` if the file system is managed by a Windows
administrator, the majority of users are SMB clients, and an application
accessing the data uses a Windows user as the service account.

- `MIXED` This is an advanced setting. For more information, see the topic
[What the security styles and their effects are](https://docs.netapp.com/us-en/ontap/nfs-admin/security-styles-their-effects-concept.html)
in the NetApp Documentation Center.

For more information, see [Volume security style](https://docs.aws.amazon.com/fsx/latest/ONTAPGuide/managing-volumes.html#volume-security-style) in the
FSx for ONTAP User Guide.

_Required_: No

_Type_: String

_Allowed values_: `UNIX | NTFS | MIXED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SizeInBytes`

Specifies the configured size of the volume, in bytes.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SizeInMegabytes`

Use `SizeInBytes` instead. Specifies the size of the volume, in megabytes (MB), that you are creating.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `2147483647`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SnaplockConfiguration`

The SnapLock configuration object for an FSx for ONTAP SnapLock volume.

_Required_: No

_Type_: [SnaplockConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-fsx-volume-snaplockconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SnapshotPolicy`

Specifies the snapshot policy for the volume. There are three built-in snapshot policies:

- `default`: This is the default policy. A maximum of six hourly snapshots taken five minutes past
the hour. A maximum of two daily snapshots taken Monday through Saturday at 10 minutes after
midnight. A maximum of two weekly snapshots taken every Sunday at 15 minutes after midnight.

- `default-1weekly`: This policy is the same as the `default` policy except
that it only retains one snapshot from the weekly schedule.

- `none`: This policy does not take any snapshots. This policy can be assigned to volumes to
prevent automatic snapshots from being taken.

You can also provide the name of a custom policy that you created with the ONTAP CLI or REST API.

For more information, see [Snapshot policies](https://docs.aws.amazon.com/fsx/latest/ONTAPGuide/snapshots-ontap.html#snapshot-policies)
in the Amazon FSx for NetApp ONTAP User Guide.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StorageEfficiencyEnabled`

Set to true to enable deduplication, compression, and compaction storage
efficiency features on the volume, or set to false to disable them.

`StorageEfficiencyEnabled` is required when creating a `RW` volume ( `OntapVolumeType` set to `RW`).

_Required_: Conditional

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StorageVirtualMachineId`

Specifies the ONTAP SVM in which to create the volume.

_Required_: Yes

_Type_: String

_Pattern_: `^(svm-[0-9a-f]{17,})$`

_Minimum_: `21`

_Maximum_: `21`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TieringPolicy`

Describes the data tiering policy for an ONTAP volume. When enabled, Amazon FSx for ONTAP's intelligent
tiering automatically transitions a volume's data between the file system's primary storage and capacity
pool storage based on your access patterns.

Valid tiering policies are the following:

- `SNAPSHOT_ONLY` \- (Default value) moves cold snapshots to the capacity pool storage tier.

- `AUTO` \- moves cold user data and snapshots to the capacity pool storage tier based on your access patterns.

- `ALL` \- moves all user data blocks in both the active file system and Snapshot copies to the storage pool tier.

- `NONE` \- keeps a volume's data in the primary storage tier, preventing it from being moved to the capacity pool tier.

_Required_: No

_Type_: [TieringPolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-fsx-volume-tieringpolicy.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VolumeStyle`

Use to specify the style of an ONTAP volume. FSx for ONTAP offers two styles of volumes that you can use for different purposes,
FlexVol and FlexGroup volumes. For more information, see
[Volume styles](https://docs.aws.amazon.com/fsx/latest/ONTAPGuide/managing-volumes.html#volume-styles) in the Amazon FSx for NetApp ONTAP User Guide.

_Required_: No

_Type_: String

_Allowed values_: `FLEXVOL | FLEXGROUP`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

NfsExports

OpenZFSConfiguration
