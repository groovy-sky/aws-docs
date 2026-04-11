This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ModelCard MetricGroup

A group of metric data that you use to initialize a metric group object.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MetricData" : [ MetricDataItems, ... ],
  "Name" : String
}

```

### YAML

```yaml

  MetricData:
    - MetricDataItems
  Name: String

```

## Properties

`MetricData`

A list of metric objects. The `MetricDataItems` list can have one of the following values:

- `bar_chart_metric`

- `matrix_metric`

- `simple_metric`

- `linear_graph_metric`

For more information about the metric schema, see the definition section of the [model card JSON schema](../../../sagemaker/latest/dg/model-cards.md#model-cards-json-schema).

_Required_: Yes

_Type_: Array of [MetricDataItems](aws-properties-sagemaker-modelcard-metricdataitems.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The metric group name.

_Required_: Yes

_Type_: String

_Pattern_: `.{1,63}`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MetricDataItems

ModelOverview

All content copied from https://docs.aws.amazon.com/.
