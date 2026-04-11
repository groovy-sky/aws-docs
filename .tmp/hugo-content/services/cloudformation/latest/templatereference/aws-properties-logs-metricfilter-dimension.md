This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Logs::MetricFilter Dimension

Specifies the CloudWatch metric dimensions to publish with this metric.

Because dimensions are part of the unique identifier for a metric, whenever a unique
dimension name/value pair is extracted from your logs, you are creating a new variation
of that metric.

For more information about publishing dimensions with metrics created by metric
filters, see [Publishing dimensions with metrics from values in JSON or space-delimited log\
events](../../../amazoncloudwatch/latest/logs/filterandpatternsyntax.md#logs-metric-filters-dimensions).

###### Important

Metrics extracted from log events are charged as custom metrics. To prevent
unexpected high charges, do not specify high-cardinality fields such as
`IPAddress` or `requestID` as dimensions. Each different
value found for a dimension is treated as a separate metric and accrues charges as a
separate custom metric.

To help prevent accidental high charges, Amazon disables a metric filter if it
generates 1000 different name/value pairs for the dimensions that you have specified
within a certain amount of time.

You can also set up a billing alarm to alert you if your charges are higher than
expected. For more information, see [Creating a Billing Alarm to Monitor Your Estimated AWS\
Charges](../../../amazoncloudwatch/latest/monitoring/monitor-estimated-charges-with-cloudwatch.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Key" : String,
  "Value" : String
}

```

### YAML

```yaml

  Key: String
  Value: String

```

## Properties

`Key`

The name for the CloudWatch metric dimension that the metric filter
creates.

Dimension names must contain only ASCII characters, must include at least one
non-whitespace character, and cannot start with a colon (:).

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The log event field that will contain the value for this dimension. This dimension
will only be published for a metric if the value is found in the log event. For example,
`$.eventType` for JSON log events, or `$server` for
space-delimited log events.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Logs::MetricFilter

MetricTransformation

All content copied from https://docs.aws.amazon.com/.
