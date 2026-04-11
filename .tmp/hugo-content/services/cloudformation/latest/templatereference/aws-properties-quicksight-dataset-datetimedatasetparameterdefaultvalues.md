This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSet DateTimeDatasetParameterDefaultValues

The `DateTimeDatasetParameterDefaultValues` property type specifies Property description not available. for an [AWS::QuickSight::DataSet](aws-resource-quicksight-dataset.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "StaticValues" : [ String, ... ]
}

```

### YAML

```yaml

  StaticValues:
    - String

```

## Properties

`StaticValues`

A list of static default values for a given date time parameter. The valid format for this property is `yyyy-MM-dd’T’HH:mm:ss’Z’`.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `32`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DateTimeDatasetParameter

DecimalDatasetParameter

All content copied from https://docs.aws.amazon.com/.
