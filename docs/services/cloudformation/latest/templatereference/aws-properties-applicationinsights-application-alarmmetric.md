This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApplicationInsights::Application AlarmMetric

The `AWS::ApplicationInsights::Application AlarmMetric` property type defines a metric to monitor for the component.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AlarmMetricName" : String
}

```

### YAML

```yaml

  AlarmMetricName: String

```

## Properties

`AlarmMetricName`

The name of the metric to be monitored for the component. For metrics supported by Application Insights, see [Logs and metrics supported by Amazon CloudWatch Application Insights](../../../amazoncloudwatch/latest/monitoring/appinsights-logs-and-metrics.md).

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Alarm

ComponentConfiguration

All content copied from https://docs.aws.amazon.com/.
