This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApplicationSignals::ServiceLevelObjective Metric

This structure defines the metric used for a service level indicator, including the metric name, namespace, and dimensions

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Dimensions" : [ Dimension, ... ],
  "MetricName" : String,
  "Namespace" : String
}

```

### YAML

```yaml

  Dimensions:
    - Dimension
  MetricName: String
  Namespace: String

```

## Properties

`Dimensions`

An array of one or more dimensions to use to define the metric that you want to use.
For more information, see
[Dimensions](../../../amazoncloudwatch/latest/monitoring/cloudwatch-concepts.md#Dimension).

_Required_: No

_Type_: Array of [Dimension](aws-properties-applicationsignals-servicelevelobjective-dimension.md)

_Minimum_: `0`

_Maximum_: `30`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MetricName`

The name of the metric to use.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Namespace`

The namespace of the metric. For more information, see
[Namespaces](../../../amazoncloudwatch/latest/monitoring/cloudwatch-concepts.md#Namespace).

_Required_: No

_Type_: String

_Pattern_: `.*[^:].*`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Interval

MetricDataQuery

All content copied from https://docs.aws.amazon.com/.
