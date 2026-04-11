This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataBrew::Job ProfileConfiguration

Configuration for profile jobs. Configuration can be used to select columns, do evaluations, and override default
parameters of evaluations. When configuration is undefined, the profile job will apply default settings to all
supported columns.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ColumnStatisticsConfigurations" : [ ColumnStatisticsConfiguration, ... ],
  "DatasetStatisticsConfiguration" : StatisticsConfiguration,
  "EntityDetectorConfiguration" : EntityDetectorConfiguration,
  "ProfileColumns" : [ ColumnSelector, ... ]
}

```

### YAML

```yaml

  ColumnStatisticsConfigurations:
    - ColumnStatisticsConfiguration
  DatasetStatisticsConfiguration:
    StatisticsConfiguration
  EntityDetectorConfiguration:
    EntityDetectorConfiguration
  ProfileColumns:
    - ColumnSelector

```

## Properties

`ColumnStatisticsConfigurations`

List of configurations for column evaluations. ColumnStatisticsConfigurations are used to
select evaluations and override parameters of evaluations for particular columns. When
ColumnStatisticsConfigurations is undefined, the profile job will profile all supported columns
and run all supported evaluations.

_Required_: No

_Type_: Array of [ColumnStatisticsConfiguration](aws-properties-databrew-job-columnstatisticsconfiguration.md)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DatasetStatisticsConfiguration`

Configuration for inter-column evaluations. Configuration can be used to select evaluations and override
parameters of evaluations. When configuration is undefined, the profile job will run all supported
inter-column evaluations.

_Required_: No

_Type_: [StatisticsConfiguration](aws-properties-databrew-job-statisticsconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EntityDetectorConfiguration`

Configuration of entity detection for a profile job. When undefined, entity detection is disabled.

_Required_: No

_Type_: [EntityDetectorConfiguration](aws-properties-databrew-job-entitydetectorconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProfileColumns`

List of column selectors. ProfileColumns can be used to select columns from the dataset. When
ProfileColumns is undefined, the profile job will profile all supported columns.

_Required_: No

_Type_: Array of [ColumnSelector](aws-properties-databrew-job-columnselector.md)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OutputLocation

Recipe

All content copied from https://docs.aws.amazon.com/.
