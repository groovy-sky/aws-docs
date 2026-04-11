This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTFleetWise::Campaign DataPartitionStorageOptions

Size, time, and location options for the data partition.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaximumSize" : StorageMaximumSize,
  "MinimumTimeToLive" : StorageMinimumTimeToLive,
  "StorageLocation" : String
}

```

### YAML

```yaml

  MaximumSize:
    StorageMaximumSize
  MinimumTimeToLive:
    StorageMinimumTimeToLive
  StorageLocation: String

```

## Properties

`MaximumSize`

The maximum storage size of the data stored in the data partition.

###### Note

Newer data overwrites older data when the partition reaches the maximum
size.

_Required_: Yes

_Type_: [StorageMaximumSize](aws-properties-iotfleetwise-campaign-storagemaximumsize.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MinimumTimeToLive`

The amount of time that data in this partition will be kept on disk.

- After the designated amount of time passes, the data can be removed, but it's
not guaranteed to be removed.

- Before the time expires, data in this partition can still be deleted if the
partition reaches its configured maximum size.

- Newer data will overwrite older data when the partition reaches the maximum
size.

_Required_: Yes

_Type_: [StorageMinimumTimeToLive](aws-properties-iotfleetwise-campaign-storageminimumtimetolive.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StorageLocation`

The folder name for the data partition under the campaign storage folder.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `4096`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataPartition

DataPartitionUploadOptions

All content copied from https://docs.aws.amazon.com/.
