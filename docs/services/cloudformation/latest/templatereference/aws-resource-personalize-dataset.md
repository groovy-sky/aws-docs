This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Personalize::Dataset

Creates an empty dataset and adds it to the specified dataset group.
Use [CreateDatasetImportJob](https://docs.aws.amazon.com/personalize/latest/dg/API_CreateDatasetImportJob.html) to import your training data to a
dataset.

There are 5 types of datasets:

- Item interactions

- Items

- Users

- Action interactions (you can't use CloudFormation to create an Action interactions dataset)

- Actions (you can't use CloudFormation to create an Actions dataset)

Each dataset type has an associated schema with required field types.
Only the `Item interactions` dataset is required in order to train a
model (also referred to as creating a solution).

A dataset can be in one of the following states:

- CREATE PENDING > CREATE IN\_PROGRESS > ACTIVE -or- CREATE
FAILED

- DELETE PENDING > DELETE IN\_PROGRESS

To get the status of the dataset, call [DescribeDataset](https://docs.aws.amazon.com/personalize/latest/dg/API_DescribeDataset.html).

###### Related APIs

- [CreateDatasetGroup](https://docs.aws.amazon.com/personalize/latest/dg/API_CreateDatasetGroup.html)

- [ListDatasets](https://docs.aws.amazon.com/personalize/latest/dg/API_ListDatasets.html)

- [DescribeDataset](https://docs.aws.amazon.com/personalize/latest/dg/API_DescribeDataset.html)

- [DeleteDataset](https://docs.aws.amazon.com/personalize/latest/dg/API_DeleteDataset.html)

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Personalize::Dataset",
  "Properties" : {
      "DatasetGroupArn" : String,
      "DatasetImportJob" : DatasetImportJob,
      "DatasetType" : String,
      "Name" : String,
      "SchemaArn" : String
    }
}

```

### YAML

```yaml

Type: AWS::Personalize::Dataset
Properties:
  DatasetGroupArn: String
  DatasetImportJob:
    DatasetImportJob
  DatasetType: String
  Name: String
  SchemaArn: String

```

## Properties

`DatasetGroupArn`

The Amazon Resource Name (ARN) of the dataset group.

_Required_: Yes

_Type_: String

_Pattern_: `arn:([a-z\d-]+):personalize:.*:.*:.+`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DatasetImportJob`

Describes a job that imports training data from a data source (Amazon S3 bucket) to an
Amazon Personalize dataset. If you specify a dataset import job as part of a dataset, all
dataset import job fields are required.

_Required_: No

_Type_: [DatasetImportJob](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-personalize-dataset-datasetimportjob.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DatasetType`

One of the following values:

- Interactions

- Items

- Users

###### Note

You can't use CloudFormation to create an Action Interactions or Actions dataset.

_Required_: Yes

_Type_: String

_Allowed values_: `Interactions | Items | Users`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the dataset.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9][a-zA-Z0-9\-_]*`

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SchemaArn`

The ARN of the associated schema.

_Required_: Yes

_Type_: String

_Pattern_: `arn:([a-z\d-]+):personalize:.*:.*:.+`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the name of the resource.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`DatasetArn`

The Amazon Resource Name (ARN) of the dataset.

## Examples

### Creating a dataset

The following example creates an Amazon Personalize dataset and a dataset import job. The dataset
import job imports data from an Amazon S3 bucket into the dataset.

#### JSON

```json

{
  "AWSTemplateFormatVersion": "2010-09-09",
  "Resources": {
    "MyDataset": {
      "Type": "AWS::Personalize::Dataset",
      "Properties": {
        "Name": "my-dataset-name",
        "DatasetType": "Interactions",
        "DatasetGroupArn": "arn:aws:personalize:us-west-2:123456789012:dataset-group/dataset-group-name",
        "SchemaArn": "arn:aws:personalize:us-west-2:123456789012:schema/schema-name",
        "DatasetImportJob": {
          "JobName": "my-import-job-name",
          "DataSource": {
            "DataLocation": "s3://bucket-name/file-name.csv"
          },
          "RoleArn": "arn:aws:iam::123456789012:role/personalize-role"
        }
      }
    }
  }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Resources:
  MyDataset:
    Type: 'AWS::Personalize::Dataset'
    Properties:
      Name: my-dataset-name
      DatasetType: Interactions
      DatasetGroupArn: 'arn:aws:personalize:us-west-2:123456789012:dataset-group/dataset-group-name'
      SchemaArn: 'arn:aws:personalize:us-west-2:123456789012:schema/schema-name'
      DatasetImportJob:
        JobName: my-import-job-name
        DataSource:
          DataLocation: 's3://bucket-name/file-name.csv'
        RoleArn: 'arn:aws:iam::123456789012:role/personalize-role'
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Amazon Personalize

DatasetImportJob
