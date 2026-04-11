This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FSx::Volume OpenZFSConfiguration

Specifies the configuration of the Amazon FSx for OpenZFS volume that you are creating.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CopyTagsToSnapshots" : Boolean,
  "DataCompressionType" : String,
  "NfsExports" : [ NfsExports, ... ],
  "Options" : [ String, ... ],
  "OriginSnapshot" : OriginSnapshot,
  "ParentVolumeId" : String,
  "ReadOnly" : Boolean,
  "RecordSizeKiB" : Integer,
  "StorageCapacityQuotaGiB" : Integer,
  "StorageCapacityReservationGiB" : Integer,
  "UserAndGroupQuotas" : [ UserAndGroupQuotas, ... ]
}

```

### YAML

```yaml

  CopyTagsToSnapshots: Boolean
  DataCompressionType: String
  NfsExports:
    - NfsExports
  Options:
    - String
  OriginSnapshot:
    OriginSnapshot
  ParentVolumeId: String
  ReadOnly: Boolean
  RecordSizeKiB: Integer
  StorageCapacityQuotaGiB: Integer
  StorageCapacityReservationGiB: Integer
  UserAndGroupQuotas:
    - UserAndGroupQuotas

```

## Properties

`CopyTagsToSnapshots`

A Boolean value indicating whether tags for the volume should be copied to snapshots.
This value defaults to `false`. If this value is set to `true`, and you do not specify any tags, all tags for the original volume are copied over to snapshots.
If this value is set to `true`, and you do specify one or more tags, only the specified tags for the original volume are copied over to snapshots. If you specify one or more tags when creating a new snapshot, no tags are copied over from the original volume, regardless of this value.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataCompressionType`

Specifies the method used to compress the data on the volume. The compression
type is `NONE` by default.

- `NONE` \- Doesn't compress the data on the volume.
`NONE` is the default.

- `ZSTD` \- Compresses the data in the volume using the Zstandard
(ZSTD) compression algorithm. Compared to LZ4, Z-Standard provides a better
compression ratio to minimize on-disk storage utilization.

- `LZ4` \- Compresses the data in the volume using the LZ4
compression algorithm. Compared to Z-Standard, LZ4 is less compute-intensive
and delivers higher write throughput speeds.

_Required_: No

_Type_: String

_Allowed values_: `NONE | ZSTD | LZ4`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NfsExports`

The configuration object for mounting a Network File System (NFS) file system.

_Required_: No

_Type_: [Array](aws-properties-fsx-volume-nfsexports.md) of [NfsExports](aws-properties-fsx-volume-nfsexports.md)

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Options`

To delete the volume's child volumes, snapshots, and clones, use the string
`DELETE_CHILD_VOLUMES_AND_SNAPSHOTS`.

_Required_: No

_Type_: Array of String

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OriginSnapshot`

The configuration object that specifies the snapshot to use as the origin of the data
for the volume.

_Required_: No

_Type_: [OriginSnapshot](aws-properties-fsx-volume-originsnapshot.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ParentVolumeId`

The ID of the volume to use as the parent volume of the volume that you are creating.

_Required_: Yes

_Type_: String

_Pattern_: `^(fsvol-[0-9a-f]{17,})$`

_Minimum_: `23`

_Maximum_: `23`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ReadOnly`

A Boolean value indicating whether the volume is read-only.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RecordSizeKiB`

Specifies the suggested block size for a volume in a ZFS dataset, in kibibytes (KiB).
For file systems using the Intelligent-Tiering storage class, valid values are 128, 256, 512, 1024, 2048, or 4096 KiB, with a default of 1024 KiB.
For all other file systems, valid values are 4, 8, 16, 32, 64, 128, 256, 512, or 1024 KiB, with a default of 128 KiB.
We recommend using the default setting for the majority of use cases. Generally, workloads that write in fixed small or large record sizes may benefit from setting a custom record size, like database workloads (small record size) or media streaming workloads (large record size).
For additional guidance on when to set a custom record size, see
[ZFS Record size](../../../fsx/latest/openzfsguide/performance.md#record-size-performance) in the _Amazon FSx for OpenZFS User Guide_.

_Required_: No

_Type_: Integer

_Minimum_: `4`

_Maximum_: `4096`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StorageCapacityQuotaGiB`

Sets the maximum storage size in gibibytes (GiB) for the volume. You can specify
a quota that is larger than the storage on the parent volume. A volume quota limits
the amount of storage that the volume can consume to the configured amount, but does not
guarantee the space will be available on the parent volume. To guarantee quota space, you must also set
`StorageCapacityReservationGiB`. To _not_ specify a storage capacity quota, set this to `-1`.

For more information, see
[Volume properties](../../../fsx/latest/openzfsguide/managing-volumes.md#volume-properties)
in the _Amazon FSx for OpenZFS User Guide_.

_Required_: No

_Type_: Integer

_Minimum_: `-1`

_Maximum_: `2147483647`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StorageCapacityReservationGiB`

Specifies the amount of storage in gibibytes (GiB) to reserve from the parent volume. Setting
`StorageCapacityReservationGiB` guarantees that the specified amount of storage space
on the parent volume will always be available for the volume.
You can't reserve more storage than the parent volume has. To _not_ specify a storage capacity
reservation, set this to `0` or `-1`. For more information, see
[Volume properties](../../../fsx/latest/openzfsguide/managing-volumes.md#volume-properties)
in the _Amazon FSx for OpenZFS User Guide_.

_Required_: No

_Type_: Integer

_Minimum_: `-1`

_Maximum_: `2147483647`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserAndGroupQuotas`

Configures how much storage users and groups can use on the volume.

_Required_: No

_Type_: [Array](aws-properties-fsx-volume-userandgroupquotas.md) of [UserAndGroupQuotas](aws-properties-fsx-volume-userandgroupquotas.md)

_Maximum_: `500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OntapConfiguration

OriginSnapshot

All content copied from https://docs.aws.amazon.com/.
