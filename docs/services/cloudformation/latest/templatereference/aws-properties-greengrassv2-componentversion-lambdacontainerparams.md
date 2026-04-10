This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GreengrassV2::ComponentVersion LambdaContainerParams

Contains information about a container in which AWS Lambda functions run on
AWS IoT Greengrass core devices.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Devices" : [ LambdaDeviceMount, ... ],
  "MemorySizeInKB" : Integer,
  "MountROSysfs" : Boolean,
  "Volumes" : [ LambdaVolumeMount, ... ]
}

```

### YAML

```yaml

  Devices:
    - LambdaDeviceMount
  MemorySizeInKB: Integer
  MountROSysfs: Boolean
  Volumes:
    - LambdaVolumeMount

```

## Properties

`Devices`

The list of system devices that the container can access.

_Required_: No

_Type_: Array of [LambdaDeviceMount](aws-properties-greengrassv2-componentversion-lambdadevicemount.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MemorySizeInKB`

The memory size of the container, expressed in kilobytes.

Default: `16384` (16 MB)

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MountROSysfs`

Whether or not the container can read information from the device's `/sys`
folder.

Default: `false`

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Volumes`

The list of volumes that the container can access.

_Required_: No

_Type_: Array of [LambdaVolumeMount](aws-properties-greengrassv2-componentversion-lambdavolumemount.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ComponentPlatform

LambdaDeviceMount

All content copied from https://docs.aws.amazon.com/.
