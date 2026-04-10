This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Cluster ClusterSlurmConfig

The Slurm configuration for an instance group in a SageMaker HyperPod cluster.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "NodeType" : String,
  "PartitionNames" : [ String, ... ]
}

```

### YAML

```yaml

  NodeType: String
  PartitionNames:
    - String

```

## Properties

`NodeType`

The type of Slurm node for the instance group. Valid values are
`Controller`, `Worker`, and `Login`.

_Required_: Yes

_Type_: String

_Allowed values_: `Controller | Login | Compute`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PartitionNames`

The list of Slurm partition names that the instance group belongs to.

_Required_: No

_Type_: Array of String

_Minimum_: `0 | 0`

_Maximum_: `1024 | 1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ClusterRestrictedInstanceGroup

DeploymentConfig

All content copied from https://docs.aws.amazon.com/.
