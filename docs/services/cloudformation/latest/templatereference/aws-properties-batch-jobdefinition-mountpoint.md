This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Batch::JobDefinition MountPoint

Details for a Docker volume mount point that's used in a job's container properties. This
parameter maps to `Volumes` in the [Create a\
container](https://docs.docker.com/engine/api/v1.43) section of the _Docker Remote API_ and the
`--volume` option to docker run.

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

The path on the container where the host volume is mounted.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReadOnly`

If this value is `true`, the container has read-only access to the volume.
Otherwise, the container can write to the volume. The default value is `false`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceVolume`

The name of the volume to mount.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LogConfiguration

MultiNodeContainerProperties

All content copied from https://docs.aws.amazon.com/.
