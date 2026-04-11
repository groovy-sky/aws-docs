This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTFleetWise::Campaign DataPartition

The configuration for signal data storage and upload options. You can only specify these
options when the campaign's spooling mode is `TO_DISK`.

###### Important

AWS IoT FleetWise will no longer be open to new customers starting April 30, 2026.
If you would like to use AWS IoT FleetWise, sign up prior to that date. Existing customers can
continue to use the service as normal. For more information, see [AWS IoT FleetWise availability change](../../../iot-fleetwise/latest/developerguide/iotfleetwise-availability-change.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Id" : String,
  "StorageOptions" : DataPartitionStorageOptions,
  "UploadOptions" : DataPartitionUploadOptions
}

```

### YAML

```yaml

  Id: String
  StorageOptions:
    DataPartitionStorageOptions
  UploadOptions:
    DataPartitionUploadOptions

```

## Properties

`Id`

The ID of the data partition. The data partition ID must be unique within a campaign.
You can establish a data partition as the default partition for a campaign by using
`default` as the ID.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9]+$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StorageOptions`

The storage options for a data partition.

_Required_: Yes

_Type_: [DataPartitionStorageOptions](aws-properties-iotfleetwise-campaign-datapartitionstorageoptions.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`UploadOptions`

The upload options for the data partition.

_Required_: No

_Type_: [DataPartitionUploadOptions](aws-properties-iotfleetwise-campaign-datapartitionuploadoptions.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataDestinationConfig

DataPartitionStorageOptions

All content copied from https://docs.aws.amazon.com/.
