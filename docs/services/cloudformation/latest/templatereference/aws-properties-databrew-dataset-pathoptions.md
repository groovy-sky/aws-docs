This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataBrew::Dataset PathOptions

Represents a set of options that define how DataBrew selects files for a
given Amazon S3 path in a dataset.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FilesLimit" : FilesLimit,
  "LastModifiedDateCondition" : FilterExpression,
  "Parameters" : [ PathParameter, ... ]
}

```

### YAML

```yaml

  FilesLimit:
    FilesLimit
  LastModifiedDateCondition:
    FilterExpression
  Parameters:
    - PathParameter

```

## Properties

`FilesLimit`

If provided, this structure imposes a limit on a number of files that should be
selected.

_Required_: No

_Type_: [FilesLimit](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-databrew-dataset-fileslimit.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LastModifiedDateCondition`

If provided, this structure defines a date range for matching Amazon S3
objects based on their LastModifiedDate attribute in Amazon S3.

_Required_: No

_Type_: [FilterExpression](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-databrew-dataset-filterexpression.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Parameters`

A structure that maps names of parameters used in the Amazon S3 path of a
dataset to their definitions.

_Required_: No

_Type_: Array of [PathParameter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-databrew-dataset-pathparameter.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Metadata

PathParameter
