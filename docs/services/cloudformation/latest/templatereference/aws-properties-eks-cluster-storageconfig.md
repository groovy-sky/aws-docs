This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EKS::Cluster StorageConfig

Request to update the configuration of the storage capability of your EKS Auto Mode
cluster. For example, enable the capability. For more information, see EKS Auto Mode
block storage capability in the _Amazon EKS User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BlockStorage" : BlockStorage
}

```

### YAML

```yaml

  BlockStorage:
    BlockStorage

```

## Properties

`BlockStorage`

Request to configure EBS Block Storage settings for your EKS Auto Mode cluster.

_Required_: No

_Type_: [BlockStorage](aws-properties-eks-cluster-blockstorage.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ResourcesVpcConfig

Tag

All content copied from https://docs.aws.amazon.com/.
