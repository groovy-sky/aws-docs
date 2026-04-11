This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Pipeline PipelineDefinition

The definition of the pipeline. This can be either a JSON string or an Amazon S3 location.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PipelineDefinitionBody" : String,
  "PipelineDefinitionS3Location" : S3Location
}

```

### YAML

```yaml

  PipelineDefinitionBody: String
  PipelineDefinitionS3Location:
    S3Location

```

## Properties

`PipelineDefinitionBody`

The [JSON pipeline definition](https://aws-sagemaker-mlops.github.io/sagemaker-model-building-pipeline-definition-JSON-schema) of the pipeline.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PipelineDefinitionS3Location`

The location of the pipeline definition stored in Amazon S3. If specified, SageMaker retrieves the pipeline
definition from this location.

_Required_: No

_Type_: [S3Location](aws-properties-sagemaker-pipeline-s3location.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ParallelismConfiguration

S3Location

All content copied from https://docs.aws.amazon.com/.
