This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis ForecastConfiguration

The forecast configuration that is used in a line chart's display properties.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ForecastProperties" : TimeBasedForecastProperties,
  "Scenario" : ForecastScenario
}

```

### YAML

```yaml

  ForecastProperties:
    TimeBasedForecastProperties
  Scenario:
    ForecastScenario

```

## Properties

`ForecastProperties`

The forecast properties setup of a forecast in the line
chart.

_Required_: No

_Type_: [TimeBasedForecastProperties](aws-properties-quicksight-analysis-timebasedforecastproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Scenario`

The forecast scenario of a forecast in the line chart.

_Required_: No

_Type_: [ForecastScenario](aws-properties-quicksight-analysis-forecastscenario.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ForecastComputation

ForecastScenario

All content copied from https://docs.aws.amazon.com/.
