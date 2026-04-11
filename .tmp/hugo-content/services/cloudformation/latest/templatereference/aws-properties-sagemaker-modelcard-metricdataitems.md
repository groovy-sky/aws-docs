This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ModelCard MetricDataItems

Metric data. The `type` determines the data types that you specify for `value`,
`XAxisName` and `YAxisName`. For information about specifying values for metrics, see [model card JSON\
schema](../../../sagemaker/latest/dg/model-cards.md#model-cards-json-schema).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MetricDataItems" : SimpleMetric
}

```

### YAML

```yaml

  MetricDataItems:
    SimpleMetric

```

## Properties

`MetricDataItems`

A list of metric data items for the model.

_Required_: No

_Type_: [SimpleMetric](aws-properties-sagemaker-modelcard-simplemetric.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IntendedUses

MetricGroup

All content copied from https://docs.aws.amazon.com/.
