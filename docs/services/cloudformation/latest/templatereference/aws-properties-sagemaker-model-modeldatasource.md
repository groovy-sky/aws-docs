This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Model ModelDataSource

Specifies the location of ML model data to deploy. If specified, you must specify one
and only one of the available data sources.

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

Specifies the S3 location of ML model data to deploy.

_Required_: Yes

_Type_: [S3DataSource](aws-properties-sagemaker-model-s3datasource.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ModelAccessConfig

MultiModelConfig

All content copied from https://docs.aws.amazon.com/.
