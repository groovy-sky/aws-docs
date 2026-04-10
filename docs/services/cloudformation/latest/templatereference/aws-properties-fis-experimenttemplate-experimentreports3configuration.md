This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FIS::ExperimentTemplate ExperimentReportS3Configuration

The S3 destination for the experiment report.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BucketName" : String,
  "Prefix" : String
}

```

### YAML

```yaml

  BucketName: String
  Prefix: String

```

## Properties

`BucketName`

The name of the S3 bucket where the experiment report will be stored.

_Required_: Yes

_Type_: String

_Minimum_: `3`

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Prefix`

The prefix of the S3 bucket where the experiment report will be stored.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataSources

ExperimentTemplateAction

All content copied from https://docs.aws.amazon.com/.
