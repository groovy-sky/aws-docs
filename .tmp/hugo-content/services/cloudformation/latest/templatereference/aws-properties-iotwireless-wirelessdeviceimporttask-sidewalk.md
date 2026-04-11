This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTWireless::WirelessDeviceImportTask Sidewalk

Sidewalk-related information about a wireless device import task.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DeviceCreationFile" : String,
  "DeviceCreationFileList" : [ String, ... ],
  "Role" : String,
  "SidewalkManufacturingSn" : String
}

```

### YAML

```yaml

  DeviceCreationFile: String
  DeviceCreationFileList:
    - String
  Role: String
  SidewalkManufacturingSn: String

```

## Properties

`DeviceCreationFile`

The CSV file contained in an S3 bucket that's used for adding devices to an import
task.

_Required_: No

_Type_: String

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeviceCreationFileList`

List of Sidewalk devices that are added to the import task.

_Required_: No

_Type_: Array of String

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Role`

The IAM role that allows AWS IoT Wireless to access the CSV file in the
S3 bucket.

_Required_: No

_Type_: String

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SidewalkManufacturingSn`

The Sidewalk manufacturing serial number (SMSN) of the Sidewalk device.

_Required_: No

_Type_: String

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::IoTWireless::WirelessDeviceImportTask

Tag

All content copied from https://docs.aws.amazon.com/.
