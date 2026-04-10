This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ModelExplainabilityJobDefinition BatchTransformInput

Input object for the batch transform job.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DataCapturedDestinationS3Uri" : String,
  "DatasetFormat" : DatasetFormat,
  "FeaturesAttribute" : String,
  "InferenceAttribute" : String,
  "LocalPath" : String,
  "ProbabilityAttribute" : String,
  "S3DataDistributionType" : String,
  "S3InputMode" : String
}

```

### YAML

```yaml

  DataCapturedDestinationS3Uri: String
  DatasetFormat:
    DatasetFormat
  FeaturesAttribute: String
  InferenceAttribute: String
  LocalPath: String
  ProbabilityAttribute: String
  S3DataDistributionType: String
  S3InputMode: String

```

## Properties

`DataCapturedDestinationS3Uri`

The Amazon S3 location being used to capture the data.

_Required_: Yes

_Type_: String

_Pattern_: `^(https|s3)://([^/]+)/?(.*)$`

_Maximum_: `512`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DatasetFormat`

The dataset format for your batch transform job.

_Required_: Yes

_Type_: [DatasetFormat](aws-properties-sagemaker-modelexplainabilityjobdefinition-datasetformat.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FeaturesAttribute`

The attributes of the input data that are the input features.

_Required_: No

_Type_: String

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`InferenceAttribute`

The attribute of the input data that represents the ground truth label.

_Required_: No

_Type_: String

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LocalPath`

Path to the filesystem where the batch transform data is available to the container.

_Required_: Yes

_Type_: String

_Pattern_: `.*`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ProbabilityAttribute`

In a classification problem, the attribute that represents the class probability.

_Required_: No

_Type_: String

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`S3DataDistributionType`

Whether input data distributed in Amazon S3 is fully replicated or sharded by an S3 key.
Defaults to `FullyReplicated`

_Required_: No

_Type_: String

_Allowed values_: `FullyReplicated | ShardedByS3Key`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`S3InputMode`

Whether the `Pipe` or `File` is used as the input mode for
transferring data for the monitoring job. `Pipe` mode is recommended for large
datasets. `File` mode is useful for small files that fit in memory. Defaults to
`File`.

_Required_: No

_Type_: String

_Allowed values_: `Pipe | File`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::SageMaker::ModelExplainabilityJobDefinition

ClusterConfig

All content copied from https://docs.aws.amazon.com/.
