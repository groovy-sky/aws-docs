This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApplicationInsights::Application Process

The `Process` property type specifies Property description not available. for an [AWS::ApplicationInsights::Application](aws-resource-applicationinsights-application.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AlarmMetrics" : [ AlarmMetric, ... ],
  "ProcessName" : String
}

```

### YAML

```yaml

  AlarmMetrics:
    - AlarmMetric
  ProcessName: String

```

## Properties

`AlarmMetrics`

Property description not available.

_Required_: Yes

_Type_: Array of [AlarmMetric](aws-properties-applicationinsights-application-alarmmetric.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProcessName`

Property description not available.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_,-]+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NetWeaverPrometheusExporter

SQLServerPrometheusExporter

All content copied from https://docs.aws.amazon.com/.
