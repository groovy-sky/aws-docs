This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataBrew::Job ColumnStatisticsConfiguration

Configuration for column evaluations for a profile job. ColumnStatisticsConfiguration can be used to select
evaluations and override parameters of evaluations for particular columns.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Selectors" : [ ColumnSelector, ... ],
  "Statistics" : StatisticsConfiguration
}

```

### YAML

```yaml

  Selectors:
    - ColumnSelector
  Statistics:
    StatisticsConfiguration

```

## Properties

`Selectors`

List of column selectors. Selectors can be used to select columns from the dataset.
When selectors are undefined, configuration will be applied to all supported columns.

_Required_: No

_Type_: Array of [ColumnSelector](aws-properties-databrew-job-columnselector.md)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Statistics`

Configuration for evaluations. Statistics can be used to select evaluations and override
parameters of evaluations.

_Required_: Yes

_Type_: [StatisticsConfiguration](aws-properties-databrew-job-statisticsconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ColumnSelector

CsvOutputOptions

All content copied from https://docs.aws.amazon.com/.
