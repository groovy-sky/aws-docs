This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApplicationInsights::Application ConfigurationDetails

The `AWS::ApplicationInsights::Application ConfigurationDetails` property type specifies the configuration settings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AlarmMetrics" : [ AlarmMetric, ... ],
  "Alarms" : [ Alarm, ... ],
  "HAClusterPrometheusExporter" : HAClusterPrometheusExporter,
  "HANAPrometheusExporter" : HANAPrometheusExporter,
  "JMXPrometheusExporter" : JMXPrometheusExporter,
  "Logs" : [ Log, ... ],
  "NetWeaverPrometheusExporter" : NetWeaverPrometheusExporter,
  "Processes" : [ Process, ... ],
  "SQLServerPrometheusExporter" : SQLServerPrometheusExporter,
  "WindowsEvents" : [ WindowsEvent, ... ]
}

```

### YAML

```yaml

  AlarmMetrics:
    - AlarmMetric
  Alarms:
    - Alarm
  HAClusterPrometheusExporter:
    HAClusterPrometheusExporter
  HANAPrometheusExporter:
    HANAPrometheusExporter
  JMXPrometheusExporter:
    JMXPrometheusExporter
  Logs:
    - Log
  NetWeaverPrometheusExporter:
    NetWeaverPrometheusExporter
  Processes:
    - Process
  SQLServerPrometheusExporter:
    SQLServerPrometheusExporter
  WindowsEvents:
    - WindowsEvent

```

## Properties

`AlarmMetrics`

A list of metrics to monitor for the component. All component types can use `AlarmMetrics`.

_Required_: No

_Type_: Array of [AlarmMetric](aws-properties-applicationinsights-application-alarmmetric.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Alarms`

A list of alarms to monitor for the component. All component types can use
`Alarm`.

_Required_: No

_Type_: Array of [Alarm](aws-properties-applicationinsights-application-alarm.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HAClusterPrometheusExporter`

The HA cluster Prometheus Exporter settings.

_Required_: No

_Type_: [HAClusterPrometheusExporter](aws-properties-applicationinsights-application-haclusterprometheusexporter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HANAPrometheusExporter`

The HANA DB Prometheus Exporter settings.

_Required_: No

_Type_: [HANAPrometheusExporter](aws-properties-applicationinsights-application-hanaprometheusexporter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`JMXPrometheusExporter`

A list of Java metrics to monitor for the component.

_Required_: No

_Type_: [JMXPrometheusExporter](aws-properties-applicationinsights-application-jmxprometheusexporter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Logs`

A list of logs to monitor for the component. Only Amazon EC2 instances can use
`Logs`.

_Required_: No

_Type_: Array of [Log](aws-properties-applicationinsights-application-log.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NetWeaverPrometheusExporter`

Property description not available.

_Required_: No

_Type_: [NetWeaverPrometheusExporter](aws-properties-applicationinsights-application-netweaverprometheusexporter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Processes`

Property description not available.

_Required_: No

_Type_: Array of [Process](aws-properties-applicationinsights-application-process.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SQLServerPrometheusExporter`

Property description not available.

_Required_: No

_Type_: [SQLServerPrometheusExporter](aws-properties-applicationinsights-application-sqlserverprometheusexporter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WindowsEvents`

A list of Windows Events to monitor for the component. Only Amazon EC2 instances
running on Windows can use `WindowsEvents`.

_Required_: No

_Type_: Array of [WindowsEvent](aws-properties-applicationinsights-application-windowsevent.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ComponentMonitoringSetting

CustomComponent

All content copied from https://docs.aws.amazon.com/.
