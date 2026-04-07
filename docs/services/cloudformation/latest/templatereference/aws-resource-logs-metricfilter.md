This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Logs::MetricFilter

The `AWS::Logs::MetricFilter` resource specifies a metric filter that
describes how CloudWatch Logs extracts information from logs and transforms it into
Amazon CloudWatch metrics. If you have multiple metric filters that are associated with
a log group, all the filters are applied to the log streams in that group.

The maximum number of metric filters that can be associated with a log group is
100.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Logs::MetricFilter",
  "Properties" : {
      "ApplyOnTransformedLogs" : Boolean,
      "EmitSystemFieldDimensions" : [ String, ... ],
      "FieldSelectionCriteria" : String,
      "FilterName" : String,
      "FilterPattern" : String,
      "LogGroupName" : String,
      "MetricTransformations" : [ MetricTransformation, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Logs::MetricFilter
Properties:
  ApplyOnTransformedLogs: Boolean
  EmitSystemFieldDimensions:
    - String
  FieldSelectionCriteria: String
  FilterName: String
  FilterPattern: String
  LogGroupName: String
  MetricTransformations:
    - MetricTransformation

```

## Properties

`ApplyOnTransformedLogs`

This parameter is valid only for log groups that have an active log transformer. For
more information about log transformers, see [PutTransformer](https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutTransformer.html).

If this value is `true`, the metric filter is applied on the transformed
version of the log events instead of the original ingested log events.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EmitSystemFieldDimensions`

The list of system fields that are emitted as additional dimensions in the generated
metrics. Returns the `emitSystemFieldDimensions` value if it was specified when the
metric filter was created.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FieldSelectionCriteria`

The filter expression that specifies which log events are processed by this metric filter
based on system fields. Returns the `fieldSelectionCriteria` value if it was
specified when the metric filter was created.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `2000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FilterName`

The name of the metric filter.

_Required_: No

_Type_: String

_Pattern_: `^[^:*]{1,512}`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FilterPattern`

A filter pattern for extracting metric data out of ingested log events. For more
information, see [Filter and Pattern\
Syntax](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/FilterAndPatternSyntax.html).

_Required_: Yes

_Type_: String

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogGroupName`

The name of an existing log group that you want to associate with this metric
filter.

_Required_: Yes

_Type_: String

_Pattern_: `^[.\-_/#A-Za-z0-9]{1,512}`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MetricTransformations`

The metric transformations.

_Required_: Yes

_Type_: Array of [MetricTransformation](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-logs-metricfilter-metrictransformation.html)

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

### Create a Metric Filter

The following example sends a value of `1` to the
`404Count` metric whenever the status code field includes a
`404` value.

#### JSON

```json

"404MetricFilter": {
    "Type": "AWS::Logs::MetricFilter",
    "Properties": {
        "LogGroupName": { "Ref": "myLogGroup" },
        "FilterPattern": "[ip, identity, user_id, timestamp, request, status_code = 404, size]",
        "MetricTransformations": [
            {
                "MetricValue": "1",
                "MetricNamespace": "WebServer/404s",
                "MetricName": "404Count"
            }
        ]
    }
}
```

#### YAML

```yaml

404MetricFilter:
  Type: AWS::Logs::MetricFilter
  Properties:
    LogGroupName:
      Ref: "myLogGroup"
    FilterPattern: "[ip, identity, user_id, timestamp, request, status_code = 404, size]"
    MetricTransformations:
      -
        MetricValue: "1"
        MetricNamespace: "WebServer/404s"
        MetricName: "404Count"
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Logs::LogStream

Dimension
