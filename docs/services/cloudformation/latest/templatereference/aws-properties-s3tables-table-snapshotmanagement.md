This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3Tables::Table SnapshotManagement

Contains details about the snapshot management settings for an Iceberg table. The oldest snapshot expires when its age exceeds the `maxSnapshotAgeHours` and the total number of snapshots exceeds the value for the minimum number of snapshots to keep `minSnapshotsToKeep`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaxSnapshotAgeHours" : Integer,
  "MinSnapshotsToKeep" : Integer,
  "Status" : String
}

```

### YAML

```yaml

  MaxSnapshotAgeHours: Integer
  MinSnapshotsToKeep: Integer
  Status: String

```

## Properties

`MaxSnapshotAgeHours`

The maximum age of a snapshot before it can be expired.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinSnapshotsToKeep`

The minimum number of snapshots to keep.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Status`

The status of the maintenance configuration.

_Required_: No

_Type_: String

_Allowed values_: `enabled | disabled`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SchemaV2Field

StorageClassConfiguration

All content copied from https://docs.aws.amazon.com/.
