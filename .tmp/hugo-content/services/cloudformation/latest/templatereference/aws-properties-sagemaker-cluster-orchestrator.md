This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Cluster Orchestrator

The orchestrator for a SageMaker HyperPod cluster.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Eks" : ClusterOrchestratorEksConfig,
  "Slurm" : ClusterOrchestratorSlurmConfig
}

```

### YAML

```yaml

  Eks:
    ClusterOrchestratorEksConfig
  Slurm:
    ClusterOrchestratorSlurmConfig

```

## Properties

`Eks`

The configuration of the Amazon EKS orchestrator cluster for the SageMaker HyperPod cluster.

_Required_: No

_Type_: [ClusterOrchestratorEksConfig](aws-properties-sagemaker-cluster-clusterorchestratoreksconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Slurm`

Configuration for using Slurm as the cluster orchestrator.

_Required_: No

_Type_: [ClusterOrchestratorSlurmConfig](aws-properties-sagemaker-cluster-clusterorchestratorslurmconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FSxLustreConfig

RollingUpdatePolicy

All content copied from https://docs.aws.amazon.com/.
