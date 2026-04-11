This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ProcessingJob ProcessingOutputsObject

Describes the results of a processing job. The processing output must specify exactly
one of either `S3Output` or `FeatureStoreOutput` types.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AppManaged" : Boolean,
  "FeatureStoreOutput" : FeatureStoreOutput,
  "OutputName" : String,
  "S3Output" : S3Output
}

```

### YAML

```yaml

  AppManaged: Boolean
  FeatureStoreOutput:
    FeatureStoreOutput
  OutputName: String
  S3Output:
    S3Output

```

## Properties

`AppManaged`

When `True`, output operations such as data upload are managed natively by
the processing job application. When `False` (default), output operations are
managed by Amazon SageMaker.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FeatureStoreOutput`

Configuration for processing job outputs in Amazon SageMaker Feature Store.

_Required_: No

_Type_: [FeatureStoreOutput](aws-properties-sagemaker-processingjob-featurestoreoutput.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`OutputName`

The name for the processing job output.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`S3Output`

Configuration for uploading output data to Amazon S3 from the processing container.

_Required_: No

_Type_: [S3Output](aws-properties-sagemaker-processingjob-s3output.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ProcessingOutputConfig

ProcessingResources

All content copied from https://docs.aws.amazon.com/.
