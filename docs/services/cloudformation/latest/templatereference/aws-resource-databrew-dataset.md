This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataBrew::Dataset

Specifies a new DataBrew dataset.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DataBrew::Dataset",
  "Properties" : {
      "Format" : String,
      "FormatOptions" : FormatOptions,
      "Input" : Input,
      "Name" : String,
      "PathOptions" : PathOptions,
      "Source" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::DataBrew::Dataset
Properties:
  Format: String
  FormatOptions:
    FormatOptions
  Input:
    Input
  Name: String
  PathOptions:
    PathOptions
  Source: String
  Tags:
    - Tag

```

## Properties

`Format`

The file format of a dataset that is created from an Amazon S3 file or folder.

_Required_: No

_Type_: String

_Allowed values_: `CSV | JSON | PARQUET | EXCEL | ORC`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FormatOptions`

A set of options that define how DataBrew interprets the data in the
dataset.

_Required_: No

_Type_: [FormatOptions](aws-properties-databrew-dataset-formatoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Input`

Information on how DataBrew can find the dataset, in either the AWS Glue Data Catalog or Amazon S3.

_Required_: Yes

_Type_: [Input](aws-properties-databrew-dataset-input.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The unique name of the dataset.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PathOptions`

A set of options that defines how DataBrew interprets an Amazon S3 path of the dataset.

_Required_: No

_Type_: [PathOptions](aws-properties-databrew-dataset-pathoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Source`

The location of the data for the dataset, either Amazon S3 or the AWS Glue Data Catalog.

_Required_: No

_Type_: String

_Allowed values_: `S3 | DATA-CATALOG | DATABASE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Metadata tags that have been applied to the dataset.

_Required_: No

_Type_: Array of [Tag](aws-properties-databrew-dataset-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref`
function, `Ref` returns the resource name. For example:

`{ "Ref": "myDataset" }`

For an AWS Glue DataBrew dataset named `myDataset`,
`Ref` returns the name of the dataset.

## Examples

### Creating datasets

The following examples create new DataBrew datasets.

#### YAML

```yaml

Resources:
  TestDataBrewDataset:
    Type: AWS::DataBrew::Dataset
    Properties:
      Name: dataset-name
      Input:
        S3InputDefinition:
          Bucket: !Join [ '', ['databrew-cfn-integration-tests-', !Ref 'AWS::Region', '-', !Ref 'AWS::AccountId' ] ]
          Key: cocktails.json
      FormatOptions:
        Json:
          MultiLine: True

```

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Description": "This CloudFormation template specifies a DataBrew Dataset",
    "Resources": {
    "TestDataBrewDataset": {
      "Type": "AWS::DataBrew::Dataset",
      "Properties": {
        "Name": "cf-test-dataset1",
        "Input": {
          "S3InputDefinition": {
            "Bucket": "test-location",
            "Key": "test.xlsx"
          }
        },
        "FormatOptions": {
          "Excel": {
            "SheetNames": ["test"]
          }
        },
        "Tags": [
                    {
                        "Key": "key00AtCreate",
                        "Value": "value001AtCreate"
                    }
                ]
      }
    }
  }
}

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS Glue DataBrew

CsvOptions

All content copied from https://docs.aws.amazon.com/.
