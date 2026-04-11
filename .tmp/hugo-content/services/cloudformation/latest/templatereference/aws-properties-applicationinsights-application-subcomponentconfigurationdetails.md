This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApplicationInsights::Application SubComponentConfigurationDetails

The `AWS::ApplicationInsights::Application SubComponentConfigurationDetails` property type specifies the configuration settings of the sub-components.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AlarmMetrics" : [ AlarmMetric, ... ],
  "Logs" : [ Log, ... ],
  "Processes" : [ Process, ... ],
  "WindowsEvents" : [ WindowsEvent, ... ]
}

```

### YAML

```yaml

  AlarmMetrics:
    - AlarmMetric
  Logs:
    - Log
  Processes:
    - Process
  WindowsEvents:
    - WindowsEvent

```

## Properties

`AlarmMetrics`

A list of metrics to monitor for the component. All component types can use `AlarmMetrics`.

_Required_: No

_Type_: Array of [AlarmMetric](aws-properties-applicationinsights-application-alarmmetric.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Logs`

A list of logs to monitor for the component. Only Amazon EC2 instances can use
`Logs`.

_Required_: No

_Type_: Array of [Log](aws-properties-applicationinsights-application-log.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Processes`

Property description not available.

_Required_: No

_Type_: Array of [Process](aws-properties-applicationinsights-application-process.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WindowsEvents`

A list of Windows Events to monitor for the component. Only Amazon EC2 instances running
on Windows can use `WindowsEvents`.

_Required_: No

_Type_: Array of [WindowsEvent](aws-properties-applicationinsights-application-windowsevent.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SQLServerPrometheusExporter

SubComponentTypeConfiguration

All content copied from https://docs.aws.amazon.com/.
