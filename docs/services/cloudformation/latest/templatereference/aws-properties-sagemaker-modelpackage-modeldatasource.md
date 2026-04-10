This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ModelPackage ModelDataSource

Specifies the location of ML model data to deploy. If specified, you must specify one
and only one of the available data sources.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "S3DataSource" : S3ModelDataSource
}

```

### YAML

```yaml

  S3DataSource:
    S3ModelDataSource

```

## Properties

`S3DataSource`

Specifies the S3 location of ML model data to deploy.

_Required_: No

_Type_: [S3ModelDataSource](aws-properties-sagemaker-modelpackage-s3modeldatasource.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ModelDataQuality

ModelInput

All content copied from https://docs.aws.amazon.com/.
