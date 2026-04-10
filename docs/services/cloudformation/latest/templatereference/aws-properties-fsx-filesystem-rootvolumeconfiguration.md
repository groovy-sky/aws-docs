This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FSx::FileSystem RootVolumeConfiguration

The configuration of an Amazon FSx for OpenZFS root volume.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CopyTagsToSnapshots" : Boolean,
  "DataCompressionType" : String,
  "NfsExports" : [ NfsExports, ... ],
  "ReadOnly" : Boolean,
  "RecordSizeKiB" : Integer,
  "UserAndGroupQuotas" : [ UserAndGroupQuotas, ... ]
}

```

### YAML

```yaml

  CopyTagsToSnapshots: Boolean
  DataCompressionType: String
  NfsExports:
    - NfsExports
  ReadOnly: Boolean
  RecordSizeKiB: Integer
  UserAndGroupQuotas:
    - UserAndGroupQuotas

```

## Properties

`CopyTagsToSnapshots`

A Boolean value indicating whether tags for the volume should be copied to snapshots
of the volume. This value defaults to `false`. If it's set to `true`,
all tags for the volume are copied to snapshots where the user doesn't specify tags. If this
value is `true` and you specify one or more tags, only the specified tags are
copied to snapshots. If you specify one or more tags when creating the snapshot, no tags
are copied from the volume, regardless of this value.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

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

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NfsExports`

The configuration object for mounting a file system.

_Required_: No

_Type_: [Array](aws-properties-fsx-filesystem-nfsexports.md) of [NfsExports](aws-properties-fsx-filesystem-nfsexports.md)

_Maximum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ReadOnly`

A Boolean value indicating whether the volume is read-only. Setting this value to
`true` can be useful after you have completed changes to a volume and no
longer want changes to occur.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RecordSizeKiB`

Specifies the record size of an OpenZFS root volume, in kibibytes (KiB). Valid values are 4, 8,
16, 32, 64, 128, 256, 512, or 1024 KiB. The default is 128 KiB. Most workloads should use the
default record size. Database workflows can benefit from a smaller record size, while streaming
workflows can benefit from a larger record size. For additional guidance on setting a custom record
size, see [Tips for maximizing performance](../../../fsx/latest/openzfsguide/performance.md#performance-tips-zfs) in the
_Amazon FSx for OpenZFS User Guide_.

_Required_: No

_Type_: Integer

_Minimum_: `4`

_Maximum_: `4096`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`UserAndGroupQuotas`

An object specifying how much storage users or groups can use on the volume.

_Required_: No

_Type_: [Array](aws-properties-fsx-filesystem-userandgroupquotas.md) of [UserAndGroupQuotas](aws-properties-fsx-filesystem-userandgroupquotas.md)

_Maximum_: `500`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ReadCacheConfiguration

SelfManagedActiveDirectoryConfiguration

All content copied from https://docs.aws.amazon.com/.
