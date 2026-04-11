This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Batch::JobDefinition EksHostPath

Specifies the configuration of a Kubernetes `hostPath` volume. A `hostPath`
volume mounts an existing file or directory from the host node's filesystem into your pod. For
more information, see [hostPath](https://kubernetes.io/docs/concepts/storage/volumes) in the _Kubernetes documentation_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Path" : String
}

```

### YAML

```yaml

  Path: String

```

## Properties

`Path`

The path of the file or directory on the host to mount into containers on the pod.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EksEmptyDir

EksMetadata

All content copied from https://docs.aws.amazon.com/.
