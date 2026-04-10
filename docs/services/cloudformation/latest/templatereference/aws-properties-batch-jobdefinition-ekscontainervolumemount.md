This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Batch::JobDefinition EksContainerVolumeMount

The volume mounts for a container for an Amazon EKS job. For more information about volumes and
volume mounts in Kubernetes, see [Volumes](https://kubernetes.io/docs/concepts/storage/volumes) in the _Kubernetes documentation_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MountPath" : String,
  "Name" : String,
  "ReadOnly" : Boolean,
  "SubPath" : String
}

```

### YAML

```yaml

  MountPath: String
  Name: String
  ReadOnly: Boolean
  SubPath: String

```

## Properties

`MountPath`

The path on the container where the volume is mounted.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name the volume mount. This must match the name of one of the volumes in the
pod.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReadOnly`

If this value is `true`, the container has read-only access to the volume.
Otherwise, the container can write to the volume. The default value is `false`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubPath`

A sub-path inside the referenced volume instead of its root.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EksContainerSecurityContext

EksEmptyDir

All content copied from https://docs.aws.amazon.com/.
