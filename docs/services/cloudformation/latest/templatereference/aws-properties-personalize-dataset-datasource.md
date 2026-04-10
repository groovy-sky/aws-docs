This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Personalize::Dataset DataSource

Describes the data source that contains the data to upload to a dataset, or the list of
records to delete from Amazon Personalize.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DataLocation" : String
}

```

### YAML

```yaml

  DataLocation: String

```

## Properties

`DataLocation`

For dataset import jobs, the path to the Amazon S3 bucket where the data that you want to upload to
your dataset is stored. For data deletion jobs, the path to the Amazon S3 bucket that stores the list of records to delete.

For example:

`s3://bucket-name/folder-name/fileName.csv`

If your CSV files are in a folder in your Amazon S3 bucket and you want your import job or data deletion job
to consider multiple files, you can specify the path to the folder. With a data deletion job, Amazon Personalize uses all files in the folder and any sub folder. Use the following syntax with a `/` after the folder
name:

`s3://bucket-name/folder-name/`

_Required_: No

_Type_: String

_Pattern_: `(s3|http|https)://.+`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DatasetImportJob

AWS::Personalize::DatasetGroup

All content copied from https://docs.aws.amazon.com/.
