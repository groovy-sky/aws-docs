This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApplicationInsights::Application JMXPrometheusExporter

The `AWS::ApplicationInsights::Application JMXPrometheusExporter` property type
defines the JMXPrometheus Exporter configuration. For more information, see the
[component configuration](../../../amazoncloudwatch/latest/monitoring/component-config-sections.md#component-configuration-prometheus) in the CloudWatch Application Insights
documentation.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "HostPort" : String,
  "JMXURL" : String,
  "PrometheusPort" : String
}

```

### YAML

```yaml

  HostPort: String
  JMXURL: String
  PrometheusPort: String

```

## Properties

`HostPort`

The host and port to connect to through remote JMX. Only one of `jmxURL` and `hostPort` can be
specified.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`JMXURL`

The complete JMX URL to connect to.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrometheusPort`

The target port to send Prometheus metrics to. If not specified, the default port `9404` is used.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HANAPrometheusExporter

Log

All content copied from https://docs.aws.amazon.com/.
