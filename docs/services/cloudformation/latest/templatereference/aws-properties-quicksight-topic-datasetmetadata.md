This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Topic DatasetMetadata

A structure that represents a dataset.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CalculatedFields" : [ TopicCalculatedField, ... ],
  "Columns" : [ TopicColumn, ... ],
  "DataAggregation" : DataAggregation,
  "DatasetArn" : String,
  "DatasetDescription" : String,
  "DatasetName" : String,
  "Filters" : [ TopicFilter, ... ],
  "NamedEntities" : [ TopicNamedEntity, ... ]
}

```

### YAML

```yaml

  CalculatedFields:
    - TopicCalculatedField
  Columns:
    - TopicColumn
  DataAggregation:
    DataAggregation
  DatasetArn: String
  DatasetDescription: String
  DatasetName: String
  Filters:
    - TopicFilter
  NamedEntities:
    - TopicNamedEntity

```

## Properties

`CalculatedFields`

The list of calculated field definitions.

_Required_: No

_Type_: Array of [TopicCalculatedField](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-topic-topiccalculatedfield.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Columns`

The list of column definitions.

_Required_: No

_Type_: Array of [TopicColumn](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-topic-topiccolumn.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataAggregation`

The definition of a data aggregation.

_Required_: No

_Type_: [DataAggregation](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-topic-dataaggregation.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DatasetArn`

The Amazon Resource Name (ARN) of the dataset.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DatasetDescription`

The description of the dataset.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DatasetName`

The name of the dataset.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Filters`

The list of filter definitions.

_Required_: No

_Type_: Array of [TopicFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-topic-topicfilter.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NamedEntities`

The list of named entities definitions.

_Required_: No

_Type_: Array of [TopicNamedEntity](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-topic-topicnamedentity.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DataAggregation

DefaultFormatting
