This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSet PhysicalTable

A view of a data source that contains information about the shape of the data in the
underlying source. This is a variant type structure. For this structure to be valid,
only one of the attributes can be non-null.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CustomSql" : CustomSql,
  "RelationalTable" : RelationalTable,
  "S3Source" : S3Source,
  "SaaSTable" : SaaSTable
}

```

### YAML

```yaml

  CustomSql:
    CustomSql
  RelationalTable:
    RelationalTable
  S3Source:
    S3Source
  SaaSTable:
    SaaSTable

```

## Properties

`CustomSql`

A physical table type built from the results of the custom SQL query.

_Required_: No

_Type_: [CustomSql](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-dataset-customsql.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RelationalTable`

A physical table type for relational data sources.

_Required_: No

_Type_: [RelationalTable](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-dataset-relationaltable.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3Source`

A physical table type for as S3 data source.

_Required_: No

_Type_: [S3Source](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-dataset-s3source.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SaaSTable`

A physical table type for Software-as-a-Service (SaaS) sources.

_Required_: No

_Type_: [SaaSTable](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-dataset-saastable.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

PerformanceConfiguration

PivotConfiguration
