This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataBrew::Project

Specifies a new AWS Glue DataBrew project.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DataBrew::Project",
  "Properties" : {
      "DatasetName" : String,
      "Name" : String,
      "RecipeName" : String,
      "RoleArn" : String,
      "Sample" : Sample,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::DataBrew::Project
Properties:
  DatasetName: String
  Name: String
  RecipeName: String
  RoleArn: String
  Sample:
    Sample
  Tags:
    - Tag

```

## Properties

`DatasetName`

The dataset that the project is to act upon.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The unique name of a project.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RecipeName`

The name of a recipe that will be developed during a project session.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The Amazon Resource Name (ARN) of the role that will be assumed for this
project.

_Required_: Yes

_Type_: String

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Sample`

The sample size and sampling type to apply to the data. If this parameter isn't
specified, then the sample consists of the first 500 rows from the dataset.

_Required_: No

_Type_: [Sample](aws-properties-databrew-project-sample.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Metadata tags that have been applied to the project.

_Required_: No

_Type_: Array of [Tag](aws-properties-databrew-project-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref`
function, `Ref` returns the resource name. For example:

`{ "Ref": "myProject" }`

For an AWS Glue DataBrew project named `myProject`,
`Ref` returns the name of the project.

## Examples

### Creating projects

The following examples create new DataBrew projects.

#### YAML

```yaml

Resources:
  TestDataBrewProject:
    Type: AWS::DataBrew::Project
    Properties:
      Name: project-name
      RecipeName: recipe-name
      DatasetName: dataset-name
      RoleArn: arn:aws:iam::12345678910:role/PassRoleAdmin
      Sample:
        Size: 500
        Type: LAST_N

```

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Description": "This CloudFormation template specifies a DataBrew Project",
    "Resources": {
        "MyDataBrewProject": {
            "Type": "AWS::DataBrew::Project",
            "Properties": {
                "Name": "test-project",
                "RecipeName": "test-project-recipe",
                "DatasetName": "test-dataset",
                "RoleArn": "arn:aws:iam::1234567891011:role/PassRoleAdmin",
                "Sample": {
                    "Size": 500,
                    "Type": "LAST_N"
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

ValidationConfiguration

Sample

All content copied from https://docs.aws.amazon.com/.
