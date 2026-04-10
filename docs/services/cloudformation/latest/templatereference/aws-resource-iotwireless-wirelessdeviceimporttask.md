This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTWireless::WirelessDeviceImportTask

Information about an import task for wireless devices. When creating the resource,
either create a single wireless device import task using the Sidewalk manufacturing serial
number (SMSN) of the wireless device, or create an import task for multiple devices by
specifying both the `DeviceCreationFile` and the `Role`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IoTWireless::WirelessDeviceImportTask",
  "Properties" : {
      "DestinationName" : String,
      "Sidewalk" : Sidewalk,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::IoTWireless::WirelessDeviceImportTask
Properties:
  DestinationName: String
  Sidewalk:
    Sidewalk
  Tags:
    - Tag

```

## Properties

`DestinationName`

The name of the destination that describes the IoT rule to route messages from the
Sidewalk devices in the import task to other applications.

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z0-9-_]+`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Sidewalk`

The Sidewalk-related information of the wireless device import task.

_Required_: Yes

_Type_: [Sidewalk](aws-properties-iotwireless-wirelessdeviceimporttask-sidewalk.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Adds to or modifies the tags of the given resource. Tags are metadata that you can use
to manage a resource.

_Required_: No

_Type_: Array of [Tag](aws-properties-iotwireless-wirelessdeviceimporttask-tag.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the wireless device import task.

### Fn::GetAtt

`Arn`

The ARN (Amazon Resource Name) of the import task.

`CreationDate`

The date and time at which the wireless device import task was created.

`FailedImportedDevicesCount`

The summary information of count of wireless devices that failed to onboard to the
import task.

`Id`

The import task ID.

`InitializedImportedDevicesCount`

The summary information of count of wireless devices that are waiting for the control
log to be added to an import task.

`OnboardedImportedDevicesCount`

The summary information of count of wireless devices that have been onboarded to the
import task.

`PendingImportedDevicesCount`

The summary information of count of wireless devices that are waiting in the queue to be
onboarded to the import task.

`Sidewalk.DeviceCreationFileList`

List of Sidewalk devices that are added to the import task.

`Status`

The status of a wireless device import task. The status can be
`INITIALIZING`, `INITIALIZED`, `PENDING`,
`COMPLETE`, `FAILED`, or `DELETING`.

`StatusReason`

The reason that provides additional information about the import task status.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Sidewalk

All content copied from https://docs.aws.amazon.com/.
