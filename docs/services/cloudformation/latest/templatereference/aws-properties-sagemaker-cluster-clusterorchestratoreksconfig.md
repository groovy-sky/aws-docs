This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Cluster ClusterOrchestratorEksConfig

The configuration for the Amazon EKS cluster that is used as the orchestrator for the SageMaker HyperPod
cluster. This includes the Amazon Resource Name (ARN) of the EKS cluster

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ClusterArn" : String
}

```

### YAML

```yaml

  ClusterArn: String

```

## Properties

`ClusterArn`

The Amazon Resource Name (ARN) of the SageMaker HyperPod cluster.

_Required_: Yes

_Type_: String

_Pattern_: `arn:aws[a-z\-]*:sagemaker:[a-z0-9\-]*:[0-9]{12}:cluster/[a-z0-9]{12}`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ClusterLifeCycleConfig

ClusterOrchestratorSlurmConfig

All content copied from https://docs.aws.amazon.com/.
