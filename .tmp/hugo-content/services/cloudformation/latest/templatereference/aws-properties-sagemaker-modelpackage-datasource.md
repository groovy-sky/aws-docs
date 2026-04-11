This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ModelPackage DataSource

Describes the location of the channel data.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "S3DataSource" : S3DataSource
}

```

### YAML

```yaml

  S3DataSource:
    S3DataSource

```

## Properties

`S3DataSource`

The S3 location of the data source that is associated with a channel.

_Required_: Yes

_Type_: [S3DataSource](aws-properties-sagemaker-modelpackage-s3datasource.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Bias

DriftCheckBaselines

All content copied from https://docs.aws.amazon.com/.
