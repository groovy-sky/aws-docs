This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Batch::JobDefinition EksPersistentVolumeClaim

A `persistentVolumeClaim` volume is used to mount a [PersistentVolume](https://kubernetes.io/docs/concepts/storage/persistent-volumes)
into a Pod. PersistentVolumeClaims are a way for users to "claim" durable storage without knowing
the details of the particular cloud environment. See the information about [PersistentVolumes](https://kubernetes.io/docs/concepts/storage/persistent-volumes)
in the _Kubernetes documentation_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ClaimName" : String,
  "ReadOnly" : Boolean
}

```

### YAML

```yaml

  ClaimName: String
  ReadOnly: Boolean

```

## Properties

`ClaimName`

The name of the `persistentVolumeClaim` bounded to a `persistentVolume`.
For more information, see [Persistent Volume Claims](https://kubernetes.io/docs/concepts/storage/persistent-volumes) in the _Kubernetes documentation_.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReadOnly`

An optional boolean value indicating if the mount is read only. Default is false. For more
information, see [Read Only Mounts](https://kubernetes.io/docs/concepts/storage/volumes) in the _Kubernetes documentation_.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EksMetadata

EksPodProperties

All content copied from https://docs.aws.amazon.com/.
