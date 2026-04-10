This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::TaskDefinition MountPoint

The details for a volume mount point that's used in a container definition.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ContainerPath" : String,
  "ReadOnly" : Boolean,
  "SourceVolume" : String
}

```

### YAML

```yaml

  ContainerPath: String
  ReadOnly: Boolean
  SourceVolume: String

```

## Properties

`ContainerPath`

The path on the container to mount the host volume at.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ReadOnly`

If this value is `true`, the container has read-only access to the volume.
If this value is `false`, then the container can write to the volume. The
default value is `false`.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SourceVolume`

The name of the volume to mount. Must be a volume name referenced in the
`name` parameter of task definition `volume`.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LogConfiguration

PortMapping

All content copied from https://docs.aws.amazon.com/.
