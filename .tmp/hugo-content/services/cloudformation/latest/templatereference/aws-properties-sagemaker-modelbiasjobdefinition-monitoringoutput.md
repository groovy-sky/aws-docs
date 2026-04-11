This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ModelBiasJobDefinition MonitoringOutput

The output object for a monitoring job.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "S3Output" : S3Output
}

```

### YAML

```yaml

  S3Output:
    S3Output

```

## Properties

`S3Output`

The Amazon S3 storage location where the results of a monitoring job are
saved.

_Required_: Yes

_Type_: [S3Output](aws-properties-sagemaker-modelbiasjobdefinition-s3output.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MonitoringGroundTruthS3Input

MonitoringOutputConfig

All content copied from https://docs.aws.amazon.com/.
