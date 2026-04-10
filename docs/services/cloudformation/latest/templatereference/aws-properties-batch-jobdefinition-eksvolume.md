This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Batch::JobDefinition EksVolume

Specifies an Amazon EKS volume for a job definition.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EmptyDir" : EksEmptyDir,
  "HostPath" : EksHostPath,
  "Name" : String,
  "PersistentVolumeClaim" : EksPersistentVolumeClaim,
  "Secret" : EksSecret
}

```

### YAML

```yaml

  EmptyDir:
    EksEmptyDir
  HostPath:
    EksHostPath
  Name: String
  PersistentVolumeClaim:
    EksPersistentVolumeClaim
  Secret:
    EksSecret

```

## Properties

`EmptyDir`

Specifies the configuration of a Kubernetes `emptyDir` volume. For more information,
see [emptyDir](https://kubernetes.io/docs/concepts/storage/volumes)
in the _Kubernetes documentation_.

_Required_: No

_Type_: [EksEmptyDir](aws-properties-batch-jobdefinition-eksemptydir.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HostPath`

Specifies the configuration of a Kubernetes `hostPath` volume. For more information,
see [hostPath](https://kubernetes.io/docs/concepts/storage/volumes)
in the _Kubernetes documentation_.

_Required_: No

_Type_: [EksHostPath](aws-properties-batch-jobdefinition-ekshostpath.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the volume. The name must be allowed as a DNS subdomain name. For more
information, see [DNS subdomain names](https://kubernetes.io/docs/concepts/overview/working-with-objects/names) in the _Kubernetes documentation_.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PersistentVolumeClaim`

Specifies the configuration of a Kubernetes `persistentVolumeClaim` bounded to a
`persistentVolume`. For more information, see [Persistent Volume Claims](https://kubernetes.io/docs/concepts/storage/persistent-volumes) in the _Kubernetes documentation_.

_Required_: No

_Type_: [EksPersistentVolumeClaim](aws-properties-batch-jobdefinition-ekspersistentvolumeclaim.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Secret`

Specifies the configuration of a Kubernetes `secret` volume. For more information, see
[secret](https://kubernetes.io/docs/concepts/storage/volumes) in the
_Kubernetes documentation_.

_Required_: No

_Type_: [EksSecret](aws-properties-batch-jobdefinition-ekssecret.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EksSecret

Environment

All content copied from https://docs.aws.amazon.com/.
