This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTSiteWise::Dataset DatasetSource

The data source for the dataset.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SourceDetail" : SourceDetail,
  "SourceFormat" : String,
  "SourceType" : String
}

```

### YAML

```yaml

  SourceDetail:
    SourceDetail
  SourceFormat: String
  SourceType: String

```

## Properties

`SourceDetail`

The details of the dataset source associated with the dataset.

_Required_: No

_Type_: [SourceDetail](aws-properties-iotsitewise-dataset-sourcedetail.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceFormat`

The format of the dataset source associated with the dataset.

_Required_: Yes

_Type_: String

_Allowed values_: `KNOWLEDGE_BASE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceType`

The type of data source for the dataset.

_Required_: Yes

_Type_: String

_Allowed values_: `KENDRA`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::IoTSiteWise::Dataset

KendraSourceDetail

All content copied from https://docs.aws.amazon.com/.
