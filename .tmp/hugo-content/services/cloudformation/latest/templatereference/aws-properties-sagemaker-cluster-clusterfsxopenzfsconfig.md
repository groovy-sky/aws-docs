This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Cluster ClusterFsxOpenZfsConfig

Defines the configuration for attaching an Amazon FSx for OpenZFS file system to
instances in a SageMaker HyperPod cluster instance group.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DnsName" : String,
  "MountPath" : String
}

```

### YAML

```yaml

  DnsName: String
  MountPath: String

```

## Properties

`DnsName`

The DNS name of the Amazon FSx for OpenZFS file system.

_Required_: Yes

_Type_: String

_Pattern_: `^((fs|fc)i?-[0-9a-f]{8,}\..{4,253})$`

_Minimum_: `16`

_Maximum_: `275`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MountPath`

The local path where the Amazon FSx for OpenZFS file system is mounted on
instances.

_Required_: No

_Type_: String

_Pattern_: `^/(?!/)(?!.*/$)[a-zA-Z0-9._-]+(/[a-zA-Z0-9._-]+)*$`

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ClusterFsxLustreConfig

ClusterInstanceGroup

All content copied from https://docs.aws.amazon.com/.
