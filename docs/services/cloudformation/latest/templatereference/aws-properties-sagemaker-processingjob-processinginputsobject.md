This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ProcessingJob ProcessingInputsObject

The inputs for a processing job. The processing input must specify exactly one of
either `S3Input` or `DatasetDefinition` types.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AppManaged" : Boolean,
  "DatasetDefinition" : DatasetDefinition,
  "InputName" : String,
  "S3Input" : S3Input
}

```

### YAML

```yaml

  AppManaged: Boolean
  DatasetDefinition:
    DatasetDefinition
  InputName: String
  S3Input:
    S3Input

```

## Properties

`AppManaged`

When `True`, input operations such as data download are managed natively by
the processing job application. When `False` (default), input operations are
managed by Amazon SageMaker.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DatasetDefinition`

Configuration for Dataset Definition inputs. The Dataset Definition input must specify
exactly one of either `AthenaDatasetDefinition` or
`RedshiftDatasetDefinition` types.

_Required_: No

_Type_: [DatasetDefinition](aws-properties-sagemaker-processingjob-datasetdefinition.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`InputName`

The name for the processing job input.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`S3Input`

Configuration for downloading input data from Amazon S3 into the processing
container.

_Required_: No

_Type_: [S3Input](aws-properties-sagemaker-processingjob-s3input.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NetworkConfig

ProcessingOutputConfig

All content copied from https://docs.aws.amazon.com/.
