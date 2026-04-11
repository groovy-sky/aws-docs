This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FSx::Volume OriginSnapshot

The configuration object that specifies the snapshot to use as the origin of the data
for the volume.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CopyStrategy" : String,
  "SnapshotARN" : String
}

```

### YAML

```yaml

  CopyStrategy: String
  SnapshotARN: String

```

## Properties

`CopyStrategy`

Specifies the strategy used when copying data from the snapshot to the new volume.

- `CLONE` \- The new volume references the data in the origin
snapshot. Cloning a snapshot is faster than copying data from the snapshot to a
new volume and doesn't consume disk throughput. However, the origin snapshot
can't be deleted if there is a volume using its copied data.

- `FULL_COPY` \- Copies all data from the snapshot to the new
volume.

Specify this option to create the volume from a snapshot on another FSx for OpenZFS file system.

###### Note

The `INCREMENTAL_COPY` option is only for updating an existing volume
by using a snapshot from another FSx for OpenZFS file system. For more
information, see [CopySnapshotAndUpdateVolume](../../../../reference/fsx/latest/apireference/api-copysnapshotandupdatevolume.md).

_Required_: Yes

_Type_: String

_Allowed values_: `CLONE | FULL_COPY | INCREMENTAL_COPY`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SnapshotARN`

Specifies the snapshot to use when creating an OpenZFS volume from a snapshot.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OpenZFSConfiguration

RetentionPeriod

All content copied from https://docs.aws.amazon.com/.
