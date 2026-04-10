This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudWatch::MetricStream

Creates or updates a metric stream. Metrics streams can automatically stream CloudWatch metrics to
AWS destinations including Amazon S3 and to many third-party solutions. For more information, see
[Metric streams](../../../amazoncloudwatch/latest/monitoring/cloudwatch-metric-streams.md).

To create a metric stream, you must be logged on to an account that has the `iam:PassRole` permission and
either the **CloudWatchFullAccess** policy or the `cloudwatch:PutMetricStream` permission.

When you create or update a metric stream, you choose one of the following:

- Stream metrics from all metric namespaces in the account.

- Stream metrics from all metric namespaces in the account, except for the namespaces that you list
in `ExcludeFilters`.

- Stream metrics from only the metric namespaces that you list in `IncludeFilters`.

When you create a metric stream, the stream is created in the `running` state. If you update an existing metric stream,
the state does not change.

If you create a metric stream in an account that has been set up as a monitoring account in CloudWatch cross-account observability,
you can choose whether to include metrics from linked source accounts in the metric stream.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CloudWatch::MetricStream",
  "Properties" : {
      "ExcludeFilters" : [ MetricStreamFilter, ... ],
      "FirehoseArn" : String,
      "IncludeFilters" : [ MetricStreamFilter, ... ],
      "IncludeLinkedAccountsMetrics" : Boolean,
      "Name" : String,
      "OutputFormat" : String,
      "RoleArn" : String,
      "StatisticsConfigurations" : [ MetricStreamStatisticsConfiguration, ... ],
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::CloudWatch::MetricStream
Properties:
  ExcludeFilters:
    - MetricStreamFilter
  FirehoseArn: String
  IncludeFilters:
    - MetricStreamFilter
  IncludeLinkedAccountsMetrics: Boolean
  Name: String
  OutputFormat: String
  RoleArn: String
  StatisticsConfigurations:
    - MetricStreamStatisticsConfiguration
  Tags:
    - Tag

```

## Properties

`ExcludeFilters`

If you specify this parameter, the stream sends metrics from all metric namespaces except
for the namespaces that you specify here. You cannot specify both `IncludeFilters`
and `ExcludeFilters` in the same metric stream.

When you modify the `IncludeFilters` or `ExcludeFilters` of an existing metric stream
in any way, the metric stream is effectively restarted, so after such a change you will get
only the datapoints that have a timestamp after the time of the update.

_Required_: No

_Type_: Array of [MetricStreamFilter](aws-properties-cloudwatch-metricstream-metricstreamfilter.md)

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FirehoseArn`

The ARN of the Amazon Kinesis Firehose delivery stream to use for this metric stream. This
Amazon Kinesis Firehose delivery stream must already exist and must be in the same account as the metric stream.

_Required_: No

_Type_: String

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncludeFilters`

If you specify this parameter, the stream sends only the metrics from the metric namespaces that you specify here.
You cannot specify both `IncludeFilters` and `ExcludeFilters` in the same metric stream.

When you modify the `IncludeFilters` or `ExcludeFilters` of an existing metric stream
in any way, the metric stream is effectively restarted, so after such a change you will get
only the datapoints that have a timestamp after the time of the update.

_Required_: No

_Type_: Array of [MetricStreamFilter](aws-properties-cloudwatch-metricstream-metricstreamfilter.md)

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncludeLinkedAccountsMetrics`

If you are creating a metric stream in a monitoring account, specify `true` to include
metrics from source accounts that are linked to this monitoring account, in the metric stream. The default is `false`.

For more information about linking accounts, see
[CloudWatch cross-account observability](../../../amazoncloudwatch/latest/monitoring/cloudwatch-unified-cross-account.md)

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

If you are creating a new metric stream, this is the name for the new stream.
The name must be different than the names of other metric streams in this account and Region.

If you are updating a metric stream, specify the name of that stream here.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`OutputFormat`

The output format for the stream. Valid values are `json`, `opentelemetry1.0` and
`opentelemetry0.7` For more information about metric stream output formats, see
[Metric streams output formats](../../../amazoncloudwatch/latest/monitoring/cloudwatch-metric-streams-formats.md).

This parameter is required.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The ARN of an IAM role that this metric stream will use to access Amazon Kinesis Firehose
resources. This IAM role must already exist and must be in the same account as the metric stream.
This IAM role must include the `firehose:PutRecord` and `firehose:PutRecordBatch`
permissions.

_Required_: No

_Type_: String

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StatisticsConfigurations`

By default, a
metric stream always sends the MAX, MIN, SUM, and SAMPLECOUNT statistics for each metric that is streamed.
You can use this parameter to have the metric stream also send additional statistics in the stream. This
array can have up to 100 members.

For each entry in this array, you specify one or more metrics and the list of additional statistics to
stream for those metrics. The additional statistics that you can stream depend on the stream's `OutputFormat`.
If the `OutputFormat` is `json`, you can stream any additional statistic that is supported by
CloudWatch, listed in
[CloudWatch statistics definitions](../../../amazoncloudwatch/latest/monitoring/statistics-definitions.md). If the `OutputFormat` is
OpenTelemetry, you can stream percentile statistics.

_Required_: No

_Type_: Array of [MetricStreamStatisticsConfiguration](aws-properties-cloudwatch-metricstream-metricstreamstatisticsconfiguration.md)

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to the metric stream.

For more information, see [Tag](../userguide/aws-properties-resource-tags.md).

_Required_: No

_Type_: Array of [Tag](aws-properties-cloudwatch-metricstream-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the name of the metric stream.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The ARN of the metric stream.

`CreationDate`

The date that the metric stream was originally created.

`LastUpdateDate`

The date that the metric stream was most recently updated.

`State`

The state of the metric stream, either `running` or `stopped`.

## Examples

### Metric stream example

The following example creates a metric stream that streams only the metrics in the
`AWS/ELB` and `AWS/EC2` namespaces.

#### YAML

```yaml

Resources:
  MyMetricStream:
    Type: 'AWS::CloudWatch::MetricStream'
    Properties:
      OutputFormat: 'json'
      FirehoseArn: 'arn:aws:firehose:us-east-1:123456789012:deliverystream/MyDeliveryStream'
      RoleArn: 'arn:aws:iam::123456789012:role/service-role/MyRole'
      IncludeFilters:
        - Namespace: AWS/ELB
        - Namespace: AWS/EC2
```

#### JSON

```json

{
  "Type" : "AWS::CloudWatch::MetricStream",
  "Properties" : {
    "FirehoseArn" : "arn:aws:firehose:us-east-1:123456789012:deliverystream/MyDeliveryStream",
    "IncludeFilters" : [
      {
        "Namespace": "AWS/ELB"
      },
      {
        "Namespace": "AWS/EC2"
      }
    ],
    "Name" : "MyMetricStream",
    "OutputFormat" : "json",
    "RoleArn" : "arn:aws:iam::123456789012:role/service-role/MyRole"
  }
}
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

MetricStreamFilter

All content copied from https://docs.aws.amazon.com/.
