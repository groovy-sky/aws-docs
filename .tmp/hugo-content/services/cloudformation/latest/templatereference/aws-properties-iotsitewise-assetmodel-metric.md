This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTSiteWise::AssetModel Metric

Contains an asset metric property. With metrics, you can calculate aggregate functions,
such as an average, maximum, or minimum, as specified through an expression. A metric maps
several values to a single value (such as a sum).

The maximum number of dependent/cascading variables used in any one metric calculation is
10\. Therefore, a _root_ metric can have
up to 10 cascading metrics in its computational dependency
tree. Additionally, a metric can only have a data type of `DOUBLE` and consume
properties with data types of `INTEGER` or `DOUBLE`.

For more information, see [Metrics](../../../iot-sitewise/latest/userguide/asset-properties.md#metrics) in the _AWS IoT SiteWise User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Expression" : String,
  "Variables" : [ ExpressionVariable, ... ],
  "Window" : MetricWindow
}

```

### YAML

```yaml

  Expression: String
  Variables:
    - ExpressionVariable
  Window:
    MetricWindow

```

## Properties

`Expression`

The mathematical expression that defines the metric aggregation function. You can specify
up to 10 variables per expression. You can specify up to 10 functions
per expression.

For more information, see [Quotas](../../../iot-sitewise/latest/userguide/quotas.md) in the _AWS IoT SiteWise User Guide_.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Variables`

The list of variables used in the expression.

_Required_: Yes

_Type_: Array of [ExpressionVariable](aws-properties-iotsitewise-assetmodel-expressionvariable.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Window`

The window (time interval) over which AWS IoT SiteWise computes the metric's aggregation expression.
AWS IoT SiteWise computes one data point per `window`.

_Required_: Yes

_Type_: [MetricWindow](aws-properties-iotsitewise-assetmodel-metricwindow.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ExpressionVariable

MetricWindow

All content copied from https://docs.aws.amazon.com/.
