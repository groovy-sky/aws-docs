This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataBrew::Job StatisticsConfiguration

Configuration of evaluations for a profile job. This configuration can be used to select
evaluations and override the parameters of selected evaluations.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IncludedStatistics" : [ String, ... ],
  "Overrides" : [ StatisticOverride, ... ]
}

```

### YAML

```yaml

  IncludedStatistics:
    - String
  Overrides:
    - StatisticOverride

```

## Properties

`IncludedStatistics`

List of included evaluations. When the list is undefined, all supported
evaluations will be included.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Overrides`

List of overrides for evaluations.

_Required_: No

_Type_: Array of [StatisticOverride](aws-properties-databrew-job-statisticoverride.md)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StatisticOverride

Tag

All content copied from https://docs.aws.amazon.com/.
