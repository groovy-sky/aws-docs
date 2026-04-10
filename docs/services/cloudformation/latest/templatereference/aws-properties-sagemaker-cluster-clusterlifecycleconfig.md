This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Cluster ClusterLifeCycleConfig

The lifecycle configuration for a SageMaker HyperPod cluster.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "OnCreate" : String,
  "SourceS3Uri" : String
}

```

### YAML

```yaml

  OnCreate: String
  SourceS3Uri: String

```

## Properties

`OnCreate`

The file name of the entrypoint script of lifecycle scripts under
`SourceS3Uri`. This entrypoint script runs during cluster creation.

_Required_: Yes

_Type_: String

_Pattern_: `^[\S\s]+$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceS3Uri`

An Amazon S3 bucket path where your lifecycle scripts are stored.

###### Important

Make sure that the S3 bucket path starts with `s3://sagemaker-`. The
[IAM role for SageMaker HyperPod](../../../sagemaker/latest/dg/sagemaker-hyperpod-prerequisites.md#sagemaker-hyperpod-prerequisites-iam-role-for-hyperpod) has the managed [`AmazonSageMakerClusterInstanceRolePolicy`](../../../sagemaker/latest/dg/security-iam-awsmanpol-cluster.md) attached, which
allows access to S3 buckets with the specific prefix `sagemaker-`.

_Required_: Yes

_Type_: String

_Pattern_: `^(https|s3)://([^/]+)/?(.*)$`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ClusterKubernetesTaint

ClusterOrchestratorEksConfig

All content copied from https://docs.aws.amazon.com/.
