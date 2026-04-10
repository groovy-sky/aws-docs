This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Cluster ClusterOrchestratorSlurmConfig

The configuration settings for the Slurm orchestrator used with the SageMaker HyperPod
cluster.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SlurmConfigStrategy" : String
}

```

### YAML

```yaml

  SlurmConfigStrategy: String

```

## Properties

`SlurmConfigStrategy`

The strategy for managing partitions for the Slurm configuration. Valid values are
`Managed`, `Overwrite`, and `Merge`.

_Required_: No

_Type_: String

_Allowed values_: `Overwrite | Managed | Merge`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ClusterOrchestratorEksConfig

ClusterRestrictedInstanceGroup

All content copied from https://docs.aws.amazon.com/.
