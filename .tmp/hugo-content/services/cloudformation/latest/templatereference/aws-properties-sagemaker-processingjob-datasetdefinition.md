This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ProcessingJob DatasetDefinition

Configuration for Dataset Definition inputs. The Dataset Definition input must specify
exactly one of either `AthenaDatasetDefinition` or
`RedshiftDatasetDefinition` types.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AthenaDatasetDefinition" : AthenaDatasetDefinition,
  "DataDistributionType" : String,
  "InputMode" : String,
  "LocalPath" : String,
  "RedshiftDatasetDefinition" : RedshiftDatasetDefinition
}

```

### YAML

```yaml

  AthenaDatasetDefinition:
    AthenaDatasetDefinition
  DataDistributionType: String
  InputMode: String
  LocalPath: String
  RedshiftDatasetDefinition:
    RedshiftDatasetDefinition

```

## Properties

`AthenaDatasetDefinition`

Configuration for Athena Dataset Definition input.

_Required_: No

_Type_: [AthenaDatasetDefinition](aws-properties-sagemaker-processingjob-athenadatasetdefinition.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DataDistributionType`

Whether the generated dataset is `FullyReplicated` or
`ShardedByS3Key` (default).

_Required_: No

_Type_: String

_Allowed values_: `FullyReplicated | ShardedByS3Key`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`InputMode`

Whether to use `File` or `Pipe` input mode. In `File`
(default) mode, Amazon SageMaker copies the data from the input source onto the local Amazon Elastic
Block Store (Amazon EBS) volumes before starting your training algorithm. This is the
most commonly used input mode. In `Pipe` mode, Amazon SageMaker streams input data from
the source directly to your algorithm without using the EBS volume.

_Required_: No

_Type_: String

_Allowed values_: `File | Pipe`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LocalPath`

The local path where you want Amazon SageMaker to download the Dataset Definition inputs to run a
processing job. `LocalPath` is an absolute path to the input data. This is a
required parameter when `AppManaged` is `False` (default).

_Required_: No

_Type_: String

_Pattern_: `.*`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RedshiftDatasetDefinition`

Configuration for Redshift Dataset Definition input.

_Required_: No

_Type_: [RedshiftDatasetDefinition](aws-properties-sagemaker-processingjob-redshiftdatasetdefinition.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ClusterConfig

ExperimentConfig

All content copied from https://docs.aws.amazon.com/.
