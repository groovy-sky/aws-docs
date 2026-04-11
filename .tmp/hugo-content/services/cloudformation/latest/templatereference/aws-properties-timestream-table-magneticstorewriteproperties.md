This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Timestream::Table MagneticStoreWriteProperties

The set of properties on a table for configuring magnetic store writes.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EnableMagneticStoreWrites" : Boolean,
  "MagneticStoreRejectedDataLocation" : MagneticStoreRejectedDataLocation
}

```

### YAML

```yaml

  EnableMagneticStoreWrites: Boolean
  MagneticStoreRejectedDataLocation:
    MagneticStoreRejectedDataLocation

```

## Properties

`EnableMagneticStoreWrites`

A flag to enable magnetic store writes.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MagneticStoreRejectedDataLocation`

The location to write error reports for records rejected asynchronously during magnetic store writes.

_Required_: No

_Type_: [MagneticStoreRejectedDataLocation](aws-properties-timestream-table-magneticstorerejecteddatalocation.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MagneticStoreRejectedDataLocation

PartitionKey

All content copied from https://docs.aws.amazon.com/.
