This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Batch::JobDefinition EksEmptyDir

Specifies the configuration of a Kubernetes `emptyDir` volume. An
`emptyDir` volume is first created when a pod is assigned to a node. It exists as
long as that pod is running on that node. The `emptyDir` volume is initially empty.
All containers in the pod can read and write the files in the `emptyDir` volume.
However, the `emptyDir` volume can be mounted at the same or different paths in each
container. When a pod is removed from a node for any reason, the data in the
`emptyDir` is deleted permanently. For more information, see [emptyDir](https://kubernetes.io/docs/concepts/storage/volumes) in the
_Kubernetes documentation_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Medium" : String,
  "SizeLimit" : String
}

```

### YAML

```yaml

  Medium: String
  SizeLimit: String

```

## Properties

`Medium`

The medium to store the volume. The default value is an empty string, which uses the storage
of the node.

""

**(Default)** Use the disk storage of the node.

"Memory"

Use the `tmpfs` volume that's backed by the RAM of the node. Contents of the
volume are lost when the node reboots, and any storage on the volume counts against the
container's memory limit.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SizeLimit`

The maximum size of the volume. By default, there's no maximum size defined.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EksContainerVolumeMount

EksHostPath

All content copied from https://docs.aws.amazon.com/.
