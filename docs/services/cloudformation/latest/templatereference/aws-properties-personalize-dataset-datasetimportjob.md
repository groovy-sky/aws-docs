This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Personalize::Dataset DatasetImportJob

Describes a job that imports training data from a data source (Amazon S3
bucket) to an Amazon Personalize dataset.

A dataset import job can be in one of the following states:

- CREATE PENDING > CREATE IN\_PROGRESS > ACTIVE -or- CREATE
FAILED

If you specify a dataset import job as part of a dataset, all
dataset import job fields are required.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DatasetArn" : String,
  "DatasetImportJobArn" : String,
  "DataSource" : DataSource,
  "JobName" : String,
  "RoleArn" : String
}

```

### YAML

```yaml

  DatasetArn: String
  DatasetImportJobArn: String
  DataSource:
    DataSource
  JobName: String
  RoleArn: String

```

## Properties

`DatasetArn`

The Amazon Resource Name (ARN) of the dataset that receives the
imported data.

_Required_: No

_Type_: String

_Pattern_: `arn:([a-z\d-]+):personalize:.*:.*:.+`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DatasetImportJobArn`

The ARN of the dataset import job.

_Required_: No

_Type_: String

_Pattern_: `arn:([a-z\d-]+):personalize:.*:.*:.+`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataSource`

The Amazon S3 bucket that contains the training data to import.

_Required_: No

_Type_: [DataSource](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-personalize-dataset-datasource.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`JobName`

The name of the import job.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9][a-zA-Z0-9\-_]*`

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The ARN of the IAM role that has permissions to read from the Amazon S3
data source.

_Required_: No

_Type_: String

_Pattern_: `arn:([a-z\d-]+):iam::\d{12}:role/?[a-zA-Z_0-9+=,.@\-_/]+`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Personalize::Dataset

DataSource
