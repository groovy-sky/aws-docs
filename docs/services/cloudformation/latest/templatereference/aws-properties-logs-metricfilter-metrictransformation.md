This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Logs::MetricFilter MetricTransformation

`MetricTransformation` is a property of the
`AWS::Logs::MetricFilter` resource that describes how to transform log
streams into a CloudWatch metric.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DefaultValue" : Number,
  "Dimensions" : [ Dimension, ... ],
  "MetricName" : String,
  "MetricNamespace" : String,
  "MetricValue" : String,
  "Unit" : String
}

```

### YAML

```yaml

  DefaultValue: Number
  Dimensions:
    - Dimension
  MetricName: String
  MetricNamespace: String
  MetricValue: String
  Unit: String

```

## Properties

`DefaultValue`

(Optional) The value to emit when a filter pattern does not match a log event. This
value can be null.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Dimensions`

The fields to use as dimensions for the metric. One metric filter can include as many as
three dimensions.

###### Important

Metrics extracted from log events are charged as custom metrics. To prevent unexpected
high charges, do not specify high-cardinality fields such as `IPAddress` or
`requestID` as dimensions. Each different value found for a dimension is
treated as a separate metric and accrues charges as a separate custom metric.

CloudWatch Logs disables a metric filter if it generates 1000 different name/value
pairs for your specified dimensions within a certain amount of time. This helps to prevent
accidental high charges.

You can also set up a billing alarm to alert you if your charges are higher than
expected. For more information, see [Creating a Billing Alarm to Monitor Your Estimated AWS Charges](../../../amazoncloudwatch/latest/monitoring/monitor-estimated-charges-with-cloudwatch.md).

_Required_: No

_Type_: Array of [Dimension](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-logs-metricfilter-dimension.html)

_Minimum_: `1`

_Maximum_: `3`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MetricName`

The name of the CloudWatch metric.

_Required_: Yes

_Type_: String

_Pattern_: `^((?![:*$])[\x00-\x7F]){1,255}`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MetricNamespace`

A custom namespace to contain your metric in CloudWatch. Use namespaces to group
together metrics that are similar. For more information, see [Namespaces](../../../amazoncloudwatch/latest/monitoring/cloudwatch-concepts.md#Namespace).

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9a-zA-Z\.\-_\/#]{1,256}`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MetricValue`

The value that is published to the CloudWatch metric. For example, if you're
counting the occurrences of a particular term like `Error`, specify 1 for the
metric value. If you're counting the number of bytes transferred, reference the value
that is in the log event by using $. followed by the name of the field that you
specified in the filter pattern, such as `$.size`.

_Required_: Yes

_Type_: String

_Pattern_: `.{1,100}`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Unit`

The unit to assign to the metric. If you omit this, the unit is set as
`None`.

_Required_: No

_Type_: String

_Allowed values_: `Seconds | Microseconds | Milliseconds | Bytes | Kilobytes | Megabytes | Gigabytes | Terabytes | Bits | Kilobits | Megabits | Gigabits | Terabits | Percent | Count | Bytes/Second | Kilobytes/Second | Megabytes/Second | Gigabytes/Second | Terabytes/Second | Bits/Second | Kilobits/Second | Megabits/Second | Gigabits/Second | Terabits/Second | Count/Second | None`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Dimension

AWS::Logs::QueryDefinition
