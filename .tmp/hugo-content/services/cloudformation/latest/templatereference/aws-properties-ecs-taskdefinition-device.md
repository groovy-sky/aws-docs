This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::TaskDefinition Device

The `Device` property specifies an object representing a container instance
host device.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ContainerPath" : String,
  "HostPath" : String,
  "Permissions" : [ String, ... ]
}

```

### YAML

```yaml

  ContainerPath: String
  HostPath: String
  Permissions:
    - String

```

## Properties

`ContainerPath`

The path inside the container at which to expose the host device.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`HostPath`

The path for the device on the host container instance.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Permissions`

The explicit permissions to provide to the container for the device. By default, the
container has permissions for `read`, `write`, and
`mknod` for the device.

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ContainerDependency

DockerVolumeConfiguration

All content copied from https://docs.aws.amazon.com/.
