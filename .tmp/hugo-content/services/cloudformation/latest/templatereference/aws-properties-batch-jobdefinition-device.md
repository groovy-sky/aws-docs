This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Batch::JobDefinition Device

An object that represents a container instance host device.

###### Note

This object isn't applicable to jobs that are running on Fargate resources and shouldn't
be provided.

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

The path inside the container that's used to expose the host device. By default, the
`hostPath` value is used.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HostPath`

The path for the device on the host container instance.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Permissions`

The explicit permissions to provide to the container for the device. By default, the
container has permissions for `read`, `write`, and `mknod` for
the device.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ContainerProperties

EcsProperties

All content copied from https://docs.aws.amazon.com/.
