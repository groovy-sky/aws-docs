This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Pipeline

The `AWS::SageMaker::Pipeline` resource creates shell scripts that run when you create and/or start a
SageMaker Pipeline. For information about SageMaker Pipelines, see [SageMaker Pipelines](../../../sagemaker/latest/dg/pipelines.md) in the _Amazon SageMaker Developer_
_Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SageMaker::Pipeline",
  "Properties" : {
      "ParallelismConfiguration" : ParallelismConfiguration,
      "PipelineDefinition" : PipelineDefinition,
      "PipelineDescription" : String,
      "PipelineDisplayName" : String,
      "PipelineName" : String,
      "RoleArn" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::SageMaker::Pipeline
Properties:
  ParallelismConfiguration:
    ParallelismConfiguration
  PipelineDefinition:
    PipelineDefinition
  PipelineDescription: String
  PipelineDisplayName: String
  PipelineName: String
  RoleArn: String
  Tags:
    - Tag

```

## Properties

`ParallelismConfiguration`

The parallelism configuration applied to the pipeline.

_Required_: No

_Type_: [ParallelismConfiguration](aws-properties-sagemaker-pipeline-parallelismconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PipelineDefinition`

The definition of the pipeline. This can be either a JSON string or an Amazon S3 location.

_Required_: Yes

_Type_: [PipelineDefinition](aws-properties-sagemaker-pipeline-pipelinedefinition.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PipelineDescription`

The description of the pipeline.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `3072`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PipelineDisplayName`

The display name of the pipeline.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9](-*[a-zA-Z0-9])*`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PipelineName`

The name of the pipeline.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9](-*[a-zA-Z0-9])*`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RoleArn`

The Amazon Resource Name (ARN) of the IAM role used to execute the pipeline.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[a-z\-]*:iam::\d{12}:role/?[a-zA-Z_0-9+=,.@\-_/]+$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags of the pipeline.

_Required_: No

_Type_: Array of [Tag](aws-properties-sagemaker-pipeline-tag.md)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the PipelineName of the pipeline.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## Examples

### SageMaker Pipeline Example

The following example creates a Pipeline with an associated lifecycle configuration.

#### JSON

```json

# Pipeline definition given as a JSON string { "Resources": {
            "MyPipeline": { "Type": "AWS::SageMaker::Pipeline", "Properties": { "PipelineName":
            "<pipeline-name>" "PipelineDisplayName": "<pipeline-display-name>",
            "PipelineDescription": "<pipeline-description>", "PipelineDefinition": {
            "PipelineDefinitionBody":
               "{\"Version\":\"2020-12-01\",\"Parameters\":[{\"Name\":\"InputDataSource\",\"DefaultValue\":\"\"},{\"Name\":\"InstanceCount\",\"Type\":\"Integer\",\"DefaultValue\":1}],\"Steps\":[{\"Name\":\"Training1\",\"Type\":\"Training\",\"Arguments\":{\"InputDataConfig\":[{\"DataSource\":{\"S3DataSource\":{\"S3Uri\":{\"Get\":\"Parameters.InputDataSource\"}}}}],\"OutputDataConfig\":{\"S3OutputPath\":\"s3://amzn-s3-demo-bucket/\"},\"ResourceConfig\":{\"InstanceType\":\"ml.m5.large\",\"InstanceCount\":{\"Get\":\"Parameters.InstanceCount\"},\"VolumeSizeInGB\":1024}}}]}"
            }, "RoleArn": "arn:aws:iam::<account-id>:root" } } } }
```

#### JSON

```json

# Pipeline definition given as an S3 string { "Resources": {
            "MyPipeline": { "Type": "AWS::SageMaker::Pipeline", "Properties": { "PipelineName":
            "<pipeline-name>", "PipelineDisplayName": "<pipeline-display-name>",
            "PipelineDescription": "<pipeline-description>", "PipelineDefinition": {
            "PipelineDefinitionS3Location": { "Bucket": "<S3-bucket-location>", "Key":
            "<S3-bucket-key>" } }, "RoleArn": "arn:aws:iam::<account-id>:root" } } }
            }
```

#### YAML

```yaml

# Pipeline definition given as a JSON string Resources:
            MyPipeline: Type: AWS::SageMaker::Pipeline Properties: PipelineName:
            "<pipeline-name>" PipelineDisplayName: "<pipeline-display-name>"
            PipelineDescription: "<pipeline-description>" PipelineDefinition:
            PipelineDefinitionBody:
               "{\"Version\":\"2020-12-01\",\"Parameters\":[{\"Name\":\"InputDataSource\",\"DefaultValue\":\"\"},{\"Name\":\"InstanceCount\",\"Type\":\"Integer\",\"DefaultValue\":1}],\"Steps\":[{\"Name\":\"Training1\",\"Type\":\"Training\",\"Arguments\":{\"InputDataConfig\":[{\"DataSource\":{\"S3DataSource\":{\"S3Uri\":{\"Get\":\"Parameters.InputDataSource\"}}}}],\"OutputDataConfig\":{\"S3OutputPath\":\"s3://amzn-s3-demo-bucket/\"},\"ResourceConfig\":{\"InstanceType\":\"ml.m5.large\",\"InstanceCount\":{\"Get\":\"Parameters.InstanceCount\"},\"VolumeSizeInGB\":1024}}}]}"
            RoleArn: "arn:aws:iam::<account-id>:root"
```

#### YAML

```yaml

# Pipeline definition given as an S3 location Resources:
            MyPipeline: Type: AWS::SageMaker::Pipeline Properties: PipelineName:
            "<pipeline-name>" PipelineDisplayName:"<pipeline-display-name>"
            PipelineDescription: "<pipeline-description>" PipelineDefinition:
            PipelineDefinitionS3Location: Bucket: "<S3-bucket-location>" Key:
            "<S3-bucket-key>" RoleArn: "arn:aws:iam::<account-id>:root"
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

ParallelismConfiguration

All content copied from https://docs.aws.amazon.com/.
